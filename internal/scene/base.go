package scene

import (
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/pkg/config"
	"github.com/majorbruteforce/clash2D/pkg/sprite"
)

type Base struct {
	title   string
	BaseMap *sprite.Cutout
}

var BaseScene = Base{
	title: "Welcome to your base!",
	BaseMap: sprite.NewCutout(
		path.Join(config.RootDir, "assets", "characters", "lucy_walk.png"),
		9, 9,
	),
}

func (b *Base) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(100, 100)
	screen.DrawImage(b.BaseMap.GetSubImageByIndex(0), op)
}

func (b *Base) Update() error {
	return nil
}
