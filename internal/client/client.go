package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetHttpTransportConfig() *http.Client {
	return &http.Client{Transport: &http.Transport{}, Timeout: 25 * time.Second}
}
func GetHttpsTransportConfig(clientCertFile, clientKeyFile, caCertFile string) *http.Client {
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		panic(err)
	}

	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
		},
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
		httpClient = GetHttpsTransportConfig("certs/client.crt", "certs/client.key", "certs/based.crt")
	} else {
		httpClient = GetHttpTransportConfig()
	}
	return &APIClient{
		httpClient: httpClient,
		basePath:   endpoint,
	}
}

func (c *APIClient) GetSimpleQuote(ctx context.Context, req RequestConfig) (QuoteResponse, error) {
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
