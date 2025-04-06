package core

type Global struct {
	frameIndex int
	tickIndex  int
	fps        int
	tps        int
	unitSize   int
	debug      bool
}

type GlobalConfig struct {
	fps      int
	tps      int
	unitSize int
	debug    bool
}

var Gb = Global{
	frameIndex: 0,
	tickIndex:  0,
	fps:        60,
	tps:        60,
	unitSize:   32,
	debug:      false,
}

func (g *Global) RunFrameIndexCycle() {
	g.frameIndex++
	g.frameIndex %= g.fps
}

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

func (g *Global) Debug() bool {
	return g.debug
}

func (g *Global) SetValues(config *GlobalConfig) {
	g.fps = config.fps
	g.tps = config.tps
	g.unitSize = config.unitSize
	g.debug = config.debug
}
