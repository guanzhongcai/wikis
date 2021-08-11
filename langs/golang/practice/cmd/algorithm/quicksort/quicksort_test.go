package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{7, 3, 2, 6, 8, 1, 9, 5, 4, 10}
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}

// output
// 第一次会变成[4 3 2 5 1 6 8 7 9] pivot是6
