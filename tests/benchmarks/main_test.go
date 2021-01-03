package benchmarks

import "testing"

func BenchmarkSort(b *testing.B) {
	arr := make([]int, 100000)
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubbleSort(arr)
	}
}

func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := length - 1; j >= i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}
