package leetcode

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
