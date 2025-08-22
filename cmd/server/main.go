package main

import (
	"log"
	"net/http"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"go-todo-api/internal/router"
)

// @title Go Todo API
// @version 1.0
// @description This is a sample Todo API built with Go + Chi
// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}))

	// Routes
	router.SetupRoutes(r)

		// --- Fancy banner & console log ---
	banner := figure.NewFigure("GO TODO API", "", true)
	banner.Print()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	log.Println()
	log.Println(bold("🚀 Go Todo API is running!"))
	log.Println(cyan("👉 Base URL: ") + green("http://localhost:8080"))
	log.Println(cyan("📂 Swagger Docs: ") + bold(green("http://localhost:8080/docs/")))
	log.Println()
	log.Fatal(http.ListenAndServe(":8080", r))
}
