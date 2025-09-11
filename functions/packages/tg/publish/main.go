package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

var (
	botToken      = os.Getenv("BOT_TOKEN")
	channelIDStr  = os.Getenv("CHANNEL_ID") // например "-1001234567890"
	apiURL        = "https://api.telegram.org/bot" + botToken
	webhookSecret = os.Getenv("WEBHOOK_SECRET")
	ownerIDs      = parseOwnerIDs(os.Getenv("OWNER_IDS"))
	// anti-replay
	maxSkew = 2 * time.Minute
)

// Main Entrypoint for DO function
func Main(in Request) (*Response, error) {
	providedApiSecret := ""

	if in.OwHeaders != nil {
		if v, ok := in.OwHeaders["x-telegram-bot-api-secret-token"]; ok {
			providedApiSecret = v
		}
	}
	if webhookSecret == "" || providedApiSecret != webhookSecret {
		return &Response{StatusCode: 401, Body: "unauthorized"}, nil
	}

	log.Printf("Request: %s", in)

	raw, err := rawJSON(in)
	if err != nil {
		return &Response{StatusCode: 400, Body: "no body: " + err.Error()}, nil
	}

	log.Printf("Raw: %s", raw)
	var upd TgUpdate
	if err := json.Unmarshal(raw, &upd); err != nil {
		log.Printf("Error found: %s", err)
		return &Response{StatusCode: 400, Body: "bad json: " + err.Error()}, nil
	}

	if upd.Message == nil || upd.Message.From == nil || upd.Message.Text == "" {
		return &Response{StatusCode: 200, Body: "skip"}, nil
	}

	ts := time.Unix(upd.Message.Date, 0)
	if time.Since(ts) > maxSkew || time.Until(ts) > maxSkew {
		return &Response{StatusCode: 200, Body: "stale"}, nil
	}

	if upd.Message.Chat.Type != "private" {
		return &Response{StatusCode: 200, Body: "forbidden"}, nil
	}
	if !ownerIDs[upd.Message.From.ID] {
		_ = sendMessage(upd.Message.Chat.ID, "⛔ Нет доступа.")
		return &Response{StatusCode: 200, Body: "forbidden"}, nil
	}

	text := strings.TrimSpace(upd.Message.Text)
	switch {
	case hasCmd(text, "/start"):
		_ = sendMessage(upd.Message.Chat.ID, "👋 Готов словить твой новый пост")

	case hasCmd(text, "/preview"):
		body := extractPayload(text)
		if body == "" {
			_ = sendMessage(upd.Message.Chat.ID, "⚠️ Пришли `/preview текст`")
			break
		}
		_ = sendMessage(upd.Message.Chat.ID, "🔎 Preview:\n\n"+body)

	case hasCmd(text, "/publish"):
		body := extractPayload(text)
		if body == "" {
			_ = sendMessage(upd.Message.Chat.ID, "⚠️ Пришли `/publish текст`")
			break
		}
		if err := sendMessageStr(channelIDStr, body); err != nil {
			_ = sendMessage(upd.Message.Chat.ID, "❌ Ошибка публикации: "+err.Error())
		} else {
			_ = sendMessage(upd.Message.Chat.ID, "✅ Опубликовано в канал")
		}

	default:
		_ = sendMessage(upd.Message.Chat.ID, "ℹ️ Доступные команды: /start, /preview, /publish")
	}

	return &Response{StatusCode: 200, Body: "ok"}, nil
}
