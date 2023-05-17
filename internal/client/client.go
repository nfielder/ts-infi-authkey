package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultAPIBase = "https://api.tailscale.com"

// Client makes API calls to the Tailscale control plane API server
//
// Use NewClient to instantiate one. Exported fields should be set before
// the client is used and not changed thereafter.
type Client struct {
	// tailnet is the globally unique identified for a Tailscale network, such
	// as "example.com" or "user@gmail.com".
	tailnet string

	// BaseURL optionally specifies and alternate API server to use.
	// If empty, "https://api.tailscale.com" is used.
	BaseURL string

	// HTTPCLient optionally specifies an alternate HTTP client to use.
	// If nil http.DefaultClient is used.
	HTTPClient *http.Client
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}

func (c *Client) baseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return defaultAPIBase
}

func NewClient(tailnet string) *Client {
	return &Client{
		tailnet: tailnet,
	}
}

func (c *Client) Tailnet() string { return c.tailnet }

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient().Do(req)
}

func (c *Client) sendRequest(req *http.Request) ([]byte, *http.Response, error) {
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	return b, resp, err
}

type ErrResponse struct {
	Status  int
	Message string
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("Status: %d, Message: %q", e.Status, e.Message)
}

func handleErrorResponse(b []byte, resp *http.Response) error {
	var errResp ErrResponse
	if err := json.Unmarshal(b, &errResp); err != nil {
		return err
	}
	errResp.Status = resp.StatusCode
	return errResp
}
