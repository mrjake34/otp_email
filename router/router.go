package router

import (
	"otpapi/controller"

	"github.com/gin-gonic/gin"
)

func SendOtp(e *gin.Engine) {
	e.GET("/send-otp-tr", controller.SendOtpHandlerTr)
	e.GET("/send-otp-en", controller.SendOtpHandlerEn)
}
