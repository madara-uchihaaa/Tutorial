package routes

import (
	"backend/db"
	"backend/model"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func register(user *model.SignUpInput) {
	fmt.Println("Inputs for register", user)
	if !user.Validate() {
		panic("Invalid Input")
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hashPassword
	newUser := &model.User{
		ID:        uuid.New().String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	db.RegisterUser(newUser)
}

func login(user *model.SignInInput) string {
	fmt.Println("Inputs for login", user)
	if !user.Validate() {
		panic("Invalid Input")
	}
	dbUser := db.LoginUser(user.Email)

	if !utils.CheckPassword(user.Password, dbUser.Password) {
		panic("Invalid Password")
	}
	token := utils.CreateToken(dbUser)
	return token
}

func getRefreshToken(initialToken string, user *model.User) string {
	fmt.Println("Inputs for refresh token", initialToken, user)
	if !utils.ValidateToken(initialToken) {
		panic("Invalid Token")
	}
	token := utils.CreateToken(user)
	return token
}

func Register(w http.ResponseWriter, r *http.Request) {
	var signInput model.SignUpInput
	json.NewDecoder(r.Body).Decode(&signInput)
	register(&signInput)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&signInput)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var signInput model.SignInInput
	json.NewDecoder(r.Body).Decode(&signInput)
	token := login(&signInput)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	token := getRefreshToken(r.Header.Get("Authorization"), &user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func GetAllChats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Chats", r.Body)
	userId := mux.Vars(r)["id"]
	chats := db.GetAllChats(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chats)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Chat", r.Body)
	id := mux.Vars(r)["id"]
	chat := db.GetChat(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Chat", r.Body)
	id := mux.Vars(r)["id"]
	db.DeleteChat(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Chat Deleted Successfully")
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Chat", r.Body)
	var chat model.Chat
	json.NewDecoder(r.Body).Decode(&chat)
	db.CreateChat(&chat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Chat Created Successfully")
}

func DeleteAllChats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete All Chats", r.Body)
	userId := mux.Vars(r)["id"]
	db.DeleteAllChats(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("All Chats Deleted Successfully")
}



func GetResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Response", r.Body)
	var chat model.QueryInput
	json.NewDecoder(r.Body).Decode(&chat)
	response := db.GetResponse(chat.Query)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}