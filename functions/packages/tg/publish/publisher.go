package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebAppInfo struct {
	Url string `json:"url"`
}

type InlineKeyboardButton struct {
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"` // кнопка-мини-аппа
	Url    string      `json:"url,omitempty"`     // обычная ссылка (не нужно для мини-аппы)
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type TgSendMessage struct {
	ChatID                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type TgResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	Parameters  *struct {
		RetryAfter      int   `json:"retry_after,omitempty"`
		MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	} `json:"parameters,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type MessageSenderHandler interface {
	sendMessage(ctx *Ctx, messageBytes *bytes.Reader) error
	postMessage(ctx *Ctx, channelId int64, messageBytes *bytes.Reader) error
	preparePayload(ctx *Ctx, channelId int64, messageText string) ([]byte, error)
}

type Publisher struct{}

func (p *Publisher) sendMessage(ctx *Ctx, messageBytes []byte) error {
	req, err := http.NewRequest("POST", ctx.APIUrl+"/sendMessage", bytes.NewReader(messageBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ctx.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var tr TgResponse
		_ = json.Unmarshal(respBody, &tr)
		desc := tr.Description
		if desc == "" {
			desc = string(respBody)
		}
		if resp.StatusCode == http.StatusTooManyRequests && tr.Parameters != nil && tr.Parameters.RetryAfter > 0 {
			return fmt.Errorf("telegram 429, retry_after=%ds: %s", tr.Parameters.RetryAfter, desc)
		}
		return fmt.Errorf("telegram http %d: %s", resp.StatusCode, desc)
	}

	var tr TgResponse
	if err := json.Unmarshal(respBody, &tr); err != nil {
		return fmt.Errorf("telegram bad json: %w", err)
	}
	if !tr.OK {
		if tr.Parameters != nil && tr.Parameters.RetryAfter > 0 {
			return fmt.Errorf("telegram error (retry_after=%ds): %s", tr.Parameters.RetryAfter, tr.Description)
		}
		return fmt.Errorf("telegram error: %s", tr.Description)
	}
	return nil
}

type TextPublisher struct{ Publisher }

func (p *TextPublisher) postMessage(ctx *Ctx, channelId int64, messageText string) error {
	payload, err := p.preparePayload(ctx, channelId, messageText)

	if err != nil {
		return err
	}

	return p.sendMessage(ctx, payload)
}

func (p *TextPublisher) preparePayload(_ *Ctx, channelId int64, messageText string) ([]byte, error) {
	req := TgSendMessage{
		ChatID:    channelId,
		Text:      messageText,
		ParseMode: "Markdown",
	}

	return json.Marshal(req)
}

type TextWithMarkupPublisher struct{ Publisher }

func (p *TextWithMarkupPublisher) postMessage(ctx *Ctx, channelId int64, messageText string) error {
	payload, err := p.preparePayload(ctx, channelId, messageText)

	if err != nil {
		return err
	}

	return p.sendMessage(ctx, payload)
}

func (p *TextWithMarkupPublisher) preparePayload(ctx *Ctx, channelId int64, messageText string) ([]byte, error) {
	if ctx.WebAppBotName == "" {
		return nil, fmt.Errorf("web app bot name is empty")
	}

	markup := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{
					Text: "Подробнее",
					Url:  fmt.Sprintf("https://t.me/%s?startapp=%s", ctx.WebAppBotName, "payload123"),
				},
			},
		},
	}

	payload := struct {
		ChatID                int64                `json:"chat_id"`
		Text                  string               `json:"text"`
		ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
		ParseMode             string               `json:"parse_mode,omitempty"`
		DisableWebPagePreview bool                 `json:"disable_web_page_preview,omitempty"`
	}{
		ChatID:                channelId,
		Text:                  messageText,
		ReplyMarkup:           markup,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

	return json.Marshal(payload)
}
