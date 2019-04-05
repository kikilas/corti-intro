# Example GraphQL API written in Go language
Start the database:

```
$ docker-compose up -d
```
Start the app:

```
$ go run main.go
```

Visit http://localhost:8085/ to interact with the GraphQL Playground!

After modifying gql/schemas/* , we need to regenerate the gqlgen code via cli:

```
$ go run github.com/99designs/gqlgen
```
More info on https://gqlgen.com/getting-started/
