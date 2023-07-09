package models

type Credit struct {
	ID         int     `json:"id"`
	Name       string  `json:"description"`
	Comment    string  `json:"comment"`
	Amount     float64 `json:"amount"`
	Payments   int     `json:"payments"`
	StartedAt  string  `json:"started_at"`
	FinishedAt string  `json:"finished_at"`
	CategoryID int     `json:"category_id"`
	IsDeleted  bool    `json:"is_deleted"`
}
