package gochannel

import (
	"fmt"
	"time"
)

func uploadFile(ch chan string) {
	fmt.Println("Uloading file....")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")

	fileUrl := "http://s3.ero45434.com"
	ch <- fileUrl // keep fileurl in channgle
}

func GoChannel() {
	ch := make(chan string) // create channel

	go uploadFile(ch) // pase channel as a argument

	fileUrl := <- ch // get file url from channel
	fmt.Println("File url:",fileUrl) // print result
}