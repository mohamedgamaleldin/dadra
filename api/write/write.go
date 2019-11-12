package main

import (
	"context"
	"fmt"
	"log"
	_ "net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Write service called.")
}*/

type Post struct {
	text string
}

func main() {
	/*http.HandleFunc("/", handler)
	fmt.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:3000", nil)*/

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection success.")

	collection := client.Database("test").Collection("posts")

	post := Post{"Hello, mongo!"}

	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted document: ", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection closed.")
}
