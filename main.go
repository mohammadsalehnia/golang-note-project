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

type outputtable interface {
	saver
	Display()
}

func printSomething(value interface{}) {
	fmt.Println(value)
}

func main() {

	printSomething("Simple Golang Note project")

	title, content := getNoteData()
	userNote := note.New(title, content)
	outputData(&userNote)

	todoText := getTodoData()
	userTodo := todo.New(todoText)
	outputData(&userTodo)
}

func outputData(data outputtable) {
	data.Display()
	err := saveData(data)
	if err != nil {
		fmt.Println(err)
	}
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
