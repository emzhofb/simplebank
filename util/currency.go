package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currencry string) bool {
	switch currencry {
	case USD, EUR, CAD:
		return true
	}

	return false
}
