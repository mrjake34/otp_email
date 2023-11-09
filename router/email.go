package router

import (
	"github.com/gorilla/mux"
	"otp_email/controller"
)

func SendOtp(r *mux.Router) {
	r.HandleFunc("/send-otp-tr", controller.SendOtpHandlerTr).Methods("POST")
	r.HandleFunc("/send-otp-en", controller.SendOtpHandlerEn).Methods("POST")
}
