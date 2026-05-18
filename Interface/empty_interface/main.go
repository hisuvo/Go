package main

import "fmt"


func print(value any){ // type any = interface{}
	fmt.Printf("%T\n",value)
}

func process(data interface{}){
	strData, ok := data.(string) // type assertion and ok adiom
	if ok {
		fmt.Println("string data is ->",strData)
	}

	intData, ok := data.(int)

	if ok {
		fmt.Println(intData+100)
	}
}


func main() {
	print("Empty interface here?")
	print(2000000000)
	print(false)

	process(2354)
}