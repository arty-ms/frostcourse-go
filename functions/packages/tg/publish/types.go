package main

import "encoding/json"

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

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url,omitempty"`
}

type TgSendMessage struct {
	ChatID                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

// Request and Response from DO middleware
type Request struct {
	Body      json.RawMessage   `json:"body,omitempty"`
	OwHeaders map[string]string `json:"__ow_headers,omitempty"`
	OwQuery   string            `json:"__ow_query,omitempty"`
	OwBody    string            `json:"__ow_body,omitempty"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}
