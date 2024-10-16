package entities

type Playlist struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	UserID       int    `json:"user_id,omitempty"`
	Creationdate string `json:"creationdate,omitempty"`
}
