package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
}

type SpotifyTrackObject struct {
	// album related fields
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesUrl   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`

	// artists related fields
	ArtistsName []string `json:"artistsName"`

	// track related fields
	Explicit bool   `json:"explicit"`
	ID       string `json:"id"`
	Name     string `json:"name"`
}
