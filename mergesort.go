package main

func MergeSort(data []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(data, p, q)
	MergeSort(data, q+1, r)

	Merge(data, p, q, r)
}

func Merge(data []int, p, q, r int) {
	i, j := 0, 0
	b := make([]int, q+1-p)
	c := make([]int, r-q)
	copy(b, data[p:q+1])
	copy(c, data[q+1:r+1])

	for k := p; k <= r; k++ {
		if j == r-q {
			data[k] = b[i]
			i++
			continue
		}
		if i == q+1-p {
			data[k] = c[j]
			j++
			continue
		}
		if b[i] <= c[j] {
			data[k] = b[i]
			i++
		} else {
			data[k] = c[j]
			j++
		}
	}
}
