package atmopshere

import (
	"clash2D/pkg/exception"
	"clash2D/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Map struct {
	*utils.TileMapJSON
	*utils.Cutout
	Sheet *ebiten.Image
}

func NewMap(jsonPath string, imagePath string, imageHeigth, imageWidth, tilesPerRow, numberOfRows int) *Map {
	tilemap := utils.ExtarctTileMapFromJSON(jsonPath)

	sheet, _, err := ebitenutil.NewImageFromFile(imagePath)
	exception.CheckFatal(err)

	cutout := utils.NewCutout(imageWidth, imageHeigth, tilesPerRow, numberOfRows)

	return &Map{
		TileMapJSON: tilemap,
		Cutout:      cutout,
		Sheet:       sheet,
	}
}

func (m *Map) Render(originX, originY int, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for _, layer := range m.Layers {
		for idx, tileId := range layer.Data {
			if tileId == 0 {
				continue
			}
			tile := m.Sheet.SubImage(m.Cutout.GetTileRectangleById(tileId - 1)).(*ebiten.Image)

			gridX := idx % layer.Width
			gridY := idx / layer.Width

			screenX := (gridX - gridY) * (m.Cutout.TileWidth / 2)
			screenY := (gridX + gridY) * (m.Cutout.TileHeight / 4)

			op.GeoM.Reset()
			op.GeoM.Translate(float64(screenX), float64(screenY))
			op.GeoM.Translate(-float64(m.Cutout.TileWidth/2), -float64(m.Cutout.TileHeight/2))
			op.GeoM.Translate(float64(originX), float64(originY))

			screen.DrawImage(tile, op)
		}
	}
}
