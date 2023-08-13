package main

import "fmt"

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

// 双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i < 0 {
			nums1[k] = nums2[j]
			k--
			j--
			continue
		}
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			k--
			j--
		} else {
			nums1[k] = nums1[i]
			i--
			k--
		}
	}
}

//插入排序
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	for i := m; i < m+n; i++ {
// 		nums1[i] = nums2[i-m]
// 	}
// 	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
// }
