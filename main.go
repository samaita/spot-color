package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	lightgrey           = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey                = color.RGBA{0xff, 0xc5, 0xc6, 0xff}
	defaultBoxSizeX     = 100
	defaultBoxSizeY     = 100
	screenSizeX     int = 600
	screenSizeY     int = 600

	defaultBoxDiffX, defaultBoxDiffY = -1, -1
)

type Game struct {
	ScreenSizeX int
	ScreenSizeY int

	indexBoxDiffX int
	indexBoxDiffY int

	posBoxDiffX int
	posBoxDiffY int

	sizeBoxX int
	sizeBoxY int

	clicked    int
	maxReduced int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

	maxBoxAmountX := g.ScreenSizeX / g.sizeBoxX
	maxBoxAmountY := g.ScreenSizeY / g.sizeBoxY

	g.setNewColorDiffIndex(maxBoxAmountX, maxBoxAmountY)

	for i := 0; i < maxBoxAmountX; i++ {
		for j := 0; j < maxBoxAmountY; j++ {
			rect := ebiten.NewImage(g.sizeBoxX-4, g.sizeBoxY-4)
			rect.Fill(g.getNewBoxColorDiff(i, j))

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(g.setBoxSize(i)), float64(g.setBoxSize(j)))
			screen.DrawImage(rect, op)

			g.setNewColorDiffPosition(i, j)
		}
	}

	if g.isClickOnBox() {
		g.clicked++
		g.reduceBoxSize()
		g.resetColorDiffPosition()
	}
}

func (g *Game) isClickOnBox() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x <= g.posBoxDiffX+g.sizeBoxX && x > g.posBoxDiffX && y <= g.posBoxDiffY+g.sizeBoxY && y > g.posBoxDiffY {
			return true
		}
	}
	return false
}

// setNewColorDiffIndex setup new box index with different color
func (g *Game) reduceBoxSize() {
	if g.clicked <= g.maxReduced {
		g.sizeBoxX = g.ScreenSizeX / g.clicked
		g.sizeBoxY = g.ScreenSizeY / g.clicked
	}
}

// setNewColorDiffIndex setup new box index with different color
func (g *Game) setNewColorDiffIndex(maxX, maxY int) {
	if g.indexBoxDiffX == -1 {
		g.indexBoxDiffX = rand.Intn(maxX)
	}
	if g.indexBoxDiffY == -1 {
		g.indexBoxDiffY = rand.Intn(maxY)
	}
}

// setNewColorDiffPosition setup new box position with different color
func (g *Game) setNewColorDiffPosition(indexX, indexY int) {
	if indexX == g.indexBoxDiffX && indexY == g.indexBoxDiffY {
		g.posBoxDiffX = (g.sizeBoxX * indexX) + 2
		g.posBoxDiffY = (g.sizeBoxY * indexY) + 2
	}
}

// getNewBoxColorDiff define a color difference betwen boxes, changes every turn, diff 40% reduce 5% each turn, stop reduced at 10%
func (g *Game) getNewBoxColorDiff(indexX, indexY int) color.RGBA {
	if indexX == g.indexBoxDiffX && indexY == g.indexBoxDiffY {
		return lightgrey
	} else {
		return grey
	}
}

// setBoxSize set box size, getting smaller each turn, stop reduced at 10x10
func (g *Game) setBoxSize(index int) int {
	return (g.sizeBoxX * index) + 2
}

// resetColorDiffPosition reset position of the diff box
func (g *Game) resetColorDiffPosition() {
	g.indexBoxDiffX, g.indexBoxDiffY = defaultBoxDiffX, defaultBoxDiffY
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenSizeX, g.ScreenSizeY
}

func main() {

	ebiten.SetWindowSize(screenSizeX, screenSizeY)
	ebiten.SetWindowTitle("Spot Different Color!")
	if err := ebiten.RunGame(&Game{
		ScreenSizeX:   screenSizeX,
		ScreenSizeY:   screenSizeY,
		indexBoxDiffX: defaultBoxDiffX,
		indexBoxDiffY: defaultBoxDiffY,
		maxReduced:    10,
		clicked:       2,
		sizeBoxX:      screenSizeX / 2,
		sizeBoxY:      screenSizeY / 2,
	}); err != nil {
		log.Fatal(err)
	}
}
