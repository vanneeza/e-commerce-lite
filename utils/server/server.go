package server

import (
	"log"
	"os"

	"github.com/vanneeza/e-commerce-lite/api"
	"github.com/vanneeza/e-commerce-lite/config"
	"github.com/vanneeza/e-commerce-lite/utils/pkg"
)

func Run() error {

	db, err := config.InitDB()
	if err != nil {
		return err
	}
	defer config.CloseDB(db)

	xenditKey := os.Getenv("XENDIT_API_KEY")
	jwtKey := os.Getenv("JWT_KEY")
	router := api.Run(db, jwtKey, xenditKey)
	serverAddress := pkg.GetEnv("SERVER_ADDRESS")
	log.Printf("Server is running on address %s\n", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		return err
	}

	return nil
}
