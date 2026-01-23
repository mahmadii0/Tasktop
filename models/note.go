package models

import "gorm.io/gorm"

type Note struct {
	NoteId   int    `gorm:"primaryKey;autoIncrement" json:"noteId"`
	Title    string `json:"title;size:200"`
	NoteText string `json:"noteText;size:2000"`
	UserID   int64  `gorm:"not null;index" json:"userId"`
	User     User   `gorm:"foreignKey:UserID;references:ID"`
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
	gorm.G[Note](db).Where("note_id=?", noteId).Delete(ctx)
	return true

}

func GetNotes(userId int64) ([]Note, error) {
	n, err := gorm.G[Note](db).Where("user_id=?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}
	return n, nil
}
