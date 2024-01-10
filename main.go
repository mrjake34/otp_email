package main

import (
	"fmt"
	"log"
	"os"
	"otpapi/router"

	"github.com/gin-gonic/gin"
)

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SendOtp(r)

	if err := r.Run(":8080"); err != nil {

		fmt.Println("Server is running on port 8080")
		log.Fatal(err)
	}

}
