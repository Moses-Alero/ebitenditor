package widget

import (
	"ebitenditor/ui/component"
	"ebitenditor/ui/entity"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/sedyh/mizu/pkg/engine"
)

type Button struct{}

type ButtonBuilder entity.Button 

func (btnBuilder *ButtonBuilder ) NewButton() *ButtonBuilder {
	return &ButtonBuilder{
		Bounds: component.Bounds{Position: component.Position{X: 120, Y: 120}, W: 15, H:15},	
		Icon: component.Icon{Img: nil},
		BackGround: component.BackGround{Color: color.RGBA{255,255,255,0}},
		Clickable: component.Clickable{IsClickable: true},
	}
}

func NewButton(w engine.World, builder *ButtonBuilder) {
	btn := new(Button)
	if builder == nil{
		builder = builder.NewButton()
	}
	btn.Setup(w, builder)
}

func (btn *Button) Update(w engine.World) {
	button := w.View(
		component.Bounds{},
		component.Icon{},
		component.BackGround{},
		component.Clickable{},
		component.ActionHandler{},
	)

	button.Each(func (entity engine.Entity){
		var bounds *component.Bounds
		var icon *component.Icon
		var bg *component.BackGround
		var clickable *component.Clickable
		var action *component.ActionHandler

		entity.Get(&bounds, &icon, &bg, &clickable, &action)
		
			
		if action.Do != nil{
			action.Do()
		}
	})
}

func (btn *Button) Draw(w engine.World, screen *ebiten.Image) {
	button := w.View(
		component.Bounds{},
		component.Icon{},
		component.BackGround{},
		component.Clickable{},
	)

	button.Each(func(entity engine.Entity){
		var bounds *component.Bounds
		var icon *component.Icon
		var bg *component.BackGround
		var clickable *component.Clickable

		entity.Get(&bounds, &icon, &bg, &clickable)

		if bounds.Position.X == 0 && bounds.Position.Y == 0 {
			bounds.Position = component.Position{350, 350}
			bounds.W, bounds.H = 20,20
		}

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(bounds.Position.X, bounds.Position.Y)
		if icon.Img != nil {
			screen.DrawImage(icon.Img, opts)
		} else if icon.Img == nil {
			vector.DrawFilledRect(screen, float32(bounds.Position.X), float32(bounds.Position.Y), float32(bounds.W), float32(bounds.H), bg.Color, false)
		}
	})	
}

func (btn *Button) Setup(w engine.World, button *ButtonBuilder) {
	w.AddComponents(
		component.Position{}, component.Icon{},
		component.BackGround{}, component.Bounds{}, 
		component.Clickable{}, component.ActionHandler{},
	)

	w.AddSystems(
		&Button{},
	)

	w.AddEntities(button)
}

