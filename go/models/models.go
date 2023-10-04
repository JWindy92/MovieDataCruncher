package models

type SearchQuery struct {
	Title   string        `json:"title"`
	Options SearchOptions `json:"options"`
}

type SearchOptions struct {
	Year string `json:"year"`
}

func (options SearchOptions) ToMap() map[string]string {
	optionsMap := map[string]string{
		"primary_release_year": options.Year,
	}
	return optionsMap
}

type SearchResult struct {
	Title string
	Id    string
}

type RedisCacheIdMsg struct {
	Id      int    `json:"id"`
	MongoId string `json:"mongoId"`
}
