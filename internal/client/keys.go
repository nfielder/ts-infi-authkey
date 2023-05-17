package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Key struct {
	ID           string          `json:"id"`
	Created      time.Time       `json:"created"`
	Expires      time.Time       `json:"expires"`
	Capabilities KeyCapabilities `json:"capabilities"`
}

type KeyCapabilities struct {
	Devices KeyDeviceCapabilities `json:"devices,omitempty"`
}

type KeyDeviceCapabilities struct {
	Create KeyDeviceCreateCapabilities `json:"create"`
}

type KeyDeviceCreateCapabilities struct {
	Reusable      bool     `json:"reusable"`
	Ephemeral     bool     `json:"ephemeral"`
	Preauthorized bool     `json:"preauthorized"`
	Tags          []string `json:"tags,omitempty"`
}

func (c *Client) CreateKey(ctx context.Context, caps KeyCapabilities) (keySecret string, keyMeta *Key, _ error) {
	return c.CreateKeyWithExpiry(ctx, caps, 0)
}

func (c *Client) CreateKeyWithExpiry(ctx context.Context, caps KeyCapabilities, expiry time.Duration) (keySecret string, keyMeta *Key, _ error) {
	expirySeconds := int64(expiry.Seconds())
	if expirySeconds < 0 {
		return "", nil, fmt.Errorf("expiry must be positive")
	}
	if expirySeconds == 0 && expiry != 0 {
		return "", nil, fmt.Errorf("non-zero expiry must be at least one second")
	}

	keyRequest := struct {
		Capabilities  KeyCapabilities `json:"capabilities"`
		ExpirySeconds int64           `json:"expirySeconds,omitempty"`
	}{caps, int64(expirySeconds)}
	bs, err := json.Marshal(keyRequest)
	if err != nil {
		return "", nil, err
	}

	path := fmt.Sprintf("%s/api/v2/tailnet/%s/keys", c.baseURL(), c.tailnet)
	req, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewReader(bs))
	if err != nil {
		return "", nil, err
	}

	b, resp, err := c.sendRequest(req)
	if err != nil {
		return "", nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", nil, handleErrorResponse(b, resp)
	}

	var key struct {
		Key
		Secret string `json:"key"`
	}
	if err := json.Unmarshal(b, &key); err != nil {
		return "", nil, err
	}
	return key.Secret, &key.Key, nil
}
