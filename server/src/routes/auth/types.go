package auth

type SessionData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
