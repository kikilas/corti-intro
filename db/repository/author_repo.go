package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kikilas/corti-intro/model"
)

// AuthorRepo interface
type AuthorRepo interface {
	Create(author *model.Author) (model.Author, error)
	GetByID(id int64) (model.Author, error)
	GetAuthors() ([]model.Author, error)
	GetByName(name string) ([]model.Author, error)
	UpdateUser(author *model.Author) error
}

//AuthorRepo PSQL implementation
type PsqlAuthorRepository struct {
	Conn *sqlx.DB
	stmt psqlAuthorRepositoryStmt
}

type psqlAuthorRepositoryStmt struct {
	createAuthor    *sqlx.NamedStmt
	getAuthors      *sqlx.Stmt
	getAuthorByID   *sqlx.Stmt
	getAuthorByName *sqlx.Stmt
	updateAuthor    *sqlx.Stmt
}

const (
	createAuthor = `
	INSERT INTO author (name)
  	VALUES (:name)
  	RETURNING id
	`

	selectAuthorBase = `
	SELECT
		id, name
	FROM author
	`
	getAuthorByID   = selectAuthorBase + `WHERE id = $1`
	getAuthorByName = selectAuthorBase + `WHERE name = $1`

	updateAuthor = `
	UPDATE author SET
		name = COALESCE($2, name)
	WHERE id = $1
	`
)

// NewPsqlAuthorRepository will create an object that represent the repository.AuthorRepo interface
func NewPsqlAuthorRepository(conn *sqlx.DB) AuthorRepo {
	stmt := &psqlAuthorRepositoryStmt{}
	err := preparePsqlAuthorRepositoryStmt(stmt, conn)
	if err != nil {
		panic(err)
	}

	return &PsqlAuthorRepository{
		Conn: conn,
		stmt: *stmt,
	}
}

func preparePsqlAuthorRepositoryStmt(stmt *psqlAuthorRepositoryStmt, db *sqlx.DB) (err error) {
	if stmt.createAuthor, err = db.PrepareNamed(createAuthor); err != nil {
		return
	}
	if stmt.getAuthors, err = db.Preparex(selectAuthorBase); err != nil {
		return
	}
	if stmt.getAuthorByID, err = db.Preparex(getAuthorByID); err != nil {
		return
	}
	if stmt.getAuthorByName, err = db.Preparex(getAuthorByName); err != nil {
		return
	}
	if stmt.updateAuthor, err = db.Preparex(updateAuthor); err != nil {
		return
	}

	return
}

func (repo PsqlAuthorRepository) Create(author *model.Author) (model.Author, error) {
	if err := repo.stmt.createAuthor.Get(author, author); err != nil {
		return *author, err
	}
	return *author, nil
}

func (repo PsqlAuthorRepository) GetByID(id int64) (author model.Author, err error) {
	if err = repo.stmt.getAuthorByID.Get(&author, id); err != nil {
		return
	}
	return
}

func (repo PsqlAuthorRepository) GetAuthors() (authors []model.Author, err error) {
	if err = repo.stmt.getAuthors.Select(&authors); err != nil {
		return
	}
	return
}

func (repo PsqlAuthorRepository) GetByName(name string) (authors []model.Author, err error) {
	if err = repo.stmt.getAuthorByName.Select(&authors, name); err != nil {
		return
	}
	return
}

func (repo PsqlAuthorRepository) UpdateUser(author *model.Author) (err error) {
	_, err = repo.stmt.updateAuthor.Exec(
		author.ID, author.Name,
	)
	return
}
