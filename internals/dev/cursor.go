package dev

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func DisplayCursorPosition(screen *ebiten.Image) {

	curX, curY := ebiten.CursorPosition()
	text.Draw(screen, fmt.Sprintf("%d, %d", curX, curY), basicfont.Face7x13, 100, 100, color.White)
}
