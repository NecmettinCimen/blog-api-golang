package models

type User struct {
	ID        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserName  string
	FirstName string
	LastName  string
	Password  string
}
