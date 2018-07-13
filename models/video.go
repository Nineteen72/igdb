package models

// Video entity
type Video struct {
	Name     string `json:"name,omitempty"`
	VideoURL string `json:"video_id,omitempty"`
	URL      string `json:"url,omitempty"`
}
