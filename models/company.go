package models

// Company entity
type Company struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
	URL  string `json:"url,omitempty"`
}
