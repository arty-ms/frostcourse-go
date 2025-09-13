package main

import "net/http"

type Ctx struct {
	BotToken      string
	PubChannelID  int64
	UserID        int64
	BotChannelID  int64
	Message       string
	WebhookSecret string
	OwnerIDsMap   map[int64]bool
	WebAppURL     string
	APIUrl        string
	HTTP          *http.Client
}

type TgUpdate struct {
	UpdateID int        `json:"update_id"`
	Message  *TgMessage `json:"message,omitempty"`
}

type TgMessage struct {
	MessageID int        `json:"message_id"`
	Text      string     `json:"text,omitempty"`
	Date      int64      `json:"date"` // unix
	From      *TgUser    `json:"from,omitempty"`
	Chat      TgChat     `json:"chat"`
	Entities  []TgEntity `json:"entities,omitempty"`
	ReplyTo   *TgMessage `json:"reply_to_message,omitempty"`
}

type TgUser struct {
	ID int64 `json:"id"`
}
type TgChat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"` // "private", "group", "supergroup", "channel"
}

type TgEntity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

// Request and Response from DO middleware
type RawRequest struct {
	HTTP struct {
		Body            string            `json:"body"`
		Headers         map[string]string `json:"headers"`
		IsBase64Encoded bool              `json:"isBase64Encoded"`
		Method          string            `json:"method"`
		Path            string            `json:"path"`
		QueryString     string            `json:"queryString"`
	} `json:"http"`

	BotToken      string `json:"BOT_TOKEN"`
	ChannelID     string `json:"CHANNEL_ID"`
	OwnerIDs      string `json:"OWNER_IDS"`
	WebhookSecret string `json:"WEBHOOK_SECRET"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}
