package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func foo3() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar3() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar ", i)
	}
}

func init() {
	fmt.Println("init func, that executes once at the beginning (without calling)")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo3()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("foo anon ", i)
		}
		wg.Done()
	}()
	bar3()
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()

	//race()	// also use -race option to check for race conditions eg: go run -race concurrency.go
	//raceFixWithMutex()
	raceFixWithAtomic()

}

func race() {
	const gs = 100 // no of go routines to create
	counter := 0
	wg.Add(gs) // wait for all the go routines to complete
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//time.Sleep(time.Second*2)
			wg.Done()
		}()
		fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Cunter value ", counter)
}

func raceFixWithMutex() {
	const gs = 100 // no of go routines to create
	counter := 0
	wg.Add(gs) // wait for all the go routines to complete
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//time.Sleep(time.Second*2)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Counter value ", counter)
}

func raceFixWithAtomic() {
	const gs = 100 // no of go routines to create
	//counter := 0
	var counter int64
	wg.Add(gs) // wait for all the go routines to complete
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			//time.Sleep(time.Second*2)
			wg.Done()
		}()
		fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Counter value ", counter)
}
