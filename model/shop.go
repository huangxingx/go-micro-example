package model

import "go-micro-example/pkg/database"

type ShopModel struct {
	ID uint32 `gorm:"primary_key" json:"id"`

	ShopName       string        `gorm:"not null" json:"shop_name"`
	RegionId       uint32        `gorm:"not null;default:0" json:"region_id"`
	Region         string        `gorm:"not null;default:''" json:"region"`
	Lon            float32       `gorm:"not null;default:0" json:"lon"`
	Lat            float32       `gorm:"not null;default:0" json:"lat"`
	Addr           string        `gorm:"not null;default:''" json:"addr"`
	Phone          string        `gorm:"not null;default:''" json:"phone"`
	RunningPeriod  database.JSON `gorm:"not null;" json:"running_period"`
	MD5List        database.JSON `gorm:"not null;" json:"md5_list"`
	UnitPrice      int64         `gorm:"not null;default:0" json:"unit_price"`
	UnitDuration   int32         `gorm:"not null;default:0" json:"unit_duration"`
	Proportion     float32       `gorm:"not null;default:0" json:"proportion"`
	ExpectIncome   float32       `gorm:"not null;default:0" json:"expect_income"`
	Heat           int8          `gorm:"not null;default:0;type:tinyint(4)" json:"head"`
	IndustryId     uint32        `gorm:"not null;default:0;" json:"industry_id"`
	Industry       string        `gorm:"not null;default:'';" json:"industry"`
	AppUserId      uint32        `gorm:"not null;default:0;" json:"app_user_id"`
	AppUserAccount string        `gorm:"not null;default:'';" json:"app_user_account"`

	DayAndNumberOfAdMap database.JSON `gorm:"not null;" json:"day_and_number_of_ad_map"`
	NumberOfAd          int32         `gorm:"not null;default:0" json:"number_of_ad"`
	BsNumber            int32         `gorm:"not null;default:0" json:"bs_number"`
	IsActive            bool          `gorm:"not null;default:0;type:tinyint(4)" json:"is_active"`

	database.Model
}
