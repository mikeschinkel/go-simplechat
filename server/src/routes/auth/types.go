package auth

type LoginReq struct {
	Email    string
	Password string
}

type UserData struct {
	ID    uint
	Email string
	Name  string
}
