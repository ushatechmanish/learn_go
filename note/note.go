package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title   string
	note    string
	created time.Time
}

func (note Note) OutputNote() {
	fmt.Printf("The note title is %v", note.title)
	fmt.Println()
	fmt.Printf("The note content  is %v", note.note)
}

func New(title, note string) (Note, error) {
	if note == "" || note == "" {
		return Note{}, errors.New("note is empty")
	}
	return Note{title: title, note: note, created: time.Now()}, nil

}
