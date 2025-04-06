package scene

import (
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/pkg/atmosphere"
	"github.com/majorbruteforce/clash2D/pkg/config"
	"github.com/majorbruteforce/clash2D/pkg/sprite"
)

type Base struct {
	title string
	Map   *atmosphere.Map
	Lucy  *sprite.Cutout
}

var BaseScene = Base{
	title: "Welcome to your base!",
	Map: atmosphere.NewMap(
		path.Join(config.RootDir, "assets", "maps", "tileset-1.png"),
		path.Join(config.RootDir, "assets", "maps", "tileset-1.json"),
		11, 11,
	),
	Lucy: sprite.NewCutout(
		path.Join(config.RootDir, "assets", "characters", "lucy_walk.png"),
		9, 9,
	),
}

func (b *Base) Draw(screen *ebiten.Image) {
	b.Map.Render(screen)
}

func (b *Base) Update() error {
	return nil
}
