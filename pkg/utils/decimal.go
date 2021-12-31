package utils

import "github.com/shopspring/decimal"

func DecimalRoundAdd(a, b float64) float64 {
	price, _ := decimal.NewFromFloat(a).
		Add(decimal.NewFromFloat(b)).Round(3).Float64()
	return price
}
