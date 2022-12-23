package entity

// type Post struct {
// 	Id    int    `json:"id"` // Все должни быть уникальными
// 	Title string `json:"title"`
// 	Text  string `json:"text"`
// }

type Post struct {
	Company string `json:"company"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Number  string `json:"number"`
}
