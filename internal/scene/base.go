package scene

import (
	"fmt"
	"image/color"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/majorbruteforce/clash2D/internal/character"
	"github.com/majorbruteforce/clash2D/pkg/atmosphere"
	"github.com/majorbruteforce/clash2D/pkg/config"
	"golang.org/x/image/font/basicfont"
)

type Base struct {
	title string
	Map   *atmosphere.Map
	Lucy  *character.Character
}

var BaseScene = Base{
	title: "Welcome to your base!",
	Map: atmosphere.NewMap(
		path.Join(config.RootDir, "assets", "maps", "tileset-1.png"),
		path.Join(config.RootDir, "assets", "maps", "tileset-1.json"),
		11, 11,
	),
	Lucy: character.NewCharacter(
		"Lucy",
		path.Join(config.RootDir, "assets", "characters", "lucy_walk.png"),
		9, 9,
	),
}

func (b *Base) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("DistX: %f, DistY: %f", b.Lucy.Dist.X, b.Lucy.Dist.Y), basicfont.Face7x13, 50, 50, color.White)
	b.Map.Render(screen)
	b.Lucy.Render(screen)
}

func (b *Base) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		b.Lucy.Dist.Y = -16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		b.Lucy.Dist.Y = 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		b.Lucy.Dist.X = -32
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		b.Lucy.Dist.X = 32
	}

	b.Lucy.MoveRemainigDist()
	b.Lucy.RunSequence(character.LucySequences["WalkS"])
	return nil
}
