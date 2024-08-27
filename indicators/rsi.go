package indicators

import (
	"errors"
	"fmt"
)

// RSI struct представляет индикатор RSI.
type RSI struct {
	period      int       // Период RSI
	prices      []float64 // Цены за определенный период
	gainTotal   float64   // Общий прирост цен за период
	lossTotal   float64   // Общий убыток цен за период
	lastPrice   float64   // Последняя цена
	initialized bool      // Флаг инициализации
}

// NewRSI создает новый экземпляр RSI с заданным периодом.
func NewRSI(period int) *RSI {
	return &RSI{
		period:      period,
		prices:      make([]float64, 0),
		gainTotal:   0,
		lossTotal:   0,
		lastPrice:   0,
		initialized: false,
	}
}

// Update добавляет новую цену в список и обновляет RSI.
func (rsi *RSI) Update(price float64) error {
	if len(rsi.prices) == rsi.period {
		// Удаляем первую цену из списка
		removedPrice := rsi.prices[0]
		rsi.prices = rsi.prices[1:]

		// Обновляем общий прирост и убыток
		diff := price - removedPrice
		if diff > 0 {
			rsi.gainTotal -= diff
		} else {
			rsi.lossTotal -= diff
		}
	} else if len(rsi.prices) > rsi.period {
		return errors.New("RSI: invalid state")
	}

	// Добавляем новую цену в список
	rsi.prices = append(rsi.prices, price)

	// Обновляем общий прирост и убыток
	if rsi.lastPrice != 0 {
		diff := price - rsi.lastPrice
		if diff > 0 {
			rsi.gainTotal += diff
		} else {
			rsi.lossTotal -= diff
		}
	}

	rsi.lastPrice = price

	// Проверяем, инициализирован ли RSI
	if len(rsi.prices) < rsi.period {
		return nil
	}

	// Вычисляем RSI
	avgGain := rsi.gainTotal / float64(rsi.period)
	avgLoss := rsi.lossTotal / float64(rsi.period)

	rs := 100 - (100 / (1 + avgGain/avgLoss))

	fmt.Printf("RSI: %.2f\n", rs)

	return nil
}

// sample
//func main() {
//	rsi := NewRSI(14)
//
//	// Пример данных (цен)
//	prices := []float64{44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64}
//
//	// Обновляем RSI для каждой цены
//	for _, price := range prices {
//		rsi.Update(price)
//	}
//}
