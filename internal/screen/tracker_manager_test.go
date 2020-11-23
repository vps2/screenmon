package screen

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeTracker struct {
	calls             int32
	shouldReturnError bool
}

func (tm *fakeTracker) TrackChanges(context.Context, time.Duration) error {
	atomic.AddInt32(&tm.calls, 1)

	if tm.shouldReturnError {
		fmt.Println("an error")
		return errors.New("an error")
	}

	return nil
}

func TestTrackerManager_Start_SimultaneousFromMultipleGoroutines(t *testing.T) {
	ft := fakeTracker{}
	mgr := NewTrackerManager(&ft, 0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mgr.Start()
		}()
	}
	wg.Wait()

	assert.Equal(t, int32(1), atomic.LoadInt32(&ft.calls))
}
func TestTrackerManager_Start_WhenTrackerReturnsAnError(t *testing.T) {
	ft := fakeTracker{
		shouldReturnError: true,
	}
	mgr := NewTrackerManager(&ft, 0)
	mgr.Start()
	<-mgr.Done()

	assert.Error(t, mgr.Error())
}

func TestTrackerManager_StopMultipleTimes(t *testing.T) {
	ft := fakeTracker{}
	mgr := NewTrackerManager(&ft, 0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			assert.NotPanics(t, mgr.Stop) //при повторном вызове метода, не должно паниковать
		}()
	}

	wg.Wait()
}
