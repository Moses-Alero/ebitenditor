package component

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bounds struct {
	W float64
	H float64
}

func (bounds *Bounds) GetBounds() (float64, float64){
	return bounds.W, bounds.H
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

func (pos *Position) GetPos() (float64, float64){
	return pos.X, pos.Y
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
	Bounds
}

type ActionHandler struct {
	Do        func()
	Event     func(arg interface{}) bool
	Tag       string
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

type Text struct {
	Text  string
}
