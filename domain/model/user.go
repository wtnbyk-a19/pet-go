package model

type User struct {
	Model
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Pets []Pet `gorm:"foreignKey:UserUUID"`
}
