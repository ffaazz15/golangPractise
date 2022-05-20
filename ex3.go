package main

import (
	"fmt"
)

func ex3() {
	score := 0
	var name string
	var age float32
	fmt.Printf("please tell us ur name:-")
	fmt.Scan(&name)
	fmt.Printf("hey %v !!\n", name)
	fmt.Println("please tell us ur age:-")
	fmt.Scan(&age)
	fmt.Printf("hey %v !!\n", name)
	if age >= 10 {

		fmt.Println("yay u can play")
		score += 1
	} else {
		fmt.Println("you cant play")

	}
	fmt.Printf("who is the Pm of india , rg or namo??")
	var answer string
	var answer2 string
	fmt.Scanln(&answer, &answer2)
	fmt.Println(answer, answer2)
	if answer+""+answer2 == "namo" {
		fmt.Println("waah ..waah ..mast jawab tha ")
		score += 1
	} else if answer+""+answer2 == "rg" {
		fmt.Println("kya phooktha hai bhai?")

	} else {
		fmt.Println("bhai kounsa wala nasha karra hai??")
	}

	fmt.Println("which year did modiji become pm?")
	var year int
	fmt.Scanln(&year)

	if year == 2014 {

		fmt.Println("you are correct")
		score++
	} else {
		fmt.Println("no your wrong")
	}

	//checking for logical operator
	var answer3 string
	fmt.Println("who is best batsman..virat kohli or sachin")
	fmt.Scanln(&answer3)
	if answer3 == "virat kohli" || answer3 == "sachin" {
		fmt.Println("good you have cricketing knowledge")
		score++
	} else {
		fmt.Println("bhai in dono me se koi ek naam bolna!!")
	}
	fmt.Printf("your score is %v \n", score)
	no_of_questions := 3
	percent := (float64(score) / float64(no_of_questions)) * 100
	fmt.Printf("you have scored:%v%% \n", percent)

}
