package debugutils

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Values struct {
	TileCoordX int
	TileCoordY int
}

var Val = Values{}

func (v *Values) ShowCoordinates(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("(%d,%d)", v.TileCoordX, v.TileCoordY), basicfont.Face7x13, 20, 20, color.White)
}
