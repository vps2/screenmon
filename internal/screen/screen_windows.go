package screen

import (
	"fmt"
	"image"

	"github.com/kbinani/screenshot"
)

//NumActiveScreens возвращает количество доступных экранов.
func NumActiveScreens() int {
	return screenshot.NumActiveDisplays()
}

//capture захватывает заданную область экрана с номером, заданным в displayIndex
func capture(displayIndex int, captureArea Area) (*image.RGBA, error) {
	if displayIndex < 0 || displayIndex >= NumActiveScreens() {
		return nil, fmt.Errorf("the screen with the number %d is missing", displayIndex)
	}

	displayBounds := screenshot.GetDisplayBounds(displayIndex)
	minX := displayBounds.Min.X + captureArea.X1
	minY := displayBounds.Min.Y + captureArea.Y1
	maxX := displayBounds.Min.X + captureArea.X2
	maxY := displayBounds.Min.Y + captureArea.Y2

	return screenshot.CaptureRect(image.Rect(minX, minY, maxX, maxY))
}
