// Package attio is a client for the Attio CRM REST API (https://docs.attio.com/rest-api),
// generated from Attio's public OpenAPI spec.
//
// Start with [NewClient]; each API group is a field of [APIClient], such as
// RecordsAPI and ListsAPI.
package attio

// This file is written by hand; every other .go file in this directory is generated.

// NewClient returns a client that authenticates every request with an Attio API key
// or OAuth access token. To pick the token per request instead, build the client with
// NewAPIClient and pass an oauth2.TokenSource in the request context under ContextOAuth2.
func NewClient(apiKey string) *APIClient {
	cfg := NewConfiguration()
	cfg.UserAgent = "libattio-go"
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	return NewAPIClient(cfg)
}
