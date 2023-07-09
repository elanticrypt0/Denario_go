package models

import (
	"github.com/elanticrypt0/denario_go/src/webcore"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Name       string   `json:"name"`
	Amount     float64  `json:"amount"`
	AmountIo   string   `json:"amount_io"`
	Comment    string   `json:"comment"`
	RecordDate string   `json:"record_date"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	IsMutable  bool     `json:"is_mutable"`
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
	feature.Db.Order("record_date ASC").Find(&records)
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

// Update a record if it is mutable
// Return the updated record
// Return an error if there was a problem
// TODO funciona pero devuelve el record actualizado. Deberìa mostrar el record viejo
func UpdateRecord(new_record Record) Record {
	feature := webcore.NewFeature()
	current_record := new(Record)
	feature.Db.First(&current_record, new_record.ID)
	if current_record.IsMutable == false {
		return *current_record
	}
	feature.Db.Save(&new_record)
	return new_record
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(id int) Record {
	feature := webcore.NewFeature()
	var record Record
	feature.Db.First(&record, id)
	feature.Db.Delete(&record)
	return record
}
