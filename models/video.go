package models

// Video entity
type Video struct {
	Name      string `json:"name,omitempty"`
	VideoID   string `json:"video_id,omitempty"`
	URL       string `json:"url,omitempty"`
	Extension string `json:"extension,omitempty"`
}
