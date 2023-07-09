package models

import "github.com/elanticrypt0/denario_go/src/webcore"

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

func FindOneCredit(id int) Credit {
	feature := webcore.NewFeature()
	var credit Credit
	feature.Db.First(&credit, id)
	return credit
}

func FindAllCredits() []Credit {
	feature := webcore.NewFeature()
	var credits []Credit
	feature.Db.Find(&credits)
	return credits
}

func CreateCredit(name string, comment string, amount float64, payments int, started_at string, finished_at string, category_id int, is_mutable bool) Credit {
	feature := webcore.NewFeature()
	credit := Credit{
		Name:       name,
		Comment:    comment,
		Amount:     amount,
		Payments:   payments,
		StartedAt:  started_at,
		FinishedAt: finished_at,
		CategoryID: category_id,
	}
	feature.Db.Create(&credit)
	return credit
}

func UpdateCredit(credit Credit) Credit {
	feature := webcore.NewFeature()
	feature.Db.Save(&credit)
	return credit
}

func DeleteCredit(id int) Credit {
	feature := webcore.NewFeature()
	var credit Credit
	feature.Db.First(&credit, id)
	credit.IsDeleted = true
	feature.Db.Save(&credit)
	return credit
}
