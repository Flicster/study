package leetcode

import (
	"slices"
)

func TwoSum(numbers []int, target int) []int {
	for x := 0; x < len(numbers); x++ {
		for i := 0; i < len(numbers); i++ {
			if i == x {
				continue
			}
			if target == numbers[x]+numbers[i] {
				return []int{x + 1, i + 1}
			}
		}
	}
	return []int{}
}

func ThreeSum(nums []int) [][]int {
	return nil
}

func Rob(nums []int) int {
	top := 0
	res := make([]int, len(nums))
	for k, n := range nums {
		res[k] = n
		if k-1 >= 0 && res[k-1] > n {
			res[k] = res[k-1]
		}
		if k-2 >= 0 && res[k-2]+n > res[k] {
			res[k] = res[k-2] + n
		}

		if res[k] > top {
			top = res[k]
		}
	}
	return top
}

func DeleteAndEarn(nums []int) int {
	top := 0
	res := make([]int, len(nums))
	slices.Sort(nums)
	for k, n := range nums {
		res[k] = n

		for x := k - 1; x >= 0; x-- {
			if n-1 == nums[x] || n+1 == nums[x] {
				continue
			}
			if res[k] < res[x]+n {
				res[k] = res[x] + n
			}
		}
		if res[k] > top {
			top = res[k]
		}
	}
	return top
}

func UniquePaths(m int, n int) int {
	top := 0
	res := make([][]int, m)
	for x := 0; x < m; x++ {
		res[x] = make([]int, n)
		for i := 0; i < n; i++ {
			ix := 0
			if x > 0 {
				ix = x - 1
			}
			ii := 0
			if i > 0 {
				ii = i - 1
			}
			if res[ix][i] == 0 && res[x][ii] == 0 {
				res[x][i] = 1
			} else {
				res[x][i] = res[ix][i] + res[x][ii]
			}
			if res[x][i] > top {
				top = res[x][i]
			}
		}
	}
	return top
}

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	res := make([][]int, len(obstacleGrid))
	for x, ogi := range obstacleGrid {
		res[x] = make([]int, len(ogi))
		for i, gridItem := range ogi {
			if gridItem == 1 {
				res[x][i] = 0
				continue
			}
			if x == 0 && i == 0 {
				res[x][i] = 1
				continue
			}
			left := 0
			if x > 0 {
				left = res[x-1][i]
			}
			up := 0
			if i > 0 {
				up = res[x][i-1]
			}
			res[x][i] = left + up
		}
	}
	last := res[len(res)-1]

	return last[len(last)-1]
}

func MinPathSum(grid [][]int) int {
	res := make([][]int, len(grid))
	for x, row := range grid {
		res[x] = make([]int, len(row))
		for i, col := range row {
			if x-1 < 0 && i-1 < 0 {
				res[x][i] = col
				continue
			}
			if x-1 < 0 {
				res[x][i] = col + res[x][i-1]
				continue
			}
			if i-1 < 0 {
				res[x][i] = col + res[x-1][i]
				continue
			}
			if col+res[x][i-1] < col+res[x-1][i] {
				res[x][i] = col + res[x][i-1]
			} else {
				res[x][i] = col + res[x-1][i]
			}
		}
	}
	last := res[len(res)-1]

	return last[len(last)-1]
}

func MinimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	res := make([][]int, len(triangle))
	for rk, r := range triangle {
		res[rk] = make([]int, len(r))
		for ck, c := range r {
			var m *int
			if rk > 0 && len(triangle[rk-1]) > ck {
				m = &res[rk-1][ck]
			}
			if rk > 0 && ck > 0 && (m == nil || res[rk-1][ck-1] < *m) {
				m = &res[rk-1][ck-1]
			}

			res[rk][ck] = c
			if m != nil {
				res[rk][ck] += *m
			}
		}
	}
	path := 0
	for k, p := range res[len(triangle)-1] {
		if k == 0 {
			path = p
			continue
		}
		if p < path {
			path = p
		}
	}
	return path
}

func MinFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	res := make([][]int, len(matrix))
	for rk, r := range matrix {
		res[rk] = make([]int, len(r))
		for ck, c := range r {
			if rk == 0 {
				res[rk][ck] = c
				continue
			}
			m := res[rk-1][ck]
			if ck-1 >= 0 && res[rk-1][ck-1] < m {
				m = res[rk-1][ck-1]
			}
			if ck+1 < len(res[rk-1]) && res[rk-1][ck+1] < m {
				m = res[rk-1][ck+1]
			}
			res[rk][ck] = c + m
		}
	}
	path := 0
	for k, p := range res[len(matrix)-1] {
		if k == 0 {
			path = p
			continue
		}
		if p < path {
			path = p
		}
	}
	return path
}

func MaximalSquare(matrix [][]byte) int {
	best := 0
	res := make([][]int, len(matrix))
	for rk, r := range matrix {
		res[rk] = make([]int, len(r))
		for ck, c := range r {
			if c == '0' {
				continue
			}
			left := 0
			if ck > 0 {
				left = res[rk][ck-1]
			}
			diag := 0
			if rk > 0 && ck > 0 {
				diag = res[rk-1][ck-1]
			}
			up := 0
			if rk > 0 {
				up = res[rk-1][ck]
			}
			res[rk][ck] = min(left, diag, up) + 1
			if res[rk][ck] > best {
				best = res[rk][ck]
			}
		}
	}
	return best * best
}

func LongestPalindrome(s string) string {
	longest := ""
	for ck := range s {
		for x := len(s) - 1; x >= 0; x-- {
			if x+1 <= ck {
				break
			}
			sub := s[ck : x+1]
			if IsPalindrome(sub) && len(sub) > len(longest) {
				if len(sub) == len(s) {
					return s
				}
				longest = sub
				break
			}
		}
	}
	return longest
}

func SetZeroes(matrix [][]int) {
	//for rk, r := range matrix {
	//	for ck, c := range r {
	//		if c != 0 {
	//			continue
	//		}
	//		for x := 0; x < len()
	//	}
	//}
}
