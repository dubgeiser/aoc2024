package algo

// This bit me sooo hard :-(
// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
func Mod(a, b int) int {
	return (a%b + b) % b
}
