package models

type Franchise struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
	URL  string `json:"url,omitempty"`
}
