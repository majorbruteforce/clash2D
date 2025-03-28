package main

import (
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"clash2D/internals/core"
	"clash2D/pkg/utils"
)

const (
	frameInterval float64 = 0.2 // in s
	FPS           float64 = 0.017
)

var (
	drawCount, frame = 0, 0
	tile             *ebiten.Image
)

type Game struct {
	mapCutout *utils.Cutout
	mapImage  *ebiten.Image
	global    *core.Global
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.global.FrameIndex % 6 == 0 {
		row := 0
		tileX := g.mapCutout.Coordinates[frame].X
		tileY := g.mapCutout.Coordinates[frame].Y + row*g.mapCutout.TileHeight
		tile = g.mapImage.SubImage(image.Rect(tileX, tileY, tileX+g.mapCutout.TileWidth, tileY+g.mapCutout.TileHeight)).(*ebiten.Image)

		frame++
		frame %= 8
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(160-g.mapCutout.TileWidth/2), float64(120-g.mapCutout.TileHeight/2))
	screen.DrawImage(tile, op)

	g.global.RunFrameIndexCycle()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {

	img, _, err := ebitenutil.NewImageFromFile("./static/assets/8Direction_TopDown_Character Sprites_ByBossNelNel/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tilemap Test")
	game := &Game{
		mapCutout: utils.NewCutout(209, 326, 9, 9),
		mapImage:  img,
		global:    core.NewGlobal(&core.DefaultGlobalConfig),
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
