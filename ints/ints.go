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

const rangeSizeDefault = 32

func FiniteRange(range0 Range, ranges ...Range) []int {
	ranges = append(ranges, range0)
	result := make([]int, 0, rangeSizeDefault)
	for _, rangeConf := range ranges {
		isInclusive := 0
		if rangeConf.IsInclusive {
			isInclusive = 1
		}
		start := rangeConf.GetStart()
		end := rangeConf.GetEnd() + isInclusive
		step := rangeConf.GetStep()
		for i := start; i < end; i += step {
			result = append(result, i)
		}
	}
	return result
}
