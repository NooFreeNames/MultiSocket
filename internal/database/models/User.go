package models

type User struct {
	Id           uint      `gorm:"primary_key"`
	Name         string    `gorm:"unique;not null;type:varchar(20)"`
	Password     string    `gorm:"not null;type:varchar(60)"`
	AccessToken  string    `gorm:"unique;not null;type:varchar(128)"`
	RefreshToken string    `gorm:"unique;not null;type:varchar(128)"`
	Channels     []Channel `gorm:"foreignkey:UserId"`
}
