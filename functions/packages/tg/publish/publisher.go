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

type MessageSenderHandler interface {
	sendMessage(ctx *Ctx, messageBytes *bytes.Reader) error
	postMessage(ctx *Ctx, channelId int64, messageBytes *bytes.Reader) error
	preparePayload(ctx *Ctx, channelId int64, messageText string) ([]byte, error)
}

type Publisher struct{}

func (p *Publisher) sendMessage(ctx *Ctx, messageBytes []byte) error {
	resp, err := http.Post(ctx.APIUrl+"/sendMessage", "application/json", bytes.NewReader(messageBytes))

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
	if ctx.WebAppURL == "" {
		return nil, fmt.Errorf("web app url is empty")
	}

	markup := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{
					Text:   "Подробнее",
					WebApp: &WebAppInfo{Url: ctx.WebAppURL},
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
