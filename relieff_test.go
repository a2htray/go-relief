package relief

import (
	"fmt"
	"testing"
)

func TestReliefF_Run(t *testing.T) {
	model := NewReliefF([][]float64{
		[]float64{0, 0, -1},
		[]float64{0, 0, -2},
		[]float64{0, 0, -3},
		[]float64{0, 0, -4},
		[]float64{1, 1, 1},
		[]float64{1, 1, 2},
		[]float64{1, 1, 3},
		[]float64{1, 1, 4},
	}, []float64{0, 0, 0, 0, 1, 1, 1, 1}, []int{
		AttributeTypeDiscrete,
		AttributeTypeDiscrete,
		AttributeTypeDiscrete,
	}, 2)
	fmt.Println(model.Run(100))
}
