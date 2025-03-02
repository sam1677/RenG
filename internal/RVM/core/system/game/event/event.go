package event

/*
#cgo CFLAGS: -I./c
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf -lSDL2_mixer

#include <SDL2/SDL.h>
*/
import "C"
import (
	"sync"
)

type Event struct {
	e C.SDL_Event

	lock sync.Mutex

	TopScreenName string

	Key        map[string][]KeyEvent
	Button     map[string][]ButtonEvent
	Bar        map[string][]BarEvent
	MouseClick map[string][]MouseClickEvent
}

func Init() *Event {
	return &Event{
		Key:        make(map[string][]KeyEvent),
		Button:     make(map[string][]ButtonEvent),
		Bar:        make(map[string][]BarEvent),
		MouseClick: make(map[string][]MouseClickEvent),
	}
}

func (e *Event) Close() {}

func (e *Event) Update() bool {
	for e.pollEvent() != 0 {
		switch e.getEventType() {
		case RENG_QUIT:
			return true
		case RENG_KEYDOWN:
			e.keyDown()
		case RENG_KEYUP:
			e.keyUp()
		case RENG_MOUSEBUTTONDOWN:
			e.buttonDown()     // 버튼 이벤트 먼저 처리.
			e.mouseClickDown() // 그후 마우스 클릭 다운을 처리해야 버튼의 press 상태를 알 수 있음.
			e.barDown()
		case RENG_MOUSEBUTTONUP:
			e.mouseClickUp() // 마우스 이벤트 먼저 처리해야 버튼이 UP하기 전에, press 상태를 알 수 있음.
			e.buttonUp()     // 그후 버튼 이벤트 처리리
			e.barUp()
		case RENG_MOUSEMOTION:
			e.buttonHover()
			e.barScroll()
		case RENG_MOUSEWHEEL:
		}
	}
	return false
}

func (e *Event) DeleteAllScreenEvent(screenName string) {
	e.lock.Lock()
	defer e.lock.Unlock()

	delete(e.Key, screenName)
	delete(e.Button, screenName)
	delete(e.Bar, screenName)
	delete(e.MouseClick, screenName)
}

func (e *Event) DeleteKeyScreenIndexEvent(screenName string, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.Key[screenName] = append(e.Key[screenName][:index], e.Key[screenName][index+1:]...)
}

func (e *Event) DeleteButtonScreenIndexEvent(screenName string, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.Button[screenName] = append(e.Button[screenName][:index], e.Button[screenName][index+1:]...)
}

func (e *Event) DeleteBarScreenIndexEvent(screenName string, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.Bar[screenName] = append(e.Bar[screenName][:index], e.Bar[screenName][index+1:]...)
}

func (e *Event) DeleteMouseClickScreenIndexEvent(screenName string, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.MouseClick[screenName] = append(e.MouseClick[screenName][:index], e.MouseClick[screenName][index+1:]...)
}
