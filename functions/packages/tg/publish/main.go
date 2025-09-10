package main

import (
	"encoding/json"
	"log"
	"os"
)

type Request struct {
	Body      json.RawMessage   `json:"body,omitempty"`
	OwHeaders map[string]string `json:"__ow_headers,omitempty"`
	OwQuery   string            `json:"__ow_query,omitempty"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	log.Println("TATATATATAT")
	want := os.Getenv("WEBHOOK_SECRET")
	got := ""

	if in.OwHeaders != nil {
		if v, ok := in.OwHeaders["x-telegram-bot-api-secret-token"]; ok {
			got = v
		}
	}
	if want == "" || got != want {
		return &Response{StatusCode: 401, Body: "unauthorized"}, nil
	}

	return &Response{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}
