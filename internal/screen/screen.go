package screen

import (
	"image"

	"github.com/kbinani/screenshot"
)

//NumActiveScreens возвращает количество доступных экранов.
func NumActiveScreens() int {
	return screenshot.NumActiveDisplays()
}

//capture захватывает всю облать экрана с номером, заданным в displayIndex
func capture(displayIndex int) (*image.RGBA, error) {
	return screenshot.CaptureDisplay(displayIndex)
}

//AlerterFunc - это адаптер для того, чтобы можно было использовать обычные функции в качестве
//проигрывателей оповещения об изменении на экране
type AlerterFunc func() error

//Play проигрывает оповещение об изменении на экране
func (a AlerterFunc) Play() error {
	return a()
}
