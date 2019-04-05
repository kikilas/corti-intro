package middleware

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/kikilas/corti-intro/gql"
	"github.com/rs/cors"
	"net/http"
)

func NewCorsHandler(r *httprouter.Router) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
	})

	return corsHandler.Handler(r)
}

func NewGraphQLHandlerFunc(graphQLConfig gql.Config) http.HandlerFunc {
	return handler.GraphQL(gql.NewExecutableSchema(graphQLConfig),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}))
}
