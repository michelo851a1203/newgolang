package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testapp/pkg/gql"
	"testapp/pkg/gql/resolvers"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// // import "github.com/99designs/gqlgen/handler"
func gqlhandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}
	h := handler.GraphQL(gql.NewExecutableSchema(c))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func backgroundhandlers(path string) gin.HandlerFunc {
	h := handler.Playground("go graphql server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Run <>
func Run() {
	r := gin.Default()
	r.GET("/", backgroundhandlers("/graphql"))
	r.POST("/graphql", gqlhandler())
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("server start error")
		}
	}()
	// server shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server error : %s", err)
	}
}
