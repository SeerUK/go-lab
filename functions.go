package main

import "log"

func add(i int) func(int) int {
	return func(j int) int {
		return i + j
	}
}

func main() {
	var add2 = add(2)

	log.Print(add2(5))
}
