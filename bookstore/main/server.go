package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gql/bookstore"
	"net/http"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: bookstore.RootQuery,
	})
	if err != nil {
		panic(err)
	}
	hand := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", hand)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
