package audio

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_mixer

#include <SDL2/SDL.h>
#include <SDL2/SDL_mixer.h>
*/
import "C"
import (
	"fmt"
)

type Music struct {
	volume int
}

func NewMusic() *Music {
	m := &Music{
		volume: 64, // 0 ~ 128
	}

	err := m.SetVolume(64)
	if err != nil {
		//TODO: Handle this error
		panic(err)
	}

	return m
}

func (m *Music) Play(music *C.Mix_Music, loop bool) {
	if m.IsPlaying() {
		m.Stop()
	}

	if loop {
		C.Mix_PlayMusic(music, C.int(-1))
	} else {
		C.Mix_PlayMusic(music, C.int(1))
	}
}

func (m *Music) PlayWithFadeIn(music *C.Mix_Music, loop bool, ms int) {
	if m.IsPlaying() {
		m.Stop()
	}

	if loop {
		C.Mix_FadeInMusic(music, C.int(-1), C.int(ms))
	} else {
		C.Mix_FadeInMusic(music, C.int(1), C.int(ms))
	}
}

func (m *Music) Stop() {
	C.Mix_HaltMusic()
}

func (m *Music) StopWithFadeOut(ms int) {
	C.Mix_FadeOutMusic(C.int(ms))
}

func (m *Music) GetVolume() (volume int) {
	return m.volume
}

func (m *Music) SetVolume(v int) error {
	if v < 0 || v > 128 {
		return fmt.Errorf("Music Volume value range 0 ~ 128, got=%d", v)
	}

	m.volume = v
	C.Mix_VolumeMusic(C.int(v))

	return nil
}

func (m *Music) IsPlaying() bool {
	return C.Mix_PlayingMusic() != 0
}
