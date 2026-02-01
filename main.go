package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct/note"
)

func main() {
	// var u1 user.User

	// u1.Age = 22
	// u1.Email = "mrvikash@example.com"
	// u1.Password = "password@123"
	// u1.Name = "Vikash Gupta"

	// fmt.Println("The First User is: ", u1)

	// var u2 user.User
	// u2.Age = 25
	// u2.Email = "mrankur@example.com"
	// u2.Password = "password@2234"
	// u2.Name = "Ankur Gupta"

	// fmt.Println("The Second User is: ", u2)

	// user1, _ := user.NewUser("Vikash Gupta", "mrvikash@example.com", "password@1234", 22)

	// fmt.Println("Before Updating ", user1)
	// user1.Age = 23
	// user1.Password = "Password@123"
	// user1.Email = "Ankur@123gmail.com"
	// user1.Name = "Ankur"

	// fmt.Println("After Updating ", user1)

	// fmt.Println("But What About User", user1)

	// isAdult := user1.IsAdult(user1.Age)

	// if isAdult {
	// 	fmt.Println("User is Adult")
	// } else {
	// 	fmt.Println("User is not Adult")
	// }

	// user1, err := user.NewUser("", "", "", 1)
	// if err != nil {
	// 	fmt.Println("Something went wrong: ", err)
	// }

	createNote()

}

func createNote() {
	fmt.Println("You can create your note Here")
	title := getUserInput("Please Enter the Title")
	fmt.Println()
	content := getUserInput("Please Enter the Content")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.CreateFile()
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + ":")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
