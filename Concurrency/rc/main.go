package rc

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var count int = 0

func increment() {
	defer wg.Done()

	for range 5 {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

var number int = 5
func increase(){
	defer wg.Done()
	number++
}

func RaceCondition() {
	wg.Add(3)
	// go increment()
	// go increment()

	go increase()
	go increase()
	go increase()

	wg.Wait()

	// fmt.Scanln()
	// fmt.Println("Final Count:", count)


	fmt.Println("Final Number:", number)
}