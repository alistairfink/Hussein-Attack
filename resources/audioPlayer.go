package resources

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type sound struct {
	index             int
	path              string
	isBackgroundMusic bool
}

// GetSounds gets all sounds for game
func GetSounds() map[string]*sound {
	return map[string]*sound{
		"menu": {
			index:             0,
			path:              menuMusicPath,
			isBackgroundMusic: true,
		},
		"game": {
			index:             1,
			path:              gameMusicPath,
			isBackgroundMusic: true,
		},
		"laser": {
			index:             2,
			path:              laserSoundPath,
			isBackgroundMusic: false,
		},
	}
}

// GetSoundFile returns streamer and format of decoded mp3 file
func GetSoundFile(soundPath string) (beep.StreamSeekCloser, beep.Format) {
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
	sounds := GetSounds()
	for k := range sounds {
		thisSound := sounds[k]
		if !thisSound.isBackgroundMusic {
			streamer, format := GetSoundFile(thisSound.path)
			if buffer == nil {
				buffer = beep.NewBuffer(format)
			}
			buffer.Append(streamer)
			streamer.Close()
		}
	}

	return buffer
}

// PlaySound plays a specific sound (background music or sound effect)
func PlaySound(context string, buffer *beep.Buffer, initializeSpeaker bool) <-chan bool {
	done := make(chan bool)
	go func() {
		thisSound := GetSounds()[context]
		// only reset audio if background music
		if thisSound.isBackgroundMusic {
			// get sound file
			streamer, format := GetSoundFile(thisSound.path)
			// create loop of new audio
			audio := beep.Loop(-1, streamer)

			// initialize speaker if needed
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
