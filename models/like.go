package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Like struct {
	ID         uint
	UserId     int
	LocationId int
	CreateAt   time.Time
}
