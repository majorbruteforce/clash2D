package utils

import (
	"image"
)

type Cutout struct {
	SheetHeight  int
	SheetWidth   int
	TilesPerRow  int
	NumberOfRows int
	Coordinates  []struct{ X, Y int }
	TileHeight   int
	TileWidth    int
}

func NewCutout(sheetWidth, sheetHeight int, tilesPerRow, numberOfRows int) *Cutout {

	tileWidth := sheetWidth / tilesPerRow
	tileHeight := sheetHeight / numberOfRows

	var coordinates []struct{ X, Y int }

	for i := range numberOfRows {
		for j := range tilesPerRow {
			x := tileWidth * j
			y := tileHeight * i

			coordinates = append(coordinates, struct {
				X int
				Y int
			}{X: x, Y: y})
		}
	}

	newCutout := &Cutout{
		SheetHeight:  sheetHeight,
		SheetWidth:   sheetWidth,
		TilesPerRow:  tilesPerRow,
		NumberOfRows: numberOfRows,
		Coordinates:  coordinates,
		TileHeight:   tileHeight,
		TileWidth:    tileWidth,
	}

	return newCutout
}

func (c *Cutout) GetTileRectById(id int) image.Rectangle {
	tileX := c.Coordinates[id].X
	tileY := c.Coordinates[id].Y
	rect := image.Rect(tileX, tileY, tileX+c.TileWidth, tileY+c.TileHeight)

	return rect
}
