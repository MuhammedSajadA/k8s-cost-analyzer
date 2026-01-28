package models

import "time"

type Cluster struct {
	ID         string    `gorm:"primaryKey"`
	UserID     string    `gorm:"index;not null"`
	Name       string    `gorm:"not null"`
	Kubeconfig string    `gorm:"type:text;not null"`
	CreatedAt  time.Time
}
