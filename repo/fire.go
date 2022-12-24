package repo

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

)

func initializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]
	opt := option.WithCredentialsFile("../key.json")
  config := &firebase.Config{ProjectID: "yoursuccess-e1c07"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}