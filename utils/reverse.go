package utils

import "fmt"

func Reverse(s string) string {
	for _, r := range s {
		fmt.Println(r);
	}
	return ""
}