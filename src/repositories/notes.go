package repositories

import (
	"database/sql"
	"modules/src/models"
)

// Notes function that represents a repository of notes
type Notes struct {
	db *sql.DB
}

func NoteRepo(db *sql.DB) *Notes {
	return &Notes{db}
}

// Create function used to insert a note into the database
func (repository Notes) Create(note models.Note) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO notes (title, description, user_id) VALUES(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(note.Title, note.Content, note.AuthorID)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil

}

// GetAll function used to retrieve all notes that meet a nickname filter
func (repository Notes) GetAll(authorID uint64) ([]models.Note, error) {
	lines, err := repository.db.Query(
		`SELECT n*, u.nick_name FROM notes n
		inner join users u on u.id = n.author_id
		WHERE n.author_id = ?`,
		authorID,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var notes []models.Note

	for lines.Next() {
		var note models.Note

		if err = lines.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.AuthorID,
			&note.CreatedAt,
		); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

// GetByID function used to retrieve a note by its ID
func (repository Notes) GetByID(noteID uint64) (models.Note, error) {
	line, err := repository.db.Query(
		`SELECT n*, u.nick_name FROM notes n inner join users u on u.id = n.author_id WHERE n.id = ?`,
		noteID,
	)
	if err != nil {
		return models.Note{}, err
	}
	defer line.Close()
	var note models.Note
	if line.Next() {
		if err = line.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.AuthorID,
			&note.CreatedAt,
		); err != nil {
			return models.Note{}, err
		}
	}
	return note, nil
}

// Update function used to change a notes information in the database
func (repository Notes) Update(ID uint64, note models.Note) error {
	statement, err := repository.db.Prepare(
		"UPDATE notes set title = ?, content = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(note.Title, note.Content, ID); err != nil {
		return err
	}

	return nil
}

// Delete function used to delete a notes information in the database
func (repository Notes) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
