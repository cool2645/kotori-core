package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Category struct {
	ID            uint       `gorm:"AUTO_INCREMENT" json:"id"`
	Name          string     `gorm:"index" json:"name"`
	Title         string     `json:"title"`
	Description   string     `sql:"type:text;" json:"description"`
	PublicPostNum uint       `gorm:"index" json:"public_post_num"`
	PostNum       uint       `gorm:"index" json:"post_num"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func (Category) TableName() string {
	return config.GlobCfg.TablePrefix + "_categories"
}

var categoryPatchableFields = []string{
	"name",
	"title",
	"description",
}

func ListCategories(db *gorm.DB, showAll bool, page uint, count int, orderBy string, order string) (categories []Category, total uint, err error) {
	q := db.Order(orderBy + " " + order)
	if !showAll {
		q = q.Where("public_post_num > 0")
	}
	if count > 0 {
		err = q.Limit(count).Offset((page - 1) * uint(count)).Find(&categories).Error
	} else {
		err = q.Find(&categories).Error
	}
	if err != nil {
		err = errors.Wrap(err, "ListCategories")
		return
	}
	err = db.Model(&Category{}).Count(&total).Error
	if err != nil {
		err = errors.Wrap(err, "ListCategories")
		return
	}
	return
}

func GetCategoryByID(db *gorm.DB, id uint) (category Category, err error) {
	err = db.Where("id = ?", id).First(&category).Error
	if err != nil {
		err = errors.Wrap(err, "GetCategoryByID")
		return
	}
	return
}

func GetCategoryByName(db *gorm.DB, name string) (category Category, err error) {
	err = db.Where("name = ?", name).First(&category).Error
	if err != nil {
		err = errors.Wrap(err, "GetCategoryByName")
		return
	}
	return
}

func StoreCategory(db *gorm.DB, category *Category) (err error) {
	if category.Name == "" {
		category.Name = category.Title
	}
	err = db.Create(category).Error
	if err != nil {
		err = errors.Wrap(err, "StoreCategory")
		return
	}
	return
}

func UpdateCategory(db *gorm.DB, category *Category, patch map[string]string) (err error) {
	err = db.Model(category).Select(categoryPatchableFields).Updates(patch).Error
	if err != nil {
		err = errors.Wrap(err, "UpdateCategory")
		return
	}
	return
}

func DeleteCategory(db *gorm.DB, id uint) (err error) {
	err = db.Delete(Category{}, "id = ?", id).Error
	if err != nil {
		err = errors.Wrap(err, "DeleteCategory")
		return
	}
	return
}
