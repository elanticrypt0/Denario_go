package models

import "github.com/elanticrypt0/denario_go/src/webcore"

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
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
	feature.Db.Find(&categories)
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
	category.IsDeleted = true
	feature.Db.Save(&category)
	return category
}
