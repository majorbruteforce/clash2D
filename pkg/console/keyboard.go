package console

import "github.com/hajimehoshi/ebiten/v2"

type KeyBuffer struct {
	MovementKeys []ebiten.Key
}

var movementControls = []ebiten.Key{
	ebiten.KeyUp,
	ebiten.KeyLeft,
	ebiten.KeyDown,
	ebiten.KeyRight,
	ebiten.KeyW,
	ebiten.KeyA,
	ebiten.KeyS,
	ebiten.KeyD,
}

func (k *KeyBuffer) checkMovementKey(key ebiten.Key) {
	for _, moveKey := range movementControls {
		if moveKey == key {
			k.MovementKeys = append(k.MovementKeys, key)
			return
		}
	}
}

func (k *KeyBuffer) HandleKeyPressEvents(key ebiten.Key) {
	k.checkMovementKey(key)
}
