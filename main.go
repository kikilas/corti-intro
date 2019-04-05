package main

import (
	"fmt"
	"github.com/kikilas/corti-intro/db"
	"github.com/kikilas/corti-intro/db/repository"
	"github.com/kikilas/corti-intro/gql"
	"github.com/kikilas/corti-intro/middleware"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//Arguments
	configPath := pflag.String("config-name", "config.local", "Name of the config file")
	pflag.Parse()

	viper.SetConfigName(*configPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	serverPort := viper.GetInt("server.port")

	//Connect to DB
	dbConfig := db.Config{
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
		viper.GetString("db.db-name"),
		viper.GetString("db.db-args"),
	}
	dbConn, err := db.Connect(dbConfig)
	if err != nil {
		panic(err)
	}
	authorRepository := repository.NewPsqlAuthorRepository(dbConn)
	articleRepository := repository.NewPsqlArticleRepository(dbConn)

	//Migration
	migrationErr := db.Migrate(dbConn.DB)
	if migrationErr != nil {
		panic(migrationErr)
	}

	//GraphQL configuration
	graphQLConfig := gql.NewConfig(&authorRepository, &articleRepository)
	graphQLHandlerFunc := middleware.NewGraphQLHandlerFunc(graphQLConfig)

	//HTTP router
	router := middleware.NewRouter(graphQLHandlerFunc)

	//HTTP handler
	handler := middleware.NewCorsHandler(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: handler,
	}

	log.Printf("Visit %s to interact with the GraphQL Playground!", "http://localhost:"+strconv.Itoa(serverPort)+"/")
	log.Printf("After gql/schemas/* schemas update, we need to regenerate the gqlgen code via cli: '%s' ", "go run github.com/99designs/gqlgen")

	log.Fatal(srv.ListenAndServe())
}
