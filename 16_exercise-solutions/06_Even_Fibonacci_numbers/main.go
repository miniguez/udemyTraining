package main
import "fmt"

func sumEvenFibonacci(limit int) int {
	a, b := 0, 1
	var sum int
	for {
		a,b = b, a+b
		if b%2 == 0 {
			sum	+= b
		}
		if b > limit {
			break
		}
	}
	return sum
}


func main() {
	fmt.Println(sumEvenFibonacci(4000000))
}
