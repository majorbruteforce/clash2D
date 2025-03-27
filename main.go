package main

import (
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	utils "clash2D/pkg"
)

type Game struct {
	mapCutout *utils.Cutout
	mapImage  *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	tileX := g.mapCutout.Coordinates[115].X
	tileY := g.mapCutout.Coordinates[115].Y

	tile := g.mapImage.SubImage(image.Rect(tileX, tileY, tileX+g.mapCutout.TileWidth, tileY+g.mapCutout.TileHeight)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(tile, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {

	img, _, err := ebitenutil.NewImageFromFile("./static/assets/tileset.png")
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Tilemap Test")
	game := &Game{
		mapCutout: utils.NewCutout(352, 352, 11, 11),
		mapImage:  img,
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
