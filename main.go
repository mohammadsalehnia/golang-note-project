package main

import (
	"bufio"
	"fmt"
	note "golang-note-project/note"
	"os"
	"strings"
)

func main() {

	fmt.Println("Simple Golang n project")

	title, content := getNoteData()
	n := note.New(title, content)
	n.PrintNoteData()
	err := n.SaveNoteInFile()
	if err != nil {
		fmt.Println("saving the file failed")
		return
	}

	fmt.Println("saved note")

	//n.ClearNoteData()
	//n.PrintNoteData()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
