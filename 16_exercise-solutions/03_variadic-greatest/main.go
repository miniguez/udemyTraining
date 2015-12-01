package main
import "fmt"


func greatest(numbers ...int) int {
	var largest int
	for _,value := range numbers {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func main() {
	max := greatest(543, 7, 9, 123, 4, 23, 435, 53, 125)
	fmt.Println(max)
}
