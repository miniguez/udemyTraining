package main

import "fmt"

//iota can only be used with const

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
