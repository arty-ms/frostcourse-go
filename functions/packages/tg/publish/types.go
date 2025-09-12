package main

type Ctx struct {
	BotToken     string
	PubChannelID string
	ApiUrl       string
	BotChannelID int64
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
