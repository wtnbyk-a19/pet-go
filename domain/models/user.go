package models

type User struct {
	Model
	UUID     string `json:"user_uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Pets []Pet `gorm:"foreignKey:UserUUID"`
}
