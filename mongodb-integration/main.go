package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title,omitempty" json:"title"`
	Body  string             `bson:"body,omitempty" json:"body"`
	Email string             `bson:"email,omitempty" json:"email"`
}

func main() {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://appAdmin:123456@localhost:27017/mydb"))
	if err != nil {
		log.Fatal(err)
	}

	//
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	/*
		Get my collection instance
	*/
	collection := client.Database("mydb").Collection("posts")

	insertMany(ctx, collection)
	posts := findAll(ctx, collection)

	updateOne(collection, posts[0].Id)
	deleteOne(collection, bson.D{{"_id", posts[0].Id}})
	deleteAll(collection)
}

func findAll(ctx context.Context, collection *mongo.Collection) []Post {
	/*
		Iterate a cursor
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var posts []Post
	if err := cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)
	return posts
}

func insertMany(ctx context.Context, collection *mongo.Collection) {
	/*
		Insert documents
	*/
	docs := []interface{}{
		bson.D{{"title", "World"}, {"body", "Hello World"}},
		bson.D{{"title", "Mars"}, {"body", "Hello Mars"}},
		bson.D{{"title", "Pluto"}, {"body", "Hello Pluto"}},
	}

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
}

func updateOne(collection *mongo.Collection, idToUpdate primitive.ObjectID) {
	filter := bson.D{{"_id", idToUpdate}}
	update := bson.D{{"$set", bson.D{{"email", "newemail@example.com"}}}}
	var updatedDocument Post
	errUpd := collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
	).Decode(&updatedDocument)

	if errUpd != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		fmt.Println("No Docs to update")
		if errUpd == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(errUpd)
	}
	formated, _ := json.Marshal(updatedDocument)
	fmt.Println("updated document", string(formated))
}

func deleteOne(coll *mongo.Collection, filter interface{}) {
	// Delete at most one document in which the "name" field is "Bob" or "bob".
	// Specify the SetCollation option to provide a collation that will ignore
	// case for string comparisons.
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	res, err := coll.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

func deleteAll(coll *mongo.Collection) {
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	res, err := coll.DeleteMany(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}
