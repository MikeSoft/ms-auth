package main

import (
	"context"
	"firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	opt := option.WithCredentialsFile("path/to/refreshToken.json")
	config := &firebase.Config{ProjectID: "my-project-id"}
	app, err := firebase.NewApp(context.Background(), config, opt)

	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	r.GET("/auth/request", func(c *gin.Context) {
		_, err := client.VerifyIDToken(context.Background(), "idToken")
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
		}
		c.String(http.StatusOK, "hello world")
	})

	r.Run()
}