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
	"go-todo-api/pkg/config"
)

// @title Go Todo API
// @version 1.0
// @description This is a sample Todo API built with Go + Chi
// @host localhost:'env APP_PORT'
// @BasePath /
func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize the router
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
	red := color.New(color.FgRed).SprintFunc()

	yellow := color.New(color.FgYellow).SprintFunc()

	var envColor func(a ...interface{}) string
	switch cfg.AppEnv {
	case "production":
		envColor = red
	case "staging":
		envColor = yellow
	default:
		envColor = green
	}

	log.Println()
	log.Println(bold("🚀 Go Todo API is running!"))

	log.Println(cyan("🌍 Environment: ") + envColor(cfg.AppEnv))
	log.Println(cyan("👉 Base URL: ") + green(`http://localhost:`+cfg.AppPort+`/`))
	log.Println(cyan("📂 Swagger Docs: ") + bold(green(`http://localhost:`+cfg.AppPort+`/docs/`)))
	log.Println()
	log.Fatal(http.ListenAndServe(`:`+cfg.AppPort+``, r))
}
