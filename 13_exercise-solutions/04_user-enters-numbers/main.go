package main
import "fmt"


func main()  {
	var (
		smallNumber int
		longNumber int
	)
	fmt.Println("Enter a small number:")
	fmt.Scan(&smallNumber)
	fmt.Println("Enter a long number:")
	fmt.Scan(&longNumber)
	fmt.Println("result ", longNumber/smallNumber)
}
