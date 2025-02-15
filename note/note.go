package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) Note {
	return Note{title, content, time.Now()}
}

func (n *Note) Display() {
	fmt.Print("----------------- Note data --------------- \n ")
	fmt.Printf("Title: %s\n", n.Title)
	fmt.Printf("Content: %s\n", n.Content)
	fmt.Printf("CreatedAt: %s\n", n.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("------------------------------------------- \n")
}

func (n *Note) ClearNoteData() {
	n.Title = ""
	n.Content = ""
}

func (n *Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}
