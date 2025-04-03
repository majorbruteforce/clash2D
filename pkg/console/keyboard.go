package console

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyBuffer struct {
	Data    map[ebiten.Key]struct{}
	Queue   []ebiten.Key
	length  int
	keyList []ebiten.Key
}

func NewKeyBuffer(length int, keyList []ebiten.Key) *KeyBuffer {
	return &KeyBuffer{
		Data:    make(map[ebiten.Key]struct{}),
		Queue:   []ebiten.Key{},
		length:  length,
		keyList: keyList,
	}
}

func (k *KeyBuffer) Load(key ebiten.Key) {

	if _, found := k.Data[key]; found {
		return
	}

	if len(k.Queue) == k.length {
		removed := k.Queue[0]
		k.Queue = k.Queue[1:]
		delete(k.Data, removed)
	}

	k.Queue = append(k.Queue, key)
	k.Data[key] = struct{}{}
}

func (k *KeyBuffer) Unload(key ebiten.Key) {
	if _, found := k.Data[key]; !found {
		return
	}

	delete(k.Data, key)

	for i, v := range k.Queue {
		if v == key {
			k.Queue = append(k.Queue[:i], k.Queue[i+1:]...)
			break
		}
	}
}

func (k *KeyBuffer) Values() []ebiten.Key {
	var values []ebiten.Key
	for value := range k.Data {
		values = append(values, value)
	}

	return values
}

func (k *KeyBuffer) MonitorKeys() {
	for _, key := range k.keyList {
		if inpututil.IsKeyJustPressed(key) {
			k.Load(key)
		}

		if inpututil.IsKeyJustReleased(key) {
			k.Unload(key)
		}
	}
}
