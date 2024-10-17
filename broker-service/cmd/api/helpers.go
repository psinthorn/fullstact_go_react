package main

type jsonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   bool   `json:"error"`
}

// readJson

// writeJson

// readError
