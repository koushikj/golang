package main

import "fmt"

func main() {
	//channelThatBlocks()
	channelWithRoutines()
	channelUsingBuffer()
	sendOnlyChannels()
	receiveOnlyChan()
	communicationsWithinGoRoutinesUsingChannels()
}

func channelThatBlocks() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}

func channelWithRoutines() {
	c := make(chan string)
	go func() {
		c <- "hello"
	}()
	fmt.Println(<-c)
}

func channelUsingBuffer() {
	c := make(chan float64, 3) // reserves buffer for 3 float values
	c <- 102.11
	c <- 2919.1
	c <- 81219.1
	//c <- 81219.1	//will block
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)	//will block
	fmt.Printf("%T \n", c)

}

func sendOnlyChannels() {
	c := make(chan<- int)
	go func() {
		c <- 102
	}()
	//fmt.Println(<-c) //error - cant receive from send only chan
	fmt.Printf("%T \n", c)
}
func receiveOnlyChan() {
	c := make(<-chan string)
	//c<-"abc" // error - cant send on receive only chan
	fmt.Printf("%T \n", c)
	//fmt.Println(<-c) // run time error - block
}

func communicationsWithinGoRoutinesUsingChannels() {
	c := make(chan int)
	go send(c) // for send channel - this has to be go routine otherwise main program will block
	receive(c) // waits for the data sent by "send" go routine
	fmt.Println("end of communicationsWithinGoRoutinesUsingChannels")
}

func send(c chan<- int) { //type conversion from bi-directional channel to send-only channel
	for i := 0; i < 5; i++ {
		fmt.Println("sending ", i, " in channel ", c)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) { //type conversion from bi-directional channel to receive-only channel
	fmt.Println("receiving int from channel ", c)
	for v := range c { // iterate over the range of received values until the channel closes.
		//fmt.Println(<-c)
		fmt.Println("received :", v)
	}
}
