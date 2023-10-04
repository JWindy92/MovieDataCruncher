package ingestion

import (
	"encoding/json"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/models"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb

func PerformSearch(s string) string {
	//TODO: currently using specific test case where only one result is returned, need to be more flexible
	tmdb_config := tmdb.Config{
		APIKey:   config.Config.Tmdb.ApiKey,
		Proxies:  nil,
		UseProxy: false,
	}
	tmdbAPI = tmdb.Init(tmdb_config)

	searchQuery := models.SearchQuery{}
	err := json.Unmarshal([]byte(s), &searchQuery)
	if err != nil {
		panic(err)
	}
	tmdbAPI = tmdb.Init(tmdb_config)

	options := searchQuery.Options.ToMap()

	searchResult, err := tmdbAPI.SearchMovie(searchQuery.Title, options)
	if err != nil {
		panic(err)
	}

	jsonData, err := tmdb.ToJSON(searchResult)

	return jsonData

}
