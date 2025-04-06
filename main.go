package main

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/majorbruteforce/clash2D/pkg/errorutils"
	"github.com/majorbruteforce/clash2D/pkg/geom"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	r  *geom.Rhombus
	r1 *geom.Rhombus
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.r.A.Y += -0.5
		g.r.B.Y += -0.5
		g.r.C.Y += -0.5
		g.r.D.Y += -0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.r.A.Y += 0.5
		g.r.B.Y += 0.5
		g.r.C.Y += 0.5
		g.r.D.Y += 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.r.A.X += 0.5
		g.r.B.X += 0.5
		g.r.C.X += 0.5
		g.r.D.X += 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.r.A.X += -0.5
		g.r.B.X += -0.5
		g.r.C.X += -0.5
		g.r.D.X += -0.5
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.r.Stroke(screen, 1, color.RGBA{255, 0, 0, 100}, false)
	g.r1.Stroke(screen, 1, color.RGBA{255, 0, 0, 100}, false)

	if g.r.Overlaps(g.r1) {
		text.Draw(screen, "Overlaps", basicfont.Face7x13, 50, 50, color.White)
	} else {
		text.Draw(screen, ":(", basicfont.Face7x13, 50, 50, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 480, 270
}

func main() {

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Tilemap Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	r := geom.NewRhombusFromTile(100, 100, 32, 32, 0, 0)
	r1 := geom.NewRhombusFromTile(101, 105, 32, 32, 0, 0)

	game := &Game{
		r:  r,
		r1: r1,
	}

	err := ebiten.RunGame(game)
	errorutils.CheckFatal(err)

}
