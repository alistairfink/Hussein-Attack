package resources

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

type sound struct {
	path              string
	isBackgroundMusic bool
}

var sounds = map[string]sound {
	"menu": {
		path:              menuMusicPath,
		isBackgroundMusic: true,
	},
	"game": {
		path:              gameMusicPath,
		isBackgroundMusic: true,
	},
	"laser": {
		path:              laserSoundPath,
		isBackgroundMusic: false,
	},
}

func getSoundFile(soundPath string) (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open(soundPath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return streamer, format
}

// GetSoundEffectsBuffer returns buffer of all game sound effects to avoid delays/load times
// (cannot do same for background music as it consumes too much memory)
func GetSoundEffectsBuffer() *beep.Buffer {
	// declare buffer
	var buffer *beep.Buffer

	// loop over and add sound effects to buffer
	for _, thisSound := range sounds {
		if !thisSound.isBackgroundMusic {
			streamer, format := getSoundFile(thisSound.path)
			if buffer == nil {
				buffer = beep.NewBuffer(format)
			}
			buffer.Append(streamer)
			streamer.Close()
		}
	}

	return buffer
}

func playSound(context string, buffer *beep.Buffer, initializeSpeaker bool) <-chan bool {
	done := make(chan bool)
	go func() {
		thisSound := sounds[context]
		if thisSound.isBackgroundMusic {
			streamer, format := getSoundFile(thisSound.path)
			audio := beep.Loop(-1, streamer)

			if initializeSpeaker {
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))
			} else {
				speaker.Clear()
			}
			speaker.Play(audio)
			defer close(done)
			select {}
		} else {
			// TODO: Make this faster, sound effect buffer doesn't seem to help much
			speaker.Play(buffer.Streamer(0, buffer.Len()))
		}
	}()

	return done
}

func PlayMenuMusic() {
	playSound("menu", nil, true)
}

func PlayGameMusic() {
	playSound("game", nil, false)
}

func PlayLaserSound(buffer *beep.Buffer) {
	playSound("laser", buffer, false)
}