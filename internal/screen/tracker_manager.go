package screen

import (
	"context"
	"sync"
	"time"
)

type tracker interface {
	TrackChanges(context.Context, time.Duration) error
}

//TrackerManager - потоко-безопасная надстройка над screen.Tracker, позволяющая запускать,
//ставить на паузу и останавливать отслеживание изменений на экране.
type TrackerManager struct {
	mu      sync.Mutex
	started bool

	tracker tracker
	timeout time.Duration
	cancel  context.CancelFunc

	errMu sync.RWMutex
	err   error

	done chan interface{}
}

func NewTrackerManager(tr tracker, timeout time.Duration) *TrackerManager {
	mgr := TrackerManager{
		tracker: tr,
		timeout: timeout,
		done:    make(chan interface{}),
	}

	return &mgr
}

//Start запускает отслеживание изменений на экране.
func (mgr *TrackerManager) Start() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	if mgr.started {
		return
	} else {
		mgr.started = true
	}

	ctx, cancel := context.WithCancel(context.Background())
	mgr.cancel = cancel

	go func() {
		err := mgr.tracker.TrackChanges(ctx, mgr.timeout)
		if err != nil {
			mgr.errMu.Lock()
			mgr.err = err
			mgr.errMu.Unlock()

			mgr.Stop()
		}
	}()
}

//Stop останавливает (с закрытием канала done) отслеживание изменений на экране.
//После остановки отслеживания, нельзя повторно вызывать метод Start.
func (mgr *TrackerManager) Stop() {
	mgr.Pause()

	// Закрываем только если канал не закрыт, т.к. повторное закрытие вызывает панику.
	select {
	case <-mgr.done:
		//эта ветка сработает только тогда, когда канал будет закрыт, т.к.
		//мы в него ничего не пишем, а сразу закрываем
	default:
		close(mgr.done)
	}
}

//Pause приостанавливает отслеживание изменений на экране.
//Для продолжения отслеживания, следует вызвать метод Start.
func (mgr *TrackerManager) Pause() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	if mgr.started {
		mgr.cancel()
		mgr.started = false
	}
}

//Error возвращает ошибку, если завершение работы произошло из-за ошибки и nil в противном случае.
//Если TrackerManager ещё работает, то возвращается nil.
func (mgr *TrackerManager) Error() error {
	mgr.errMu.RLock()
	defer mgr.errMu.RUnlock()

	return mgr.err
}

//Done - канал, сигнализирующий о завершении работы.
func (mgr *TrackerManager) Done() <-chan interface{} {
	return mgr.done
}
