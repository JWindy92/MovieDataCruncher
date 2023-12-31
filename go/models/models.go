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
