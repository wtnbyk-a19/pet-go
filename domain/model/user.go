package model

import "github.com/google/uuid"

type User struct {
	Model
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Pets []Pet `gorm:"foreignKey:UserUUID"`
}

func (u *User) GenerateUUID() {
	uuidobj, _ := uuid.NewUUID()
	uuid := uuidobj.String()
	u.UUID = uuid
}
