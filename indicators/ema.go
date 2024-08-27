package indicators

import "fmt"

// EMA представляет индикатор экспоненциальной скользящей средней.
type EMA struct {
	period      int       // Период EMA
	prices      []float64 // Цены за определенный период
	emaValue    float64   // Значение EMA
	alpha       float64   // Коэффициент сглаживания
	initialized bool      // Флаг инициализации
}

// NewEMA создает новый экземпляр EMA с заданным периодом.
func NewEMA(period int) *EMA {
	return &EMA{
		period:      period,
		prices:      make([]float64, 0),
		emaValue:    0,
		alpha:       0,
		initialized: false,
	}
}

// Update добавляет новую цену в список и обновляет EMA.
func (ema *EMA) Update(price float64) {
	if !ema.initialized {
		ema.prices = append(ema.prices, price)
		if len(ema.prices) == ema.period {
			ema.initialized = true
			ema.calculateAlpha()
			ema.calculateInitialEMA()
		}
	} else {
		ema.calculateEMA(price)
	}
}

// calculateAlpha вычисляет коэффициент сглаживания alpha для EMA.
func (ema *EMA) calculateAlpha() {
	ema.alpha = 2.0 / float64(ema.period+1)
}

// calculateInitialEMA вычисляет начальное значение EMA на основе первых N цен.
func (ema *EMA) calculateInitialEMA() {
	sum := 0.0
	for _, price := range ema.prices {
		sum += price
	}
	ema.emaValue = sum / float64(ema.period)
	fmt.Printf("Initial EMA: %.2f\n", ema.emaValue)
}

// calculateEMA вычисляет текущее значение EMA на основе новой цены.
func (ema *EMA) calculateEMA(price float64) {
	ema.emaValue = (price * ema.alpha) + (ema.emaValue * (1 - ema.alpha))
	fmt.Printf("EMA: %.2f\n", ema.emaValue)
}

// sample
//func main() {
//	ema := NewEMA(14)
//
//	// Пример данных (цен)
//	prices := []float64{44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64}
//
//	// Обновляем EMA для каждой цены
//	for _, price := range prices {
//		ema.Update(price)
//	}
//}
