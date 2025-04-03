package character

var StellaSequences = map[string]AnimationSequence{
	"WalkS": AnimationSequence{
		Start: 0,
		End:   3,
		Step:  1,
		Speed: 10,
	},
	"WalkW": AnimationSequence{
		Start: 4,
		End:   7,
		Step:  1,
		Speed: 10,
	},
	"WalkE": AnimationSequence{
		Start: 8,
		End:   11,
		Step:  1,
		Speed: 10,
	},
	"WalkN": AnimationSequence{
		Start: 12,
		End:   15,
		Step:  1,
		Speed: 10,
	},
	"WalkSE": AnimationSequence{
		Start: 16,
		End:   19,
		Step:  1,
		Speed: 10,
	},
	"WalkSW": AnimationSequence{
		Start: 20,
		End:   23,
		Step:  1,
		Speed: 10,
	},
	"WalkNW": AnimationSequence{
		Start: 24,
		End:   27,
		Step:  1,
		Speed: 10,
	},
	"WalkNE": AnimationSequence{
		Start: 28,
		End:   31,
		Step:  1,
		Speed: 10,
	},
}
