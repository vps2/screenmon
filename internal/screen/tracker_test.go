package screen

import (
	"context"
	"image"
	"image/color"
	"image/draw"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var img1, img2 *image.RGBA

func init() {
	img1 = image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img1, img1.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	for x := 0; x < img1.Rect.Dx(); x++ {
		img1.Set(x, img1.Rect.Dy()/2, color.Black)
	}

	img2 = image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img2, img2.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	for y := 0; y < img2.Rect.Dy(); y++ {
		img2.Set(img2.Rect.Dx()/2, y, color.Black)
	}
}

func TestTracker_TrackChanges_WithCorrectDisplayNumber(t *testing.T) {
	tracker := NewTracker(0, Area{0, 0, 100, 100})

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond+10)
	defer cancel()

	assert.NoError(t, tracker.TrackChanges(ctx, time.Millisecond), "The code return error!")
}

func TestTracker_TrackChanges_WithWrongDisplayNumber(t *testing.T) {
	displayNumbers := []int{-1, 100}

	for i, display := range displayNumbers {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tracker := NewTracker(display, Area{0, 0, 100, 100})
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond+10)
			defer cancel()
			assert.Error(t, tracker.TrackChanges(ctx, time.Millisecond), "The code did not return error!")
		})
	}
}

func TestTracker_TrackChanges_DifferentScreenCapture(t *testing.T) {
	var captureFuncCallCounter int = 0
	capture := func(int, Area) (*image.RGBA, error) {
		captureFuncCallCounter++

		if captureFuncCallCounter%2 == 0 {
			return img1, nil
		}

		return img2, nil
	}

	playAlertFuncCallCounter := 0
	playAlert := func() error {
		playAlertFuncCallCounter++
		return nil
	}

	tracker := NewTracker(0, Area{0, 0, 100, 100}).WithAlert(playAlert)
	tracker.capture = capture

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond+10)
	defer cancel()

	_ = tracker.TrackChanges(ctx, time.Millisecond)

	assert.GreaterOrEqual(t, playAlertFuncCallCounter, 1)
}

func TestTracker_TrackChanges_SameScreenCapture(t *testing.T) {
	capture := func(int, Area) (*image.RGBA, error) {
		return img1, nil
	}

	playAlertFuncCallCounter := 0
	playAlert := func() error {
		playAlertFuncCallCounter++
		return nil
	}

	tracker := NewTracker(0, Area{0, 0, 100, 100}).WithAlert(playAlert)
	tracker.capture = capture

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond+10)
	defer cancel()

	_ = tracker.TrackChanges(ctx, time.Millisecond)

	assert.Equal(t, 0, playAlertFuncCallCounter)
}
