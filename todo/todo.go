package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Todo struct {
	Text string
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("Please insert the todo content")
	}
	return &Todo{
		Text: text,
	}, nil
}

func (todo Todo) Save() {
	todoContent, _ := json.Marshal(todo)
	fileName := strings.ReplaceAll(todo.Text, " ", "-")
	os.WriteFile(fileName+".json", todoContent, 0644)
}
