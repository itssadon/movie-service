package film

type Film struct {
	Title        string   `json:"title"`
	EpisodeId    string   `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	ReleaseDate  string   `json:"release_date"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	Created      string   `json:"created"`
	Updated      string   `json:"updated"`
	Url          string   `json:"url"`
}
