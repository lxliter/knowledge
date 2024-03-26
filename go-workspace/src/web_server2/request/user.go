package request

type SignUpReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}