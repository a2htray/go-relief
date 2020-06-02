package ints

import (
	"math/rand"
	"time"
)

// 打乱原 slice
func Shuffle(s []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// 在 int slice 中删除指定值，并返回
func Remove(s []int, value int) []int {
	return Filter(s, func(v int) bool {
		return value != v
	})
}

// Int slice 过滤操作
func Filter(s []int, filter func(v int) bool) []int {
	ret := make([]int, 0)
	for _, v := range s {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Iter(10) => [0,1,2,3,4,5,6,7,8,9]
func Iter(n int) []int {
	return Range(0, n - 1, 1)
}

// Python range 的 go 实现
func Range(start, end, step int) []int {
	ret := make([]int, 0)
	for {
		if start > end {
			break
		}
		ret = append(ret, start)
		start += step
	}
	return ret
}

// 整型序列逆序
func Reverse(s []int) {
	length := len(s)
	i := 0;
	j := length-1
	for ; i <= j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}