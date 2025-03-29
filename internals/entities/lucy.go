package entities

import (
	"clash2D/internals/core"
	"clash2D/pkg/config"
	"clash2D/pkg/utils"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadLucy() *core.Entity {

	path := filepath.Join(config.RootDir, "assets", "8Direction_TopDown_Character Sprites_ByBossNelNel", "SpriteSheet.png")
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	Lucy := core.NewEntity(
		"Lucy",
		img,
		utils.NewCutout(209, 326, 9, 9),
		9,
	)

	return Lucy
}
