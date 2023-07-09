package features

import (
	"github.com/elanticrypt0/denario_go/models"
	"github.com/elanticrypt0/denario_go/webcore"
)

// Find a record by ID
// Return the record
// Return an error if there was a problem
func FindOneRecord(id int) models.Record {
	feature := webcore.NewFeature()
	var record models.Record
	feature.Db.First(&record, id)
	return record
}

// Find all records
// Return the records
// Return an error if there was a problem
func FindAllRecords() []models.Record {
	feature := webcore.NewFeature()
	var records []models.Record
	feature.Db.Find(&records)
	return records
}

// Create a new record
// Return the new record
// Return an error if there was a problem
func CreateRecord(name string, amount float64, amount_io string, comment string, record_date string, category_id int, is_mutable bool) models.Record {
	feature := webcore.NewFeature()
	record := models.Record{
		Name:       name,
		Amount:     amount,
		AmountIo:   amount_io,
		Comment:    comment,
		RecordDate: record_date,
		CategoryID: category_id,
		IsMutable:  is_mutable,
	}
	feature.Db.Create(&record)
	return record
}

// Update a record
// Return the updated record
// Return an error if there was a problem
func UpdateRecord(record models.Record) models.Record {
	feature := webcore.NewFeature()
	feature.Db.Save(&record)
	return record
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(id int) models.Record {
	feature := webcore.NewFeature()
	var record models.Record
	feature.Db.First(&record, id)
	record.IsDeleted = true
	feature.Db.Save(&record)
	return record
}
