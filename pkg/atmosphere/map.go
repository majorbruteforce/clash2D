package atmosphere

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/majorbruteforce/clash2D/internal/core"
	"github.com/majorbruteforce/clash2D/pkg/debugutils"
	"github.com/majorbruteforce/clash2D/pkg/geom"
	"github.com/majorbruteforce/clash2D/pkg/sprite"
	"github.com/majorbruteforce/clash2D/pkg/utils"
)

type Map struct {
	*sprite.Cutout
	*utils.TileMapJSON
}

func NewMap(imgPath, jsonPath string, tilesPerRow, numberOfRows int) *Map {
	return &Map{
		Cutout:      sprite.NewCutout(imgPath, tilesPerRow, numberOfRows),
		TileMapJSON: utils.ExtarctTileMapFromJSON(jsonPath),
	}
}

func (m *Map) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	originX, originY := core.Sh.Origin()

	for _, layer := range m.Layers {
		for idx, tileId := range layer.Data {
			if tileId == 0 {
				continue
			}

			gridX := idx % layer.Width
			gridY := idx / layer.Width

			projX := (gridX - gridY) * (m.Cutout.TileWidth / 2)
			projY := (gridX + gridY) * (m.Cutout.TileHeight / 4)

			screenX := float64(projX) + float64(originX) - float64(m.Cutout.TileWidth/2)
			screenY := float64(projY) + float64(originY) - float64(m.Cutout.TileHeight/2)

			op.GeoM.Reset()
			// op.ColorM.Reset()
			op.GeoM.Translate(screenX, screenY)

			tile := m.Cutout.GetTileByIndex(tileId - 1)
			screen.DrawImage(tile, op)

			if core.Gb.Debug() {
				curX, curY := ebiten.CursorPosition()
				rh := geom.NewRhombusFromTile(int(screenX), int(screenY), m.TileWidth, m.TileHeight, 0, -1)

				if (rh.IsPointInside(geom.Point{X: float32(curX), Y: float32(curY)})) {
					m.HighlightTile(screen, op)
					debugutils.Val.ShowCoordinates(screen)
					debugutils.Val.TileCoordX, debugutils.Val.TileCoordY = gridX, gridY
				}
			}

		}
	}
}

func (m *Map) HighlightTile(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	screen.DrawImage(m.GetTileByIndex(97), op)
}
