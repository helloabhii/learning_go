package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/learning_go/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://helloabhii@cluster0.vwqrixa.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

//connect with mongoDB

func init() { //initial method run initally one at one time
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success !!!")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection Instance is ready")
}

// MongoDB Helpers - file

// insert  1 record

func insertOneMovie(movie model.Netflix) { //insert in lower case because we are not exporting this anywhere
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err) // more control version of the panic
	}
	fmt.Println("Inserted 1 movie in db with id : ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count : ", result.ModifiedCount)
}

//delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count :", deleteCount)
}

// delete all record from mongoDB
func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delted : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
