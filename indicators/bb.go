package indicators

import (
	"math"
)

// BollingerBands представляет индикатор Bollinger Bands.
type BollingerBands struct {
	period       int       // Период BB
	stdDevFactor float64   // Коэффициент стандартного отклонения
	prices       []float64 // Цены за определенный период
}

// NewBollingerBands создает новый экземпляр Bollinger Bands с заданным периодом и коэффициентом стандартного отклонения.
func NewBollingerBands(period int, stdDevFactor float64) *BollingerBands {
	return &BollingerBands{
		period:       period,
		stdDevFactor: stdDevFactor,
		prices:       make([]float64, 0),
	}
}

// Update добавляет новую цену в список и обновляет Bollinger Bands.
func (bb *BollingerBands) Update(price float64) {
	bb.prices = append(bb.prices, price)
	if len(bb.prices) > bb.period {
		bb.prices = bb.prices[1:]
	}
}

// Calculate возвращает верхнюю, среднюю и нижнюю полосы Bollinger Bands.
func (bb *BollingerBands) Calculate() (upperBand, middleBand, lowerBand float64) {
	if len(bb.prices) < bb.period {
		return 0, 0, 0
	}

	// Вычисляем среднее значение
	sum := 0.0
	for _, price := range bb.prices {
		sum += price
	}
	middleBand = sum / float64(bb.period)

	// Вычисляем стандартное отклонение
	variance := 0.0
	for _, price := range bb.prices {
		variance += math.Pow(price-middleBand, 2)
	}
	stdDev := math.Sqrt(variance / float64(bb.period))

	// Вычисляем верхнюю и нижнюю полосы
	upperBand = middleBand + (bb.stdDevFactor * stdDev)
	lowerBand = middleBand - (bb.stdDevFactor * stdDev)

	return upperBand, middleBand, lowerBand
}

// sample
//func main() {
//	bb := NewBollingerBands(20, 2)
//
//	// Пример данных (цен)
//	prices := []float64{44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64}
//
//	// Обновляем Bollinger Bands для каждой цены
//	for _, price := range prices {
//		bb.Update(price)
//		upper, middle, lower := bb.Calculate()
//		fmt.Printf("Upper Band: %.2f, Middle Band: %.2f, Lower Band: %.2f\n", upper, middle, lower)
//	}
//}
