package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("x: %v, y: %v\n", x, y)
}
