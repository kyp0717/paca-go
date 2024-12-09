package db

import (
  "time"
)

type MinuteBar struct {
    Time   time.Time `gorm:"index;not null" json:"time"`         // Indexed for faster queries
    Open   float64   `gorm:"type:decimal(10,2);not null" json:"open"`
    High   float64   `gorm:"type:decimal(10,2);not null" json:"high"`
    Low    float64   `gorm:"type:decimal(10,2);not null" json:"low"`
    Close  float64   `gorm:"type:decimal(10,2);not null" json:"close"`
    Volume int64     `gorm:"not null" json:"volume"`
}

