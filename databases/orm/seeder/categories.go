package seeder

import "github.com/aldiramdan/go-backend/databases/orm/models"

var CategorySeed = models.Categories{
	{
		Name: "Cars",
	},
	{
		Name: "Motorbike",
	},
	{
		Name: "Bike",
	},
}
