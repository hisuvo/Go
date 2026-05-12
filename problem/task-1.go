package main

import "fmt"

func displayMenu() {
	fmt.Println("Welcome to the grade calculator")
	fmt.Println("Chose an options __ ")
	fmt.Println("1.calculate grade")
	fmt.Println("2.Chack pass/fail status")
	fmt.Println("3.Exit the program")
	fmt.Println("No:")
}

func calculateGrade(score int) string{
	if(score > 90){
		return "A"
	} else if( score > 80){
		return "B"
	} else if (score > 60){
		return "C"
	} else if (score > 40){
		return "D"
	} else {
		return "F"
	}
}

func main() {
	
	var score int
	var chose int
	running := true
	
	for running{
		displayMenu()
		fmt.Scan(&chose)
		
		switch chose{
		case 1:
			fmt.Println("You chose grade calculater")
			fmt.Println("Enter your number")
			fmt.Scan(&score)
			fmt.Println("CGPA:", calculateGrade(score))
			fmt.Println("------------------------")
			
		case 2:
			fmt.Println("check pase and fail status")
			fmt.Println("Enter your number")
			fmt.Scan(&score)
			if(score > 40){
				fmt.Println("Your are passed")
				} else{
					fmt.Println("Your are failed")
				}
				
			fmt.Println("------------------------")
		case 3:
			fmt.Println("Exit...");
			running=false

		default:
			fmt.Println("Enter currect number")
		}

	}
}

