package routes

import (
	"encoding/json"
	"net/http"
	"server/firebase/entity"
	"server/firebase/repo"

)

var (
	repos repo.PostRepo = repo.NewRepository()
)

func GetPosts(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")
	
	posts, err := repos.FindlAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Getting the post"`))
		return
	}

	res.WriteHeader(http.StatusOK)
	
	
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post 
	err := json.NewDecoder(req.Body).Decode(&post) //что бы прочитать  боди или ответ пустой прийдет

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Error marshaling"`))
		return
	}
	

	repos.Save(&post)
	res.WriteHeader(http.StatusOK)
	
	json.NewEncoder(res).Encode(post)
}
