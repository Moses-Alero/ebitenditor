package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets
var asset embed.FS

var PlayerSprite = LoadAsset("assets/Player/Player_Ship_Ant_00.png")
var PlayerBulletSprite = LoadAsset("assets/Player/Player_Shot_27.png")
var EnemySprite = LoadAsset("assets/Enemy/Enemy_Ship_Scout_21.png")

func LoadAsset(path string) *ebiten.Image {
	//read the file
	f, err := asset.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadFont(name string) font.Face {
	f, err := asset.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}
