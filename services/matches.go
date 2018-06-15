package services

import "github.com/dghubble/sling"

func NewMatchService(sling *sling.Sling) *MatchService {
	return &MatchService{
		sling: sling.Path("matches/"),
	}
}

type MatchService struct {
	sling *sling.Sling
}

func (m *MatchService) Fetch() (Matches, error) {
	matches := new(Matches)

	m.sling.New().Get("").ReceiveSuccess(matches)

	return *matches, nil
}

type Matches []Match

type Match struct {
	Venue string
}
