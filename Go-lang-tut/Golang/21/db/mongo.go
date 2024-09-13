package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restApi/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoURL = `mongodb+srv://test123:test123@cluster0.dvkdg4a.mongodb.net/`

const DBName = `Database`
const CollectionName = `platforms`

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(MongoURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(DBName).Collection(CollectionName)
}

func insertOne(plateform model.Platform) {
	inserted, err := collection.InsertOne((context.Background()), plateform)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", inserted.InsertedID)
}

func updatePlatform(platformId string) {
	id, _ := primitive.ObjectIDFromHex(platformId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// bson.M vs bson.D:

	// bson.M is an unordered, unsorted map. It is the equivalent of a JSON object.
	// bson.D is an ordered, sorted map. It is the equivalent of a JSON array.
	// The MongoDB documentation recommends using bson.D for structured documents.
	// https://docs.mongodb.com/manual/core/document/#document-bson-type

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func deletePlatform(platformId string) {
	id, _ := primitive.ObjectIDFromHex(platformId)

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func deleteAllPlatforms() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func getAllPlatforms() []model.Platform {
	var results []model.Platform

	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var elem model.Platform
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func GetAllPlateforms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	platforms := getAllPlatforms()
	json.NewEncoder(w).Encode(platforms)
}

func InsertPlateform(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var platform model.Platform
	json.NewDecoder(r.Body).Decode(&platform)
	insertOne(platform)
	json.NewEncoder(w).Encode(platform)
}

func UpdatePlateform(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	updatePlatform(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeletePlateform(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	deletePlatform(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllPlateforms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	deleteAllPlatforms()
}
