package igdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nineteen72/igdb/services"

	"github.com/nineteen72/igdb/models"
)

const (
	fields = "id,name,summary,storyline,updated_at,created_at,collection.slug,collection.url,franchise,hypes,popularity,rating,aggregated_rating,aggregated_rating_count,total_rating,total_rating_count,game,cover,screenshots,artworks,time_to_beat,genres,videos,platforms,game_modes"
	expand = "collection,franchise,genres,platforms,game_modes"
)

// GetGame returns a game object
func (c *Client) GetGame(id int64) (models.Game, error) {
	gameURL := fmt.Sprintf("%s/games/%d", c.URI, id)
	params := url.Values{}
	params.Set("fields", fields)
	params.Set("expand", expand)

	resp, err := services.PerformRequest(c.APIKey, http.MethodGet, gameURL, params)
	if err != nil {
		return models.Game{}, err
	}

	var game []models.Game
	if err = json.Unmarshal(resp, &game); err != nil {
		return models.Game{}, err
	}

	parseVideos(&game[0].Videos)

	return game[0], nil
}
