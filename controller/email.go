package controller

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/smtp"
	"otp_email/config"
	"otp_email/model"

	"github.com/jordan-wright/email"
)

// SendOtpHandler is used to send OTP to user's email address.
func SendOtpHandler(w http.ResponseWriter, r *http.Request) {
	userEmail := &model.Email{}
	fmt.Print(r.Body)
	err := json.NewDecoder(r.Body).Decode(&userEmail)
	if err != nil {
		fmt.Println("Decoder hatası:", err)
	}
	// E-posta ayarları
	from := config.GetConfig().Email
	password := config.GetConfig().Password
	to := userEmail.Email
	smtpServer := config.GetConfig().Host
	port := config.GetConfig().Port
	serverAndPort := smtpServer + ":" + port

	// OTP oluştur
	otp := generateOTP(5)

	// E-posta gönder
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = "E-posta doğrulama"
	e.HTML = []byte("<!DOCTYPE html> <html> <head> <meta charset=\"UTF-8\"> <title>OTP Email</title> </head> <body> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"100%\" style=\"background-color: #f5f5f5; font-family: Arial, sans-serif;\"> <tr> <td align=\"center\"> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"600\" style=\"background-color: #ffffff;\"> <tr> <td align=\"center\" style=\"padding: 20px;\"> <h1 style=\"color: #333333;\">Tek Kullanımlık Şifre (OTP)</h1> <p style=\"color: #666666;\">Hesabınıza giriş yapabilmek için aşağıdaki OTP kodunu kullanın:</p> <div style=\"background-color: #007BFF; color: #ffffff; font-size: 24px; padding: 10px; text-align: center; margin: 20px 0;\"> " + otp + "</div> <p style=\"color: #666666;\">Bu OTP kodunu kimseyle paylaşmayın. Hesabınızı güvende tutmak için kullanılır. </p> <p style=\"color: #666666;\">Eğer bu e-posta size ait değilse, lütfen dikkate almayın.</p> </td> </tr> </table> </td> </tr> </table> </body> </html>")

	err = e.Send(serverAndPort, smtp.PlainAuth("", from, password, smtpServer))
	if err != nil {
		fmt.Println("E-posta gönderme hatası:", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println("E-posta başarıyla gönderildi.")
		w.WriteHeader(http.StatusOK)
		otpJson := &model.Otp{Otp: otp}
		json.NewEncoder(w).Encode(otpJson)
	}

}
func generateOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		idx, _ := rand.Int(rand.Reader, max)
		otp[i] = charset[idx.Int64()]
	}

	return string(otp)
}
