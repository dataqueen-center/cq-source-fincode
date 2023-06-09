package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

type Client struct {
	APIKey  string
	client  *http.Client
	baseUrl string
	Logger  zerolog.Logger
}

func (c *Client) ID() string {
	return "fincode"
}

func (c *Client) Execute(method, endpoint string) ([]byte, int, error) {
	const payType = "Card"
	const processDateFrom = "2023/01/01"
	URL := c.baseUrl + endpoint
	req, err := http.NewRequest(
		method,
		URL,
		nil,
	)
	if err != nil {
		return nil, -1, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	params := req.URL.Query()
	params.Add("pay_type", payType)
	params.Add("process_date_from", processDateFrom)
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	return bodyBytes, resp.StatusCode, nil
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	logger.Info().Msg("Initializing client")

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	pluginSpec.SetDefaults()
	if err := pluginSpec.Validate(); err != nil {
		return nil, err
	}

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	return &Client{
		Logger:  logger,
		APIKey:  pluginSpec.APIKey,
		baseUrl: environments[pluginSpec.Environment],
		client:  retryClient.StandardClient(),
	}, nil
}
