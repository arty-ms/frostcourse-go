package main

var apiURL = "https://api.telegram.org/bot"

// Main Entrypoint for DO function
func Main(req RawRequest) (*Response, error) {
	ctx := getInitialContext()

	// Add botToken to ApiUrl
	apiURL += ctx.BotToken

	// Handlers execution direction is from the bottom to top
	handler := Chain(
		MainLogic,
		AuthorizeMW(),
		ParseTGMessageMW(),
		AuthenticateMW(),
	)

	return handler(ctx, req)
}
