package core

import "clash2D/pkg/console"

type Global struct {
	frameIndex int
	tickIndex  int
	fps        int
	tps        int
	*console.KeyBuffer
}

type GlobalConfig struct {
	fps int
	tps int
}

var (
	DefaultGlobalConfig = GlobalConfig{
		fps: 60,
		tps: 60,
	}
)

func NewGlobal(config *GlobalConfig) *Global {

	global := &Global{
		frameIndex: 0,
		tickIndex:  0,
		fps:        config.fps,
		tps:        config.tps,
	}

	return global
}

// Invoke for every draw iteration to maintain draw count
func (g *Global) RunFrameIndexCycle() {
	g.frameIndex++
	g.frameIndex %= g.fps
}

// Invoke for every update iteration to maintain draw count
func (g *Global) RunTickIndexCycle() {
	g.tickIndex++
	g.tickIndex %= g.tps
}

func (g *Global) FrameIndex() int {
	return g.frameIndex
}

func (g *Global) TickIndex() int {
	return g.tickIndex
}


