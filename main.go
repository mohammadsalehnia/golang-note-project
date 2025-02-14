package main

import (
	"bufio"
	"fmt"
	"golang-note-project/note"
	"golang-note-project/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {

	fmt.Println("Simple Golang n project")
	title, content := getNoteData()
	userNote := note.New(title, content)
	userNote.PrintNoteData()

	err := saveData(&userNote)

	if err != nil {
		fmt.Println(err)
	}

	todoText := getTodoData()
	userTodo := todo.New(todoText)
	userTodo.PrintTodoData()
	err = saveData(&userTodo)

	if err != nil {
		fmt.Println(err)
	}

	//n.ClearNoteData()
	//n.PrintNoteData()
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving the file failed")
		return err
	}
	fmt.Println("file saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo Text:")
	return text
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
