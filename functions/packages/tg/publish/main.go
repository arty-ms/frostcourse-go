package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var (
	botToken      = os.Getenv("BOT_TOKEN")
	channelIDStr  = os.Getenv("CHANNEL_ID")
	apiURL        = "https://api.telegram.org/bot" + botToken
	webhookSecret = os.Getenv("WEBHOOK_SECRET")
	ownerIDs      = parseOwnerIDs(os.Getenv("OWNER_IDS"))
)

// Main Entrypoint for DO function
func Main(req RawRequest) (*Response, error) {
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
	message := strings.TrimSpace(upd.Message.Text)

	if !ownerIDs[userID] {
		_ = sendMessage(chatID, "⛔ Нет доступа")
		return &Response{StatusCode: 200, Body: "forbidden"}, nil
	}

	ctx := &Ctx{
		BotToken:     botToken,
		BotChannelID: chatID,
		PubChannelID: channelIDStr,
		ApiUrl:       apiURL,
	}

	err = processMessage(ctx, message)

	if err != nil {
		log.Printf("Error while processing message: %s", err)

		return &Response{StatusCode: 500, Body: "Internal Error"}, nil
	}

	return &Response{StatusCode: 200, Body: "ok"}, nil
}
