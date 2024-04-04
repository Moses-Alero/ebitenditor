package component

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bounds struct {
	Position
	W float64
	H float64
}

type Icon struct {
	Img *ebiten.Image
}

type BackGround struct {
	Color color.RGBA
}

type Targetable struct {
	IsTargetable bool
}

type Clickable struct {
	IsClickable bool
}

type Draggable struct {
	IsDraggable bool
}

type Shape struct {
	Transform
	Bounds
	Visibility bool
}

type Position struct {
	X, Y float64
}

type Scale struct {
	X, Y float64
}

type Rotation struct {
	X, Y int64
}

type Transform struct {
	Position
	Scale
	Rotation
}

type ActionHandler struct {
	Do        func(params ...interface{})
}

type Input struct {
	LeftMouseButtonPressed   bool
	MiddleMouseButtonPressed bool
	RightMouseButtonPressed  bool
	CursorPos                Position  
	WheelX                   float64
	WheelY                   float64

	LeftMouseButtonJustPressed   bool
	MiddleMouseButtonJustPressed bool
	RightMouseButtonJustPressed  bool

	LastLeftMouseButtonPressed   bool
	LastMiddleMouseButtonPressed bool
	LastRightMouseButtonPressed  bool

	InputChars     []rune
	KeyPressed     map[ebiten.Key]bool
	AnyKeyPressed  bool
	isTouched      bool
	cursorImages   map[string]*ebiten.Image
	cursorOffset   map[string]image.Point

	//touchscreenPlatform bool
}
