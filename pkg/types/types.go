package types
type Error struct {
	Msg string `json:"msg"`
}

type User struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Admin bool `json:"admin"`
}
type Userlist struct{
	Users []User `json:"users"`
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


type Request struct{
	Id string `json:"id"`
	BookId int `json:"bookid"`
	Username string `json:"username"`
	Status string `json:"status"`
}

type RequestLists struct{
	Requests []Request `json:"requests"`
}

type Data struct {
	Books []Book
	Requests []Request
	Error string
}

type AdminRequest struct {
	Adminrequest bool `json:"adminrequest"`
}