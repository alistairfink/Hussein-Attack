package resources

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

var soundEffectPaths = [1]string{laserSoundPath}

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
	for _, soundEffectPath := range soundEffectPaths {
		streamer, format := getSoundFile(soundEffectPath)
		if buffer == nil {
			buffer = beep.NewBuffer(format)
		}
		buffer.Append(streamer)
		streamer.Close()
	}

	return buffer
}

func playSound(filePath string, buffer *beep.Buffer, initializeSpeaker bool, isEffect bool) <-chan bool {
	done := make(chan bool)
	go func() {
		if !isEffect {
			streamer, format := getSoundFile(filePath)
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
			index := 0
			for i, soundEffectPath := range soundEffectPaths {
				if soundEffectPath == filePath {
					index = i
					break
				}
			}

			// TODO: Make this faster, sound effect buffer doesn't seem to help much
			speaker.Play(buffer.Streamer(index, buffer.Len()))
		}
	}()

	return done
}

func PlayMenuMusic() {
	playSound(menuMusicPath, nil, true, false)
}

func PlayGameMusic() {
	playSound(gameMusicPath, nil, false, false)
}

func PlayLaserSound(buffer *beep.Buffer) {
	playSound(laserSoundPath, buffer, false, true)
}
