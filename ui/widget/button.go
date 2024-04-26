package widget

import (
	"bytes"
	"ebitenditor/ui/component"
	"ebitenditor/ui/entity"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sedyh/mizu/pkg/engine"

	_ "embed"
)

var (
	//go:embed mplus-1p-regular.ttf
	MPlus1pRegular_ttf []byte
)

type Button struct{
}

type ButtonBuilder entity.Button 

func (btnBuilder *ButtonBuilder ) NewButton() *ButtonBuilder {
	return &ButtonBuilder{
		Transform: component.Transform{
			Position: component.Position{X: 120, Y: 120},
			Bounds: component.Bounds{W: 50, H: 35},
		},	
		Icon: component.Icon{Img: nil},
		BackGround: component.BackGround{Color: color.RGBA{8,8,8,0}},
		Clickable: component.Clickable{IsClickable: true},
		Text: component.Text{ Text: "Button"},
	}
}
 

func NewButton(w engine.World, builder *ButtonBuilder) *Button {
	btn := new(Button)
	if builder == nil{
		builder = builder.NewButton()
	}
	btn.Setup(w, builder)
	return btn
}

func (btn *Button) Update(w engine.World) {
	btn.candidates = btn.candidates[:0]
	button := w.View(
		component.Transform{},
		component.Icon{},
		component.BackGround{},
		component.Clickable{},
		component.ActionHandler{},
	)

	button.Each(func (entity engine.Entity){
		var transform *component.Transform
		var action *component.ActionHandler

		entity.Get(&transform, &action)

		if action.Event != nil{
			if action.Event(transform){
				if action.Do != nil{
					action.Do()
				}
			}
		}
	})
}

func (btn *Button) Draw(w engine.World, screen *ebiten.Image) {

	s, _ := text.NewGoTextFaceSource(bytes.NewReader(MPlus1pRegular_ttf))

	button := w.View(
		component.Transform{},
		component.Icon{},
		component.BackGround{},
		component.Clickable{},
		component.Text{},
	)

	button.Each(func(entity engine.Entity){

		var transform *component.Transform
		var icon *component.Icon
		var bg *component.BackGround
		var clickable *component.Clickable
		var btnText *component.Text

		entity.Get(&transform, &icon, &bg, &clickable, &btnText)

		opts := &ebiten.DrawImageOptions{}
		txtOpts := &text.DrawOptions{}
		txtOpts.GeoM.Translate(transform.Position.X + (transform.Bounds.W/2) - (transform.Bounds.W/3), transform.Position.Y + transform.Bounds.H/2 - transform.Bounds.H/5)
		face := &text.GoTextFace{
			Source: s,
			Size: transform.W/5,
		}
		opts.GeoM.Translate(transform.Position.X, transform.Position.Y)
		if icon.Img != nil {
			screen.DrawImage(icon.Img, opts)
		} else if icon.Img == nil {
			text.Draw(screen, btnText.Text, face, txtOpts)
			vector.DrawFilledRect(screen, float32(transform.Position.X), float32(transform.Position.Y), float32(transform.W), float32(transform.H), bg.Color, false)
		}
	})	
}

func (btn *Button) Setup(w engine.World, button *ButtonBuilder) {
	w.AddComponents(
		component.Position{}, component.Icon{},
		component.BackGround{}, component.Transform{}, 
		component.Clickable{}, component.ActionHandler{},
		component.Text{},
	)

	w.AddSystems(
		&Button{},
	)

	w.AddEntities(button)
}
