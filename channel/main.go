package main

import (
	"fmt"

	"example.com/channel/buffer_channel"
	"example.com/channel/unbuffer_channel"
)

func sender(ch chan string) {
	email := "email send successfully"
	ch <- email
}

func OutPut(){
	// create channel
	ch := make(chan string)

	go sender(ch)

	msg := <-ch

	fmt.Println("Sender infor:", msg)
}

func main() {
	buffer_channel.Buffer()
	unbuffer_channel.UnBuffer()
}