package main

import "fmt"

type customerError struct {
	message string
	status  int
}

func (ce *customerError) Error() string {
	return ce.message
}

func login(password string) error {
	if password != "1234" {
		return &customerError{
			message: "Pawword worng",
			status:  404,
		}
	}
	return nil
}

func main() {
	err := login("12345")
	if err != nil {

		// type assertion
		if customErr, ok := err.(*customerError); ok {
			fmt.Println("Message:", customErr.message)
			fmt.Println("Status:", customErr.status)
		}

		return
	}

	fmt.Println("Login success")
}