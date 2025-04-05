package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) /*here we are attaching this function to user struct*/ outputUserDetails() {
	// *user is not required to dereference . shortcut is allowed by go to use user directly  only for structs
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// This will help clarify the assignment of which value to which field
	//appUser := user{
	//	firstName: userFirstName,
	//	lastName:  userLastName,
	//	birthDate: userBirthDate,
	//	createdAt: time.Now(),
	//}

	appUser := user{userFirstName, userLastName, userBirthDate, time.Now()}
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
