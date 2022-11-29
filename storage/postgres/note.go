package postgres

import (
	"github.com/burxondv/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewNote(db *sqlx.DB) repo.NoteStorageI {
	return &noteRepo{
		db: db,
	}
}

func (ur *noteRepo) Create(note *repo.Note) (*repo.Note, error) {
	return nil, nil
}

func (ur *noteRepo) Get(id int64) (*repo.Note, error) {
	return nil, nil
}

func (ur *noteRepo) GetAll(params *repo.GetAllNotesParams) (*repo.GetAllNotesResult, error) {
	return nil, nil
}

func (ur *noteRepo) Update(note *repo.Note) (*repo.Note, error) {
	return nil, nil
}

func (ur *noteRepo) Delete(id int64) error {
	return nil
}