package audio

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func play(audioData []byte) (err error) {
	streamer, format, err := mp3.Decode(io.NopCloser(bytes.NewReader(audioData)))
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
