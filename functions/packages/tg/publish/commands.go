package main

type Command string

type Handler interface {
	Run(ctx *Ctx, raw string) error
}

type BaseCommand struct {
	Syntax      string
	SuccessText string
	WarnText    string
	ErrorText   string
}

// StartCommand implements the /start command
type StartCommand struct{ BaseCommand }

func (c *StartCommand) Run(ctx *Ctx, _ string) error {
	return sendMessage(ctx.BotChannelID, "👋 Готов словить твой новый пост")
}

// PreviewCommand implements the /publish command
type PreviewCommand struct{ BaseCommand }

func (c *PreviewCommand) Run(ctx *Ctx, message string) error {
	body := extractPayload(message)

	if body == "" {
		return sendMessage(ctx.BotChannelID, c.WarnText)
	}

	return sendMessage(ctx.BotChannelID, c.SuccessText+body)
}

// PublishCommand implements the /publish command
type PublishCommand struct{ BaseCommand }

func (c *PublishCommand) Run(ctx *Ctx, message string) error {
	body := extractPayload(message)

	if body == "" {
		return sendMessage(ctx.BotChannelID, c.WarnText)
	}

	if err := sendMessageStr(ctx.PubChannelID, body); err != nil {
		_ = sendMessage(ctx.BotChannelID, c.ErrorText+err.Error())
	}

	return sendMessage(ctx.BotChannelID, c.SuccessText)
}

const (
	CmdStart   Command = "start"
	CmdPreview Command = "preview"
	CmdPublish Command = "publish"
)

var commands = map[Command]Handler{
	CmdStart: &StartCommand{BaseCommand{Syntax: "/start"}},
	CmdPreview: &PreviewCommand{BaseCommand{
		Syntax:      "/preview",
		WarnText:    "⚠️ Пришли `/preview текст`",
		SuccessText: "🔎 Preview:\n\n",
	}},
	CmdPublish: &PublishCommand{BaseCommand{
		Syntax:      "/publish",
		WarnText:    "⚠️ Пришли `/publish текст`",
		ErrorText:   "❌ Ошибка публикации: ",
		SuccessText: "✅ Опубликовано в канал",
	}},
}
