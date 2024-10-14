/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.



Constraints:

    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106


*/

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merged []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}
	for j < len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}
	if len(merged)%2 == 0 {
		return float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2
	} else {
		return float64(merged[len(merged)/2])
	}
}
