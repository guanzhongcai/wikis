package quicksort

import "fmt"

// https://www.bilibili.com/video/BV1S5411P7ky?p=21&spm_id_from=pageDriver

func QuickSort00(arr [] int) (sorts []int) {

	// TODO 递归

	size := len(arr)
	middle := size / 2
	pivot := arr[middle]
	var smalls, bigs []int
	for i := 0; i < size; i++ {
		if arr[i] > pivot {
			bigs = append(bigs, arr[i])
		} else {
			smalls = append(smalls, arr[i])
		}
	}
	fmt.Printf("pivot=%v, smalls=%v, bigs=%v, arr=%v", pivot, smalls, bigs, arr)
	return arr
}

// 数组分区
// leftBound是数组左边的边界索引
// rightBound是数组右边的边界索引
func partition(arr []int, leftBound, rightBound int) (posPivot int) {
	// 以右边的做轴
	pivot := arr[rightBound]
	// 左边和右边的指针
	left := leftBound
	right := rightBound - 1

	// 以pivot为轴，小的放左边，大的放右边，两边同时找，找到不满足排序规则的索引位置后，交换位置上的值，放入轴的左边和右边
	for ; left <= right; {
		// left < right：如果最右边就是最小值，此处会出现数组越界bug
		for ; left <= right && arr[left] <= pivot; {
			left++
		}
		for ; left <= right && arr[right] > pivot; {
			right--
		}

		fmt.Printf("left=%d, right=%d\n", left, right)
		if left < right {
			swap(arr, left, right)
		}
		fmt.Printf(">> %v\n", arr)
	}

	// 把pivot轴换到合适的中间的分割位置：左边比其小，右边比其大
	swap(arr, left, rightBound)
	posPivot = left
	return
}

func QuickSort(arr []int, leftBound, rightBound int) {
	if leftBound >= rightBound {
		return
	}
	mid := partition(arr, leftBound, rightBound)
	// 以轴pivot为届，分为两个区以后，再分别对两个区递归的进行排序
	QuickSort(arr, leftBound, mid-1)
	QuickSort(arr, mid+1, rightBound)
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
