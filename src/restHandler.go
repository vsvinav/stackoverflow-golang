package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goji/httpauth"
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
	http.Handle("/", httpauth.SimpleBasicAuth("someuser", "somepassword")(http.HandlerFunc(PrintHello))) // router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	// http.Handle("/", router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
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

func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}

		handler(w, r)
	}
}
