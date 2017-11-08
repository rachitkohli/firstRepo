// Concurrency1
package main

import (
	"fmt"
	"time"
)

func main() {
	// testGoRoutine()
	// BasicChannel()
	// FirstChannel()
	// BufferedChannel()
	channelSynchronization()
	
}

func testGoRoutine() {
	//Running GoRoutine in seperate memories
	LoopMe := func (wrd string) () {
		for i:=1; i<=5; i++ {
		fmt.Println(wrd)
		}
	}
	go LoopMe("Hi from here!") //Called as a Go Routine
	
	LoopMe("Hi from another process")
}

func BasicChannel() {
	//Basic example of creating a channel, passing string to channel via Go Routine.
	chn := make(chan string)
	go func() { chn <- "ping" }()
	msg := <-chn
	fmt.Println(msg)
}

func FirstChannel() {
	//Basic channel routine
	sum := func(numbers []int, c chan int) {
		var total int
		for _, v:= range numbers {
			total += v
		}
		c <- total //send total to c
	}
	
	nos := []int{4,3,6,1,7,5,2,8}
	chn := make(chan int)
	go sum(nos[:len(nos)/2],chn)
	go sum(nos[len(nos)/2:],chn)
	go sum(nos, chn)
	
	x,y,z := <-chn, <-chn, <-chn
	fmt.Println(x, y, z, x+y)
}

func BufferedChannel() {
	//Example on creating buffered Channel. Values are transmitted sequentially & received too
	c := make(chan int, 3) //Second argument is the length of buffer
	c <- 6
	c <- 3
	c <- 5
	//c3 := func() { c <-5 }
	//go c3()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	
}

func channelSynchronization() {
	//Synchronize execution across go routines.
	
	worker := func(notify chan bool) {
				fmt.Print("Working....")	
				time.Sleep(time.Second)
				fmt.Println("Done")
				
				notify <- true
			}
		
	done := make(chan bool, 1)	//Channel used to notify that this function is completed. Sends a value to notify
	go worker(done) //Starts worker goroutine, passing a channel where it will be notified if its done
	<- done		//Worker wont execute until it is returned
}