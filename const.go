package main

const (
	AppName    = "ngocok"
	AppVersion = "1.0.0"
	AppAuthor  = "dwisiswant0"
)

const (
	infGetAuthtokenMsg = "Get your token"
	infGetAuthtokenRef = "https://dashboard.ngrok.com/get-started/your-authtoken"
	sessOK             = "client session established"
	authFailed         = "authentication failed"
	reqReceived        = "Request received"
)

const (
	banner = `
                       _   
   ___ ___ ___ ___ ___| |_
  |   | . | . |  _| . | '_|
  |_|_|_  |___|___|___|_,_|
      |___|   v` + AppVersion + `
  --
  made with ♥ by @` + AppAuthor + `
`
	usage = `
Usage:
  ngocok [OPTIONS...]

Options:
  -e, --endpoint string     ngrok tunnel endpoint ("http" or "tcp") (default: "http")
  -t, --token string        ngrok authentication token
      --unstrip             Unstrip X-Forwarded-{For,Host,Proto} headers
  -o, --output string       Log incoming requests to a file instead of stdout

Examples:
  ngocok --endpoint tcp
  ngocok --token [authtoken]
  ngocok --output /path/to/requests.log
`
)

const (
	errGetAuthtokenOpt = "cannot get authtoken anywhere"
	errGetTunnelCustom = "cannot get tunnel for '%s' endpoint"
	errGetTunnelEmpty  = "cannot get tunnel for empty endpoint"
	errRuntime         = "runtime error"
)

const (
	ngrokAuthtoken    = "authtoken"
	ngrokAuthtokenEnv = "NGROK_AUTHTOKEN"
)
