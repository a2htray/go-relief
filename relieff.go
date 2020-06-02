package relief

import (
	"github.com/a2htray/relief/floats"
)

func _sort(t []float64) func([]float64) float64 {
	return func(v []float64) float64 {
		return floats.EuclideanDistance(v, t)
	}
}

type ReliefF struct {
	*Relief
	k int
}

func NewReliefF(values [][]float64, targets []float64, attributeTypes []int, k int) *ReliefF {
	relief := NewRelief(values, targets, attributeTypes)
	return &ReliefF{
		Relief: relief,
		k:      k,
	}
}

func (r *ReliefF) Run(m int) []float64 {
	return r.runK(m, r.k)
}
