package floats

import (
	"fmt"
	"testing"
)

func TestArgSort(t *testing.T) {
	argSort := ArgSort([][]float64{
		{1, 1},
		{0, 0},
		{1, 3},
		{1, 4},
		{2, 2},
	}, func(v []float64) float64 {
		return EuclideanDistance(v, []float64{0, 0})
	}, false)

	fmt.Println(argSort)
}
