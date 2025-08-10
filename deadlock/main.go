package main

import "fmt"

func main(){
	c := make(chan int)
	// c<-5 // it'll blocked here and not reach line 8 since it didn't find reciever
	// fmt.Println(<-c) 

	go func(){
		c<-5 // Now this goroutine will launch and give value to reciever
	}()
	fmt.Println(<-c)
}

// why is there a deadlock??
//  Channels communicate through goroutines