package internal

func IsPositiveNumber(n int) bool {
	return n > 0
}

func ValidateRange(n, min, max int) bool {
	return n >= min && n <= max
}
