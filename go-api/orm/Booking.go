package orm

import (
	"time"

	"gorm.io/gorm" // framwork ต่อกับ database
)

type Booking struct { //สร้างตารางใน database ชื่อ Booking
	gorm.Model
	UserID string
	CarID string
	Start time.Time
	End time.Time
}