package character

var LucySequences = map[string]AnimationSequence{
	"WalkS": AnimationSequence{
		Start: 0,
		End:   7,
		Step:  1,
		Speed: 6,
	},
	"WalkSE": AnimationSequence{
		Start: 9,
		End:   16,
		Step:  1,
		Speed: 6,
	},
	"WalkE": AnimationSequence{
		Start: 18,
		End:   25,
		Step:  1,
		Speed: 6,
	},
	"WalkNE": AnimationSequence{
		Start: 27,
		End:   34,
		Step:  1,
		Speed: 6,
	},
	"WalkN": AnimationSequence{
		Start: 36,
		End:   43,
		Step:  1,
		Speed: 6,
	},
	"WalkNW": AnimationSequence{
		Start: 45,
		End:   52,
		Step:  1,
		Speed: 6,
	},
	"WalkW": AnimationSequence{
		Start: 54,
		End:   61,
		Step:  1,
		Speed: 6,
	},
	"WalkSW": AnimationSequence{
		Start: 63,
		End:   70,
		Step:  1,
		Speed: 6,
	},
}

func (c *Character) Walk(seq string) {

	switch seq {
	case "WalkNE":
		c.Seq = "WalkNE"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = 16
		c.Dist.Y = -8

	case "WalkSW":
		c.Seq = "WalkSW"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = -16
		c.Dist.Y = 8

	case "WalkNW":
		c.Seq = "WalkNW"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = -16
		c.Dist.Y = -8

	case "WalkSE":
		c.Seq = "WalkSE"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = 16
		c.Dist.Y = 8
	case "WalkN":
		c.Seq = "WalkN"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.Y = -16

	case "WalkS":
		c.Seq = "WalkS"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.Y = 16

	case "WalkW":
		c.Seq = "WalkW"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = -32

	case "WalkE":
		c.Seq = "WalkE"
		c.SetFrame(LucySequences[c.Seq].Start)
		c.Dist.X = 32

	}
}