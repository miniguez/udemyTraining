package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	fmt.Println(x) // x is still 5
}
