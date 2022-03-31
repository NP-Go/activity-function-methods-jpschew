package main

import "fmt"

//Declare a struct
type customer struct {
	fName    string
	lName    string
	userName string
	password string
	email    string
	phone    int
	address  string
}

//Create Methods
func (c customer) userCredentials() (string, string) {
	return c.userName, c.password
}

func (c customer) userAddress() string {
	return c.address
}

func (c customer) printUserInformation() {
	fmt.Println(c)
}
