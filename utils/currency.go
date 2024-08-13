package utils

const (
	USD = "USD"
	INR = "INR"
	GBP = "GBP"
)

// Exported function with uppercase 'I'
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, GBP:
		return true
	}
	return false
}
