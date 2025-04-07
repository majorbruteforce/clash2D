package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/pkg/debugutils"
)

type Debug struct {
	*debugutils.Grid
}

var DebugScene = Debug{
	Grid: debugutils.NewGrid(
		0, 0,
		32, 32,
		1, color.RGBA{0, 255, 255, 100},
	),
}

func (d *Debug) Draw(screen *ebiten.Image) {
	d.Grid.Render(screen)
}

func (d *Debug) Update() error {
	return nil
}
