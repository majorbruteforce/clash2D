package utils

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
