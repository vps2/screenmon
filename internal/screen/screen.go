package screen

//AlerterFunc - это адаптер для того, чтобы можно было использовать обычные функции в качестве
//проигрывателей оповещения об изменении на экране
type AlerterFunc func() error

//Play проигрывает оповещение об изменении на экране
func (a AlerterFunc) Play() error {
	return a()
}
