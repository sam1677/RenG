package image

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
*/
import "C"
import "internal/RVM/core/globaltype"

func (i *Image) GetImgaeTextureName(t *globaltype.SDL_Texture) string {
	i.lock.Lock()
	defer i.lock.Unlock()

	for name, image := range i.images {
		if image.texture == t {
			return name
		}
	}

	return ""
}

func (i *Image) GetImageTexture(name string) *globaltype.SDL_Texture {
	i.lock.Lock()
	defer i.lock.Unlock()

	if image, ok := i.images[name]; !ok {
		return nil
	} else {
		return image.texture
	}
}

func (i *Image) GetImageWidth(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()

	if image, ok := i.images[name]; !ok {
		return 0
	} else {
		return int(image.surface.w)
	}
}

func (i *Image) GetImageHeight(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()

	if image, ok := i.images[name]; !ok {
		return 0
	} else {
		return int(image.surface.h)
	}
}
