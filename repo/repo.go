package repo

import (
	"context"
	"log"
	"server/firebase/entity"

	"cloud.google.com/go/firestore"

)

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindlAll() ([]entity.Post, error)
}

type repo struct{}

// new Repo
func NewRepository() PostRepo {
	return &repo{}
}

var (
	projectId      string = "yoursuccess-e1c07"
	collectionName string = "users"
	posts                 = []entity.Post{}
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Messasing err!!!", err)
		return nil, err
	}

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"company": post.Company,
		"email":   post.Email,
		"name":    post.Name,
		"message": post.Message,
		"number":  post.Number,
	})
	if err != nil {
		log.Fatalln("Failed adding the new post", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindlAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalln("Mi FATAL", err)
		return nil, err
	}
	defer client.Close()
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		v, err := iterator.Next()
		if err != nil {
			log.Fatalln("Failed adding the FinalAll", err)
			return nil, err
		}

		post := entity.Post{
			Company: v.Data()["company"].(string),
			Email:   v.Data()["email"].(string),
			Name:    v.Data()["name"].(string),
			Message: v.Data()["message"].(string),
			Number:  v.Data()["number"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil

}
