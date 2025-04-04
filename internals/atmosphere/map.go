package atmosphere

import (
	"clash2D/internals/core"
	"clash2D/pkg/exception"
	"clash2D/pkg/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Map struct {
	*utils.TileMapJSON
	*utils.Cutout
	Sheet *ebiten.Image
	Debug bool
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
		Debug:       true,
	}
}

func (m *Map) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	originX, originY := core.Gb.Origin()

	for _, layer := range m.Layers {
		for idx, tileId := range layer.Data {
			if tileId == 0 {
				continue
			}

			highlight := m.Sheet.SubImage(m.Cutout.GetTileRectangleById(97)).(*ebiten.Image)

			tile := m.Sheet.SubImage(m.Cutout.GetTileRectangleById(tileId - 1)).(*ebiten.Image)

			gridX := idx % layer.Width
			gridY := idx / layer.Width

			projX := (gridX - gridY) * (m.Cutout.TileWidth / 2)
			projY := (gridX + gridY) * (m.Cutout.TileHeight / 4)

			screenX := float64(projX) + float64(originX)
			// - float64(m.Cutout.TileWidth/2)
			screenY := float64(projY) + float64(originY)
			// - float64(m.Cutout.TileHeight/2)

			op.GeoM.Reset()
			op.GeoM.Translate(screenX, screenY)

			if false {
				screen.DrawImage(highlight, op)
				text.Draw(screen, "Focused!", basicfont.Face7x13, 100, 900, color.White)
			} else {
				screen.DrawImage(tile, op)
			}
			if m.Debug {
				pointA := struct{ x, y float32 }{
					x: float32(screenX),
					y: float32(screenY) + float32(m.TileHeight/2) - 4,
				}
				pointB := struct{ x, y float32 }{
					x: float32(screenX) + float32(m.TileWidth/2),
					y: float32(screenY) + float32(m.TileHeight/4) - 4,
				}
				pointC := struct{ x, y float32 }{
					x: float32(screenX) + float32(m.TileWidth),
					y: float32(screenY) + float32(m.TileHeight/2) - 4,
				}
				pointD := struct{ x, y float32 }{
					x: float32(screenX) + float32(m.TileWidth/2),
					y: float32(screenY) + float32(m.TileHeight*3/4) - 4,
				}

				vector.StrokeLine(screen, pointA.x, pointA.y, pointB.x, pointB.y, 1, color.RGBA{255, 0, 0, 100}, false)
				vector.StrokeLine(screen, pointC.x, pointC.y, pointB.x, pointB.y, 1, color.RGBA{255, 0, 0, 100}, false)
				vector.StrokeLine(screen, pointC.x, pointC.y, pointD.x, pointD.y, 1, color.RGBA{255, 0, 0, 100}, false)
				vector.StrokeLine(screen, pointA.x, pointA.y, pointD.x, pointD.y, 1, color.RGBA{255, 0, 0, 100}, false)
			}
		}
	}
}
