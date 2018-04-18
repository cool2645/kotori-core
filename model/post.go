package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
)

type Post struct {
	ID          uint       `gorm:"AUTO_INCREMENT" json:"id"`
	Edition     uint       `json:"edition"`
	Title       string     `gorm:"index" json:"title"`
	Summary     string     `sql:"type:text;" json:"summary"`
	CategoryID  uint       `json:"category_id"`
	Category    Category   `json:"category"`
	Tags        []Tag      `gorm:"many2many:post_tags;"`
	Private     bool       `gorm:"index" json:"private"`
	Hide        bool       `gorm:"index" json:"hide"`
	CommentMode int        `json:"comment_mode"`
	CommentNum  uint       `json:"comment_num"`
	VisitNum    uint64     `json:"visit_num"`
	Content     string     `sql:"type:text;" json:"content"`
	Format      string     `json:"format"`
	License     string     `json:"license"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (Post) TableName() string {
	return config.GlobCfg.TablePrefix + "_posts"
}

const (
	CommentModePermitAll      = iota
	CommentModeNeedAdmission
	CommentModeDisableComment
)
