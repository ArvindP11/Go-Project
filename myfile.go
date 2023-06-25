package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Quiz game!")
	fmt.Printf("Enter your name:")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to this game\n", name)

	fmt.Println("Enter your age")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play the game")
	} else {
		fmt.Println("You cannot play")
		return
	}

	score := 0
	ques := 2
	fmt.Printf("What is better Pizza 10 or Burger 20? ")
	var ans string
	var ans2 string
	fmt.Scan(&ans, &ans2)

	if ans+" "+ans2 == "Burger 20" || ans+" "+ans2 == "burger 20" {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("How many slices does a Pizza have?")
	var slices uint
	fmt.Scan(&slices)

	if slices == 6 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("You have scored %v out of %v.\n", score, ques)
	percent := (float64(score) / float64(ques)) * 100
	fmt.Printf("You score percentage is %v %%.", percent)
}
