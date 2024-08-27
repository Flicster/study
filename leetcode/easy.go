package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if target <= nums[0] {
			return 0
		} else {
			return 1
		}
	}
	m := len(nums) / 2
	return searchInsert(nums[:m], target) + searchInsert(nums[m:], target)
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	mem := 1
	for x := len(digits) - 1; mem != 0; x-- {
		if digits[x] == 9 {
			if x == 0 {
				addDigit := []int{1}
				digits[x] = 0
				digits = append(addDigit, digits...)
				mem = 0
			} else {
				digits[x] = 0
				continue
			}
		}
		digits[x] += mem
		mem = 0
	}
	return digits
}

func rotateMatrix(matrix [][]int) {
	var initialMatrix [][]int
	initialMatrix = append(initialMatrix, matrix...)
	for x := 0; x < len(initialMatrix); x++ {
		row := make([]int, 0, len(initialMatrix))
		for i := len(initialMatrix) - 1; i >= 0; i-- {
			row = append(row, initialMatrix[i][x])
		}
		matrix[x] = row
	}
}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1 := 1
	n2 := 2
	for x := 3; x <= n; x++ {
		prevN1 := n1
		n1 = n2
		prevN2 := n2
		n2 = prevN1 + prevN2
	}
	return n2
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	var full int
	for i := 1; i <= x; i++ {
		q := i * i
		if q > x {
			break
		}
		full = i
	}
	diff := x - full*full
	if diff == 0 {
		return full
	}
	t := float64(diff) / float64(diff*full)
	fl := full + int(t)
	return fl
}

func isSubsequence(s string, t string) bool {
	q := make([]string, 0, len(s))
	for _, l := range s {
		q = append(q, fmt.Sprintf("%c", l))
	}
	if len(q) == 0 {
		return true
	}
	for _, l := range t {
		if len(q) == 0 {
			break
		}
		if fmt.Sprintf("%c", l) == q[0] {
			q = q[1:]
		}
	}
	return len(q) == 0
}

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	res := calcHappy(n)
	return res == 1 || res == 7
}

func calcHappy(n int) int {
	res := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		res += digit * digit
	}
	if res < 10 {
		return res
	}
	return calcHappy(res)
}

type ListNode struct {
	Prev *ListNode
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
		if fast == head {
			return true
		}
	}
	return false
}

func rotate(nums []int, k int) {
	delta := k
	if delta == len(nums) {
		return
	}
	if k > len(nums) {
		delta = k - (k / len(nums) * len(nums))
	}
	for x := 0; x < delta; x++ {
		last := nums[len(nums)-1]
		var old int
		newRota := true
		for i := range nums {
			if newRota {
				old = last
				newRota = false
			}
			o := old
			old, nums[i] = nums[i], o
		}
	}
}

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	profit := 0
	minPrice := prices[0]
	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
			continue
		}
		if p-minPrice > profit {
			profit = p - minPrice
		}
	}
	return profit
}

func majorityElement(nums []int) int {
	c := make(map[int]int, 0)
	for _, n := range nums {
		_, ok := c[n]
		if !ok {
			c[n] = 1
			continue
		}
		c[n] += 1
	}
	most := 0
	r := 0
	for n, count := range c {
		if count > most {
			most = count
			r = n
		}
	}
	return r
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	firstStack := make([]int, 0, m)
	secondStack := make([]int, 0, n)
	if m > 0 {
		for x := 0; x < m; x++ {
			firstStack = append(firstStack, nums1[x])
		}
	}
	if n > 0 {
		for x := 0; x < n; x++ {
			secondStack = append(secondStack, nums2[x])
		}
	}
	for k := range nums1 {
		first := 0
		if len(firstStack) > 0 {
			first = firstStack[0]
		}
		second := 0
		if len(secondStack) > 0 {
			second = secondStack[0]
		}
		if len(firstStack) == 0 {
			nums1[k] = second
			if len(secondStack) != 0 {
				secondStack = secondStack[1:]
			}
			continue
		}
		if len(secondStack) == 0 {
			nums1[k] = first
			if len(firstStack) != 0 {
				firstStack = firstStack[1:]
			}
			continue
		}
		if first >= second {
			nums1[k] = second
			if len(secondStack) != 0 {
				secondStack = secondStack[1:]
			}
		} else {
			nums1[k] = first
			if len(firstStack) != 0 {
				firstStack = firstStack[1:]
			}
			continue
		}
	}
}

func IsPalindrome(s string) bool {
	for x := 0; x < len(s)/2; x++ {
		first := fmt.Sprintf("%c", s[x])
		last := fmt.Sprintf("%c", s[len(s)-1-x])
		if last != first {
			return false
		}
	}
	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineCounter := make(map[string]int, len(magazine))
	for _, c := range magazine {
		magazineCounter[fmt.Sprintf("%c", c)] += 1
	}
	for _, c := range ransomNote {
		count, ok := magazineCounter[fmt.Sprintf("%c", c)]
		if !ok || count == 0 {
			return false
		}
		magazineCounter[fmt.Sprintf("%c", c)] -= 1
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	patternInts := make(map[string]int, len(pattern))
	for _, c := range pattern {
		_, ok := patternInts[fmt.Sprintf("%c", c)]
		if !ok {
			patternInts[fmt.Sprintf("%c", c)] = len(patternInts) + 1
		}
	}
	strSlice := strings.Split(s, " ")
	strInts := make(map[string]int, len(strSlice))
	for _, word := range strSlice {
		_, ok := strInts[word]
		if !ok {
			strInts[word] = len(strInts) + 1
		}
	}

	patternIntsStr := ""
	for _, c := range pattern {
		i := patternInts[fmt.Sprintf("%c", c)]
		if i != 0 {
			patternIntsStr += fmt.Sprintf("%d", i)
		}
	}
	strIntsStr := ""
	for _, word := range strSlice {
		i := strInts[word]
		if i != 0 {
			strIntsStr += fmt.Sprintf("%d", i)
		}
	}
	return patternIntsStr == strIntsStr
}

func isAnagram(s string, t string) bool {
	sCounter := make(map[string]int, len(s))
	for _, c := range s {
		sCounter[fmt.Sprintf("%c", c)] += 1
	}
	tCounter := make(map[string]int, len(t))
	for _, c := range t {
		tCounter[fmt.Sprintf("%c", c)] += 1
	}

	if len(sCounter) != len(tCounter) {
		return false
	}
	for sl, sc := range sCounter {
		tc, ok := tCounter[sl]
		if !ok {
			return false
		}
		if sc != tc {
			return false
		}
	}
	return true
}

func singleNumber(nums []int) int {
	res := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		_, ok := res[n]
		if ok {
			delete(res, n)
		} else {
			res[n] = struct{}{}
		}
	}
	for k := range res {
		return k
	}
	return 0
}

func isIsomorphic(s string, t string) bool {
	lnln := make(map[string]int, len(s))
	r := 0
	ss := ""
	for _, l := range s {
		ln, ok := lnln[fmt.Sprintf("%c", l)]
		if !ok {
			r++
			lnln[fmt.Sprintf("%c", l)] = r
			ln = r
		}
		ss += fmt.Sprintf("%d", ln)
	}

	lnln = make(map[string]int, len(t))
	r = 0
	ts := ""
	for _, l := range t {
		ln, ok := lnln[fmt.Sprintf("%c", l)]
		if !ok {
			r++
			lnln[fmt.Sprintf("%c", l)] = r
			ln = r
		}
		ts += fmt.Sprintf("%d", ln)
	}
	return ss == ts
}

func Tribonacci(n int) int {
	q := make([]int, 0)
	for x := 0; x <= n; x++ {
		if x == 0 {
			q = append(q, 0)
			continue
		}
		if x == 1 || x == 2 {
			q = append(q, 1)
			continue
		}
		last := 0
		for _, v := range q {
			last += v
		}
		q = q[1:]
		q = append(q, last)
	}
	return q[len(q)-1]
}

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	last := 0
	first, second := 1, 1
	for x := 1; x < n; x++ {
		last = first + second
		first = second
		second = last
	}
	return last
}

func MinCostClimbingStairs(cost []int) int {
	res := make([]int, len(cost))
	for k, c := range cost {
		if k == 0 {
			res[k] = c
			continue
		}
		if k == 1 {
			res[k] = c
			continue
		}
		first := res[k-2] + c
		second := res[k-1] + c
		res[k] = first
		if res[k] > second {
			res[k] = second
		}
	}
	r := res[len(res)-1]
	if r > res[len(res)-2] {
		r = res[len(res)-2]
	}

	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

//func HasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//	if root.Right == nil && root.Left == nil {
//		return false
//	}
//	curr := 0
//	sums := []int{}
//	for root.Right != nil && root.Left != nil {
//
//	}
//}

func addBinary(a string, b string) string {
	diff := len(a) - len(b)
	if diff < 0 {
		diff *= -1
	}
	for i := 0; i < diff; i++ {
		if len(a) < len(b) {
			a = "0" + a
			continue
		}
		if len(a) > len(b) {
			b = "0" + b
			continue
		}
	}
	mem := 0
	res := ""
	for i := len(a) - 1; i >= 0; i-- {
		up, _ := strconv.Atoi(string(a[i]))
		down, _ := strconv.Atoi(string(b[i]))
		switch {
		case up+down+mem == 0:
			res = "0" + res
			mem = 0
		case up+down+mem == 1:
			res = "1" + res
			mem = 0
		case up+down+mem == 2:
			res = "0" + res
			mem = 1
		case up+down+mem == 3:
			res = "1" + res
			mem = 1
		}
		if i == 0 && mem != 0 {
			res = "1" + res
		}
	}
	return res
}

func deleteDuplicates(head *ListNode) *ListNode {
	for node := head; node != nil; node = node.Next {
		for node.Next != nil && node.Next.Val == node.Val {
			node.Next = node.Next.Next
		}
	}
	return head
}

func Walk(t *ListNode, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Val
	if t.Prev != nil {
		Walk(t.Prev, ch)
	}
	if t.Next != nil {
		Walk(t.Next, ch)
	}
}

func encode(s string) string {
	res := ""
	counter := 1
	for k, l := range s {
		if k == 0 {
			res += fmt.Sprintf("%c", l)
			continue
		}
		prevL := fmt.Sprintf("%c", s[k-1])
		currL := fmt.Sprintf("%c", l)
		if prevL == currL {
			counter += 1
		} else {
			if counter != 1 {
				res += fmt.Sprintf("%d", counter)
				counter = 1
			}
			res += currL
		}
	}
	return res
}
