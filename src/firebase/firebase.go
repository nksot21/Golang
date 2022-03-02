package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

//$env:GOOGLE_APPLICATION_CREDENTIALS="D:\myedu\gopath\src\chatsdemo\src\key\firebasekey.json"

type FireApp struct {
	App  *firebase.App
	Auth *auth.Client
	Db   *firestore.Client
}

var FirebaseApp = FireApp{}
var Ctx = context.Background()

func ConnectFirebase() {
	opt := option.WithCredentialsFile("src/key/firebasekey.json")
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
	client, err := firestore.NewClient(Ctx, "chatdemo-bfd28")
	if err != nil {
		fmt.Println(err)
	}
	FirebaseApp.Db = client
}
