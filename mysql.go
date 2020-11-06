package main
import (
  "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"

  
)
type Post struct {
  ID string `json:"id"`
  Title string `json:"name"`
  CITY string `json:"city"`
}
var db *sql.DB
var err error
func main() {
db, err = sql.Open("mysql", "root:143114@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  router := mux.NewRouter()
  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  router.HandleFunc("/posts/{id}", getPosts).Methods("GET")
//   router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
//   router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
  http.ListenAndServe(":8080", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var posts []Post
  result, err := db.Query("SELECT id, name, city from employee")
  fmt.Print(err,result)
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title, &post.CITY)
    if err != nil {
      panic(err.Error())
    }
    posts = append(posts, post)
  }
  json.NewEncoder(w).Encode(posts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  stmt, err := db.Prepare("INSERT INTO employee(id,name,city) VALUES(?,?,?)")
  fmt.Println(err,stmt)
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  fmt.Print(err,body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  fmt.Println(keyVal,stmt)
  id:=keyVal["id"]
  name:=keyVal["name"]
  city:=keyVal["city"]
  _, err = stmt.Exec(id,name,city)
  fmt.Println(err)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New post was created")
  }


