package todo

var todos []Todo
var nextID = 1

func GetAll() []Todo {
	result := make([]Todo, len(todos))
	copy(result, todos)
	return result
}

func GetByID(id int) (Todo, bool) {
	for _, t := range todos {
		if t.ID == id {
			return t, true
		}
	}
	return Todo{}, false
}

func Create(title string) Todo {
	t := Todo{ID: nextID, Title: title, Done: false}
	nextID++
	todos = append(todos, t)
	return t
}

func Update(id int, updated Todo) (Todo, bool) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updated.Title
			todos[i].Done = updated.Done
			return todos[i], true
		}
	}
	return Todo{}, false
}

func Delete(id int) bool {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
