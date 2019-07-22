package kwaymerge

import (
	"math/rand"
	"runtime"
	"sort"
	"testing"
)

func checkEqual(array []int, target []int) bool {
	if len(array) != len(target) {
		return false
	}
	for i := 0; i < len(array); i++ {
		if array[i] != target[i] {
			return false
		}
	}

	return true
}

func generateRandomArray(n int) []int {
	r := make([]int, n)
	rand.Seed(int64(n))
	for i := 0; i < n; i++ {
		r[i] = rand.Int()
	}

	return r
}

func expectSortedResult(source []int) []int {
	sort.Ints(source)
	return source
}

func TestSort(t *testing.T) {
	sourceArray := generateRandomArray(100)
	array := Sort(sourceArray)
	targetArray := expectSortedResult(sourceArray)
	if !checkEqual(array, targetArray) {
		t.Errorf("wrong case: %v, expect %v, but got %v", sourceArray, targetArray, array)
	}

	array = TwoSort(sourceArray)
	if !checkEqual(array, targetArray) {
		t.Errorf("wrong two way sort case: %v, expect %v, but got %v", sourceArray, targetArray, array)
	}
}

// 数据量在 200000 左右开始产生差异
func BenchmarkSort(b *testing.B) {
	sourceArray := generateRandomArray(1000000)
	b.StartTimer()
	Sort(sourceArray)
	b.StopTimer()
}

// 这种方式比较依赖 CPU 核心数，如果空闲 CPU 不足 2 个，就会等待，导致性能较差，大部分时间性能都比 Sort 快
func BenchmarkTwoSort(b *testing.B) {
	sourceArray := generateRandomArray(1000000)
	b.StartTimer()
	TwoSort(sourceArray)
	b.StopTimer()
}

func BenchmarkTwoSortLimitCPU(b *testing.B) {
	runtime.GOMAXPROCS(1)
	sourceArray := generateRandomArray(1000000)
	b.StartTimer()
	TwoSort(sourceArray)
	b.StopTimer()
}
