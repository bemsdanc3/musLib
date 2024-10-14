package entities

type Track struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	UploadDate  string `json:"upload_date,omitempty"`
	FileLink    string `json:"file_link,omitempty"`
}
