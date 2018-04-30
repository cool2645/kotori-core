package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
)

type Comment struct {
	ID              uint       `gorm:"AUTO_INCREMENT" json:"id"`
	PostID          uint       `gorm:"index" json:"post_id"`
	Post            Post
	FatherCommentID uint       `gorm:"index" json:"father_comment_id"`
	FatherUserID    uint       `json:"father_user_id"`
	FatherUser      User       `json:"father_user"`
	Children        []Comment  `json:"children"`
	ChildrenNum     uint       `json:"children_num"`
	ReplyCommentID  uint       `gorm:"index" json:"reply_comment_id"`
	ReplyUserID     uint       `json:"reply_user_id"`
	ReplyUser       User
	UserID          uint       `json:"user_id"`
	User            User
	Content         string     `sql:"type:text;" json:"content"`
	Format          string     `json:"format"`
	Encoding        string     `json:"encoding"`
	IsPublic        bool       `gorm:"index" json:"is_public"`
	Public          bool       `json:"public"`
	Private         bool       `json:"private"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func (Comment) TableName() string {
	return config.GlobCfg.TablePrefix + "_comments"
}
