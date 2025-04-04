package core

type GlobalConfig struct {
	fps      int
	tps      int
	originX  int
	originY  int
	unitSize int
}
type Global struct {
	frameIndex int
	tickIndex  int
	fps        int
	tps        int
	originX    int
	originY    int
	unitSize   int // in px
}

var (
	DefaultGlobalConfig = GlobalConfig{
		fps:      60,
		tps:      60,
		unitSize: 32,
		originX:  100,
		originY:  100,
	}
)

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

func (g *Global) UnitSize() int {
	return g.unitSize
}

func (g *Global) Origin() (int, int) {
	return g.originX, g.originY
}

func (g *Global) UpdateOrigin(x, y int) {
	g.originX += x
	g.originY += y
}

func (g *Global) SetValues(config *GlobalConfig) {
	g.fps = config.fps
	g.tps = config.tps
	g.unitSize = config.unitSize
	g.originX = config.originX
	g.originY = config.originY
}

var Gb = Global{
	frameIndex: 0,
	tickIndex:  0,
	fps:        DefaultGlobalConfig.fps,
	tps:        DefaultGlobalConfig.tps,
	originX:    DefaultGlobalConfig.originX,
	originY:    DefaultGlobalConfig.originY,
	unitSize:   DefaultGlobalConfig.unitSize,
}
