package features

import (
	"github.com/elanticrypt0/denario_go/src/models"
	"github.com/elanticrypt0/denario_go/src/webcore"
)

func FindOneCategory(id int) models.Category {
	feature := webcore.NewFeature()
	var category models.Category
	feature.Db.First(&category, id)
	return category
}

func FindAllCategories() []models.Category {
	feature := webcore.NewFeature()
	var categories []models.Category
	feature.Db.Find(&categories)
	return categories
}

func CreateCategory(name string, comment string) models.Category {
	feature := webcore.NewFeature()
	category := models.Category{
		Name: name,
	}
	feature.Db.Create(&category)
	return category
}

func UpdateCategory(category models.Category) models.Category {
	feature := webcore.NewFeature()
	feature.Db.Save(&category)
	return category
}

func DeleteCategory(id int) models.Category {
	feature := webcore.NewFeature()
	var category models.Category
	feature.Db.First(&category, id)
	category.IsDeleted = true
	feature.Db.Save(&category)
	return category
}
