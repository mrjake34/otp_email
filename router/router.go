package router

import (
	"otpapi/controller"

	"github.com/gorilla/mux"
)

func SendOtp(r *mux.Router) {
	r.HandleFunc("/send-otp-tr", controller.SendOtpHandlerTr).Methods("POST")
	r.HandleFunc("/send-otp-en", controller.SendOtpHandlerEn).Methods("POST")
}
