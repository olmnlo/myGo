package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	os.WriteFile(filename, json, 0644)
	return nil
}

func New(inContent string) (*Todo, error) {

	if inContent == "" {
		return &Todo{}, errors.New("Invalid")
	}

	return &Todo{
		Text: inContent,
	}, nil
}
