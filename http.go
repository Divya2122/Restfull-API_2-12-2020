package main
import (
    
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)
type User struct{
     FullName string `json:"fullName"`
     Username string `json:"username"`
     Email string  `json:"email"`
}

 type Post struct {
      Title string `json:"title"`
      Body string `json:"body"`
      Author User `json:"author"`
}
 var Posts []Post = []Post{}
func main() {
   router := mux.NewRouter()
   
   router.HandleFunc("/posts", addItem).Methods("POST")
   router.HandleFunc("/posts", getAllPosts).Methods("GET")
   http.ListenAndServe(":5000", router)
}
func getAllPosts(w http.ResponseWriter, r *http.Request) {
     json.NewEncodeder(w).Encode(posts)

func addItem(w http.ResponseWriter, r *http.Request) {
  var newPost Post
   json.NewDecoder(r.Body).Decode(&newPost)
   Posts = append(Posts, newPost)
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(Posts)
}