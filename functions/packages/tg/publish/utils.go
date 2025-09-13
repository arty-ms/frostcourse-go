package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	pubChannelIdStr := os.Getenv("CHANNEL_ID")
	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	ownersIds := os.Getenv("OWNER_IDS")
	WebAppBotName := os.Getenv("WEB_APP_BOT_NAME")

	if botToken == "" {
		log.Panic("Error: BOT_TOKEN is not set")
	}

	if pubChannelIdStr == "" {
		log.Panic("Error: CHANNEL_ID is not set")
	}
	if webhookSecret == "" {
		log.Panic("Error: WEBHOOK_SECRET is not set")
	}

	if ownersIds == "" {
		log.Panic("Error: OWNER_IDS is not set")
	}

	if WebAppBotName == "" {
		log.Panic("Error: WEB_APP_BOT_NAME is not set")
	}

	pubChannelId, err := strconv.ParseInt(pubChannelIdStr, 10, 64)
	if err != nil {
		log.Panicf("Error: invalid CHANNEL_ID: %v", err)
	}

	ownersIdsMap := parseOwnerIDs(ownersIds)

	return &Ctx{
		BotToken:      botToken,
		PubChannelID:  pubChannelId,
		WebhookSecret: webhookSecret,
		OwnerIDsMap:   ownersIdsMap,
		WebAppBotName: WebAppBotName,
		APIUrl:        "https://api.telegram.org/bot" + botToken,
		HTTP:          &http.Client{Timeout: 10 * time.Second},
	}
}
