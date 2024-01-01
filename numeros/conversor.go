package numeros

import "math"

func ConverteInt64ToInt8(value int64) int8 {
	if value > math.MaxInt8 {
		return math.MaxInt8
	} else if value < math.MinInt8 {
		return math.MinInt8
	}
	return int8(value)
}
