package main

import (
	"encoding/json"
	"log"
	"strings"
)

type WrappedHandler func(ctx *Ctx, req RawRequest) (*Response, error)
type Middleware func(WrappedHandler) WrappedHandler

// Chain applies middlewares to a handler in the order they are provided.
func Chain(h WrappedHandler, ms ...Middleware) WrappedHandler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}
	return h
}

// MainLogic is the main logic for processing messages.
func MainLogic(ctx *Ctx, _ RawRequest) (*Response, error) {
	if err := processMessage(ctx); err != nil {
		log.Printf("Error while processing message: %s", err)

		return &Response{StatusCode: 500, Body: "Internal Error"}, nil
	}

	return &Response{StatusCode: 200, Body: "ok"}, nil
}

// AuthenticateMW checks the secret token in the request headers.
func AuthenticateMW() Middleware {
	return func(next WrappedHandler) WrappedHandler {
		return func(ctx *Ctx, req RawRequest) (*Response, error) {
			providedSecret := header(req.HTTP.Headers, "x-telegram-bot-api-secret-token")

			if providedSecret == "" || providedSecret != ctx.WebhookSecret {
				return &Response{StatusCode: 401, Body: "unauthorized"}, nil
			}

			return next(ctx, req)
		}
	}
}

// AuthorizeMW checks if the user is authorized to use the bot.
func AuthorizeMW() Middleware {
	return func(next WrappedHandler) WrappedHandler {
		return func(ctx *Ctx, req RawRequest) (*Response, error) {
			if !ctx.OwnerIDsMap[ctx.UserID] {
				_ = sendForbiddenMessage(ctx)
				return &Response{StatusCode: 200, Body: "forbidden"}, nil
			}
			return next(ctx, req)
		}
	}
}

// ParseTGMessageMW parses the incoming Telegram message and extracts relevant fields.
func ParseTGMessageMW() Middleware {
	return func(next WrappedHandler) WrappedHandler {
		return func(ctx *Ctx, req RawRequest) (*Response, error) {
			body, err := getRawBody(req)

			if err != nil {
				log.Printf("Error while reading body: %s", err)

				return &Response{StatusCode: 400, Body: "bad body: " + err.Error()}, nil
			}

			var upd TgUpdate

			if err := json.Unmarshal(body, &upd); err != nil {
				log.Printf("Error while parsing json: %s", err)

				return &Response{StatusCode: 400, Body: "bad json: " + err.Error()}, nil
			}

			messageText := strings.TrimSpace(upd.Message.Text)

			if upd.Message == nil || upd.Message.From == nil || messageText == "" {
				log.Println("Skipping non-message or empty update")

				return &Response{StatusCode: 200, Body: "skip"}, nil
			}

			ctx.Message = messageText
			ctx.UserID = upd.Message.From.ID
			ctx.BotChannelID = upd.Message.Chat.ID

			return next(ctx, req)
		}
	}
}
