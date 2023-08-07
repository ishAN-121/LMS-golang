package types



type Error struct {
	Msg string `json:"msg"`
}

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

