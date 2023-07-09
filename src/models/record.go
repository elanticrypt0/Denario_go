package models

import "github.com/elanticrypt0/denario_go/src/webcore"

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

// Find a record by ID
// Return the record
// Return an error if there was a problem
func FindOneRecord(id int) Record {
	feature := webcore.NewFeature()
	var record Record
	feature.Db.First(&record, id)
	return record
}

// Find all records
// Return the records
// Return an error if there was a problem
func FindAllRecords() []Record {
	feature := webcore.NewFeature()
	var records []Record
	feature.Db.Find(&records)
	return records
}

// Create a new record
// Return the new record
// Return an error if there was a problem
func CreateRecord(name string, amount float64, amount_io string, comment string, record_date string, category_id int, is_mutable bool) Record {
	feature := webcore.NewFeature()
	record := Record{
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
func UpdateRecord(record Record) Record {
	feature := webcore.NewFeature()
	feature.Db.Save(&record)
	return record
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(id int) Record {
	feature := webcore.NewFeature()
	var record Record
	feature.Db.First(&record, id)
	record.IsDeleted = true
	feature.Db.Save(&record)
	return record
}
