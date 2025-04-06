package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/majorbruteforce/clash2D/pkg/errorutils"
)

type Cutout struct {
	Img          *ebiten.Image
	TilesPerRow  int
	NumberOfRows int
	Coordinates  []struct{ X, Y int }
	TileHeight   int
	TileWidth    int
}

func NewCutout(path string, tilesPerRow, numberOfRows int) *Cutout {

	img, _, err := ebitenutil.NewImageFromFile(path)
	errorutils.CheckFatal(err)

	sheetWidth, sheetHeight := img.Bounds().Dx(), img.Bounds().Dy()

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
		Img:          img,
		TilesPerRow:  tilesPerRow,
		NumberOfRows: numberOfRows,
		Coordinates:  coordinates,
		TileHeight:   tileHeight,
		TileWidth:    tileWidth,
	}

	return newCutout
}

func (c *Cutout) GetSubImageByIndex(idx int) *ebiten.Image {
	tileX := c.Coordinates[idx].X
	tileY := c.Coordinates[idx].Y
	rect := image.Rect(tileX, tileY, tileX+c.TileWidth, tileY+c.TileHeight)

	subImage := c.Img.SubImage(rect).(*ebiten.Image)

	return subImage
}
