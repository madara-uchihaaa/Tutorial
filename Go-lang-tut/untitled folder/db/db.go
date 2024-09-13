package db

import (
	"backend/model"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Mongo_Url = `mongodb+srv://test123:test123@cluster0.dvkdg4a.mongodb.net/`
const DBName = `Database`
const CollectionName = `user`
const ChatCollectionName = `chat`

var Collection *mongo.Collection
var ChatCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(Mongo_Url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	Collection = client.Database(DBName).Collection(CollectionName)
	ChatCollection = client.Database(DBName).Collection(ChatCollectionName)
}

func RegisterUser(model *model.User) {
	_, err := Collection.InsertOne(context.Background(), model)
	if err != nil {
		panic(err)
	}
}

func LoginUser(email string) *model.User {
	var user model.User
	filter := bson.D{{Key: "email", Value: email}}
	err := Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	return &user
}

func GetAllChats(id string) []model.Chat {
	var chats []model.Chat
	filter := bson.D{{Key: "userid", Value: id}}
	cursor, err := ChatCollection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var chat model.Chat
		cursor.Decode(&chat)
		chats = append(chats, chat)
	}
	return chats
}

func GetChat(id string) *model.Chat {
	var chat model.Chat
	filter := bson.D{{Key: "userid", Value: id}}
	err := ChatCollection.FindOne(context.Background(), filter).Decode(&chat)
	if err != nil {
		panic(err)
	}
	return &chat
}

func DeleteChat(id string) {
	filter := bson.D{{Key: "userid", Value: id}}
	_, err := ChatCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
}

func CreateChat(chat *model.Chat) {
	_, err := ChatCollection.InsertOne(context.Background(), chat)
	if err != nil {
		panic(err)
	}
}

func DeleteAllChats(id string) {
	filter := bson.D{{Key: "userid", Value: id}}
	_, err := ChatCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
}

func GetResponse(query string) string {
	err := godotenv.Load(".env")
	if err != nil{
	log.Fatalf("Error loading .env file: %s", err)
	}

	client := openai.NewClient(os.Getenv("TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		return "Error while getting response from GPT-3"
	}
	return resp.Choices[0].Message.Content
}
