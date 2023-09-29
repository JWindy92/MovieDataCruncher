package ingestion

import (
	"encoding/json"
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/models"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb

func PerformSearch(s string) {
	tmdb_config := tmdb.Config{
		APIKey:   config.Config.Tmdb.ApiKey,
		Proxies:  nil,
		UseProxy: false,
	}
	searchQuery := models.SearchQuery{}
	err := json.Unmarshal([]byte(s), &searchQuery)
	if err != nil {
		panic(err)
	}
	tmdbAPI = tmdb.Init(tmdb_config)

	api_key := config.Config.Tmdb.ApiKey
	log.Printf("TMDB API KEY IS: %s", api_key)

	// options := map[string]string{
	// 	"primary_release_year": "2002",
	// }
	options := searchQuery.Options.ToMap()
	log.Printf("Options: %s", options)
	// fightClubInfo, err := tmdbAPI.GetMovieInfo(550, nil)
	searchResult, err := tmdbAPI.SearchMovie(searchQuery.Title, options)
	if err != nil {
		panic(err)
	}

	gangsJson, err := tmdb.ToJSON(searchResult)

	log.Printf("Got Fight Club INfo: %s", gangsJson)

}
