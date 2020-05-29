package alert

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go -prefix "../../assets" ../../assets/

//Play проигрывает сэмпл
func Play() error {
	data, err := Asset("alert.mp3")
	if err != nil {
		return err
	}
	mp3Reader := ioutil.NopCloser(bytes.NewReader(data))

	streamer, format, err := mp3.Decode(mp3Reader)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/5))
	defer speaker.Close()

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}
