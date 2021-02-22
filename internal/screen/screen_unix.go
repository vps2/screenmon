// +build linux freebsd netbsd openbsd solaris

package screen

import (
	"fmt"
	"image"
	"image/color"

	"github.com/BurntSushi/xgb"
	mshm "github.com/BurntSushi/xgb/shm"
	"github.com/BurntSushi/xgb/xinerama"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/gen2brain/shm"
)

var conn *xgb.Conn

func init() {
	var err error
	conn, err = xgb.NewConn()
	if err != nil {
		panic(fmt.Errorf("screen: init error: %w", err))
	}

	err = xinerama.Init(conn)
	if err != nil {
		panic(fmt.Errorf("screen: init error: %w", err))
	}
}

//NumActiveScreens возвращает количество доступных экранов.
func NumActiveScreens() int {
	reply, err := xinerama.QueryScreens(conn).Reply()
	if err != nil {
		return 0
	}

	return int(reply.Number)
}

//capture захватывает заданную область экрана с номером, заданным в displayIndex
func capture(displayIndex int, captureArea Area) (*image.RGBA, error) {
	if displayIndex < 0 || displayIndex >= NumActiveScreens() {
		return nil, fmt.Errorf("the screen with the number %d is missing", displayIndex)
	}

	displayBounds := getDisplayBounds(displayIndex)
	minX := displayBounds.Min.X + captureArea.X1
	minY := displayBounds.Min.Y + captureArea.Y1
	maxX := displayBounds.Min.X + captureArea.X2
	maxY := displayBounds.Min.Y + captureArea.Y2

	return captureRect(image.Rect(minX, minY, maxX, maxY))
}

func getDisplayBounds(displayIndex int) (rect image.Rectangle) {
	reply, err := xinerama.QueryScreens(conn).Reply()
	if err != nil {
		return image.Rectangle{}
	}

	primary := reply.ScreenInfo[0]
	x0 := int(primary.XOrg)
	y0 := int(primary.YOrg)

	screen := reply.ScreenInfo[displayIndex]
	x := int(screen.XOrg) - x0
	y := int(screen.YOrg) - y0
	w := int(screen.Width)
	h := int(screen.Height)
	rect = image.Rect(x, y, x+w, y+h)

	return rect
}

func captureRect(rect image.Rectangle) (img *image.RGBA, e error) {
	return captureArea(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
}

func captureArea(x, y, width, height int) (img *image.RGBA, e error) {
	reply, err := xinerama.QueryScreens(conn).Reply()
	if err != nil {
		return nil, err
	}

	primary := reply.ScreenInfo[0]
	x0 := int(primary.XOrg)
	y0 := int(primary.YOrg)

	useShm := true
	err = mshm.Init(conn)
	if err != nil {
		useShm = false
	}

	screen := xproto.Setup(conn).DefaultScreen(conn)
	wholeScreenBounds := image.Rect(0, 0, int(screen.WidthInPixels), int(screen.HeightInPixels))
	targetBounds := image.Rect(x+x0, y+y0, x+x0+width, y+y0+height)
	intersect := wholeScreenBounds.Intersect(targetBounds)

	rect := image.Rect(0, 0, width, height)
	img = image.NewRGBA(rect)

	// Paint with opaque black
	index := 0
	for iy := 0; iy < height; iy++ {
		j := index
		for ix := 0; ix < width; ix++ {
			img.Pix[j+3] = 255
			j += 4
		}
		index += img.Stride
	}

	if !intersect.Empty() {
		var data []byte

		if useShm {
			shmSize := intersect.Dx() * intersect.Dy() * 4
			shmId, err := shm.Get(shm.IPC_PRIVATE, shmSize, shm.IPC_CREAT|0777)
			if err != nil {
				return nil, err
			}

			seg, err := mshm.NewSegId(conn)
			if err != nil {
				return nil, err
			}

			data, err = shm.At(shmId, 0, 0)
			if err != nil {
				return nil, err
			}

			mshm.Attach(conn, seg, uint32(shmId), false)

			defer mshm.Detach(conn, seg)
			defer shm.Rm(shmId)
			defer shm.Dt(data)

			_, err = mshm.GetImage(conn, xproto.Drawable(screen.Root),
				int16(intersect.Min.X), int16(intersect.Min.Y),
				uint16(intersect.Dx()), uint16(intersect.Dy()), 0xffffffff,
				byte(xproto.ImageFormatZPixmap), seg, 0).Reply()
			if err != nil {
				return nil, err
			}
		} else {
			xImg, err := xproto.GetImage(conn, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root),
				int16(intersect.Min.X), int16(intersect.Min.Y),
				uint16(intersect.Dx()), uint16(intersect.Dy()), 0xffffffff).Reply()
			if err != nil {
				return nil, err
			}

			data = xImg.Data
		}

		// BitBlt by hand
		offset := 0
		for iy := intersect.Min.Y; iy < intersect.Max.Y; iy++ {
			for ix := intersect.Min.X; ix < intersect.Max.X; ix++ {
				r := data[offset+2]
				g := data[offset+1]
				b := data[offset]
				img.SetRGBA(ix-(x+x0), iy-(y+y0), color.RGBA{r, g, b, 255})
				offset += 4
			}
		}
	}

	return img, e
}
