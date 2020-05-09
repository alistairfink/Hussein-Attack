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
	filePath string
	isMusic  bool
}

type AudioPlayer struct {
	menuMusic   audioFile
	gameMusic   audioFile
	laserEffect audioFile
}

func NewAudioPlayer() AudioPlayer {
	obj := AudioPlayer{}
	obj.menuMusic = audioFile{
		filePath: menuMusicPath,
		isMusic:  true,
	}

	obj.gameMusic = audioFile{
		filePath: gameMusicPath,
		isMusic:  true,
	}

	obj.laserEffect = audioFile{
		filePath: laserSoundPath,
		isMusic:  false,
	}

	_, tempFormat := obj.menuMusic.getSoundFile()
	speaker.Init(tempFormat.SampleRate, tempFormat.SampleRate.N(time.Second/20))

	return obj
}

func (this *audioFile) getSoundFile() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open(this.filePath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return streamer, format
}

func (this *AudioPlayer) playSound(soundFile audioFile) <-chan bool {
	done := make(chan bool)
	go func() {
		streamer, format := soundFile.getSoundFile()
		defer streamer.Close()
		var audio beep.Streamer
		if soundFile.isMusic {
			speaker.Clear()
			audio = beep.Loop(-1, streamer)
		} else {
			buffer := beep.NewBuffer(format)
			buffer.Append(streamer)
			audio = buffer.Streamer(0, buffer.Len())
		}

		speaker.Play(audio)
		defer close(done)
		select {}
	}()

	return done
}

func (this *AudioPlayer) PlayMenuMusic() {
	this.playSound(this.menuMusic)
}

func (this *AudioPlayer) PlayGameMusic() {
	this.playSound(this.gameMusic)
}

func (this *AudioPlayer) PlayLaserSound() {
	this.playSound(this.laserEffect)
}
