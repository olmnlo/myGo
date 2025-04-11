package note

import (
	"bufio"
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

func NewNote(inTitle, inContent string) *Note {

	return &Note{
		Title:     inTitle,
		Content:   inContent,
		CreatedAt: time.Now(),
	}
}

func (note Note) Display() {
	fmt.Println("title: ", note.Title, " content: ", note.Content)
}

func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, err

}

func GetNoteData() (string, string, error) {
	title, err := GetUserInput("Note title: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := GetUserInput("Note contetnt: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return title, content, err
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	os.WriteFile(filename, json, 0644)
	return nil
}
