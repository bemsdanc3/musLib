package entities

type Likes struct {
	ID       int `json:"id,omitempty"`
	UserID   int `json:"user_id,omitempty"`
	TrackID  int `json:"track_id,omitempty"`
	LikeDate int `json:"like_date,omitempty"`
}
