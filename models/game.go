package models

// GameResponse entity.
type Game struct {
	ID                    uint64     `json:"id,omitempty"`
	Name                  string     `json:"name,omitempty"`
	Summary               string     `json:"summary,omitempty"`
	UpdatedAt             int        `json:"updated_at,omitempty"`
	CreatedAt             int        `json:"created_at,omitempty"`
	Storyline             string     `json:"storyline,omitempty"`
	Collection            Collection `json:"collection,omitempty"`
	Franchise             Franchise  `json:"franchise,omitempty"`
	Hypes                 int        `json:"hypes,omitempty"`
	Popularity            float64    `json:"popularity,omitempty"`
	Rating                float64    `json:"rating,omitempty"`
	RatingCount           int        `json:"rating_count,omitempty"`
	AggregatedRating      float64    `json:"aggregated_rating,omitempty"`
	AggregatedRatingCount float64    `json:"aggregated_rating_count,omitempty"`
	TotalRating           float64    `json:"total_rating,omitempty"`
	TotalRatingCount      int        `json:"total_rating_count,omitempty"`
	Game                  int        `json:"game,omitempty"`
	Cover                 Image      `json:"cover,omitempty"`
	Screenshots           []Image    `json:"screenshots,omitempty"`
	Artworks              Image      `json:"artwork,omitempty"`
	TimeToBeat            TimeToBeat `json:"time_to_beat,omitempty"`
	Genres                []Genre    `json:"genres,omitempty"`
	Videos                []Video    `json:"videos,omitempty"`
	Platforms             []Platform `json:"platforms,omitempty"`
	GameModes             []GameMode `json:"game_modes,omitempty"`
}
