package audio

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var once sync.Once

func Play(audioData []byte) {
	streamer, format, err := mp3.Decode(io.NopCloser(bytes.NewReader(audioData)))
	if err != nil {
		fmt.Println(err)
		return
	}
	once.Do(func() {
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			fmt.Println(err)
		}
	})
	speaker.Play(streamer)
}
