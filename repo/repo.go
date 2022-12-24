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

const (
	projectId      string = "yoursuccess-e1c07"
	collectionName string = "users"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Messasing err!!!", err)
		return nil, err
	}

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Company": post.Company,
		"Email":   post.Email,
		"Name":    post.Name,
		"Message": post.Message,
		"Number":  post.Number,
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
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		log.Println("DAAAATA",doc.Data())
		if err != nil {
			log.Fatalln("Failed adding the FinalAll", err)
			return nil, err
		}
		post := entity.Post{
			Company: doc.Data()["company"].(string),
			Email:   doc.Data()["email"].(string),
			Name:    doc.Data()["name"].(string),
			Message: doc.Data()["message"].(string),
			Number:  doc.Data()["number"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
