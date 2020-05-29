package screen

import (
	"context"
	"fmt"
	"image"
	"os"
	"time"

	"github.com/corona10/goimagehash"
)

//Порог выше которого будет считаться, что изображение изменилось
const threshold = 1

//Tracker предназначен для отслеживания изменений на экране.
type Tracker struct {
	display        int
	captureDisplay func(int) (*image.RGBA, error)
	alerter        AlerterFunc
}

//NewTracker возвращает новый Tracker для экрана с номером display.
//Нумерация должна начинаться с 0 для первого активного экрана и по возрастающей: 1,2,3 ...
//Если будет передан номер экрана, которого не существует в системе, то при попытке
//вызвать метод Tracker.TrackChanges() он возвратит ошибку.
func NewTracker(display int) *Tracker {
	return &Tracker{
		display:        display,
		captureDisplay: capture,
		alerter: func() error {
			fmt.Printf("%s the screen has changed.\n", time.Now().Format("2006/02/01 15:04:05"))

			return nil
		},
	}
}

//WithAlert возвращает Tracker, c алгоритмом оповещения об изменении на экране, заданным в параметре alerter.
func (t *Tracker) WithAlert(alerter AlerterFunc) *Tracker {
	t.alerter = alerter

	return t
}

//TrackChanges отслеживает в цикле изменения на экране с интервалом, заданным в timeout.
//Если на экране произошли какие-то изменения, то происходит оповещение.
//ctx используется для того, чтобы завершить отслеживание изменений.
func (t Tracker) TrackChanges(ctx context.Context, timeout time.Duration) error {
	ticker := time.NewTicker(timeout)

	img1, err := t.captureDisplay(t.display)
	if err != nil {
		return err
	}

loop:
	for {
		select {
		case <-ticker.C:
			img2, err := t.captureDisplay(t.display)
			if err != nil {
				return err
			}

			if distance := compareImages(img1, img2); distance > threshold {
				if err := t.alerter.Play(); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "play alert error: %s\n", err)
				}
			}

			img1 = img2
		case <-ctx.Done():
			break loop
		}
	}

	return nil
}

func compareImages(img1, img2 image.Image) int {
	width := 80
	height := img1.Bounds().Dy() / (img1.Bounds().Dx() / width)

	img1Hash, _ := goimagehash.ExtDifferenceHash(img1, width, height)
	img2Hash, _ := goimagehash.ExtDifferenceHash(img2, width, height)
	distance, _ := img1Hash.Distance(img2Hash)

	return distance
}
