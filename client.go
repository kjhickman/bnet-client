package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	oauth *clientcredentials.Config
}

type Config struct {
	ClientID     string
	ClientSecret string
}

func NewClient(config Config) (*Client, error) {
	oauth := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     "https://us.battle.net/oauth/token",
	}

	return &Client{oauth: oauth}, nil
}

type WoWToken struct {
	Price       int64 `json:"price"`                  // Price in gold
	LastUpdated int64 `json:"last_updated_timestamp"` // When price was last updated (Unix timestamp in ms)
}

func (c *Client) GetWoWToken(ctx context.Context) (*WoWToken, error) {
	httpClient := c.oauth.Client(ctx)
	url := "https://us.api.blizzard.com/data/wow/token/index?namespace=dynamic-us&locale=en_US"

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result WoWToken
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
