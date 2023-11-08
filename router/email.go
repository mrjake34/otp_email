package router

import (
	"github.com/gorilla/mux"
	"otp_email/controller"
)

func SendOtp(r *mux.Router) {
	r.HandleFunc("/send-otp", controller.SendOtpHandler).Methods("POST")
}
