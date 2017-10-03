package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Comment struct {
	gorm.Model
	Content    string
	UserId     int
	LocationId int
	At         int
}
