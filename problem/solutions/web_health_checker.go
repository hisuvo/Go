package problem

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct{
	Url string
	Status string
	Err error
}

func checkWebsiteUrl(url string, ch chan Result){
	res, err := http.Get(url)
		
		if err != nil {
			// fmt.Println(url,"is down")
			ch <- Result{
				Url: url,
				Status: "is down",
				Err: err,
			}
		} else {
			// fmt.Println(url,"is up")

			ch <- Result{
				Url: url,
				Status: "is down",
				Err:nil,
			}

			res.Body.Close()
		}
}

func WebHealthChecker() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://suvodatta.com",
		"https://github.com",
		"https://linksdin.com",
	}

	ch := make(chan Result)
	start := time.Now()

	for _, url := range urls {
		go checkWebsiteUrl(url,ch)
	}

	for range urls{
		data := <- ch

		if data.Err != nil {
			fmt.Println(data.Url, data.Status, `Error:`,data.Err)
			continue
		}

		fmt.Println(data)
	}

	fmt.Println("All url checked successfully")
	
	fmt.Println("Take Time:", time.Since(start))
}