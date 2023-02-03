package seeder

import "github.com/aldiramdan/go-backend/databases/orm/models"

var UserSeed = models.Users{
	{
		Username:  "admin",
		Email:     "admin@email.com",
		Role:      "admin",
		Password:  "$2a$12$LzPmxfEZoVbCpGUGticqreZHbKLJICuXHPjwOPMZ9OFrmSDWHyPQW",
		Name:      "admin",
		Gender:    "Male",
		Address:   "localhost",
		Phone:     "08123456789",
		BirthDate: "02/01/1999",
		IsActive:  true,
		Picture:   "public/default_image.jpg",
	},
	{
		Username:  "user",
		Email:     "user@email.com",
		Role:      "user",
		Password:  "$2a$12$wcGtHuywxUX8fvxYqv8aJ.A0JcasSMqFglWtoIjNQxNRQlPQ/ChGO",
		Name:      "user",
		Gender:    "Male",
		Address:   "localhost",
		Phone:     "08123456789",
		BirthDate: "02/01/1999",
		IsActive:  true,
		Picture:   "public/default_image.jpg",
	},
}
