package router

import (
	"github.com/gin-gonic/gin"
	"otpapi/controller"
)

func SendOtp(e *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	e.GET("/send-otp-tr", controller.SendOtpHandlerTr)
	e.GET("/send-otp-en", controller.SendOtpHandlerEn)
}
