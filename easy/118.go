//Given an integer numRows, return the first numRows of Pascal's triangle.

package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, temp)
	}
	return res
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	res := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 1; j < i; j++ {
			temp[j] = res[j-1] + res[j]
		}
		res = temp
	}
	return res
}
