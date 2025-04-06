package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Base struct {
	title string
}

var BaseScene = Base{
	title: "Welcome to your base!",
}

func (b *Base) Draw(screen *ebiten.Image) {
	text.Draw(screen, b.title, basicfont.Face7x13, 100, 100, color.White)
}

func (b *Base) Update() error {
	return nil
}
