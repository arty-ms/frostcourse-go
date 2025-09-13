package main

var apiURL = "https://api.telegram.org/bot"

// Main Entrypoint for DO function
func Main(req RawRequest) (*Response, error) {
	ctx := getInitialContext()

	// Add botToken to ApiUrl
	apiURL += ctx.BotToken

	handler := Chain(
		MainLogic,
		AuthenticateMW(),
		ParseTGMessageMW(),
		AuthorizeMW(),
	)

	return handler(ctx, req)
}
