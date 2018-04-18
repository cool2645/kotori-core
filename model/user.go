package model

import (
	"time"
	"github.com/cool2645/kotori-core/config"
)

type User struct {
	ID        uint       `gorm:"AUTO_INCREMENT" json:"id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Website   string     `json:"website"`
	Rank      int64      `json:"rank"`
	Honor     string     `json:"honor"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (User) TableName() string {
	return config.GlobCfg.TablePrefix + "_users"
}
