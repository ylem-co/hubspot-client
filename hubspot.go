package hubspotclient

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"net/http"
)

var cfg *Config

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string
}

type Hubspot struct {
	Client *resty.Client
}

func (h *Hubspot) CreateTicket(request CreateTicketRequest) error {
	resp, err := h.Client.R().
		SetBody(request).
		Post("/crm/v3/objects/tickets")

	if err != nil {
		return fmt.Errorf("hubspot: %s", err.Error())
	}

	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf(
			"hubspot: Expected: %d http response got %s. Body: %s",
			http.StatusCreated,
			resp.Status(),
			string(resp.Body()),
		)
	}

	return nil
}

func (h *Hubspot) Token() (*oauth2.Token, error) {
	httpClient := h.Client.GetClient()
	transport := httpClient.Transport.(*oauth2.Transport)

	token, err := transport.Source.Token()

	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateInstance(ctx context.Context, token *oauth2.Token) (*Hubspot, error) {
	conf := createConfig()
	if conf == nil {
		return nil, fmt.Errorf("hubspot: please inititate the Config")
	}

	httpClient := conf.Client(ctx, token)

	return &Hubspot{
		Client: resty.NewWithClient(httpClient).SetBaseURL("https://api.hubapi.com"),
	}, nil
}

func CreateGrantLink(State string) (string, error) {
	conf := createConfig()
	if conf == nil {
		return "", fmt.Errorf("hubspot: please inititate the Config")
	}

	return conf.AuthCodeURL(State), nil
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	conf := createConfig()
	if conf == nil {
		return nil, fmt.Errorf("hubspot: please inititate the Config")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func Initiate(config Config) {
	cfg = &config
}

func createConfig() *oauth2.Config {
	if cfg == nil {
		return nil
	}

	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL:  "https://api.hubspot.com/oauth/v1/token",
			AuthURL:   "https://app.hubspot.com/oauth/authorize",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: cfg.RedirectUrl,
	}
}
