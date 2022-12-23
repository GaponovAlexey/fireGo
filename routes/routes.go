package routes

import (
	"encoding/json"
	"net/http"

)

type Post struct {
	Id    int    `json:"id"`
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

	// res.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Error marshaling"`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
