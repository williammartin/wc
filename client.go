package wc

import (
	"github.com/dghubble/sling"
)

type Client struct {
	sling *sling.Sling

	MatchService *MatchService
}

func NewClient(api string) *Client {
	base := sling.New().Base(api)

	return &Client{
		sling:        base,
		MatchService: newMatchService(base.New()),
	}
}
