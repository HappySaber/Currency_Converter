package converter

func Converter(val1, val2, quantity float64) float64 {
	res := val1 * quantity
	return res / val2
}
