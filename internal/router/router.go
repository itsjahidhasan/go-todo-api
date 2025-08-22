package router

import (
	"net/http"

	"go-todo-api/internal/todo"
	"go-todo-api/swagger/routeJson"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	// Root endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo API"))
	})

	// Todos routes
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", todo.GetTodos)
		r.Post("/", todo.CreateTodoHandler)
		r.Get("/{id}", todo.GetTodoByID)
		r.Put("/{id}", todo.UpdateTodoHandler)
		r.Delete("/{id}", todo.DeleteTodoHandler)
	})

	// Serve swagger.json
	r.Get("/swagger.json", routeJson.SwaggerJSONHandler)

	// Serve Swagger UI
	fs := http.FileServer(http.Dir("./swagger/swagger-ui"))
	r.Handle("/docs/*", http.StripPrefix("/docs", fs))
}
