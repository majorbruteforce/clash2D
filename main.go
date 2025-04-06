package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/majorbruteforce/clash2D/internal/scene"
	"github.com/majorbruteforce/clash2D/pkg/errorutils"
)

type Game struct {
	SceneManager *scene.Manager
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.SceneManager.SetScene(&scene.BaseScene)
	}
	return g.SceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 480, 270
}

func main() {

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Tilemap Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &Game{
		SceneManager: scene.NewManager(),
	}

	game.SceneManager.SetScene(&scene.TitleScene)

	err := ebiten.RunGame(game)
	errorutils.CheckFatal(err)

}
