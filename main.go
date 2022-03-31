package main

import "fmt"

func main() {
	//Declare new item customer
	customer1 := customer{
		fName:    "Michael",
		lName:    "Jordan",
		userName: "MJ2020",
		password: "1234567",
		email:    "MJ2020@gmail.com",
		phone:    12345678,
		address:  "18227 Capstan Green Road Cornelius, NC28031.",
	}
	//Insert Code here
	customer1.printUserInformation()
	fmt.Println(customer1.userCredentials())
	fmt.Println(customer1.userAddress())
}
