package main

import (
	"encoding/json"
)

type Request struct {
	Body json.RawMessage `json:"body,omitempty"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	// TODO: твоя логика (/preview, /publish и т.д.)
	return &Response{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}
