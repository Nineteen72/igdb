package models

// Image entity
type Image struct {
	URL    string `json:"url,omitempty"`
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`
}
