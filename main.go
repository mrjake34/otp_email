package main

import (
	"log"
	"os"
	"otpapi/router"

	"github.com/gin-gonic/gin"
)

// GetConfig is used to get all configuration data.

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	r := gin.Default()

	router.SendOtp(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
