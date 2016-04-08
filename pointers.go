package main

import "fmt"

func main() {
	x, y := 20, 40
	swap(&x, &y)

	fmt.Println(x)
	fmt.Println(y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
