package ints

type Range struct {
	Start, End, Step *int
	IsInclusive      bool
}

func (r *Range) GetStart() int {
	if r.Start == nil {
		return 0
	}
	return *r.Start
}

func (r *Range) GetEnd() int {
	if r.End == nil {
		return 0
	}
	return *r.End
}

func (r *Range) GetStep() int {
	if r.Step == nil {
		return 1
	}
	return *r.Step
}

func FiniteRange(rangeConf Range) []int {
	isInclusive := func() int {
		if rangeConf.IsInclusive {
			return 1
		}
		return 0
	}
	start := rangeConf.GetStart()
	end := rangeConf.GetEnd() + isInclusive()
	step := rangeConf.GetStep()
	size := 1 + ((rangeConf.GetEnd() + isInclusive() - rangeConf.GetStart()) / rangeConf.GetStep())
	if size <= 0 {
		return []int{}
	}

	intRange := make([]int, 0, size)
	for i := start; i < end; i += step {
		intRange = append(intRange, i)
	}
	return intRange
}
