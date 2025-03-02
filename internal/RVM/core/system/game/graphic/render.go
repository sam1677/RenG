package graphic

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>

SDL_Rect* CreateRect(int x, int y, int w, int h)
{
	SDL_Rect* Quad = (SDL_Rect*)malloc(sizeof(SDL_Rect));
	Quad->x = x;
	Quad->y = y;
	Quad->w = w;
	Quad->h = h;
	return Quad;
}

void FreeRect(SDL_Rect* r)
{
	free(r);
}
*/
import "C"
import (
	"internal/RVM/core/globaltype"
	"internal/RVM/core/object"
	"strconv"
	"strings"
)

func (g *Graphic) Render() {
	g.userLock.Lock()
	g.sayLock.Lock()
	g.lock.Lock()
	g.Video_Manager.Lock()

	C.SDL_RenderClear((*C.SDL_Renderer)(g.renderer))
	C.SDL_SetRenderDrawColor(
		(*C.SDL_Renderer)(g.renderer),
		C.uchar(0x13), C.uchar(0x13), C.uchar(0x12), C.uchar(0xFF),
	)

	x, y := g.GetCurrentWindowSize()

	for i := 0; i < len(g.renderBuffer); i++ {
		for j := 0; j < len(g.renderBuffer[i]); j++ {
			r1 := C.CreateRect(
				C.int(float32(g.renderBuffer[i][j].transform.Pos.X)*float32(x)/float32(g.width)),
				C.int(float32(g.renderBuffer[i][j].transform.Pos.Y)*float32(y)/float32(g.height)),
				C.int(float32(g.renderBuffer[i][j].transform.Size.X)*float32(x)/float32(g.width)),
				C.int(float32(g.renderBuffer[i][j].transform.Size.Y)*float32(y)/float32(g.height)),
			)
			if g.renderBuffer[i][j].transform.Flip.X != 0 || g.renderBuffer[i][j].transform.Flip.Y != 0 {
				r2 := C.CreateRect(
					C.int(0),
					C.int(0),
					C.int(g.renderBuffer[i][j].transform.Flip.X),
					C.int(g.renderBuffer[i][j].transform.Flip.Y),
				)
				C.SDL_RenderCopyEx(
					(*C.SDL_Renderer)(g.renderer),
					(*C.SDL_Texture)(g.renderBuffer[i][j].texture),
					r2,
					r1,
					C.double(g.renderBuffer[i][j].transform.Rotate),
					nil,
					C.SDL_FLIP_NONE,
				)
				C.FreeRect(r2)
			} else {
				C.SDL_RenderCopyEx(
					(*C.SDL_Renderer)(g.renderer),
					(*C.SDL_Texture)(g.renderBuffer[i][j].texture),
					nil,
					r1,
					C.double(g.renderBuffer[i][j].transform.Rotate),
					nil,
					C.SDL_FLIP_NONE,
				)
			}
			C.FreeRect(r1)
		}
	}

	C.SDL_RenderPresent((*C.SDL_Renderer)(g.renderer))

	g.Video_Manager.Unlock()
	g.lock.Unlock()
	g.sayLock.Unlock()
	g.userLock.Unlock()
}

func (g *Graphic) AddScreenRenderBuffer() {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.renderBuffer = append(g.renderBuffer, []struct {
		texture   *globaltype.SDL_Texture
		transform object.Transform
	}{})
}

func (g *Graphic) DeleteScreenRenderBuffer(bps int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.renderBuffer = append(g.renderBuffer[:bps], g.renderBuffer[bps+1:]...)
}

func (g *Graphic) AddScreenTextureRenderBuffer(
	bps int,
	texture *globaltype.SDL_Texture,
	transform object.Transform,
) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.renderBuffer[bps] = append(
		g.renderBuffer[bps],
		struct {
			texture   *globaltype.SDL_Texture
			transform object.Transform
		}{
			texture,
			transform,
		})
}

func (g *Graphic) DeleteScreenTextureRenderBuffer(bps, index int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.renderBuffer[bps] = append(g.renderBuffer[bps][:index], g.renderBuffer[bps][index+1:]...)
}

func (g *Graphic) GetCurrentRenderBufferTextureNameANDTransformByBPS(bps int) []string {
	g.lock.Lock()
	defer g.lock.Unlock()

	var ret []string

	for _, t := range g.renderBuffer[bps] {
		name := "I#" + g.Image_Manager.GetImgaeTextureName(t.texture)
		if name == "I#" {
			name = "V#"
			v, l := g.Video_Manager.GetVideoNameANDLoopByTexture(t.texture)
			name += v + "?" + strconv.Itoa(l)
		}

		format := strings.Join(
			[]string{
				name,
				formatFloat(t.transform.Pos.X),
				formatFloat(t.transform.Pos.Y),
				formatFloat(t.transform.Size.X),
				formatFloat(t.transform.Size.Y),
				formatFloat(t.transform.Rotate),
			},
			"?",
		)

		if name != "I#" && name != "V#" {
			ret = append(ret, format)
		}
	}

	return ret
}

func (g *Graphic) GetCurrentTopRenderBps() int {
	g.lock.Lock()
	defer g.lock.Unlock()

	return len(g.renderBuffer) - 1
}

func (g *Graphic) GetCurrentTopScreenIndexByBps(bps int) int {
	g.lock.Lock()
	defer g.lock.Unlock()

	return len(g.renderBuffer[bps]) - 1
}

func (g *Graphic) GetCurrentWindowSize() (x, y int) {
	var xsize, ysize C.int
	C.SDL_GetWindowSize((*C.SDL_Window)(g.window), &xsize, &ysize)
	return int(xsize), int(ysize)
}

func formatFloat(f float32) string {
	return strconv.FormatFloat(float64(f), 'G', 3, 32)
}
