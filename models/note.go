package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	NoteId   int    `gorm:"primaryKey;autoIncrement" json:"noteId"`
	Title    string `json:"title;size:200"`
	NoteText string `json:"noteText;size:2000"`
	UserName string `json:"username;size:100"`
	User     User   `gorm:"foreignKey:UserName" json:"user"`
}

func DeleteAllNotes() bool {
	gorm.G[Note](db).Delete(ctx)
	return true
}

func AddNote(note *Note) bool {
	gorm.G[Note](db).Create(ctx, note)
	return true
}

func DeleteNote(noteId int) bool {
	gorm.G[Note](db).Where("noteId=?", noteId).Delete(ctx)
	return true

}

func GetNotes(username string) ([]Note, error) {
	n, err := gorm.G[Note](db).Where("username=?", username).Find(ctx)
	if err != nil {
		return nil, err
	}
	return n, nil
}
