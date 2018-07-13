package models

// Genre data transfer object used by IGDB.
type Genre struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
