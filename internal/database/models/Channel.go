// Package models contains the structures used to represent data in the
// application.
package models

// Channel represents a websocket channel in the database
type Channel struct {
	Id       uint    `gorm:"primary_key"`
	Name     string  `gorm:"unique;not null;type:varchar(20)"`
	Password *string `gorm:"null;type:varchar(60)"`
	UserId   uint    `gorm:"not null"`
}
