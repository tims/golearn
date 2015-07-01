package utils

func Reverse(s string) string {
	reversed := make([]rune, len(s))
	for i, r := range s {
		reversed[len(s) - i - 1] = r
	}
	return string(reversed)
}