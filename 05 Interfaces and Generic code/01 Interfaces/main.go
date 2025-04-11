package main

import (
	"fmt"

	"example.com/myNote/note"
)

func main() {
	title, content, err := note.GetNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	myNote := note.NewNote(title, content)
	myNote.Display()
	myNote.Save()
}
