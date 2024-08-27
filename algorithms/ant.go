package algorithms

import (
	"math"
)

// TSPSolver представляет алгоритм обхода муравья для решения задачи коммивояжера.
type TSPSolver struct {
	addresses []string
	distances [][]float64
	times     [][]float64
}

// NewTSPSolver создает новый экземпляр TSPSolver.
func NewTSPSolver(addresses []string, distances [][]float64, times [][]float64) *TSPSolver {
	return &TSPSolver{
		addresses: addresses,
		distances: distances,
		times:     times,
	}
}

// AntColonyTSP решает задачу коммивояжера с использованием алгоритма обхода муравья.
func (t *TSPSolver) AntColonyTSP() ([]string, float64, float64) {
	bestPath := make([]string, len(t.addresses))
	minDistance := math.Inf(1)
	minTime := math.Inf(1)

	// Перебираем все возможные перестановки адресов
	for perm := range rangePermutations(t.addresses) {
		distance, time := t.calculateDistanceAndTime(perm)
		if distance < minDistance {
			minDistance = distance
			minTime = time
			copy(bestPath, perm)
		}
	}

	return bestPath, minDistance, minTime
}

// rangePermutations генерирует все возможные перестановки элементов в слайсе.
func rangePermutations(arr []string) <-chan []string {
	c := make(chan []string)
	go func() {
		defer close(c)
		permute(arr, c, 0)
	}()
	return c
}

// permute генерирует все возможные перестановки элементов в слайсе.
func permute(arr []string, c chan<- []string, i int) {
	if i == len(arr)-1 {
		tmp := make([]string, len(arr))
		copy(tmp, arr)
		c <- tmp
		return
	}
	for j := i; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		permute(arr, c, i+1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// calculateDistanceAndTime вычисляет расстояние и время для данного маршрута.
func (t *TSPSolver) calculateDistanceAndTime(path []string) (float64, float64) {
	distance := 0.0
	time := 0.0
	for i := 0; i < len(path)-1; i++ {
		from := t.findIndex(path[i])
		to := t.findIndex(path[i+1])
		distance += t.distances[from][to]
		time += t.times[from][to]
	}
	return distance, time
}

// findIndex находит индекс элемента в списке адресов.
func (t *TSPSolver) findIndex(addr string) int {
	for i, a := range t.addresses {
		if a == addr {
			return i
		}
	}
	return -1
}

//func main() {
//	addresses := []string{"Address 1", "Address 2", "Address 3", "Address 4"}
//	distances := [][]float64{
//		{0, 10, 15, 20},
//		{10, 0, 35, 25},
//		{15, 35, 0, 30},
//		{20, 25, 30, 0},
//	}
//	times := [][]float64{
//		{0, 5, 10, 15},
//		{5, 0, 20, 25},
//		{10, 20, 0, 30},
//		{15, 25, 30, 0},
//	}
//
//	solver := NewTSPSolver(addresses, distances, times)
//	bestPath, minDistance, minTime := solver.AntColonyTSP()
//
//	fmt.Println("Optimal Path:", bestPath)
//	fmt.Println("Minimum Distance:", minDistance)
//	fmt.Println("Minimum Time:", minTime)
//}
