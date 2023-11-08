package main

import (
	"net/http"
	"otp_email/router"

	"github.com/gorilla/mux"
)

func main() {
	// E-posta ayarları
	r := mux.NewRouter()

	router.SendOtp(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
