package model

type User struct {
	Email string `json:"email"`
}

type Key struct {
	Key string `json:"key"`
}

type Otp struct {
	Otp string `json:"otp"`
}

type Config struct {
	Host     string
	Port     string
	From     string
	Password string
	Key      string
}
type Message struct {
	From    string
	To      []string
	Subject string
	Body    []byte
}
