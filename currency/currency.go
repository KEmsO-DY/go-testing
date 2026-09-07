package currency

// Converter отвечает за пересчёт сумм между валютами.
type Converter interface {
	Convert(amount float64, from, to string) float64
}

// PriceIn пересчитывает цену через переданный Converter.
func PriceIn(amount float64, from, to string, c Converter) float64 {
	return c.Convert(amount, from, to)
}
