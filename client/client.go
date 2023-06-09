package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

type Client struct {
	APIKey     string
	HttpClient *http.Client
	BaseUrl    string
	Logger     zerolog.Logger
}

func (c *Client) ID() string {
	return "fincode"
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
		Logger:     logger,
		APIKey:     pluginSpec.APIKey,
		BaseUrl:    environments[pluginSpec.Environment],
		HttpClient: retryClient.StandardClient(),
	}, nil
}
