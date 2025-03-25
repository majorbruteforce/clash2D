package main

import (
	"image"
	"log"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	spriteSheetWidth  = 352
	spriteSheetHeight = 352

	tilesPerRowInSpriteSheet = 11
	rowsInSpriteSheet        = 11

	tileWidth  = spriteSheetWidth / tilesPerRowInSpriteSheet
	tileHeight = spriteSheetHeight / rowsInSpriteSheet

	screenHeight = 270
	screenWidth  = 480

	originX = screenWidth / 2
	originY = screenHeight / 8
)

type Game struct {
	baseTileSet *ebiten.Image
	baseTileMap TileMapJSON
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	for i, layer := range g.baseTileMap.Layers {
		for idx, data := range layer.Data {
			if data == 0 {
				continue
			}

			// Get tile in tileset
			tileX := ((data - 1) % tilesPerRowInSpriteSheet) * tileWidth
			tileY := ((data - 1) / tilesPerRowInSpriteSheet) * tileHeight
			tile := g.baseTileSet.SubImage(image.Rect(tileX, tileY, tileX+tileWidth, tileY+tileHeight)).(*ebiten.Image)

			// Grid coords (3,4)
			gridX := idx % g.baseTileMap.Layers[i].Width
			gridY := idx / g.baseTileMap.Layers[i].Width

			// Isometric screen coords
			screenX := (gridX - gridY) * (tileWidth / 2)
			screenY := (gridX + gridY) * (tileHeight / 4)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Reset()
			op.GeoM.Translate(float64(screenX), float64(screenY))
			op.GeoM.Translate(-float64(tileWidth/2), -float64(tileHeight/2))

			op.GeoM.Translate(originX, originY)
			screen.DrawImage(tile, op)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	var baseTileMap TileMapJSON

	const baseTileMapJSONPath string = "./assets/test.json"

	err := LoadTileMapFromJSON(baseTileMapJSONPath, &baseTileMap)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := ebitenutil.NewImageFromFile("./assets/tileset.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		baseTileSet: img,
		baseTileMap: baseTileMap,
	}

	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Tilemap Test")
	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
