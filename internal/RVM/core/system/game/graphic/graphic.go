package graphic

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>
*/
import "C"
import (
	"internal/RVM/core/globaltype"
	"internal/RVM/core/object"
	"internal/RVM/core/system/game/graphic/image"
	"internal/RVM/core/system/game/graphic/sprite"
	"internal/RVM/core/system/game/graphic/video"
	"sync"
	"time"
	"unsafe"
)

type Graphic struct {
	window   *globaltype.SDL_Window
	renderer *globaltype.SDL_Renderer
	Cursor   *C.SDL_Surface

	path          string
	width, height int

	lock     sync.Mutex
	userLock sync.Mutex
	sayLock  sync.Mutex

	// screenBps -> targetScreentextures > texture
	renderBuffer [][]struct {
		texture   *globaltype.SDL_Texture
		transform object.Transform
	}

	Image_Manager  *image.Image
	Video_Manager  *video.Video
	Sprite_Manager *sprite.Sprite

	fonts map[string]struct {
		Font        *globaltype.TTF_Font
		Size        int
		LimitPixels int
	}
	textMemPool map[string][]*globaltype.SDL_Texture
	typingFXs   map[string][]struct {
		Data []struct {
			Texture   *globaltype.SDL_Texture
			Transform object.Transform
		}
		Duration  float64
		StartTime time.Time
		Bps       int
		Index     int
	}
	animations map[string]map[int][]struct {
		Anime *object.Anime
		Bps   int
	}
	sprites map[string][]struct {
		Name      string
		Bps       int
		Index     int
		Duration  float64
		Loop      bool
		StartTime time.Time
	}
}

func Init(window *globaltype.SDL_Window, r *globaltype.SDL_Renderer, p string, w, h int) *Graphic {
	return &Graphic{
		window:   window,
		renderer: r,
		renderBuffer: [][]struct {
			texture   *globaltype.SDL_Texture
			transform object.Transform
		}{},
		path:           p,
		width:          w,
		height:         h,
		Image_Manager:  image.Init(r),
		Video_Manager:  video.Init(),
		Sprite_Manager: sprite.Init(r),
		fonts: make(map[string]struct {
			Font        *globaltype.TTF_Font
			Size        int
			LimitPixels int
		}),
		textMemPool: make(map[string][]*globaltype.SDL_Texture),
		typingFXs: make(map[string][]struct {
			Data []struct {
				Texture   *globaltype.SDL_Texture
				Transform object.Transform
			}
			Duration  float64
			StartTime time.Time
			Bps       int
			Index     int
		}),
		animations: make(map[string]map[int][]struct {
			Anime *object.Anime
			Bps   int
		}),
		sprites: make(map[string][]struct {
			Name      string
			Bps       int
			Index     int
			Duration  float64
			Loop      bool
			StartTime time.Time
		}),
	}
}

func (g *Graphic) Close() {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(g.renderer))

	g.Image_Manager.Close()
	g.Video_Manager.Close()
	g.Sprite_Manager.Close()

	C.SDL_FreeSurface(g.Cursor)
}

func (g *Graphic) Update() {
	g.UpdateAnimation()
	g.UpdateTypingFX()
	g.UpdateSprite()
}

func (g *Graphic) RegisterCursor(path string) {
	cpath := C.CString(g.path + path)
	defer C.free(unsafe.Pointer(cpath))

	surface := C.IMG_Load(cpath)
	cursor := C.SDL_CreateColorCursor(surface, 0, 0)
	C.SDL_SetCursor(cursor)
	g.Cursor = surface
}

func (g *Graphic) RegisterImages(images *map[string]string) {
	for name, path := range *images {
		g.Image_Manager.RegisterImage(name, g.path+path)
	}
}

func (g *Graphic) RegisterVideos(videos *map[string]string) {
	for name, path := range *videos {
		g.Video_Manager.RegisterVideo(name, g.path+path, g.renderer)
	}
}

func (g *Graphic) RegisterSprites(sprites *map[string]string) {
	for name, path := range *sprites {
		g.Sprite_Manager.RegisterSprite(name, g.path+path)
	}
}

func (g *Graphic) RegisterFonts(
	fonts *map[string]struct {
		Path        string
		Size        int
		LimitPixels int
	},
) {
	for name, font := range *fonts {
		cpath := C.CString(g.path + font.Path)
		defer C.free(unsafe.Pointer(cpath))
		g.fonts[name] = struct {
			Font        *globaltype.TTF_Font
			Size        int
			LimitPixels int
		}{
			(*globaltype.TTF_Font)(C.TTF_OpenFont(cpath, C.int(font.Size))),
			font.Size,
			font.LimitPixels,
		}
	}
}

func (g *Graphic) SayLock() {
	g.sayLock.Lock()
}

func (g *Graphic) SayUnlock() {
	g.sayLock.Unlock()
}

func (g *Graphic) RenderLock() {
	g.userLock.Lock()
}

func (g *Graphic) RenderUnlock() {
	g.userLock.Unlock()
}
