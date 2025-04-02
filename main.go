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
)

type Game struct {
	Lucy    *character.Character
	BaseMap *atmopshere.Map
	*core.Global
}

func (g *Game) Update() error {
	g.RunTickIndexCycle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.BaseMap.Render(200, 0, screen)

	g.Lucy.RunSequence(g.TickIndex(), character.LucySequences["WalkSW"])

	op := &ebiten.DrawImageOptions{}
	lucySubImage := g.Lucy.Sheet.SubImage(g.Lucy.Cutout.GetTileRectangleById(g.Lucy.Frame())).(*ebiten.Image)

	op.GeoM.Translate(160, 120)
	screen.DrawImage(lucySubImage, op)

	g.RunFrameIndexCycle()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tilemap Test")

	game := &Game{
		Lucy:   character.LoadLucy(),
		Global: core.NewGlobal(&core.DefaultGlobalConfig),
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
