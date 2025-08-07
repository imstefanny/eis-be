package main

import (
	"eis-be/database"
	"eis-be/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.InitDB()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://landing-page-school.web.app", "https://eis-letjen.web.app", "http://localhost:5174", "http://localhost:5173", "https://currently-big-flea.ngrok-free.app"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "ngrok-skip-browser-warning"},
		AllowCredentials: true,
	}))

	route.Route(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
