package schema

type UserSchema struct {
	Id                   int    `json:"id" form:"id"`
	Name                 string `json:"name" form:"name" validate:"required,min=5,max=40"`
	Email                string `json:"email" form:"email" validate:"required,email"`
	Phone                string `json:"phone" form:"phone" validate:"required,min=10,max=10"`
	Password             string `json:"password" form:"password" validate:"required,min=8,max=20"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,min=8,max=20"`
}

type UserLoginSchema struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type TokenUsers struct {
	Token string             `json:"token"`
	Users UserSchemaResponse `json:"user"`
}

type UserSchemaResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
