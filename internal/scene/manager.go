package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Draw(screen *ebiten.Image)
	Update() error
}

type Manager struct {
	currentScene Scene
}

func NewManager() *Manager{
	return &Manager{}
}

func (m *Manager) SetScene(scene Scene) {
	m.currentScene = scene
}

func (m *Manager) Update() error {
	if m.currentScene != nil {
		return m.currentScene.Update()
	}
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	if m.currentScene != nil {
		m.currentScene.Draw(screen)
	}
}
