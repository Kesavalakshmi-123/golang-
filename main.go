package main

import (
	"fmt"
)

func Hi(a int, b int) int {
	c := a + b
	return c
}

func main() {
	num := Hi(10, 20)
	fmt.Println(num)
}
