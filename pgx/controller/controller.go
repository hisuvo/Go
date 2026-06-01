package controller

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"get users")
}
