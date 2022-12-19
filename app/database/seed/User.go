package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"backend_api/app/models"
)

var users = []models.User{
	models.User{
		FirstName: "Steven",
		LastName: "victor",
		Username: "Steven",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		FirstName: "Martin",
		LastName: "Luther",
		Username: "Martin",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}


func Load(db *gorm.DB) {
	var err error
	if db.Debug().HasTable(&models.User{}) != true {
		log.Fatalf("do not exist table: %v", err)
	}
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}