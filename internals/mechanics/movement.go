package mechanics

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var translationKeyMap = map[ebiten.Key][2]int{
	ebiten.KeyUp:    {0, -1},
	ebiten.KeyLeft:  {-1, 0},
	ebiten.KeyDown:  {0, 1},
	ebiten.KeyRight: {1, 0},
	ebiten.KeyW:     {0, -1},
	ebiten.KeyA:     {-1, 0},
	ebiten.KeyS:     {0, 1},
	ebiten.KeyD:     {1, 0},
}

var translationValueMap = map[[2]int][2]float64{
	{0, 0}:   {0, 0},                             // No movement
	{0, 1}:   {0, 1},                             // Down
	{0, -1}:  {0, -1},                            // Up
	{1, 0}:   {1, 0},                             // Right
	{-1, 0}:  {-1, 0},                            // Left
	{1, 1}:   {1 / math.Sqrt2, 1 / math.Sqrt2},   // Down-Right
	{-1, 1}:  {-1 / math.Sqrt2, 1 / math.Sqrt2},  // Down-Left
	{1, -1}:  {1 / math.Sqrt2, -1 / math.Sqrt2},  // Up-Right
	{-1, -1}: {-1 / math.Sqrt2, -1 / math.Sqrt2}, // Up-Left
}

func GetTransltionXY(speed int, key1 ebiten.Key, opKey ...ebiten.Key) (x, y float64) {
	if len(opKey) == 0 {
		t := translationKeyMap[key1]
		return float64(speed * t[0]), float64(speed * t[1])
	}

	netX, netY := translationKeyMap[key1][0]+translationKeyMap[opKey[0]][0], translationKeyMap[key1][1]+translationKeyMap[opKey[0]][1]

	t := translationValueMap[[2]int{netX, netY}]

	return float64(speed) * t[0], float64(speed) * t[1]
}

// Make a channel to capture key presses in each tick, store in a container
// at the end of the tick, evaluate the last two values captured
// cleanup the container