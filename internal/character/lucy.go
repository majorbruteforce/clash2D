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
