package utils

// Strdef returns non nil value if available
func Strdef(a, b string) string {
	if len(a) > 0 {
		return a
	}
	return b
}

// Uintdef returns non nil value if available
func Uintdef(a, b uint) uint {
	if a > 0 {
		return a
	}
	return b
}
