package main

import (
	"eis-be/database"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDB()
	fmt.Println(db)

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
