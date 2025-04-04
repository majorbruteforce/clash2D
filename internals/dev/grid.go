package dev

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Grid struct {
	StartX     float32
	StartY     float32
	CellHeight float32
	CellWidth  float32
	Stroke     float32
	Color      color.Color
}

func NewGrid(startX, startY, cellWidth, cellHeight, stroke float32, clr color.Color) *Grid {
	return &Grid{
		StartX:     startX,
		StartY:     startY,
		CellHeight: cellHeight,
		CellWidth:  cellWidth,
		Stroke:     stroke,
		Color:      clr,
	}
}

func (g *Grid) Render(screen *ebiten.Image) {
	var screenWidth, screenHeight float32 = 960, 540

	for x := g.StartX; x <= screenWidth; x += g.CellWidth {
		for y := g.StartY; y <= screenHeight; y += g.CellHeight {
			vector.StrokeLine(screen, x, y, x, y+screenHeight, g.Stroke, g.Color, true)
			vector.StrokeLine(screen, x, y, x+screenWidth, y, g.Stroke, g.Color, true)
		}
	}
}
