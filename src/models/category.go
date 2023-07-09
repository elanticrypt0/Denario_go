package models

import (
	"github.com/elanticrypt0/denario_go/src/webcore"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func FindOneCategory(id int) Category {
	feature := webcore.NewFeature()
	var category Category
	feature.Db.First(&category, id)
	return category
}

func FindAllCategories() []Category {
	feature := webcore.NewFeature()
	var categories []Category
	// TODO revisar esta parte
	feature.Db.Order("created_at ASC").Find(&categories)
	return categories
}

func CreateCategory(name string) Category {
	feature := webcore.NewFeature()
	category := Category{
		Name: name,
	}
	feature.Db.Create(&category)
	return category
}

func UpdateCategory(category Category) Category {
	feature := webcore.NewFeature()
	feature.Db.Save(&category)
	return category
}

func DeleteCategory(id int) Category {
	feature := webcore.NewFeature()
	var category Category
	feature.Db.First(&category, id)
	feature.Db.Delete(&category)
	return category
}
