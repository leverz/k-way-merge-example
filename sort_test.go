package sort

import (
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

func TestSort(t *testing.T) {
	sourceArray := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	array := Sort(sourceArray)
	targetArray := []int{1, 4, 9, 11, 12, 32, 39, 42, 53, 55, 64, 70, 77,  87, 94}
	if !checkEqual(array, targetArray) {
		t.Errorf("wrong case: %v, expect %v, but got %v", sourceArray, array, targetArray)
	}
}
