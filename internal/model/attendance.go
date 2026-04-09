package models

import "time"

type Attendance struct {
	AttendanceID       uint `gorm:"primaryKey"`
	UserID             uint
	ReportID           uint
	AttendanceTime     time.Time
	AttendanceDeadline time.Time

	User   User
	Report Report
}
