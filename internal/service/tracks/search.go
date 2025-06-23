package tracks

import (
	"context"
	"log"

	"github.com/azybk/music-catalog/internal/models/spotify"
	spotifyRepo "github.com/azybk/music-catalog/internal/repository/spotify"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Println("error search track to spotify")
		return nil, err
	}
}

func modelToSpotify(data *spotifyRepo.SpotifySearchResponse) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  nil,
		Total:  data.Tracks.Total,
	}
}
