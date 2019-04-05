package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kikilas/corti-intro/model"
)

type ArticleRepo interface {
	GetByID(id int64) (model.Article, error)
	FindAll() ([]model.Article, error)
	Create(author *model.Article) (model.Article, error)
	Update(author *model.Article) error
}

//AuthorRepo PSQL implementation
type PsqlArticleRepository struct {
	Conn *sqlx.DB
	stmt psqlArticleRepositoryStmt
}

type psqlArticleRepositoryStmt struct {
	getByID *sqlx.Stmt
	findAll *sqlx.Stmt
	update  *sqlx.Stmt
	create  *sqlx.NamedStmt
}

const (
	createArticle = `
	INSERT INTO article (text, author_id)
  	VALUES (:text, :author_id)
  	RETURNING id
	`

	selectArticleBase = `
	SELECT
		id, text, author_id
	FROM article
	`
	getArticleByID = selectArticleBase + `WHERE id = $1`

	updateArticle = `
	UPDATE article SET
		text = COALESCE($2, text),
	 	author_id = COALESCE($3, author_id)
	WHERE id = $1
	`
)

// NewPsqlArticleRepository will create an object that represent the repository.ArticleRepo interface
func NewPsqlArticleRepository(conn *sqlx.DB) ArticleRepo {
	stmt := &psqlArticleRepositoryStmt{}
	err := preparePsqlArticleRepositoryStmt(stmt, conn)
	if err != nil {
		panic(err)
	}

	return &PsqlArticleRepository{
		Conn: conn,
		stmt: *stmt,
	}
}

func preparePsqlArticleRepositoryStmt(stmt *psqlArticleRepositoryStmt, db *sqlx.DB) (err error) {
	if stmt.create, err = db.PrepareNamed(createArticle); err != nil {
		return
	}
	if stmt.findAll, err = db.Preparex(selectArticleBase); err != nil {
		return
	}
	if stmt.getByID, err = db.Preparex(getArticleByID); err != nil {
		return
	}
	if stmt.update, err = db.Preparex(updateArticle); err != nil {
		return
	}

	return
}

func (repo PsqlArticleRepository) Create(article *model.Article) (model.Article, error) {
	//TODO: author_id
	if err := repo.stmt.create.Get(article, article); err != nil {
		return *article, err
	}

	return *article, nil
}

func (repo PsqlArticleRepository) GetByID(id int64) (article model.Article, err error) {
	if err = repo.stmt.getByID.Get(&article, id); err != nil {
		return
	}
	return
}

func (repo PsqlArticleRepository) FindAll() (articles []model.Article, err error) {
	if err = repo.stmt.findAll.Select(&articles); err != nil {
		return
	}
	return
}

func (repo PsqlArticleRepository) Update(article *model.Article) (err error) {
	//TODO: author_id
	_, err = repo.stmt.update.Exec(
		article.ID, article.Text, article.Author.ID,
	)
	return
}
