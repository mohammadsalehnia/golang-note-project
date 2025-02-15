package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) Todo {
	return Todo{text}
}

func (t *Todo) Display() {
	fmt.Print("----------------- Todo data --------------- \n ")
	fmt.Printf("Text: %s\n", t.Text)
	fmt.Printf("------------------------------------------- \n")
}

func (t *Todo) ClearTodoData() {
	t.Text = ""
}

func (t *Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}
