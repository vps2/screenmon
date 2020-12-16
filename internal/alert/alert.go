package alert

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go -prefix "../../assets" ../../assets/

type readSeekNopCloser struct {
	io.ReadSeeker
}

func (readSeekNopCloser) Close() error { return nil }

type Player struct {
	stream beep.StreamSeekCloser
}

func NewPlayer() *Player {
	const errTmpl = "failed to create the player: %w"

	data, err := Asset("alert.mp3")
	if err != nil {
		panic(fmt.Errorf(errTmpl, err))
	}

	sampleReader := readSeekNopCloser{bytes.NewReader(data)}
	stream, format, err := mp3.Decode(sampleReader)
	if err != nil {
		panic(fmt.Errorf(errTmpl, err))
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/5)); err != nil {
		panic(fmt.Errorf(errTmpl, err))
	}

	player := Player{
		stream: stream,
	}

	return &player
}

func (p *Player) Play() error {
	if err := p.stream.Seek(0); err != nil {
		return err
	}

	done := make(chan bool)

	speaker.Play(p.stream, beep.Callback(func() {
		done <- true
	}))

	<-done

	return nil
}

func (p *Player) Close() {
	defer p.stream.Close()
	defer speaker.Close()
}
