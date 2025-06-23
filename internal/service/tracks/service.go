package tracks

import (
	"context"

	"github.com/azybk/music-catalog/internal/repository/spotify"
)

type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type service struct {
	spotifyOutbound spotifyOutbound
}

func NewService(spotifyOutbound spotifyOutbound) *service {
	return &service{
		spotifyOutbound: spotifyOutbound,
	}
}
