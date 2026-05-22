package main

import (
	"concurrency/group"
	"concurrency/rc"
	"fmt"
	"net"
	"sync"
	"time"
)

func uploadFile() {
	fmt.Println("File Upload start....")
	time.Sleep(3*time.Second)
	fmt.Println("File Upload done....")
	// wg.Add(-1)
	wg.Done()
}

func saveDB() {
	fmt.Println("saveing to db....")
	time.Sleep(2*time.Second)
	fmt.Println("save to done....")
	// wg.Add(-1)
	wg.Done()
}

func sendEmail() {
	fmt.Println("email sending start....")
	time.Sleep(3*time.Second)
	fmt.Println("email send successfull....")
	// wg.Add(-1)
	wg.Done()
}

func getMAC() string {

	interfaces, _ := net.Interfaces()

	for _, i := range interfaces {

		if i.HardwareAddr.String() != "" {
			fmt.Println(i)
			return i.HardwareAddr.String()
		}
	}

	return ""
}

var wg sync.WaitGroup

func main() {
	start := time.Now()

	wg.Add(1)
	go uploadFile()
	// wg.Go(uploadFile())
	
	
	wg.Add(1)
	go saveDB()

	wg.Add(1)
	go sendEmail()

	wg.Wait() // wait untils counter is 0

	
	

	//basic go
	fmt.Println("--------------------------------------")
	fmt.Println("Basic function here")
	fmt.Println("--------------------------------------")

	group.Owner()

	fmt.Println("-------------------------")
	fmt.Println("all tasks completed")
	fmt.Println("take time",time.Since(start))
	
	fmt.Println("-------------------------")
	fmt.Println("Race Condition")
	fmt.Println("-------------------------")
	rc.RaceCondition()

}