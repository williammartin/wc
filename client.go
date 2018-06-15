package wc

import (
	"github.com/dghubble/sling"
	"github.com/williammartin/wc/services"
)

//go:generate counterfeiter . MatchService
type MatchService interface {
	Fetch() (services.Matches, error)
}

type Client struct {
	sling *sling.Sling

	MatchService MatchService
}

func NewClient(api string) *Client {
	base := sling.New().Base(api)

	return &Client{
		sling:        base,
		MatchService: services.NewMatchService(base.New()),
	}
}
