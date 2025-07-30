package models

type Note struct {
	NoteId   int    `json:"noteId"`
	Title    string `json:"title`
	NoteText string `json:"noteText"`
	UserName string `json:"username`
}

func AddNote(note *Note) bool {
	status := true
	query := `INSERT INTO notes(title,note_text,username) VALUES(?,?,?)`
	_, err := db.Exec(query, note.Title, note.NoteText, note.UserName)
	if err != nil {
		status = false
	}
	return status
}

func GetNotes(username string) ([]*Note, error) {
	ns := make([]*Note, 0)
	query := `SELECT * FROM notes WHERE username=?`
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		n := new(Note)
		if err := rows.Scan(&n.NoteId, &n.Title, &n.NoteText, &n.UserName); err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, err
}
