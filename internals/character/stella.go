package character

import (
	"clash2D/pkg/config"
	"clash2D/pkg/utils"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadStella() *Character {

	path := filepath.Join(config.RootDir, "assets", "stella_walk_1.png")
	sheet, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	Stella := NewCharacter(
		"Stella",
		sheet,
		utils.NewCutout(256, 512, 4, 8),
		0,
	)

	return Stella
}
