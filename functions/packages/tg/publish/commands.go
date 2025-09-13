package main

import "log"

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
type StartCommand struct {
	BaseCommand
	TextPublisher
}

func (c *StartCommand) Run(ctx *Ctx, _ string) error {
	return c.postMessage(ctx, ctx.BotChannelID, "👋 Готов словить твой новый пост")
}

// PreviewCommand implements the /publish command
type PreviewCommand struct {
	BaseCommand
	TextPublisher
}

func (c *PreviewCommand) Run(ctx *Ctx, message string) error {
	body := extractPayload(message)

	if body == "" {
		return c.postMessage(ctx, ctx.BotChannelID, c.WarnText)
	}

	return c.postMessage(ctx, ctx.BotChannelID, c.SuccessText+body)
}

// PublishCommand implements the /publish command
type PublishCommand struct {
	BaseCommand
	TextPublisher
}

func (c *PublishCommand) Run(ctx *Ctx, message string) error {
	body := extractPayload(message)

	if body == "" {
		return c.postMessage(ctx, ctx.BotChannelID, c.WarnText)
	}

	if err := c.postMessage(ctx, ctx.PubChannelID, body); err != nil {
		_ = c.postMessage(ctx, ctx.BotChannelID, c.ErrorText+err.Error())
	}

	return c.postMessage(ctx, ctx.BotChannelID, c.SuccessText)
}

// PublishWithLinkCommand implements the /publish command
type PublishWithLinkCommand struct {
	BaseCommand
	TextWithMarkupPublisher
}

func (c *PublishWithLinkCommand) Run(ctx *Ctx, message string) error {
	body := extractPayload(message)

	if body == "" {
		return c.postMessage(ctx, ctx.BotChannelID, c.WarnText)
	}

	if err := c.postMessage(ctx, ctx.PubChannelID, body); err != nil {
		log.Printf("Error while publishing with inline button: %s", err)

		return c.postMessage(ctx, ctx.BotChannelID, c.ErrorText+err.Error())
	}

	return c.postMessage(ctx, ctx.BotChannelID, c.SuccessText)
}

const (
	CmdStart                   Command = "start"
	CmdPreview                 Command = "preview"
	CmdPublish                 Command = "publish"
	CmdPublishWithInlineButton Command = "publish_with_inline_button"
)

var commands = map[Command]Handler{
	CmdStart: &StartCommand{
		BaseCommand{Syntax: "/start"},
		TextPublisher{},
	},
	CmdPreview: &PreviewCommand{
		BaseCommand{
			Syntax:      "/preview",
			WarnText:    "⚠️ Пришли `/preview текст`",
			SuccessText: "🔎 Preview:\n\n",
		},
		TextPublisher{},
	},
	CmdPublish: &PublishCommand{
		BaseCommand{
			Syntax:      "/publish",
			WarnText:    "⚠️ Пришли `/publish текст`",
			ErrorText:   "❌ Ошибка публикации: ",
			SuccessText: "✅ Опубликовано в канал",
		},
		TextPublisher{},
	},

	// Temporary fields
	CmdPublishWithInlineButton: &PublishWithLinkCommand{
		BaseCommand{
			Syntax:      "/publish_with_inline_button",
			ErrorText:   "❌ Inline Link-a не установлена. Отправка отменена",
			SuccessText: "✅ Опубликовано в канал с прикрепленной Inline Link-ой",
		},
		TextWithMarkupPublisher{},
	},
}

func sendForbiddenMessage(ctx *Ctx) error {
	publisher := TextPublisher{Publisher{}}

	return publisher.postMessage(ctx, ctx.BotChannelID, "⛔ Нет доступа")
}

func sendHelpMessage(ctx *Ctx) error {
	helpText := "ℹ️ Доступные команды:\n" +
		"/start - начать работу с ботом\n" +
		"/preview <текст> - посмотреть, как будет выглядеть пост\n" +
		"/publish <текст> - опубликовать пост в канал" +
		"/publish_with_inline_button <текст> - опубликовать пост в канал c прикрепленной Inline Link-ой"
	publisher := TextPublisher{Publisher{}}

	return publisher.postMessage(ctx, ctx.BotChannelID, helpText)
}
