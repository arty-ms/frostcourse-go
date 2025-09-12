package main

import "strings"

func parseCommand(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || s[0] != '/' {
		return ""
	}

	s = s[1:]

	for i, r := range s {
		if r == ' ' {
			return s[:i]
		}
	}
	return s
}

func processMessage(ctx *Ctx, message string) error {
	cmd := parseCommand(message)

	currentCommand, ok := commands[Command(cmd)]

	if !ok {
		return sendMessage(ctx.BotChannelID, "ℹ️ Доступные команды: /start, /preview, /publish")
	}

	return currentCommand.Run(ctx, message)
}
