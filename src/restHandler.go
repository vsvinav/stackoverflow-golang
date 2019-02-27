package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Api() {
	router := mux.NewRouter()
	// router.HandleFunc("/", PrintHello).Methods("GET")
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	router.HandleFunc("/comments", GetComments).Methods("GET")
	router.HandleFunc("/comments/{id}", GetComment).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/answers", GetAnswers).Methods("GET")
	router.HandleFunc("/upvote/{id}", Upvote).Methods("PUT")
	router.HandleFunc("/downvote/{id}", Downvote).Methods("PUT")

	// router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Listening on port 8000")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}

func PrintHello(w http.ResponseWriter, r *http.Request) {

}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := getPosts()
	json.NewEncoder(w).Encode(posts)
}
func GetComments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getComments())
}
func GetComment(w http.ResponseWriter, req *http.Request) {
	comments := getComments()
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		// handle error
		fmt.Println(err)
	}
	// y := getComment(comments[id].PostID)
	// json.NewEncoder(w).Encode(y)
	for _, item := range comments {

		if item.PostID == id {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getUsers())
}

func GetAnswers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getAnswers())
}

func GetPost(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	post := getPost(id)
	json.NewEncoder(w).Encode(post)
	// for _, item := range posts {

	// 	if item.ID == id {
	// 		json.NewEncoder(w).Encode(item)
	// 		fmt.Println("I am working in encoder")
	// 		return
	// 	}

	// }

}

func Upvote(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	upvote(id)
	post := getPost(id)
	json.NewEncoder(w).Encode(post)

}
func Downvote(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	downvote(id)
	post := getPost(id)
	json.NewEncoder(w).Encode(post)

}

// func BasicAuth(pass handler) handler {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

// 		if len(auth) != 2 || auth[0] != "Basic" {
// 			http.Error(w, "authorization failed", http.StatusUnauthorized)
// 			return
// 		}

// 		payload, _ := base64.StdEncoding.DecodeString(auth[1])
// 		pair := strings.SplitN(string(payload), ":", 2)

// 		if len(pair) != 2 || !validate(pair[0], pair[1]) {
// 			http.Error(w, "authorization failed", http.StatusUnauthorized)
// 			return
// 		}

// 		pass(w, r)
// 	}
// }

// func validate(username, password string) bool {
// 	if username == "test" && password == "test" {
// 		return true
// 	}
// 	return false
// }

// func GetPost(w http.ResponseWriter, r *http.Request) {
// 	var posts []Post
// 	params := mux.Vars(r)
// 	for _, item := range posts {

// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Post{})

// }
