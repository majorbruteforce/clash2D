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
	Moves []string
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
	Moves: []string{
		"WalkS",
		"WalkS",
		"WalkS",
		"WalkS",
		"WalkS",
		"WalkE",
		"WalkSE",
		"WalkE",
		"WalkNE",
	},
}

func (b *Base) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0, 200, 210, 50})

	text.Draw(screen, fmt.Sprintf("DistX: %f, DistY: %f", b.Lucy.Dist.X, b.Lucy.Dist.Y), basicfont.Face7x13, 50, 50, color.White)
	b.Map.Render(screen)
	b.Lucy.Render(screen)
}

func (b *Base) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		b.Lucy.Walk("WalkN")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		b.Lucy.Walk("WalkS")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		b.Lucy.Walk("WalkW")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		b.Lucy.Walk("WalkE")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		b.Lucy.Walk("WalkNE")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		b.Lucy.Walk("WalkSW")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		b.Lucy.Walk("WalkSE")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		b.Lucy.Walk("WalkNW")
	}

	if len(b.Moves) > 0 && b.Lucy.Dist.X == 0 && b.Lucy.Dist.Y == 0 {
		b.Lucy.Walk(b.Moves[0])
		b.Moves = b.Moves[1:]
	}

	b.Lucy.RunSequence()
	b.Lucy.MoveRemainigDist()
	return nil
}
