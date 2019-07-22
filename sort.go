package kwaymerge

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	i, j, lLen, rLen := 0, 0, len(left), len(right)
	for lLen > i && rLen > j {
		if left[i] >= right[rLen - 1] {
			result = append(result, right[j:]...)
			result = append(result, left[i:]...)
			return result
		}
		if right[j] >= left[lLen - 1] {
			result = append(result, left[i:]...)
			result = append(result, right[j:]...)
			return result
		}

		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func Sort(list []int) []int {
	listLength := len(list)
	if listLength < 2 {
		return list
	}

	mid := listLength / 2
	left := list[:mid]
	right := list[mid:]

	return merge(Sort(left), Sort(right))
}

func TwoSort(list []int) []int {
	listLength := len(list)
	if listLength < 2 {
		return list
	}

	mid := listLength / 2
	left := list[:mid]
	right := list[mid:]

	sortChan := make(chan []int, 2)
	defer close(sortChan)
	go func() {
		sortChan <- Sort(left)
	}()
	go func() {
		sortChan <- Sort(right)
	}()
	return merge(<- sortChan, <-sortChan)
}
