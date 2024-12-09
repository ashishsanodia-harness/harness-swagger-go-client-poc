package utils

import (
	"github.com/harness/harness-openapi-go-client/harness"
	"github.com/harness/harness-openapi-go-client/logging"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

var (
	defaultRetryMax = 10
)

func GetDefaultHttpClient(logger *log.Logger) *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.Logger = retryablehttp.LeveledLogger(&logging.Logger{Logger: logger})
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, logger, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = defaultRetryMax
	return httpClient
}
