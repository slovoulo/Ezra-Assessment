package models

import "time"

type Elevator struct {
	LogID        int `gorm:"primary_key"`
	ElevatorID   int `gorm:"not null"`
	CalledAt     time.Time
	CurrentFloor int    `gorm:"not null"`
	TargetFloor  int    `gorm:"not null"`
	State        string `gorm:"not null"`
	Direction    string `gorm:"not null"`
	CallerName   string `gorm:"not null"`
	CallerID     string `gorm:"not null"`
}