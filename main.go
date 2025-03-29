package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"clash2D/internals/core"
	"clash2D/internals/entities"
)

type Game struct {
	Lucy   *core.Entity
	global *core.Global
}

func (g *Game) Update() error {

	g.Lucy.UpdateFrame(9, 16, 1, 6, g.global.TickIndex)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Lucy.Position.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Lucy.Position.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Lucy.Position.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Lucy.Position.X += 2
	}

	g.global.RunTickIndexCycle()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	lucyFrame := g.Lucy.Sheet.SubImage(g.Lucy.Cutout.GetTileRectById(g.Lucy.Frame)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Lucy.Position.X), float64(g.Lucy.Position.Y))
	screen.DrawImage(lucyFrame, op)
	g.global.RunFrameIndexCycle()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tilemap Test")
	game := &Game{
		Lucy:   entities.LoadLucy(),
		global: core.NewGlobal(&core.DefaultGlobalConfig),
	}

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}

}
