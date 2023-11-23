package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"otpapi/config"
	"otpapi/model"
	"otpapi/utils"

	"github.com/gin-gonic/gin"
)

// SendOtpHandler is used to send OTP to user's email address.
func SendOtpHandlerTr(c *gin.Context) {
	key := &model.Key{}
	if err := c.BindJSON(&key); err != nil {
		log.Println("Key is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required", "message": err.Error()})
		return
	}
	// E-posta ayarları
	cfg := config.GetConfig()
	from := cfg.From
	password := cfg.Password
	smtpServer := cfg.Host
	port := cfg.Port
	serverAndPort := smtpServer + ":" + port
	savedKey := cfg.Key
	to := c.Query("email")

	if to == "" {
		log.Println("Email adresi boş olamaz")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email adresi boş olamaz"})
		return
	}
	if key.Key != savedKey {
		log.Println("Yetkisiz erişim")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim"})
		return
	}

	// OTP oluştur
	otp := utils.GenerateOTP(5)

	// E-posta gönder
	body := []byte("<!DOCTYPE html> <html> <head> <meta charset=\"UTF-8\"> <title>OTP Email</title> </head> <body> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"100%\" style=\"background-color: #f5f5f5; font-family: Arial, sans-serif;\"> <tr> <td align=\"center\"> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"600\" style=\"background-color: #ffffff;\"> <tr> <td align=\"center\" style=\"padding: 20px;\"> <h1 style=\"color: #333333;\">Tek Kullanımlık Şifre (OTP)</h1> <p style=\"color: #666666;\">İşleminize devam etmek için kullanacağınız doğrulama kodu:</p> <div style=\"background-color: #007BFF; color: #ffffff; font-size: 24px; padding: 10px; text-align: center; margin: 20px 0;\"> " + otp + "</div> <p style=\"color: #666666;\">Bu OTP kodunu kimseyle paylaşmayın. Hesabınızı güvende tutmak için kullanılır. </p> <p style=\"color: #666666;\">Eğer bu e-posta size ait değilse, lütfen dikkate almayın.</p> </td> </tr> </table> </td> </tr> </table> </body> </html>")
	subject := "Doğrulama kodu"
	msg := fmt.Sprintf("Subject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", subject, body)
	if err := smtp.SendMail(serverAndPort, smtp.PlainAuth("", from, password, smtpServer), from, []string{to}, []byte(msg)); err != nil {
		log.Println("E-posta gönderme hatası:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-posta gönderme hatası", "message": err.Error()})
	} else {
		log.Println("E-posta başarıyla gönderildi.")
		c.JSON(http.StatusOK, gin.H{"otp": otp})
	}
}

func SendOtpHandlerEn(c *gin.Context) {
	key := &model.Key{}
	if err := c.BindJSON(&key); err != nil {
		log.Println("Key is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required", "message": err.Error()})
		return
	}

	// Get config data
	config := config.GetConfig()

	// Email settings
	from := config.From
	password := config.Password
	to := c.Query("email")
	if to == "" {
		log.Println("Email does not empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not empty"})
		return
	}
	if key.Key != config.Key {
		log.Println("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	smtpServer := config.Host
	port := config.Port
	serverAndPort := smtpServer + ":" + port

	// Generate OTP
	otp := utils.GenerateOTP(5)

	// Create the email body
	body := []byte(
		"<!DOCTYPE html> <html> <head> <meta charset=\"UTF-8\"> <title>OTP Email</title> </head> <body> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"100%\" style=\"background-color: #f5f5f5; font-family: Arial, sans-serif;\"> <tr> <td align=\"center\"> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"600\" style=\"background-color: #ffffff;\"> <tr> <td align=\"center\" style=\"padding: 20px;\"> <h1 style=\"color: #333333;\">One-Time Password (OTP)</h1> <p style=\"color: #666666;\">The verification code you will use to continue your transaction:</p> <div style=\"background-color: #007BFF; color: #ffffff; font-size: 24px; padding: 10px; text-align: center; margin: 20px 0;\"> " + otp + "</div> <p style=\"color: #666666;\">Do not share this OTP code with anyone. It is used to keep your account secure. </p> <p style=\"color: #666666;\">If this email does not belong to you, please disregard it.</p> </td> </tr> </table> </td> </tr> </table> </body> </html>")
	subject := "Verification code"
	msg := fmt.Sprintf("Subject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", subject, body)
	// Send the email and handle errors
	if err := smtp.SendMail(serverAndPort, smtp.PlainAuth("", from, password, smtpServer), from, []string{to}, []byte(msg)); err != nil {
		log.Println("E-posta gönderme hatası:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email error", "message": err.Error()})
	} else {
		log.Println("E-posta başarıyla gönderildi.")
		c.JSON(http.StatusOK, gin.H{"otp": otp})
	}
}
