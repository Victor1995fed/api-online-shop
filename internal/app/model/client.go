package model

//Client ...
type Client struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	Phone             int    `json:"phone"`
	EncryptedPassword string `json:"-"`
	DateCreate        string `json:"-"`
	DateUpdade        string `json:"-"`
}
