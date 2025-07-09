package main

import (
	"back/infrastructure"
	"back/infrastructure/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	plainPassword := "katedegree"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	orm := infrastructure.NewGorm()

	user := model.UserModel{
		Name:        "katedegree",
		AccountCode: "katedegree",
		Password:    string(hashedPassword),
	}

	if err := orm.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
}
