package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

//$env:GOOGLE_APPLICATION_CREDENTIALS="pkg\firebase\firebasekey.json"

type FireApp struct {
	App  *firebase.App
	Auth *auth.Client
	Db   *firestore.Client
}

var FirebaseApp = FireApp{}
var Ctx = context.Background()

func ConnectFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("KEY_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error auth: %v", err)
	}
	FirebaseApp.App = app
	FirebaseApp.Auth = auth
}

func ConnectFirestore() {
	client, err := firestore.NewClient(Ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		fmt.Println(err)
	}
	FirebaseApp.Db = client
}
