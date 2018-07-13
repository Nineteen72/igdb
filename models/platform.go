package models

// Platform data transfer object used by IGDB.
type Platform struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
