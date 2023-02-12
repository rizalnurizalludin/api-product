package main

import (
	"github.com/rizalnurizalludin/api-product/entity"
	"github.com/rizalnurizalludin/api-product/routes"
)

func main() {
	db := entity.ConnectDatabase()
	db.AutoMigrate(&entity.Product{})

	r := routes.SetupRoutes(db)
	r.Run(":3030")
}
