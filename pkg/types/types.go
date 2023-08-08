package types



type Error struct {
	Msg string `json:"msg"`
}

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type Book struct{
	Id string  `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Copies int `json:"copies"`
	Totalcount int `json:"totalcount"`	
}

type ListBooks struct {
	Books []Book `json:"books"`
}
