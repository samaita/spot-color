package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lightgrey           = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey                = color.RGBA{0xff, 0xc5, 0xc6, 0xff}
	defaultBoxSizeX     = 100
	defaultBoxSizeY     = 100
	screenSizeX     int = 600
	screenSizeY     int = 600

	defaultDiffX, defaultDiffY = -1, -1
)

type Game struct {
	ScreenSizeX int
	ScreenSizeY int

	diffX int
	diffY int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

	maxBoxAmountX := g.ScreenSizeX / defaultBoxSizeX
	maxBoxAmountY := g.ScreenSizeY / defaultBoxSizeY

	if g.diffX == -1 {
		g.diffX = rand.Intn(maxBoxAmountX)
	}

	if g.diffY == -1 {
		g.diffY = rand.Intn(maxBoxAmountX)
	}

	for i := 0; i < maxBoxAmountX; i++ {
		for j := 0; j < maxBoxAmountY; j++ {
			rect := ebiten.NewImage(defaultBoxSizeX-4, defaultBoxSizeY-4)

			if i == g.diffX && j == g.diffY {
				rect.Fill(grey)
			} else {
				rect.Fill(lightgrey)
			}

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
		diffX:       defaultDiffX,
		diffY:       defaultDiffY,
	}); err != nil {
		log.Fatal(err)
	}
}
