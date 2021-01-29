package model

//User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	DateCreate        string `json:"-"`
	DateUpdade        string `json:"-"`
}

// Validate ...
// func (u *User) Validate() error {
// 	return validation.ValidateStruct(u,
// 		validation.Field(
// 			&u.Email,
// 			validation.Required,
// 			is.Email,
// 		),
// 		validation.Field(
// 			&u.Password,
// 			validation.By(requiredIf(u.EncryptedPassword == "")),
// 			validation.Length(6, 100),
// 		),
// 	)
// }

// // BeforeCreate ...
// func (u *User) BeforeCreate() error {
// 	if len(u.Password) > 0 {
// 		enc, err := encryptString(u.Password)
// 		if err != nil {
// 			return nil
// 		}

// 		u.EncryptedPassword = enc
// 	}
// 	return nil
// }

// // Sanitaze ...
// func (u *User) Sanitaze() {
// 	u.Password = ""
// }

// // ComparePassword ...
// func (u *User) ComparePassword(password string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
// }

// func encryptString(s string) (string, error) {
// 	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(b), nil
// }
