package spotify

import (
	"time"

	"github.com/azybk/music-catalog/internal/configs"
	"github.com/azybk/music-catalog/pkg/httpclient"
)

type outbound struct {
	cfg *configs.Config
	client httpclient.HTTPClient
	AccessToken string
	TokenType string
	ExpiredAt time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, client httpclient.HTTPClient) *outbound {
	return &outbound{
		cfg: cfg,
		client: client,
	}
}