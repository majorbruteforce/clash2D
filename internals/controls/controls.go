package controls

import (
	"clash2D/pkg/console"
)

type ControlKeyBuffer struct {
	Movement *console.KeyBuffer
}

var ControlBuffer = ControlKeyBuffer{
	Movement: console.NewKeyBuffer(2, MovementControls),
}
