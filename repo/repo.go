package repo

import (
	"context"
	"log"
	"server/firebase/entity"

	// "cloud.google.com/go/firestore"

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
	app := initializeAppWithServiceAccount()

	app.Firestore()



	_, _, err = app.Collection(collectionName).Add(ctx, map[string]interface{}{
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
	
	client, err := firestore.NewClient(ctx, "../key.json")
	if err != nil {
		log.Fatalf("new Client don't find %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	log.Println("iterator", iterator)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalln("Failed adding the FinalAll", err)
			return nil, err
		}
		post := entity.Post{
			Company: doc.Data()["Company"].(string),
			Email:   doc.Data()["Email"].(string),
			Name:    doc.Data()["Name"].(string),
			Message: doc.Data()["Message"].(string),
			Number:  doc.Data()["Number"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
