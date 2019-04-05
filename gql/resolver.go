package gql

import (
	"context"

	"github.com/kikilas/corti-intro/db/repository"
	"github.com/kikilas/corti-intro/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	authorRepo  *repository.AuthorRepo
	articleRepo *repository.ArticleRepo
}

func (r *Resolver) Article() ArticleResolver {
	return &articleResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type articleResolver struct{ *Resolver }

func (r *articleResolver) Author(ctx context.Context, obj *model.Article) (*model.Author, error) {
	author, e := r.Query().Author(ctx, obj.AuthorID)

	return author, e
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.InputNewArticle) (*model.Article, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.InputOldArticle) (*model.Article, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Article(ctx context.Context, id int64) (*model.Article, error) {
	article, e := (*r.articleRepo).GetByID(id)

	return &article, e
}
func (r *queryResolver) Articles(ctx context.Context) ([]model.Article, error) {
	return (*r.articleRepo).FindAll()
}
func (r *queryResolver) Author(ctx context.Context, id int64) (*model.Author, error) {
	author, e := (*r.authorRepo).GetByID(id)

	return &author, e
}
func (r *queryResolver) Authors(ctx context.Context, name *string) ([]model.Author, error) {
	if name == nil {
		return (*r.authorRepo).GetAuthors()
	}

	return (*r.authorRepo).GetByName(*name)
}
