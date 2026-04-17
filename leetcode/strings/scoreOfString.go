package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(scoreOfString(s))
}
func scoreOfString(s string) int {
	var result int
	for i := 0; i < len(s)-1; i++ {
		res := int(s[i]) - int(s[i+1])
		if res < 0 {
			result += res * -1
		} else {
			result += int(s[i]) - int(s[i+1])
		}
	}
	return result
}
