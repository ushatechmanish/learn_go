package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string    `json:"title"`
	Note    string    `json:"note"`
	Created time.Time `json:"created_time"`
}

func (note Note) OutputNote() {
	fmt.Printf("The note title is %v", note.Title)
	fmt.Println()
	fmt.Printf("The note content  is %v", note.Note)
}

func New(title, note string) (Note, error) {
	if note == "" || note == "" {
		return Note{}, errors.New("note is empty")
	}
	return Note{Title: title, Note: note, Created: time.Now()}, nil

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
