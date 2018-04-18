package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
)

type Comment struct {
	ID             uint       `gorm:"AUTO_INCREMENT" json:"id"`
	PostID         uint       `gorm:"index" json:"post_id"`
	Post           Post
	ReplyCommentID uint       `gorm:"index" json:"reply_comment_id"`
	ReplyNum       uint       `json:"reply_num"`
	ReplyUserID    uint       `json:"reply_user_id"`
	ReplyUser      User
	UserID         uint       `json:"user_id"`
	User           User
	Content        string     `sql:"type:text;" json:"content"`
	Format         string     `json:"format"`
	Public         bool       `gorm:"index" json:"public"`
	Admit          bool       `json:"admit"`
	Ban            bool       `json:"ban"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func (Comment) TableName() string {
	return config.GlobCfg.TablePrefix + "_comments"
}
