package audio

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_mixer

#include <SDL2/SDL.h>
#include <SDL2/SDL_mixer.h>
*/
import "C"

type Channel struct {
	chanIndex int
	volume    int
}

var (
	assignedChannel = 0
)

func NewChannel() *Channel {
	ch := &Channel{
		chanIndex: assignedChannel,
		volume:    64, // 0 ~ 128
	}

	assignedChannel++

	return ch
}

func (c *Channel) Play(chunk *C.Mix_Chunk) {
	C.Mix_PlayChannelTimed(C.int(c.chanIndex), chunk, C.int(0), C.int(-1))
}

func (c *Channel) GetVolume() (volume int) {
	return c.volume
}

func (c *Channel) SetVolume(volume int) {
	if volume < 0 || volume > 128 {
		return
	}

	c.volume = volume

	C.Mix_Volume(C.int(c.chanIndex), C.int(volume))
}
