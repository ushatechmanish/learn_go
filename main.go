package main

import (
	"bufio"
	"examples.com/note/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title := getUserInput("Title:")
	noteContent := getUserInput("Note:")
	note, err := note.New(title, noteContent)
	if err != nil {
		fmt.Println(err)
	}
	note.OutputNote()
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	//var value string
	//fmt.Scan(&value) // only for single words no spaces
	buf := bufio.NewReader(os.Stdin)
	input, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}
