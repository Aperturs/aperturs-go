package main

import (
	"apertursGin/database"
	"apertursGin/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"
func main() {
	err := godotenv.Load()

	if err!=nil{
		log.Fatal("Error Loading .env file")
	}
	
	port := os.Getenv("PORT")
	log.Println(port," This is the Port")
	if port == "" {
		port = defaultPort
	}
	db := database.Connect();
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB:db,
	}}))
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
