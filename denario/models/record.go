package models

type Record struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	AmountIo   string  `json:"amount_io"`
	Comment    string  `json:"comment"`
	RecordDate string  `json:"record_date"`
	CategoryID int     `json:"category_id"`
	IsMutable  bool    `json:"is_mutable"`
	IsDeleted  bool    `json:"is_deleted"`
}
