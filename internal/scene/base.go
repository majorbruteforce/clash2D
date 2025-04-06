package scene

import (
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/internal/character"
	"github.com/majorbruteforce/clash2D/pkg/atmosphere"
	"github.com/majorbruteforce/clash2D/pkg/config"
)

type Base struct {
	title string
	Map   *atmosphere.Map
	Lucy  *character.Character
}

var BaseScene = Base{
	title: "Welcome to your base!",
	Map: atmosphere.NewMap(
		path.Join(config.RootDir, "assets", "maps", "tileset-1.png"),
		path.Join(config.RootDir, "assets", "maps", "tileset-1.json"),
		11, 11,
	),
	Lucy: character.NewCharacter(
		"Lucy",
		path.Join(config.RootDir, "assets", "characters", "lucy_walk.png"),
		9, 9,
	),
}

func (b *Base) Draw(screen *ebiten.Image) {
	b.Map.Render(screen)
	b.Lucy.Render(screen)
}

func (b *Base) Update() error {
	b.Lucy.RunSequence(character.LucySequences["WalkS"])
	return nil
}
