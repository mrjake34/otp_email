package router

import (
	"otpapi/controller"

	"github.com/gin-gonic/gin"
)

func SendOtp(e *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	e.GET("/send-otp-tr", controller.SendOtpHandlerTr)
	e.GET("/send-otp-en", controller.SendOtpHandlerEn)
}
