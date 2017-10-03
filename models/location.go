package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Location struct {
	gorm.Model
	//LType      []int
	Imgs     []string `gorm:"type:varchar(64)[]"`
	Point    GeoPoint
	PointL   []int `gorm:"type:varchar(100)[]"`
	Content  string
	IsDelete bool
	//Liked      []Like
	LikedNum int
	//Comments   []Comment
	CommentNum int
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v, %v", p.Lat, p.Lng)
}
