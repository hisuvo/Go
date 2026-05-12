package main

import "fmt"

type Skill struct{
	name string
	experience int
}

type User struct{
	name string
	email string
	age int
	country string
	skills Skill
	info func()
}

func about() User {
	return User{
		name: "suvo datta",
		email: "suvo@gmail.com",
		age: 24,
		country: "Bangladesh",
		skills: Skill{
			name: "GoLang",
			experience: 2,
		},
		info: func() {
			fmt.Println(about().name,"is a good developer.He is expert on",about().skills.name)
		},
	}
}


func main() {
	
}
