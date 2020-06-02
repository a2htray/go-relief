package relief

import (
	"fmt"
	"github.com/a2htray/relief/floats"
	"github.com/a2htray/relief/ints"
	"math"
	"math/rand"
	"strings"
)

const (
	AttributeTypeDiscrete = iota
	AttributeTypeContinuous
)

func discreteDiff(v1, v2 float64) float64 {
	if v1 == v2 {
		return 1.0
	}
	return .0
}

func continuousDiff(v1, v2, min, max float64) float64 {
	return math.Abs(v1-v2) / (max - min)
}

type Relief struct {
	Matrix  *floats.Matrix // 保存原始数据
	Targets []float64      // 每一个样本对应的标签 slice
	// 每一维特征的类型集合
	// 离散型 AttributeTypeDiscrete
	// 连续型 attributeTypeContinuous
	AttributeTypes []int
	// 保存每一维的最小值与最大值
	minMaxes     []floats.MinMax
}

func (r *Relief) Count() int {
	return r.Matrix.M
}

func (r *Relief) Dim() int {
	return r.Matrix.N
}

func NewRelief(values [][]float64, targets []float64, attributeTypes []int) *Relief {
	if len(values) == 0 {
		panic("the number of samples must not be 0")
	}
	if len(values) != len(targets) {
		panic("the numbers of samples and targets do not match")
	}
	if len(values[0]) != len(attributeTypes) {
		panic("the numbers of features and attribute types do not match")
	}

	if !ints.AllIn([]int{AttributeTypeDiscrete, AttributeTypeContinuous}, attributeTypes) {
		panic(fmt.Sprintf("attribute types must in array [%s]", strings.Join([]string{"0", "1"}, ",")))
	}

	// 保存每一维的最大值与最小值
	minMaxes := make([]floats.MinMax, 0)
	for j := 0; j < len(attributeTypes); j++ {
		ts := make([]float64, 0)
		for _, i := range ints.Iter(len(targets)) {
			ts = append(ts, values[i][j])
		}
		minMaxes = append(minMaxes, floats.NewMinMax(floats.Min(ts), floats.Max(ts)))
	}

	return &Relief{
		Matrix:         floats.NewMatrix(values),
		Targets:        targets,
		AttributeTypes: attributeTypes,
		minMaxes:       minMaxes,
	}
}

func (r *Relief) runK(m, k int) []float64 {
	w := make([]float64, r.Dim())
	for iter := 0; iter < m; iter++ {
		var R []float64
		// 随机生成一个下标，并选择该样本
		rIndex := rand.Intn(r.Count())
		R = r.Matrix.Row(rIndex)
		target := r.Targets[rIndex]

		// 相同标签的样本下标 slice
		equalIndexes := ints.Remove(floats.EqualIndexes(r.Targets, target), rIndex)
		equalMatrix := r.Matrix.Rows(equalIndexes)
		equalSortK := floats.ArgSort(equalMatrix.Data, _sort(R), false)[:k]
		equalMatrix = equalMatrix.Rows(equalSortK)

		// 不同标签的样本下标 slice
		notEqualIndexes := floats.NotEqualIndexes(r.Targets, target)
		notEqualMatrix := r.Matrix.Rows(notEqualIndexes)
		notEqualSortK := floats.ArgSort(notEqualMatrix.Data, _sort(R), false)[:k]
		notEqualMatrix = notEqualMatrix.Rows(notEqualSortK)

		for i := 0; i < r.Dim(); i++ {
			diffH, diffM := 0.0, 0.0
			if r.AttributeTypes[i] == AttributeTypeDiscrete {
				for _, v := range equalMatrix.Col(i) {
					diffH += discreteDiff(R[i], v)
				}
				for _, v := range notEqualMatrix.Col(i) {
					diffM += discreteDiff(R[i], v)
				}
			} else {
				min, max := r.minMaxes[i][0], r.minMaxes[i][1]
				for _, v := range equalMatrix.Col(i) {
					diffH += continuousDiff(R[i], v, min, max)
				}
				for _, v := range notEqualMatrix.Col(i) {
					diffM += continuousDiff(R[i], v, min, max)
				}
			}
			w[i] = w[i] - diffH/float64(k)/float64(m) + diffM/float64(k)/float64(m)
		}
	}
	return w
}

func (r *Relief) Run(m int) []float64 {
	return r.runK(m, 1)
}
