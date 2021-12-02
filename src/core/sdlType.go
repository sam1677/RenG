package core

/*
#cgo CFLAGS: -I./sdl/include
#cgo LDFLAGS: -L./sdl/lib -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf -lSDL2_mixer

#include <SDL.h>
#include <SDL_image.h>
#include <SDL_ttf.h>
#include <SDL_mixer.h>
*/
import "C"

type SDL_Window C.SDL_Window
type SDL_Renderer C.SDL_Renderer
type SDL_Event C.SDL_Event

type SDL_Texture struct {
	Texture *C.SDL_Texture
	Xpos    int
	Ypos    int
	Width   int
	Height  int
	Alpha   uint8
	Degree  float64
}
type SDL_Surface C.SDL_Surface
type SDL_Rect C.SDL_Rect
type SDL_Color C.SDL_Color

type Mix_Music C.Mix_Music
type Mix_Chunk C.Mix_Chunk

type TTF_Font C.TTF_Font

const (
	SDL_INIT_TIMER = C.SDL_INIT_TIMER
	SDL_INIT_AUDIO = C.SDL_INIT_AUDIO
	SDL_INIT_VIDEO = C.SDL_INIT_VIDEO
)

const (
	SDL_WINDOW_SHOWN        = C.SDL_WINDOW_SHOWN
	SDL_WINDOW_RESIZABLE    = C.SDL_WINDOW_RESIZABLE
	SDL_WINDOWPOS_UNDEFINED = C.SDL_WINDOWPOS_UNDEFINED
)

const (
	SDL_HINT_RENDER_SCALE_QUALITY = C.SDL_HINT_RENDER_SCALE_QUALITY
)

const (
	SDL_RENDERER_ACCELERATED   = C.SDL_RENDERER_ACCELERATED
	SDL_RENDERER_PRESENTVSYNC  = C.SDL_RENDERER_PRESENTVSYNC
	SDL_RENDERER_TARGETTEXTURE = C.SDL_RENDERER_TARGETTEXTURE
)

const (
	SDL_TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC
	SDL_TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING
	SDL_TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET
)

const (
	SDL_PIXELFORMAT_YV12 = C.SDL_PIXELFORMAT_YV12
)

const (
	SDL_BLENDMODE_BLEND = C.SDL_BLENDMODE_BLEND
)

const (
	IMG_INIT_JPG  = C.IMG_INIT_JPG
	IMG_INIT_PNG  = C.IMG_INIT_PNG
	IMG_INIT_TIF  = C.IMG_INIT_TIF
	IMG_INIT_WEBP = C.IMG_INIT_WEBP
)

const (
	MIX_DEFAULT_FORMAT = C.MIX_DEFAULT_FORMAT
)

// event type
const (
	SDL_QUIT = C.SDL_QUIT

	SDL_WINDOWEVENT              = C.SDL_WINDOWEVENT
	SDL_WINDOWEVENT_SIZE_CHANGED = C.SDL_WINDOWEVENT_SIZE_CHANGED
	SDL_WINDOWEVENT_EXPOSED      = C.SDL_WINDOWEVENT_EXPOSED
	SDL_WINDOWEVENT_MINIMIZED    = C.SDL_WINDOWEVENT_MINIMIZED
	SDL_WINDOWEVENT_MAXIMIZED    = C.SDL_WINDOWEVENT_MAXIMIZED
	SDL_WINDOWEVENT_RESTORED     = C.SDL_WINDOWEVENT_RESTORED

	SDL_KEYDOWN = C.SDL_KEYDOWN
	SDL_KEYUP   = C.SDL_KEYUP

	SDL_MOUSEMOTION     = C.SDL_MOUSEMOTION
	SDL_MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	SDL_MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	SDL_MOUSEWHEEL      = C.SDL_MOUSEWHEEL
)

const (
	SDLK_0 = C.SDLK_0
	SDLK_1 = C.SDLK_1
	SDLK_2 = C.SDLK_2
	SDLK_3 = C.SDLK_3
	SDLK_4 = C.SDLK_4
	SDLK_5 = C.SDLK_5
	SDLK_6 = C.SDLK_6
	SDLK_7 = C.SDLK_7
	SDLK_8 = C.SDLK_8
	SDLK_9 = C.SDLK_9

	SDLK_a = C.SDLK_a
	SDLK_b = C.SDLK_b
	SDLK_c = C.SDLK_c
	SDLK_d = C.SDLK_d
	SDLK_e = C.SDLK_e
	SDLK_f = C.SDLK_f
	SDLK_g = C.SDLK_g
	SDLK_h = C.SDLK_h
	SDLK_i = C.SDLK_i
	SDLK_j = C.SDLK_j
	SDLK_k = C.SDLK_k
	SDLK_l = C.SDLK_l
	SDLK_m = C.SDLK_m
	SDLK_n = C.SDLK_n
	SDLK_o = C.SDLK_o
	SDLK_p = C.SDLK_p
	SDLK_q = C.SDLK_q
	SDLK_r = C.SDLK_r
	SDLK_s = C.SDLK_s
	SDLK_t = C.SDLK_t
	SDLK_u = C.SDLK_u
	SDLK_v = C.SDLK_v
	SDLK_w = C.SDLK_w
	SDLK_x = C.SDLK_x
	SDLK_y = C.SDLK_y
	SDLK_z = C.SDLK_z

	SDLK_DOWN = C.SDL_KEYDOWN

	SDLK_COMMA  = C.SDLK_COMMA
	SDLK_EQUALS = C.SDLK_EQUALS

	SDLK_BACKSPACE = C.SDLK_BACKSPACE
	SDLK_CAPSLOCK  = C.SDLK_CAPSLOCK

	SDLK_HOME   = C.SDLK_HOME
	SDLK_END    = C.SDLK_END
	SDLK_DELETE = C.SDLK_DELETE

	SDLK_ESCAPE = C.SDLK_ESCAPE

	SDLK_F1  = C.SDLK_F1
	SDLK_F2  = C.SDLK_F2
	SDLK_F3  = C.SDLK_F3
	SDLK_F4  = C.SDLK_F4
	SDLK_F5  = C.SDLK_F5
	SDLK_F6  = C.SDLK_F6
	SDLK_F7  = C.SDLK_F7
	SDLK_F8  = C.SDLK_F8
	SDLK_F9  = C.SDLK_F9
	SDLK_F10 = C.SDLK_F10
	SDLK_F11 = C.SDLK_F11
	SDLK_F12 = C.SDLK_F12
	SDLK_F13 = C.SDLK_F13
	SDLK_F14 = C.SDLK_F14
	SDLK_F15 = C.SDLK_F15
	SDLK_F16 = C.SDLK_F16
	SDLK_F17 = C.SDLK_F17
	SDLK_F18 = C.SDLK_F18
	SDLK_F19 = C.SDLK_F19
	SDLK_F20 = C.SDLK_F20
	SDLK_F21 = C.SDLK_F21
	SDLK_F22 = C.SDLK_F22
	SDLK_F23 = C.SDLK_F23
	SDLK_F24 = C.SDLK_F24
)

const (
	SDL_BUTTON_LEFT   = C.SDL_BUTTON_LEFT
	SDL_BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE
	SDL_BUTTON_RIGHT  = C.SDL_BUTTON_RIGHT
)

const (
	SDL_TRUE  = C.SDL_TRUE
	SDL_FALSE = C.SDL_FALSE
)

const (
	SDL_FLIP_NONE       = C.SDL_FLIP_NONE
	SDL_FLIP_HORIZONTAL = C.SDL_FLIP_HORIZONTAL
	SDL_FLIP_VERTICAL   = C.SDL_FLIP_VERTICAL
)
