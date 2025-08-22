package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-todo-api/pkg/response"

	"github.com/go-chi/chi/v5"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get the list of todos
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, GetAll())
}

// GetTodoByID godoc
// @Summary Get a todo by ID
// @Description Get a todo item by its ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	t, found := GetByID(id)
	if !found {
		response.JSON(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
		return
	}
	response.JSON(w, http.StatusOK, t)
}

// CreateTodoHandler godoc
// @Summary Create a new todo
// @Description Add a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "Todo item"
// @Success 201 {object} Todo
// @Router /todos [post]
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)
	newTodo := Create(t.Title)
	response.JSON(w, http.StatusCreated, newTodo)
}

// UpdateTodoHandler godoc
// @Summary Update a todo
// @Description Update an existing todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body Todo true "Updated Todo"
// @Success 200 {object} Todo
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)
	updated, ok := Update(id, t)
	if !ok {
		response.JSON(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
		return
	}
	response.JSON(w, http.StatusOK, updated)
}

// DeleteTodoHandler godoc
// @Summary Delete a todo
// @Description Delete a todo by ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if !Delete(id) {
		response.JSON(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}
