package controls

import (
	"clash2D/internals/character"
	"clash2D/internals/mechanics"

	"github.com/hajimehoshi/ebiten/v2"
)

var MovementControls = []ebiten.Key{
	ebiten.KeyUp,
	ebiten.KeyLeft,
	ebiten.KeyDown,
	ebiten.KeyRight,
	ebiten.KeyW,
	ebiten.KeyA,
	ebiten.KeyS,
	ebiten.KeyD,
}

func ActuateTranslation(speed float64, c *character.Character) {

	keys := ControlBuffer.Movement.Values()

	if len(keys) == 0 {
		return
	} else if len(keys) == 1 {
		c.Position.Dx, c.Position.Dy = mechanics.GetTransltionXY(speed, keys[0])
	} else {
		c.Position.Dx, c.Position.Dy = mechanics.GetTransltionXY(speed, keys[0], keys[1])
	}

}

func ActuateAnimation(c *character.Character){
	
}
