package entity

import (
	"ebitenditor/ui/component"
	"image/color"
)

type Button struct {
	component.Bounds
	component.Icon
	component.BackGround
	component.Clickable
}

func NewButton() *Button {
	return &Button{
		Bounds: component.Bounds{Position: component.Position{X: 120, Y: 120}, W: 15, H:15},	
		Icon: component.Icon{Img: nil},
		BackGround: component.BackGround{Color: color.RGBA{255,255,255,0}},
		Clickable: component.Clickable{IsClickable: true},
	}
}
