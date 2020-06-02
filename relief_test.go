package relief

import (
	"fmt"
	"testing"
)

func TestRelief_Run(t *testing.T) {
	model := NewRelief([][]float64{
		[]float64{0, 0},
		[]float64{0, 1},
		[]float64{1, 0},
		[]float64{1, 1},
		[]float64{1, 1},
	}, []float64{0, 0, 0, 1, 1}, []int{
		AttributeTypeDiscrete,
		AttributeTypeDiscrete,
	})
	fmt.Println(model.Run(20))
}
