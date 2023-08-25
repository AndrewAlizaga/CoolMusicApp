package http

type SpotifyTrackResponse struct {
	//@inject_tag: json:"result"
	Result SpotifyResult
}

type SpotifyResult struct {
	//@inject_tag: json:"tracks"
	Tracks SpotifyTracks
}

type SpotifyTracks struct {
	//@inject_tag: json:"items"
	Items []SpotifyTracksItem
}

type SpotifyTracksItem struct {
	//@inject_tag: json:"artists"
	Artists []SpotifyArtist
	//@inject_tag: json:"popularity"
	Popularity int32
	//@inject_tag: json:"name"
	Name string
	//@inject_tag: json:"album"
	Album SpotifyAlbum
}

type SpotifyArtist struct {
	//@inject_tag: json:"name"
	Name string
}

type SpotifyAlbum struct {
	//@inject_tag: json:"images"
	Images string
}
