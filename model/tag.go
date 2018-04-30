package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
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
