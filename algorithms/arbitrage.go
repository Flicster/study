package algorithms

import "fmt"

// Функция для проверки возможности арбитража между тремя валютами
func findArbitrageOpportunity(graph [][]float64) {
	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				if graph[i][j] < graph[i][k]*graph[k][j] {
					// Если найдено улучшение, выводим возможность арбитража
					fmt.Printf("Найдена возможность арбитража: %d -> %d -> %d -> %d\n", i, k, j, i)
				}
			}
		}
	}
}

func main() {
	// Пример графа с курсами валютных пар
	graph := [][]float64{
		{1, 1.5, 0.8},   // USD -> USD, USD -> EUR, USD -> GBP
		{0.67, 1, 0.55}, // EUR -> USD, EUR -> EUR, EUR -> GBP
		{1.25, 1.82, 1}, // GBP -> USD, GBP -> EUR, GBP -> GBP
	}

	// Поиск возможностей арбитража
	findArbitrageOpportunity(graph)
}
