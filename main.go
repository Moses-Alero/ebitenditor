package main

import (
	"ebitenditor/event"
	"ebitenditor/ui/component"
	"ebitenditor/ui/widget"
	"fmt"

	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

var (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

func (g *Game) Setup(w engine.World) {

	widget.NewButton(w, &widget.ButtonBuilder{
		Transform: component.Transform{
			Position: component.Position{X: 350, Y: 350}, 
			Bounds: component.Bounds{W: 50, H: 20},
		},
		BackGround: component.BackGround{
			Color: color.RGBA{0, 0, 255, 0},
		},
		Text: component.Text{
			Text: "Submit",
		},
		ActionHandler: component.ActionHandler{
			Do: func(){
				fmt.Println("I am doing something else")
			},
			Event: event.OnClick,
			Tag: "Button 1",
		},

	})

	widget.NewButton(w, &widget.ButtonBuilder{
		Transform: component.Transform{
			Position: component.Position{X: 450, Y: 450}, 
			Bounds: component.Bounds{W: 50, H: 20},
		},
		BackGround: component.BackGround{
			Color: color.RGBA{255, 0, 0, 0},
		},
		Text: component.Text{
			Text: "Submit",
		},
		ActionHandler: component.ActionHandler{
			Do: func (){
				fmt.Println("This is probably working")
			},
			Event: event.OnClick,
			Tag: "Button 2",

		},
	})

	widget.NewButton(w, &widget.ButtonBuilder{
		Transform: component.Transform{
			Position: component.Position{X: 250, Y: 250}, 
			Bounds: component.Bounds{W: 50, H: 20},
		},
		BackGround: component.BackGround{
			Color: color.RGBA{255, 0, 255, 0},
		},
		Text: component.Text{
			Text: "Submit",
		},
		ActionHandler: component.ActionHandler{
			Do: func(){
				fmt.Println("I am doing something")
			},
			Event: event.OnClick,
			Tag: "Button 3",
		},
	})

	widget.NewButton(w, (&widget.ButtonBuilder{}).NewButton())
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowTitle("Ebitenditor")
	g := engine.NewGame(&Game{})

	fmt.Println("This is ebitenditor")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
