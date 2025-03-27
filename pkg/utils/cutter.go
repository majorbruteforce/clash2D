package main

type Cutout struct {
	SheetHeight  int16
	SheetWidth   int16
	TilesPerRow  int8
	NumberOfRows int8
	Coordinates  []struct{ x, y int16 }
	TileHeight   int16
	TileWidth    int16
}

func NewCutout(sheetHeight, sheetWidth int16, tilesPerRow, numberOfRows int8) *Cutout {

	tileWidth := sheetWidth / int16(tilesPerRow)
	tileHeight := sheetHeight / int16(numberOfRows)

	var coordinates []struct{ x, y int16 }

	for i := range numberOfRows {
		for j := range tilesPerRow {
			x := tileWidth * int16(j)
			y := tileHeight * int16(i)

			coordinates = append(coordinates, struct {
				x int16
				y int16
			}{x: x, y: y})
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
