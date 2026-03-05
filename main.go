package main

import (
	"fmt"
)

func main() {
	
	var score int
	for {
		fmt.Print("Enter your score: ")
		fmt.Scan(&score) 
	 
		switch {
		case score >= 90 && score <= 100:
			fmt.Println("Excellent!")
		case score >= 75 && score < 90:
			fmt.Println("Good!")
		case score >= 60 && score < 75:
			fmt.Println("Pass!")
		default:
			fmt.Println("Fail")
		}
	}
}