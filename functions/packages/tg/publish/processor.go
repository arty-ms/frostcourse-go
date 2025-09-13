package main

import "strings"

// processMessage processes an incoming message and executes the corresponding command.
func processMessage(ctx *Ctx, message string) error {
	cmd := parseCommand(message)

	currentCommand, ok := commands[Command(cmd)]

	if !ok {
		return sendHelpMessage(ctx)
	}

	return currentCommand.Run(ctx, message)
}

// parseCommand extracts the command from the message text.
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
