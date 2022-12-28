package routes

import (
	"encoding/json"
	"log"
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
		log.Println("FINDALL DON't WORK")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Getting the post"`))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}


func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post) //что бы прочитать  боди или ответ пустой прийдет
	
	
	log.Println("BODYYYY",post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error"}: "Error marshaling"`))
		return
	}

	


	log.Println("POSTs", post)

	repos.Save(&post)
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(post)
}
