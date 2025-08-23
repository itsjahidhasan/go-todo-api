package paths

var TodoPaths = map[string]interface{}{
	"/todos": map[string]interface{}{
		"get": map[string]interface{}{
			"summary": "Get all todos",
			"responses": map[string]interface{}{
				"200": map[string]interface{}{
					"description": "List of todos",
					"content": map[string]interface{}{
						"application/json": map[string]interface{}{
							"schema": map[string]interface{}{
								"type":  "array",
								"items": map[string]interface{}{"$ref": "#/components/schemas/Todo"},
							},
						},
					},
				},
			},
		},
		"post": map[string]interface{}{
			"summary": "Create a new todo",
			"requestBody": map[string]interface{}{
				"required": true,
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/Todo",
						},
					},
				},
			},
			"responses": map[string]interface{}{
				"201": map[string]interface{}{
					"description": "Todo created",
				},
			},
		},
	},
	"/todos/{id}": map[string]interface{}{
		"get": map[string]interface{}{
			"summary": "Get todo by ID",
			"parameters": []map[string]interface{}{
				{
					"in":       "path",
					"name":     "id",
					"required": true,
					"schema":   map[string]interface{}{"type": "string"},
				},
			},
			"responses": map[string]interface{}{
				"200": map[string]interface{}{
					"description": "Todo found",
					"content": map[string]interface{}{
						"application/json": map[string]interface{}{
							"schema": map[string]interface{}{
								"$ref": "#/components/schemas/Todo",
							},
						},
					},
				},
				"404": map[string]interface{}{
					"description": "Todo not found",
				},
			},
		},
		"put": map[string]interface{}{
			"summary": "Update todo",
			"parameters": []map[string]interface{}{
				{
					"in":       "path",
					"name":     "id",
					"required": true,
					"schema":   map[string]interface{}{"type": "string"},
				},
			},
			"requestBody": map[string]interface{}{
				"required": true,
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/Todo",
						},
					},
				},
			},
			"responses": map[string]interface{}{
				"200": map[string]interface{}{
					"description": "Todo updated",
				},
			},
		},
		"delete": map[string]interface{}{
			"summary": "Delete todo",
			"parameters": []map[string]interface{}{
				{
					"in":       "path",
					"name":     "id",
					"required": true,
					"schema":   map[string]interface{}{"type": "string"},
				},
			},
			"responses": map[string]interface{}{
				"204": map[string]interface{}{
					"description": "Todo deleted",
				},
			},
		},
	},
}
