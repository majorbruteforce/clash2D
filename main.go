package main

import (
	"clash2D/internals/atmosphere"
	"clash2D/internals/character"
	"clash2D/internals/core"
	"clash2D/internals/dev"
	"clash2D/pkg/config"
	"image/color"
	_ "image/png"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Lucy    *character.Character
	Stella  *character.Character
	BaseMap *atmosphere.Map
	Grid    *dev.Grid
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ToggleFullscreen()
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		core.Gb.UpdateOrigin(0, 1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		core.Gb.UpdateOrigin(0, -1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		core.Gb.UpdateOrigin(-1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		core.Gb.UpdateOrigin(1, 0)
	}

	// core.Gb.RunTickIndexCycle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.BaseMap.Render(screen)
	// g.Lucy.Render(screen)
	g.Stella.Render(screen)

	// g.Grid.Render(screen)
	dev.DisplayCursorPosition(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 480, 270
}

func ToggleFullscreen() {
	ebiten.SetFullscreen(!ebiten.IsFullscreen())
}

func main() {

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Tilemap Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &Game{
		Lucy:   character.LoadLucy(),
		Stella: character.LoadStella(),
		BaseMap: atmosphere.NewMap(
			path.Join(config.RootDir, "assets", "test.json"),
			path.Join(config.RootDir, "assets", "tileset-1.png"),
			352, 352,
			11, 11,
		),
		Grid: dev.NewGrid(
			0, 0,
			8, 8,
			0.5, color.RGBA{0, 0, 255, 100}),
	}

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}

}
