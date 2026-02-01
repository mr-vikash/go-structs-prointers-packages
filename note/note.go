package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Note struct {
	Title   string
	Content string
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Please Enter the Title and Content to create the Note")
	}
	return Note{
		Title:   title,
		Content: content,
	}, nil
}

func (note Note) CreateFile() {
	fileName := strings.ReplaceAll(note.Title, " ", "-")
	content, _ := json.Marshal(note)
	os.WriteFile(fileName+".json", content, 0644)
}
