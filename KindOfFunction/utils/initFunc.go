// init function - you can not call
// this, computer calls this automatically

package utils

import "fmt"

var myInfo = "My name is suvo datta.currently i am learnig Golang"

func InitFunc() {
	fmt.Println("InitFunc()",myInfo)
}

// func init(){
// 	fmt.Println("Hello from init() ->", myInfo)
// 	myInfo = "What's your future goal" 
// }