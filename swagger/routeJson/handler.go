package routeJson

import (
	"encoding/json"
	"go-todo-api/pkg/config"
	"go-todo-api/swagger/routeJson/paths"
	"maps"
	"net/http"
)

func mergePaths(pathsList ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, p := range pathsList {
		maps.Copy(merged, p)
	}
	return merged
}

func SwaggerJSONHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.LoadConfig()

	swagger := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "Go Todo API",
			"description": "This is a sample Todo API built with Go + Chi",
			"version":     "1.0.0",
		},
		"servers": []map[string]string{
			{"url": "http://localhost:" + cfg.AppPort},
		},
		"paths":      mergePaths(paths.TodoPaths),
		"components": SwaggerComponents(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(swagger)
}
