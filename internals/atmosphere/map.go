package atmosphere

import (
	"clash2D/internals/core"
	"clash2D/pkg/exception"
	"clash2D/pkg/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

			screen.DrawImage(tile, op)

			if m.Debug {
				cursorX, cursorY := ebiten.CursorPosition()

				pointA, pointB, pointC, pointD := m.getIsometricCoordinates(screenX, screenY)

				// vector.StrokeLine(screen, pointA.x, pointA.y, pointB.x, pointB.y, 1, color.RGBA{255, 0, 0, 100}, false)
				// vector.StrokeLine(screen, pointC.x, pointC.y, pointB.x, pointB.y, 1, color.RGBA{255, 0, 0, 100}, false)
				// vector.StrokeLine(screen, pointC.x, pointC.y, pointD.x, pointD.y, 1, color.RGBA{255, 0, 0, 100}, false)
				// vector.StrokeLine(screen, pointA.x, pointA.y, pointD.x, pointD.y, 1, color.RGBA{255, 0, 0, 100}, false)

				if isPointInRhombus(float64(cursorX), float64(cursorY), pointA, pointB, pointC, pointD) {
					screen.DrawImage(highlight, op)
					if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
						layer.Data[idx] = 67
					}
				}
			}
		}
	}
}

func isPointInTriangle(px, py float64, A, B, C struct{ x, y float32 }) bool {
	ax, ay := float64(A.x), float64(A.y)
	bx, by := float64(B.x), float64(B.y)
	cx, cy := float64(C.x), float64(C.y)

	denominator := (by-cy)*(ax-cx) + (cx-bx)*(ay-cy)
	alpha := ((by-cy)*(px-cx) + (cx-bx)*(py-cy)) / denominator
	beta := ((cy-ay)*(px-cx) + (ax-cx)*(py-cy)) / denominator
	gamma := 1 - alpha - beta

	return alpha >= 0 && beta >= 0 && gamma >= 0
}

func isPointInRhombus(px, py float64, A, B, C, D struct{ x, y float32 }) bool {
	return isPointInTriangle(px, py, A, B, C) || isPointInTriangle(px, py, A, C, D)
}

func (m *Map) getIsometricCoordinates(screenX, screenY float64) (a, b, c, d struct{ x, y float32 }) {
	pointA := struct{ x, y float32 }{
		x: float32(screenX),
		y: float32(screenY) + float32(m.TileHeight/2) - 3,
	}
	pointB := struct{ x, y float32 }{
		x: float32(screenX) + float32(m.TileWidth/2),
		y: float32(screenY) + float32(m.TileHeight/4) - 3,
	}
	pointC := struct{ x, y float32 }{
		x: float32(screenX) + float32(m.TileWidth),
		y: float32(screenY) + float32(m.TileHeight/2) - 3,
	}
	pointD := struct{ x, y float32 }{
		x: float32(screenX) + float32(m.TileWidth/2),
		y: float32(screenY) + float32(m.TileHeight*3/4) - 3,
	}

	return pointA, pointB, pointC, pointD

}
