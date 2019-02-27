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
	router.HandleFunc("/posts", GetPostsJSON).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPostJSON).Methods("GET")
	router.HandleFunc("/comments", GetCommentsJSON).Methods("GET")
	router.HandleFunc("/comments/{id}", GetCommentJSON).Methods("GET")
	router.HandleFunc("/users", GetUsersJSON).Methods("GET")
	router.HandleFunc("/answers", GetAnswersJSON).Methods("GET")
	router.HandleFunc("/upvote/{id}", UpvoteJSON).Methods("PUT")
	router.HandleFunc("/downvote/{id}", DownvoteJSON).Methods("PUT")
	// http.Handle("/", httpauth.SimpleBasicAuth("someuser", "somepassword")(http.HandlerFunc(PrintHello))) // router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	http.Handle("/", router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("listening on port 8000")
}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func GetPostsJSON(w http.ResponseWriter, r *http.Request) {
	posts := GetPosts()
	json.NewEncoder(w).Encode(posts)
}
func GetCommentsJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetComments())
}
func GetCommentJSON(w http.ResponseWriter, req *http.Request) {
	comments := GetComments()
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		// handle error
		fmt.Println(err)
	}
	// y := GetComment(comments[id].PostID)
	// json.NewEncoder(w).Encode(y)
	for _, item := range comments {

		if item.PostID == id {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func GetUsersJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetUsers())
}

func GetAnswersJSON(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetAnswers())
}

func GetPostJSON(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	post := GetPost(id)
	json.NewEncoder(w).Encode(post)
	// for _, item := range posts {

	// 	if item.ID == id {
	// 		json.NewEncoder(w).Encode(item)
	// 		fmt.Println("I am working in encoder")
	// 		return
	// 	}

	// }

}

func UpvoteJSON(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	upvote(id)
	post := GetPost(id)
	json.NewEncoder(w).Encode(post)

}
func DownvoteJSON(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	downvote(id)
	post := GetPost(id)
	json.NewEncoder(w).Encode(post)

}
