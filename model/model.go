package model

type User struct {
	Email string `json:"email"`
}

type Otp struct {
	Otp string `json:"otp"`
}

type Config struct {
	Host     string
	Port     string
	From     string
	Password string
}
