package audio

import (
	"bytes"
	"fmt"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type reader struct {
	*bytes.Reader
}

func (*reader) Close() error { return nil }

func play(audioData []byte) (err error) {
	r := &reader{
		Reader: bytes.NewReader(audioData),
	}
	streamer, format, err := mp3.Decode(r)
	if err != nil {
		return err
	}
	if err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return err
	}
	speaker.Play(streamer)
	return nil
}

func Play(audioData []byte) {
	if err := play(audioData); err != nil {
		fmt.Println(err)
	}
}
