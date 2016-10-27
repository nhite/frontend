package config

var (
	// FBRedirectURI is the url of the callback
	FBRedirectURI = "http://localhost:8000/#/auth/facebook"
	//FBClientID holds the client ID
	FBClientID = "1751884195071396"
	// FBBackendAuth is the backend url where the code is exchanged for a token
	FBBackendAuth = "http://localhost:3000/auth/facebook"
	// BAckendBaseURL is the backend url
	BackendBaseURL = "http://localhost:3000"
)
