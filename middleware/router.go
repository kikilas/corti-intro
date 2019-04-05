package middleware

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(graphQLHandlerFunc http.HandlerFunc) *httprouter.Router {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler.Playground("GraphQL Playground", "/query").ServeHTTP(w, r)
	})
	r.POST("/query", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		graphQLHandlerFunc.ServeHTTP(w, r)
	})

	return r
}
