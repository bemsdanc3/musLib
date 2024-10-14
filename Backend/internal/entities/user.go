package entities

type User struct {
	ID    int    `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Email string `json:"email,omitempty"`
	Pass  string `json:"pass,omitempty"`
}
