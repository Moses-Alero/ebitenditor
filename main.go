package main

import (
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
		Bounds: component.Bounds{
			Position: component.Position{X: 250, Y: 250}, 
			W: 10, H: 25,
		},
		BackGround: component.BackGround{
			Color: color.RGBA{255, 0, 255, 1},
		},
		ActionHandler: component.ActionHandler{
			Do: func(d ...interface{}){
				if ebiten.IsKeyPressed(ebiten.KeyDown){
					fmt.Println(d)		
				}
			},
		},	
	})

	widget.NewButton(w, &widget.ButtonBuilder{
		Bounds: component.Bounds{
			Position: component.Position{X: 250, Y: 250}, 
			W: 50, H: 50,
		},
		BackGround: component.BackGround{
			Color: color.RGBA{0, 0, 255, 1},
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
