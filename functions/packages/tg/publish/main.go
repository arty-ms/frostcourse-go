package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var (
	botToken      = os.Getenv("BOT_TOKEN")
	channelIDStr  = os.Getenv("CHANNEL_ID") // например "-1001234567890"
	apiURL        = "https://api.telegram.org/bot" + botToken
	webhookSecret = os.Getenv("WEBHOOK_SECRET")
	ownerIDs      = parseOwnerIDs(os.Getenv("OWNER_IDS"))
	// anti-replay
	//maxSkew = 2 * time.Minute
)

// Main Entrypoint for DO function
func Main(req RawRequest) (*Response, error) {

	log.Printf("botToken: %s", botToken)
	log.Printf("channelIDStr: %s", channelIDStr)
	log.Printf("apiURL: %s", channelIDStr)
	log.Printf("webhookSecret: %s", webhookSecret)
	log.Printf("ownerIDs: %s", ownerIDs)
	providedApiSecret := header(req.HTTP.Headers, "x-telegram-bot-api-secret-token")
	if providedApiSecret == "" || providedApiSecret != webhookSecret {
		return &Response{StatusCode: 401, Body: "unauthorized"}, nil
	}

	log.Printf("Req: %s", req)

	body, err := getRawBody(req)
	if err != nil {
		return &Response{StatusCode: 400, Body: "bad body: " + err.Error()}, nil
	}

	var upd TgUpdate
	if err := json.Unmarshal(body, &upd); err != nil {
		return &Response{StatusCode: 400, Body: "bad json: " + err.Error()}, nil
	}

	if upd.Message == nil || upd.Message.From == nil || upd.Message.Text == "" {
		return &Response{StatusCode: 200, Body: "skip"}, nil
	}

	chatID := upd.Message.Chat.ID
	userID := upd.Message.From.ID
	text := strings.TrimSpace(upd.Message.Text)

	if !ownerIDs[userID] {
		_ = sendMessage(chatID, "⛔ Нет доступа")
		return &Response{StatusCode: 200, Body: "forbidden"}, nil
	}

	//ts := time.Unix(upd.Message.Date, 0)
	//if time.Since(ts) > maxSkew || time.Until(ts) > maxSkew {
	//	return &Response{StatusCode: 200, Body: "stale"}, nil
	//}

	//if upd.Message.Chat.Type != "private" {
	//	return &Response{StatusCode: 200, Body: "forbidden"}, nil
	//}
	if !ownerIDs[upd.Message.From.ID] {
		_ = sendMessage(chatID, "⛔ Нет доступа.")
		return &Response{StatusCode: 200, Body: "forbidden"}, nil
	}

	switch {
	case hasCmd(text, "/start"):
		_ = sendMessage(chatID, "👋 Готов словить твой новый пост")

	case hasCmd(text, "/preview"):
		body := extractPayload(text)
		if body == "" {
			_ = sendMessage(chatID, "⚠️ Пришли `/preview текст`")
			break
		}
		_ = sendMessage(chatID, "🔎 Preview:\n\n"+body)

	case hasCmd(text, "/publish"):
		body := extractPayload(text)
		if body == "" {
			_ = sendMessage(chatID, "⚠️ Пришли `/publish текст`")
			break
		}
		if err := sendMessageStr(channelIDStr, body); err != nil {
			_ = sendMessage(chatID, "❌ Ошибка публикации: "+err.Error())
		} else {
			_ = sendMessage(chatID, "✅ Опубликовано в канал")
		}

	default:
		_ = sendMessage(upd.Message.Chat.ID, "ℹ️ Доступные команды: /start, /preview, /publish")
	}

	return &Response{StatusCode: 200, Body: "ok"}, nil
}
