package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
)

type Edition struct {
	ID        uint       `gorm:"AUTO_INCREMENT" json:"id"`
	Edition   uint       `gorm:"index" json:"edition"`
	PostID    uint       `gorm:"index" json:"post_id"`
	Post      Post
	Title     string     `json:"title"`
	Summary   string     `sql:"type:text;" json:"summary"`
	Content   string     `sql:"type:text;" json:"content"`
	Format    string     `json:"format"`
	Encoding  string     `json:"encoding"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Edition) TableName() string {
	return config.GlobCfg.TablePrefix + "_editions"
}
