package main

import (
	"fmt"
	"sync"
)

func main() {
	produce := producer(2, 3, 4, 5, 6, 7)
	consumer1 := consumer(produce)
	consumer2 := consumer(produce)
	consumer3 := consumer(produce)

	// // 1 producer 3 consumers - FAN-OUT
	// for v := range consumer1 {
	// 	fmt.Println("square of consumer1:", v)
	// }
	// for val := range consumer2 {
	// 	fmt.Println("square of consumer2:", val)
	// }
	// for va := range consumer3 {
	// 	fmt.Println("square of consumer3:", va)
	// }

	finalResult := merge(consumer1, consumer2, consumer2, consumer3)
	for v := range finalResult {
		fmt.Println("fan in val:", v)
	}
}

// This produces or gets numbers to perform any operation
func producer(nums ...int) chan int {
	var in = make(chan int)
	go func() {
		for _, v := range nums {
			in <- v
		}
		close(in)
	}()
	return in
}

func consumer(c chan int) chan int {
	var out = make(chan int)
	go func() {
		for v := range c {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func merge(c ...chan int) chan int {
	var fanIN = make(chan int)
	var wg sync.WaitGroup
	output := func(c chan int) {
		for v := range c {
			fanIN <- v
		}
		wg.Done()
	}
	wg.Add(len(c))
	for _, chans := range c {
		go output(chans)
	}
	go func() {
		wg.Wait()
		close(fanIN)
	}()
	return fanIN
}
