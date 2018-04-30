package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Tag struct {
	ID            uint       `gorm:"AUTO_INCREMENT" json:"id"`
	Name          string     `gorm:"index" json:"name"`
	PublicPostNum uint       `gorm:"index" json:"public_post_num"`
	PostNum       uint       `gorm:"index" json:"post_num"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func (Tag) TableName() string {
	return config.GlobCfg.TablePrefix + "_tags"
}

func ListTags(db *gorm.DB, showAll bool, since int, until int, page uint, count int, orderBy string, order string) (tags []Tag, total uint, err error) {
	q := db.Order(orderBy + " " + order)
	if !showAll {
		q = q.Where("public_post_num > 0")
	}
	if since >= 0 {
		q = q.Where("id > ?", since)
	}
	if until >= 0 {
		q = q.Where("id < ?", until)
	}
	if count >= 0 {
		err = q.Limit(count).Offset((page - 1) * uint(count)).Find(&tags).Error
	} else {
		err = q.Find(&tags).Error
	}
	if err != nil {
		err = errors.Wrap(err, "ListTags")
		return
	}
	err = db.Model(&Tag{}).Count(&total).Error
	if err != nil {
		err = errors.Wrap(err, "ListTags")
		return
	}
	return
}

func GetTagByID(db *gorm.DB, id uint) (tag Tag, err error) {
	err = db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		err = errors.Wrap(err, "GetTagByID")
		return
	}
	return
}

func GetTagByName(db *gorm.DB, name string) (tag Tag, err error) {
	err = db.Where("name = ?", name).First(&tag).Error
	if err != nil {
		err = errors.Wrap(err, "GetTagByName")
		return
	}
	return
}

func StoreTag(db *gorm.DB, tag *Tag) (err error) {
	err = db.Create(tag).Error
	if err != nil {
		err = errors.Wrap(err, "StoreTag")
		return
	}
	return
}

func CreateTag(db *gorm.DB, name string) (tag Tag, err error) {
	tag, err = GetTagByName(db, name)
	if err != nil && err.Error() == "GetTagByName: record not found" {
		tag.Name = name
		err = StoreTag(db, &tag)
	}
	return
}

func DeleteTagByID(db *gorm.DB, id uint) (err error) {
	err = db.Delete(Tag{}, "id = ?", id).Error
	if err != nil {
		err = errors.Wrap(err, "DeleteTagByID")
		return
	}
	return
}

func DeleteTagByName(db *gorm.DB, name string) (err error) {
	err = db.Delete(Tag{}, "id = ?", name).Error
	if err != nil {
		err = errors.Wrap(err, "DeleteTagByName")
		return
	}
	return
}
