package utils

type Animator struct {
	Sequences map[string][]int
}

type AnimatorConfig struct {
	SequenceIntervals map[string][]struct{ start, end int }
}

func NewAnimator() *Animator {

	m := make(map[string][]int)

	animator := &Animator{
		Sequences: m,
	}

	return animator
}
