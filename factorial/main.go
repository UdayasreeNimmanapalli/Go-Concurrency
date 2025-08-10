package main

import "fmt"

func main() {
	//total := factorial(5)
	total := concurrentFactorial(5)
	fmt.Println(<-total)
}

// func factorial(n int) int {
// 	var total = 1
// 	for i := n; i >= 1; i-- {
// 		total *= i
// 	}
// 	return total
// }

// Use Goroutines and Channels to calculate factorial it's helpful when you have huge number of calculations. These utilizes all the cores and cpu accordingly
func concurrentFactorial(n int) chan int{
	var out  = make(chan int)
	go func ()  {
		var total = 1
		for i:=n;i>=1;i--{
			total *= i
		}
		out <- total
		close(out)
	}()
	return  out
}