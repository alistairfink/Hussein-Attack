package resources

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

type audioFile struct {
	filePath     string
	streamSeeker *beep.StreamSeeker
	loopFile     bool
}

type audioPlayer struct {
	menuMusic   audioFile
	gameMusic   audioFile
	laserEffect audioFile
}

func NewAudioPlayer() audioPlayer {
	obj := audioPlayer{}
	obj.menuMusic = audioFile{
		filePath:     menuMusicPath,
		streamSeeker: nil,
		loopFile:     true,
	}

	obj.gameMusic = audioFile{
		filePath:     gameMusicPath,
		streamSeeker: nil,
		loopFile:     true,
	}

	obj.laserEffect = audioFile{
		filePath:     laserSoundPath,
		streamSeeker: initEffectBuffer(laserSoundPath),
		loopFile:     false,
	}

	_, tempFormat := getSoundFile(menuMusicPath)
	speaker.Init(tempFormat.SampleRate, tempFormat.SampleRate.N(time.Second/2))

	return obj
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

func initEffectBuffer(filePath string) *beep.StreamSeeker {
	streamSeekCloser, format := getSoundFile(filePath)
	defer streamSeekCloser.Close()
	buffer := beep.NewBuffer(format)
	buffer.Append(streamSeekCloser)
	streamer := buffer.Streamer(0, buffer.Len())
	return &streamer
}

func (this *audioPlayer) playSound(soundFile audioFile) <-chan bool {
	done := make(chan bool)
	go func() {
		if soundFile.streamSeeker == nil {
			streamer, _ := getSoundFile(soundFile.filePath)
			audio := beep.Loop(-1, streamer)
			speaker.Clear()
			speaker.Play()
			speaker.Play(audio)
		} else {
			speaker.Play(*soundFile.streamSeeker)
		}

		defer close(done)

	}()

	return done
}

func (this *audioPlayer) PlayMenuMusic() {
	this.playSound(this.menuMusic)
}

func (this *audioPlayer) PlayGameMusic() {
	this.playSound(this.gameMusic)
}

func (this *audioPlayer) PlayLaserSound() {
	this.playSound(this.laserEffect)
}

// GetSoundEffectsBuffer returns buffer of all game sound effects to avoid delays/load times
// (cannot do same for background music as it consumes too much memory)
// func GetSoundEffectsBuffer() *beep.Buffer {
// 	var buffer *beep.Buffer

// 	// loop over and add sound effects to buffer
// 	for _, soundEffectPath := range soundEffectPaths {
// 		streamer, format := getSoundFile(soundEffectPath)
// 		if buffer == nil {
// 			buffer = beep.NewBuffer(format)
// 		}
// 		buffer.Append(streamer)
// 		streamer.Close()
// 	}

// 	return buffer
// }

// func playSound(filePath string, buffer *beep.Buffer, initializeSpeaker bool, isEffect bool) <-chan bool {
// 	done := make(chan bool)
// 	go func() {
// 		if !isEffect {
// 			streamer, format := getSoundFile(filePath)
// 			audio := beep.Loop(-1, streamer)

// 			if initializeSpeaker {
// 				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))
// 			} else {
// 				speaker.Clear()
// 			}
// 			speaker.Play(audio)
// 			defer close(done)
// 			select {}
// 		} else {
// 			index := 0
// 			for i, soundEffectPath := range soundEffectPaths {
// 				if soundEffectPath == filePath {
// 					index = i
// 					break
// 				}
// 			}

// 			println(1)
// 			// TODO: Make this faster, sound effect buffer doesn't seem to help much
// 			println("", buffer.Len())
// 			buffer := buffer.Streamer(index, buffer.Len())
// 			println(2)
// 			speaker.Play(buffer)
// 			println(3)
// 		}
// 	}()

// 	return done
// }

// func PlayMenuMusic() {
// 	playSound(menuMusicPath, nil, true, false)
// }

// func PlayGameMusic() {
// 	playSound(gameMusicPath, nil, false, false)
// }

// var played bool = false

// func PlayLaserSound(buffer *beep.Buffer) {
// 	if !played {
// 		println("Playing...")
// 		playSound(laserSoundPath, buffer, false, true)
// 	}
// 	played = true
// }
