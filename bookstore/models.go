package bookstore

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    int    `json:"author"`
	Price     int    `json:"price"`
	Publisher int    `json:"publisher"`
	Published string `json:"published"`
}

type Author struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
type Publisher struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Bookstore struct {
	Id     int `json:"id"`
	BookId int `json:"book_id"`
	UserId int `json:"user_id"`
}
