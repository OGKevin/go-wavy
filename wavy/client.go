package wavy

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	wavyBaseUrl = "https://wavy.fm/api/v1beta"
)

type Client interface {
	MetricsService() MetricsService
}

type authContext struct {
	bearerToken string
	tokenType   string
	expiresAt   time.Time

	clientID      string
	clienttSecret string
}

type client struct {
	c *http.Client
	*authContext
	logger hclog.Logger

	mService MetricsService
}

func (c *client) MetricsService() MetricsService {
	return c.mService
}

func NewClient(ctx context.Context, logger hclog.Logger, clientID, clientSecret string) Client {
	if logger == nil {
		logger = hclog.New(&hclog.LoggerOptions{
			Name:  "go-wavy",
			Level: hclog.Warn,
		})
	} else {
		logger = logger.Named("go-wavy")
	}

	c := &client{
		logger: logger,
	}

	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     fmt.Sprintf("%s/token", wavyBaseUrl),
		Scopes:       []string{},
		EndpointParams: map[string][]string{
			"": {},
		},
		AuthStyle: oauth2.AuthStyleInHeader,
	}

	httpClient := conf.Client(ctx)
	c.c = httpClient

	c.mService = newMetricsService(c, logger)

	return c
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	c.logger.Trace("processing request", "url", req.URL.String())
	defer c.logger.Trace("finished processing request", "url", req.URL.String())

	url, err := url.Parse(fmt.Sprintf("%s%s", wavyBaseUrl, req.URL.Path))
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse request url: %w", c.logger.Name(), err)
	}
	req.URL = url

	return c.c.Do(req)
}

func (c *client) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}
