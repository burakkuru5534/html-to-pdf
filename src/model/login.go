package model

type LoginData struct {
	UpassFromDb  *string
	UserID       int64
	UserName     string
	UserFullName string
	UserEmail    string
}

type LoginReqData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
