package models

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
}
