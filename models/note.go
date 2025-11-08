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

//
//func DeleteAllNotes() bool {
//	status := true
//	query := `DELETE FROM notes`
//	_, err := db.Exec(query)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func AddNote(note *Note) bool {
//	status := true
//	query := `INSERT INTO notes(title,note_text,username) VALUES(?,?,?)`
//	_, err := db.Exec(query, note.Title, note.NoteText, note.UserName)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func DeleteNote(noteId int) bool {
//	status := true
//	query := `DELETE FROM notes WHERE noteId=?`
//	if _, err := db.Exec(query, noteId); err != nil {
//		status = false
//	}
//	return status
//
//}
//
//func GetNotes(username string) ([]*Note, error) {
//	ns := make([]*Note, 0)
//	query := `SELECT * FROM notes WHERE username=?`
//	rows, err := db.Query(query, username)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		n := new(Note)
//		if err := rows.Scan(&n.NoteId, &n.Title, &n.NoteText, &n.UserName); err != nil {
//			return nil, err
//		}
//		ns = append(ns, n)
//	}
//	return ns, err
//}
