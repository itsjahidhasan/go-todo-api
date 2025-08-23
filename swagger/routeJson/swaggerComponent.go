package routeJson

func SwaggerComponents() map[string]interface{} {
	return map[string]interface{}{
		"schemas": map[string]interface{}{
			"Todo": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type":    "string",
						"example": "1",
					},
					"title": map[string]interface{}{
						"type":    "string",
						"example": "Sample Todo",
					},
					"completed": map[string]interface{}{
						"type":    "boolean",
						"example": false,
					},
				},
			},
		},
	}
}
