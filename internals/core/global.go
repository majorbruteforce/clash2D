package core

type Global struct {
	FrameIndex int
	FPS        int
}

type GlobalConfig struct {
	FPS int
}

var (
	DefaultGlobalConfig = GlobalConfig{
		FPS: 60,
	}
)

func NewGlobal(config *GlobalConfig) *Global {

	global := &Global{
		FrameIndex: 0,
		FPS:        config.FPS,
	}

	return global
}

// Call for every Draw iteration to cycle through
func (g *Global) RunFrameIndexCycle() {
	g.FrameIndex++
	g.FrameIndex %= g.FPS
}
