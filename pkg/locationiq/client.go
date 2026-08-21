package locationiq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey  string
	BaseURL string
}

type SearchResult struct {
	PlaceID     string `json:"place_id"`
	Licence     string `json:"licence"`
	OSMType     string `json:"osm_type"`
	OSMID       string `json:"osm_id"`
	BoundingBox []string `json:"boundingbox"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Class       string `json:"class"`
	Type        string `json:"type"`
	Importance  float64 `json:"importance"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://us1.locationiq.com/v1",
	}
}

func (c *Client) Search(query string) ([]SearchResult, error) {
	endpoint := fmt.Sprintf("%s/search", c.BaseURL)
	reqURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("key", c.APIKey)
	params.Add("q", query)
	params.Add("format", "json")

	reqURL.RawQuery = params.Encode()

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var results []SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}
