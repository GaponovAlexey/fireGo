package routes

import (
	"encoding/json"
	"net/http"

)

type Post struct {
	Id    int    `json:"id"` // Все должни быть уникальными 
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "title 1", Text: "text 1"}}
}

func GetPosts(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Error marshaling"`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	var post Post                                  // наследуется от массива
	err := json.NewDecoder(req.Body).Decode(&post) //что бы прочитать  боди или ответ пустой прийдет
	
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Error marshaling"`))
		return
	}

	post.Id = len(posts) + 1 // проверяет длину и добавлят + 1 это если не указывать в боди
	posts = append(posts, post) // вот тут записывает в сам обжект

	res.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(post) // розпечатывает для отправки 
	res.Write(result)

}
