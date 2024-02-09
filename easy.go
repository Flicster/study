package main

import (
	"fmt"
	"strings"
	"unicode"
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

func isPalindrome(s string) bool {
	str := make([]string, 0, len(s))
	for _, c := range s {
		char := fmt.Sprintf("%c", c)
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			str = append(str, strings.ToLower(char))
		}
	}
	forward := ""
	backward := ""
	for x := 0; x < len(str); x++ {
		forward += str[x]
	}
	for x := len(str) - 1; x >= 0; x-- {
		backward += str[x]
	}
	return forward == backward
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
