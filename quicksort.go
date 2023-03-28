package main

func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	pivot := (len(data) / 2) - 1
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, d := range data {
		if d < data[pivot] {
			less = append(less, d)
		} else if d > data[pivot] {
			greater = append(greater, d)
		}
	}
	return append(append(QuickSort(less), data[pivot]), QuickSort(greater)...)
}
