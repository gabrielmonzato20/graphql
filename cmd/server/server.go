package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabrielmonzato20/graphql/graph"
	"github.com/gabrielmonzato20/graphql/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CategoryDB: categoryDb, CourseDB: courseDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
