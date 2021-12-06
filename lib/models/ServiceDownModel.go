package model

type ServiceDownload struct {
	ID int `json:"id" db:"id"`
	DownloadURL string `json:"download_url" db:"download_url"`
	Title string `json:"title" db:"title"`
	Version string `json:"version" db:"version"`
	ParentID string `json:"parent_id" db:"parent_id"`
}

