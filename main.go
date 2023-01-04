package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lightgrey           = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	defaultBoxSizeX     = 100
	defaultBoxSizeY     = 100
	screenSizeX     int = 600
	screenSizeY     int = 600
)

type Game struct {
	ScreenSizeX int
	ScreenSizeY int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

	maxBoxAmountX := g.ScreenSizeX / defaultBoxSizeX
	maxBoxAmountY := g.ScreenSizeY / defaultBoxSizeY

	for i := 0; i < maxBoxAmountX; i++ {
		for j := 0; j < maxBoxAmountY; j++ {
			rect := ebiten.NewImage(defaultBoxSizeX-4, defaultBoxSizeY-4)
			rect.Fill(lightgrey)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((defaultBoxSizeX*i)+2), float64((defaultBoxSizeY*j)+2))
			screen.DrawImage(rect, op)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenSizeX, g.ScreenSizeY
}

func main() {

	ebiten.SetWindowSize(screenSizeX, screenSizeY)
	ebiten.SetWindowTitle("Spot Different Color!")
	if err := ebiten.RunGame(&Game{
		ScreenSizeX: screenSizeX,
		ScreenSizeY: screenSizeY,
	}); err != nil {
		log.Fatal(err)
	}
}
