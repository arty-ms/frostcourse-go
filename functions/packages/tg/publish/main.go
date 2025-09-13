package main

// Main Entrypoint for DO function
func Main(req RawRequest) (*Response, error) {
	ctx := getInitialContext()

	handler := Chain(
		MainLogic,
		AuthenticateMW(),
		ParseTGMessageMW(),
		AuthorizeMW(),
	)

	return handler(ctx, req)
}
