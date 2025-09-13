package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parseOwnerIDs(raw string) map[int64]bool {
	m := map[int64]bool{}
	for _, p := range strings.Split(raw, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if id, err := strconv.ParseInt(p, 10, 64); err == nil {
			m[id] = true
		}
	}
	return m
}

func extractPayload(s string) string {
	// supports /command <payload>
	if i := strings.IndexAny(s, " \n"); i >= 0 && i+1 < len(s) {
		return strings.TrimSpace(s[i+1:])
	}
	return ""
}

func sendMessage(chatID int64, text string) error {
	return sendMessageStr(strconv.FormatInt(chatID, 10), text)
}

func sendMessageStr(chatIDStr string, text string) error {
	id, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("bad chat id")
	}

	req := TgSendMessage{
		ChatID:    id,
		Text:      text,
		ParseMode: "Markdown",
	}
	b, _ := json.Marshal(req)
	resp, err := http.Post(apiURL+"/sendMessage", "application/json", bytes.NewReader(b))

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("telegram status %d", resp.StatusCode)
	}
	return nil
}

func header(h map[string]string, name string) string {
	for k, v := range h {
		if strings.EqualFold(k, name) {
			return v
		}
	}
	return ""
}

func getRawBody(req RawRequest) ([]byte, error) {
	if req.HTTP.Body == "" {
		return nil, fmt.Errorf("empty body")
	}
	if req.HTTP.IsBase64Encoded {
		return base64.StdEncoding.DecodeString(req.HTTP.Body)
	}
	return []byte(req.HTTP.Body), nil
}

func getInitialContext() *Ctx {
	botToken := os.Getenv("BOT_TOKEN")
	pubChannelId := os.Getenv("CHANNEL_ID")
	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	ownersIds := os.Getenv("OWNER_IDS")

	if botToken == "" {
		log.Panic("Error: BOT_TOKEN is not set")
	}

	if pubChannelId == "" {
		log.Panic("Error: CHANNEL_ID is not set")
	}
	if webhookSecret == "" {
		log.Panic("Error: WEBHOOK_SECRET is not set")
	}

	if ownersIds == "" {
		log.Panic("Error: OWNER_IDS is not set")
	}

	ownersIdsMap := parseOwnerIDs(ownersIds)

	return &Ctx{
		BotToken:      botToken,
		PubChannelID:  pubChannelId,
		WebhookSecret: webhookSecret,
		OwnerIDsMap:   ownersIdsMap,
	}
}
