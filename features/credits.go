package features

import (
	"github.com/elanticrypt0/denario_go/models"
	"github.com/elanticrypt0/denario_go/webcore"
)

func FindOneCredit(id int) models.Credit {
	feature := webcore.NewFeature()
	var credit models.Credit
	feature.Db.First(&credit, id)
	return credit
}

func FindAllCredits() []models.Credit {
	feature := webcore.NewFeature()
	var credits []models.Credit
	feature.Db.Find(&credits)
	return credits
}

func CreateCredit(name string, comment string, amount float64, payments int, started_at string, finished_at string, category_id int, is_mutable bool) models.Credit {
	feature := webcore.NewFeature()
	credit := models.Credit{
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

func UpdateCredit(credit models.Credit) models.Credit {
	feature := webcore.NewFeature()
	feature.Db.Save(&credit)
	return credit
}

func DeleteCredit(id int) models.Credit {
	feature := webcore.NewFeature()
	var credit models.Credit
	feature.Db.First(&credit, id)
	credit.IsDeleted = true
	feature.Db.Save(&credit)
	return credit
}
