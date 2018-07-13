package models

// TimeToBeat entity
type TimeToBeat struct {
	Hastly     uint64 `json:"hastly,omitempty"`
	Normally   uint64 `json:"normally,omitempty"`
	Completely uint64 `json:"completely,omitempty"`
}
