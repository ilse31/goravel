package models

type Animes struct {
	Data       []Datum    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Datum struct {
	MalID          int64            `json:"mal_id"`
	URL            string           `json:"url"`
	Images         map[string]Image `json:"images"`
	Trailer        Trailer          `json:"trailer"`
	Approved       bool             `json:"approved"`
	Titles         []Title          `json:"titles"`
	Title          string           `json:"title"`
	TitleEnglish   string           `json:"title_english"`
	TitleJapanese  string           `json:"title_japanese"`
	TitleSynonyms  []string         `json:"title_synonyms"`
	Type           string           `json:"type"`
	Source         string           `json:"source"`
	Episodes       int64            `json:"episodes"`
	Status         string           `json:"status"`
	Airing         bool             `json:"airing"`
	Aired          Aired            `json:"aired"`
	Duration       string           `json:"duration"`
	Rating         string           `json:"rating"`
	Score          int64            `json:"score"`
	ScoredBy       int64            `json:"scored_by"`
	Rank           int64            `json:"rank"`
	Popularity     int64            `json:"popularity"`
	Members        int64            `json:"members"`
	Favorites      int64            `json:"favorites"`
	Synopsis       string           `json:"synopsis"`
	Background     string           `json:"background"`
	Season         string           `json:"season"`
	Year           int64            `json:"year"`
	Broadcast      Broadcast        `json:"broadcast"`
	Producers      []Demographic    `json:"producers"`
	Licensors      []Demographic    `json:"licensors"`
	Studios        []Demographic    `json:"studios"`
	Genres         []Demographic    `json:"genres"`
	ExplicitGenres []Demographic    `json:"explicit_genres"`
	Themes         []Demographic    `json:"themes"`
	Demographics   []Demographic    `json:"demographics"`
}

type Aired struct {
	From string `json:"from"`
	To   string `json:"to"`
	Prop Prop   `json:"prop"`
}

type Prop struct {
	From   From   `json:"from"`
	To     From   `json:"to"`
	String string `json:"string"`
}

type From struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

type Broadcast struct {
	Day      string `json:"day"`
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
	String   string `json:"string"`
}

type Demographic struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Image struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}

type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Trailer struct {
	YoutubeID string `json:"youtube_id"`
	URL       string `json:"url"`
	EmbedURL  string `json:"embed_url"`
}

type Pagination struct {
	LastVisiblePage int64 `json:"last_visible_page"`
	HasNextPage     bool  `json:"has_next_page"`
	Items           Items `json:"items"`
}

type Items struct {
	Count   int64 `json:"count"`
	Total   int64 `json:"total"`
	PerPage int64 `json:"per_page"`
}
