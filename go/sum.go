package go_test

import "fmt"

func main() {
	fmt.Println(Sum(2, 2))
	fmt.Println(Subtract(2, 2))
}

func Sum(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}
