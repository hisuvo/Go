package unbuffer_channel

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


func UnBuffer() {
	fmt.Println("------------------")
	fmt.Println("Un Buffer channel")
	fmt.Println("------------------")

	ch := make(chan string)
	go sender(ch)
	go payer(ch)

	fmt.Println("capacity is ch",cap(ch)) // alawys cap: 0

	for range 2{
		data := <-ch
		fmt.Println(data)
	}
}