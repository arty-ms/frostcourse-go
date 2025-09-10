package main

import (
	"encoding/json"
	"log/slog"
	_ "log/slog"
	"os"

	slogbetterstack "github.com/samber/slog-betterstack"
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
	want := os.Getenv("WEBHOOK_SECRET")
	got := ""

	logger := slog.New(
		slogbetterstack.Option{
			Token:    "$LOGGER_TOKEN",
			Endpoint: "https://$INGESTING_HOST/",
		}.NewBetterstackHandler(),
	)

	logger.Info("Hello from Function DO!")
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
