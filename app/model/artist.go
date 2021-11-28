package model

type Artist struct {
	ID      int64  `json:"id,omitempty" db:"id"`
	Name    string `json:"name" db:"name"`
	Genre   string `json:"genre" db:"genre"`
	Country string `json:"country" db:"country"`
}

type ArtistGet struct {
	Name    string `json:"name"`
	Genre   string `json:"genre"`
	Country string `json:"country"`
}
