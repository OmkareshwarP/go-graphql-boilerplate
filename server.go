package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"go-graphql-boilerplate/graph"
	"go-graphql-boilerplate/pkg/utils"

	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err.Error())
	}
	if os.Getenv("SM") != "" {
		var env map[string]string
		json.Unmarshal([]byte(os.Getenv("SM")), &env)
		for k, v := range env {
			os.Setenv(k, v)
		}
	}

	router := chi.NewRouter()
	port := os.Getenv("GO_PORT")
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		debug.PrintStack()
		er := gqlerror.Errorf("Internal server error")
		ctxString := fmt.Sprintf("%v", ctx)
		errorMessage := fmt.Sprintf("%v", err)

		utils.LogError(errorMessage, "internalServerError", 5, er, map[string]interface{}{
			"ctx": ctxString,
		})
		return er
	})

	router.Use(middleware.Recoverer)
	router.Use(utils.AuthMiddleware)
	// health check route
	router.HandleFunc("/health", ApiToGetHealth)

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			playground.Handler("GraphQL playground", "/").ServeHTTP(w, r)
			return
		} else if r.Method == "POST" {
			srv.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("server started at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Printf("Server failed to start: %s\n", err)
	}
}

func ApiToGetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
