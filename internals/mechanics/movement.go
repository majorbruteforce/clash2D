package mechanics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var keyMap = map[ebiten.Key][2]int{
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
	{0, 0}:   {0, 0},   // No movement
	{0, 1}:   {0, 1},   // Down
	{0, -1}:  {0, -1},  // Up
	{1, 0}:   {1, 0},   // Right
	{-1, 0}:  {-1, 0},  // Left
	{1, 1}:   {1, 1},   // Down-Right
	{-1, 1}:  {-1, 1},  // Down-Left
	{1, -1}:  {1, -1},  // Up-Right
	{-1, -1}: {-1, -1}, // Up-Left
}

var directionValueMap = map[[2]int]string{
	{0, 0}:   "",
	{0, 1}:   "S",
	{0, -1}:  "N",
	{1, 0}:   "E",
	{-1, 0}:  "W",
	{1, 1}:   "SE",
	{-1, 1}:  "SW",
	{1, -1}:  "NE",
	{-1, -1}: "NW",
}

func GetDirection(key1 ebiten.Key, opKey ...ebiten.Key) string {
	if len(opKey) == 0 {
		d := keyMap[key1]
		return directionValueMap[d]
	}

	netX, netY := keyMap[key1][0]+keyMap[opKey[0]][0], keyMap[key1][1]+keyMap[opKey[0]][1]

	return directionValueMap[[2]int{netX, netY}]
}

func GetTransltionXY(speed float64, key1 ebiten.Key, opKey ...ebiten.Key) (x, y float64) {
	if len(opKey) == 0 {
		t := keyMap[key1]
		return float64(t[0]) * speed, float64(t[1]) * speed
	}

	netX, netY := keyMap[key1][0]+keyMap[opKey[0]][0], keyMap[key1][1]+keyMap[opKey[0]][1]

	t := translationValueMap[[2]int{netX, netY}]

	return float64(speed) * t[0], float64(speed) * t[1]
}
