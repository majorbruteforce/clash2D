package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screnHeight int) {
	return 480, 240
}

func main() {

	err := ebiten.RunGame(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
