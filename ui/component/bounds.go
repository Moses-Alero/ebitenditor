package component

import (
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

func NewIcon(img *ebiten.Image) Icon {
	return Icon{
		Img: img,
	}
}
