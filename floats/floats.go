package floats

import (
	"fmt"
	"github.com/a2htray/relief/ints"
	"math"
)

func SliceDeepCopy(s []float64) []float64 {
	ret := make([]float64, 0)
	for _, v := range s {
		ret = append(ret, v)
	}
	return ret
}

type Matrix struct {
	Data [][]float64
	M, N int
}

func NewMatrix(data [][]float64) *Matrix {
	m := len(data)
	if m == 0 {
		panic("matrix m could not be 0")
	}
	n := len(data[0])
	return &Matrix{
		Data: data,
		M: m,
		N: n,
	}
}

func (m *Matrix) Row(index int) []float64 {
	return SliceDeepCopy(m.Data[index])
}

func (m *Matrix) Col(index int) []float64 {
	ret := make([]float64, 0)
	for _, row := range m.Data {
		ret = append(ret, row[index])
	}
	return ret
}

func (m *Matrix) Rows(rowIndexes []int) *Matrix {
	data := make([][]float64, 0)
	for _, i := range rowIndexes {
		if i >= m.M {
			panic(fmt.Sprintf("row index %d must less then %d", i, m.M))
		}
		rowCopy := SliceDeepCopy(m.Data[i])
		data = append(data, rowCopy)
	}
	return NewMatrix(data)
}

func (m *Matrix) Cols(colIndexes []int) *Matrix {
	data := make([][]float64, 0)
	for _, j := range colIndexes {
		col := make([]float64, 0)
		for _, row := range m.Data {
			col = append(col, row[j])
		}
	}
	return NewMatrix(data)
}

// 最小值与最大值对
type MinMax [2]float64

func NewMinMax(min, max float64) MinMax {
	return MinMax{min, max}
}

func (m MinMax) Min() float64 {
	return m[0]
}

func (m MinMax) SetMin(min float64) {
	m[0] = min
}

func (m MinMax) Max() float64 {
	return m[1]
}

func (m MinMax) SetMax(max float64) {
	m[1] = max
}

type sliceWithIndex struct {
	index int
	data []float64
}

// 计算欧式距离
func EuclideanDistance(s, t []float64) float64 {
	if len(s) != len(t) {
		panic("floats: slices lengths do not match")
	}

	norm := 0.0
	for i, v := range s {
		norm += math.Pow(v-t[i], 2)
	}
	return math.Sqrt(norm)
}

// 计算序列中的最大值
func Max(s []float64) float64 {
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// 计算序列中的最小值
func Min(s []float64) float64 {
	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// 返回不同值的下标 slice
func NotEqualIndexes(s []float64, value float64) []int {
	return FloatFilterKey(s, func(v float64) bool {
		return value != v
	})
}

// 返回相同值的下标 slice
func EqualIndexes(s []float64, value float64) []int {
	return FloatFilterKey(s, func(v float64) bool {
		return value == v
	})
}

// 返回符合条件的元素下标 slice
func FloatFilterKey(s []float64, filter func(v float64) bool) []int {
	ret := make([]int, 0)
	for i, v := range s {
		if filter(v) {
			ret = append(ret, i)
		}
	}
	return ret
}

func ArgSort(matrix [][]float64, sort func([]float64) float64, descent bool) []int {
	s := make([]sliceWithIndex, 0)
	for i, row := range matrix {
		s = append(s, sliceWithIndex{
			data: row,
			index: i,
		})
	}
	ret := make([]int, 0)
	length := len(s)

	for i := 0; i < length; i++ {
		metric := sort(s[i].data)
		for j := i+1; j < length; j++ {
			newMetric :=  sort(s[j].data)
			if newMetric < metric {
				s[i], s[j] = s[j], s[i]
				metric = newMetric
			}
		}
	}
	for _, v := range s {
		ret = append(ret, v.index)
	}

	if descent {
		ints.Reverse(ret)
	}

	return ret
}