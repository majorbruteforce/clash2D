package core

type Global struct {
	FrameIndex int
	TickIndex  int
	FPS        int
	TPS        int
}

type GlobalConfig struct {
	FPS int
	TPS int
}

var (
	DefaultGlobalConfig = GlobalConfig{
		FPS: 60,
		TPS: 60,
	}
)

func NewGlobal(config *GlobalConfig) *Global {

	global := &Global{
		FrameIndex: 0,
		TickIndex:  0,
		FPS:        config.FPS,
		TPS:        config.TPS,
	}

	return global
}

// Invoke for every draw iteration to maintain draw count
func (g *Global) RunFrameIndexCycle() {
	g.FrameIndex++
	g.FrameIndex %= g.FPS
}

// Invoke for every update iteration to maintain draw count
func (g *Global) RunTickIndexCycle() {
	g.TickIndex++
	g.TickIndex %= g.TPS
}
