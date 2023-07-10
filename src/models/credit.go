package models

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/src/webcore"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type Credit struct {
	gorm.Model
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	Amount     float64  `json:"amount"`
	Payments   int      `json:"payments"`
	StartedAt  string   `json:"started_at"`
	FinishedAt string   `json:"finished_at"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
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

func CreateCredit(name string, comment string, amount float64, payments int, started_at string, finished_at string, category_id int) Credit {
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
	CreateRecordsForCredit(credit)
	return credit
}

func UpdateCredit(credit Credit) Credit {
	feature := webcore.NewFeature()
	feature.Db.Save(&credit)
	deleteAllRecordsOfThisCredit(int(credit.ID))
	CreateRecordsForCredit(credit)
	return credit
}

func DeleteCredit(id int) Credit {
	feature := webcore.NewFeature()
	var credit Credit
	feature.Db.First(&credit, id)
	feature.Db.Delete(&credit)
	// also remote all records associated with this credit
	deleteAllRecordsOfThisCredit(id)
	return credit
}

// CreateRecordsForCredit
// This function creates the records for a credit
// It is called from the controller
func CreateRecordsForCredit(credit Credit) {
	// default values for the "credit" records
	id := strconv.Itoa(int(credit.ID))
	comment := "[CREDIT] " + id
	amount_io := "out"
	base_date := carbon.Parse(credit.StartedAt)

	// Por cada uno de los payments, crear un record
	for i := 0; i < (credit.Payments); i++ {
		// primer dia del mes
		record_date := base_date.AddMonths(i).StartOfMonth().StartOfDay().ToDateString()
		CreateRecord(credit.Name, credit.Amount, amount_io, comment, record_date, credit.CategoryID, false)
	}
}

// DeleteAllRecordsOfThisCredit
// This function deletes all the records associated with a credit
// It is called from the controller
func deleteAllRecordsOfThisCredit(credit_id int) {
	feature := webcore.NewFeature()
	var records []Record
	feature.Db.Where("comment = '[CREDIT] ?' and is_mutable = ?", credit_id, false).Find(&records)
	// for _, record := range records {
	// 	feature.Db.Delete(&record)
	// }
	feature.Db.Delete(&records)
}
