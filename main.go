package main

import (
	"clash2D/internals/atmopshere"
	"clash2D/internals/character"
	"clash2D/internals/core"
	"clash2D/pkg/config"
	_ "image/png"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Lucy    *character.Character
	Stella  *character.Character
	BaseMap *atmopshere.Map
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ToggleFullscreen()
	}

	core.Gb.RunTickIndexCycle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.BaseMap.Render(screen)
	// g.Lucy.Render(screen)
	g.Stella.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 960, 540
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
		BaseMap: atmopshere.NewMap(
			path.Join(config.RootDir, "assets", "test.json"),
			path.Join(config.RootDir, "assets", "tileset.png"),
			352, 352,
			11, 11,
		),
	}

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}

}
