package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetHttpTransportConfig() *http.Client {
	transport := &http.Transport{
		IdleConnTimeout:   50 * time.Second,
		DisableKeepAlives: false,
	}
	// Enable HTTP/2
	return &http.Client{
		Transport: transport,
		Timeout:   25 * time.Second,
	}
}
func GetHttpsTransportConfig() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ForceAttemptHTTP2: true,
	}
	return &http.Client{Transport: t, Timeout: 25 * time.Second}
}

type APIClient struct {
	httpClient *http.Client
	basePath   string
}

func NewAPIClient(endpoint string, withCerts bool) *APIClient {
	var httpClient *http.Client
	if withCerts {
		httpClient = GetHttpsTransportConfig()
	} else {
		httpClient = GetHttpTransportConfig()
	}

	return &APIClient{
		httpClient: httpClient,
		basePath:   endpoint,
	}
}

func (c *APIClient) GetSimpleQuote(ctx context.Context, req QuoteRequest) (QuoteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return QuoteResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.basePath+"/getSimpleQuote", bytes.NewBuffer(body))
	if err != nil {
		return QuoteResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return QuoteResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return QuoteResponse{}, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var quote QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&quote)
	return quote, err
}

func (c *APIClient) GetTokenPrice(ctx context.Context, req PriceRequest) (PriceResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return PriceResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.basePath+"/getTokenPrice", bytes.NewBuffer(body))
	if err != nil {
		return PriceResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return PriceResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PriceResponse{}, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var quote PriceResponse
	err = json.NewDecoder(resp.Body).Decode(&quote)
	return quote, err
}
