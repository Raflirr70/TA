package models

import "time"

type Attendance struct {
	AttendanceID       uint `gorm:"primaryKey"`
	UserID             uint
	ReportID           uint
	QRID               uint
	AttendanceTime     time.Time
	AttendanceDeadline time.Time

	User   User
	Report Report
	QR     QR
}
