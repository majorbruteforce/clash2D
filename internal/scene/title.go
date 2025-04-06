package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Title struct {
	title string
}

var TitleScene = Title{
	title: "Clash2D",
}

func (t *Title) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.title, basicfont.Face7x13, 100, 100, color.White)
}

func (t *Title) Update() error {
	return nil
}
