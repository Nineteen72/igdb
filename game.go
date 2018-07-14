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
	fields = "id,name,summary,storyline,updated_at,created_at,collection.slug,collection.url,franchise," +
		"hypes,popularity,rating,aggregated_rating,aggregated_rating_count,total_rating,total_rating_count," +
		"game,cover,screenshots,artworks,time_to_beat,genres,videos,platforms,game_modes,publishers,games"
	expand = "collection,franchise,genres,platforms,game_modes,publishers,games"
)

// GetGame returns a game object
func (c *Client) GetGame(id int64, client *http.Client) (models.Game, error) {
	gameURL := fmt.Sprintf("%s/games/%d", c.URI, id)
	params := url.Values{}
	params.Set("fields", fields)
	params.Set("expand", expand)

	resp, err := services.PerformRequest(c.APIKey, gameURL, params, client)
	if err != nil {
		return models.Game{}, err
	}

	var game []models.Game
	if err = json.Unmarshal(resp, &game); err != nil {
		return models.Game{}, err
	}

	game[0].Videos = parseVideos(game[0].Videos)
	game[0].Cover = parseImage(game[0].Cover)[0]
	game[0].Screenshots = parseImage(game[0].Screenshots...)
	game[0].Artworks = parseImage(game[0].Artworks)[0]

	return game[0], nil
}
