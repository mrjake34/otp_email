package main

import (
	"log"
	"net/http"
	"os"
	"otpapi/router"

	"github.com/gorilla/mux"
)

// GetConfig is used to get all configuration data.

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// E-posta ayarları
	r := mux.NewRouter()

	router.SendOtp(r)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("ListenAndServe: ", err)
		panic(err)
	} else {
		log.Println("Server is running on port 8080")
	}
}
