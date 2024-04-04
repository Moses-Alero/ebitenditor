package input

import (
	"ebitenditor/ui/component"
	"image"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputHandler component.Input

var DefaultInputHandler *InputHandler = &InputHandler{

	KeyPressed:  make(map[ebiten.Key]bool),
	cursorImages:   make(map[string]*ebiten.Image),
	cursorOffset:   make(map[string]image.Point),
}


func (handler *InputHandler) AfterDraw(screen *ebiten.Image) {
	handler.InputChars = handler.InputChars[:0]
	handler.WheelX, handler.WheelY = 0, 0
}

func (handler *InputHandler) MouseButtonPressed(b ebiten.MouseButton) bool {
	switch b {
	case ebiten.MouseButtonLeft:
		return handler.LeftMouseButtonPressed
	case ebiten.MouseButtonMiddle:
		return handler.MiddleMouseButtonPressed
	case ebiten.MouseButtonRight:
		return handler.RightMouseButtonPressed
	default:
		return false
	}
}
func (handler *InputHandler) MouseButtonJustPressed(b ebiten.MouseButton) bool {
	switch b {
	case ebiten.MouseButtonLeft:
		return handler.LeftMouseButtonJustPressed
	case ebiten.MouseButtonMiddle:
		return handler.MiddleMouseButtonJustPressed
	case ebiten.MouseButtonRight:
		return handler.RightMouseButtonJustPressed
	default:
		return false
	}
}

func (handler *InputHandler) CursorPosition() (int, int) {
	return handler.CursorX, handler.CursorY
}

func (handler *InputHandler) GetCursorImage(name string) *ebiten.Image {
	return handler.cursorImages[name]
}

func (handler *InputHandler) GetCursorOffset(name string) image.Point {
	return handler.cursorOffset[name]
}
func (handler *InputHandler) SetCursorImage(name string, cursorImage *ebiten.Image, offset image.Point) {
	handler.cursorImages[name] = cursorImage
	handler.cursorOffset[name] = offset
}

