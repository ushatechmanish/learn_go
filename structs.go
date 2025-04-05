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

func newUser(firstName, lastName, birthDate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Time{},
	}
}

func (u *user) outputUserDetails() {
	// *user is not required to dereference . shortcut is allowed by go to use user directly  only for structs
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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

	appUser := newUser(userFirstName, userLastName, userBirthDate)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
