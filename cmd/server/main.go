package main

import (
	"IDS/api/internal/db"
	"IDS/api/internal/handlers"

	// "IDS/api/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	defer database.Close()

	router := gin.Default()
	// userHandler := handlers.NewUserHandler(database)
	roomHandler := handlers.NewRoomHandler(database)
	roomTypeHandler := handlers.NewRoomTypeHandler(database)

	router.GET("/rooms", roomHandler.GetRooms)
	router.GET("/roomtypes", roomTypeHandler.GetRoomTypes)
	// router.POST("/users", userHandler.CreateUser)
	// router.GET("/users/:id", userHandler.GetUser)

	router.Run(":8080")
}
