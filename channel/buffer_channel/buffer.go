package buffer_channel

import (
	"fmt"
	"time"
)

func sender(ch chan string){
	time.Sleep(1 * time.Second)
	ch <- "Email send successfully"
}

func payer(ch chan string){
	time.Sleep(1 * time.Second)
	ch <- "Payment successfull"
}

func Buffer() {
	fmt.Println("------------------")
	fmt.Println("Buffer channel")
	fmt.Println("------------------")

	ch := make(chan string, 2)
	go sender(ch)
	go payer(ch)

	fmt.Println("capacity is ch",cap(ch))
	len := cap(ch)

	for range len{
		data := <-ch
		fmt.Println(data)
	}
}