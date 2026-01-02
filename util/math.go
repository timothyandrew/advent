package util

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Abs(n int) int {
	return max(n, -n)
}
