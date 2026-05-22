package group

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Owner() {
	wg.Add(1)
	go worker1()

	wg.Add(1)
	go worker2()

	wg.Add(1)
	go worker3()

	wg.Wait()
	// time.Sleep(3 * time.Second)
}

func worker1() {
	defer wg.Done()
	fmt.Println("Worker one start.....",)
	time.Sleep(3 * time.Second)
	fmt.Println("Worker one finished ",)
}

func worker2() {
	defer wg.Done()
	fmt.Println("Worker two start.....",)
	time.Sleep(3 * time.Second)
	fmt.Println("Worker two finished ",)
}

func worker3() {
	defer wg.Done()
	fmt.Println("Worker three start.....",)
	time.Sleep(3 * time.Second)
	fmt.Println("Worker three finished ",)
}

