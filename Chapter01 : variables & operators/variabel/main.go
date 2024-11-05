package main

import (
	"fmt"
)

func main() {
	var firstName string = "Iqbal"
	var familyName string = "Athorid"
	var age int = 24
	var peanutAllergy bool = false

	//define multiple variables at once with var
	var (
		work string = "Software Engineer"
		company string = "Google"
		isMarried bool = false
	)

	//type inference
	var city = "Jakarta"

	//short hand declaration
	country := "Indonesia"

	//multiple variable declaration in one line
	var test1,test12, end float32
	var day, month, year int = 1, 1, 2020

	fmt.Println("First Name: ", firstName)
	fmt.Println("Family Name: ", familyName)
	fmt.Println("Age: ", age)
	fmt.Println("Peanut Allergy: ", peanutAllergy)
	fmt.Println("Work: ", work)
	fmt.Println("Company: ", company)
	fmt.Println("Is Married: ", isMarried)
	fmt.Println("City: ", city)
	fmt.Println("Country: ", country)

}