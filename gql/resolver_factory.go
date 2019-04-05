package gql

import (
	"github.com/kikilas/corti-intro/db/repository"
)

func NewConfig(authorRepo *repository.AuthorRepo, articleRepo *repository.ArticleRepo) Config {
	return Config{
		Resolvers: &Resolver{
			authorRepo:  authorRepo,
			articleRepo: articleRepo,
		},
	}
}
