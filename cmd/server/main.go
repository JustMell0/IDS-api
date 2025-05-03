package main

import (
	"IDS/api/internal/db"
	"IDS/api/internal/handlers"

	// "IDS/api/internal/handlers"
	"log"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	roomHandler := handlers.NewRoomHandler(database)
	roomTypeHandler := handlers.NewRoomTypeHandler(database)
	bookingHandler := handlers.NewBookingHandler(database)
	reservationHandler := handlers.NewReservationHandler(database)
	loginHandler := handlers.NewLoginHandler(database)
	conciergeHandler := handlers.NewConciergeHandler(database)
	userHandler := handlers.NewUserHandler(database)

	router.GET("/rooms", roomHandler.GetRooms)
	router.GET("/roomtypes", roomTypeHandler.GetRoomTypes)

	router.POST("/booking", bookingHandler.CreateBooking)

	router.GET("/reservations", reservationHandler.GetReservations)

	router.POST("/login", loginHandler.Login)

	router.GET("/users/:id/reservations", reservationHandler.GetUserReservations)

	router.GET("/concierge/requests", conciergeHandler.GetRequests)
	router.POST("/concierge/requests/:id", conciergeHandler.AcceptRequest)
	router.DELETE("/concierge/requests/:id", conciergeHandler.RejectRequest)

	router.GET("/users/:id/requests", userHandler.GetUserRequests)
	// router.POST("/users", userHandler.CreateUser)
	// router.GET("/users/:id", userHandler.GetUser)

	router.Run(":8080")
}
