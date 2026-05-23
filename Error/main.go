package main

import (
	"errors"
	"fmt"
)

var err error

func calculate(a int, b int) (int, error){
	if b <= 0{
		return 0, errors.New("Value must be b > 0")
	}
	return a/b, nil
}

func main() {
	result, err := calculate(8,-10)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(result)
}