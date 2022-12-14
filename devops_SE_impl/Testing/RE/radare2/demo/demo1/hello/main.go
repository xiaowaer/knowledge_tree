package main

//
import (
	"example/httprouter"
	"fmt"
	"log"
	"net/http"
)

//
  func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      fmt.Fprint(w, "Welcome!\n")
  }
//
  func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
      fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
  }
//
  func main() {
      router := httprouter.New()
      router.GET("/", Index)
      router.GET("/hello/:namb", Hello)
      fmt.Printf("aaa")
//
      log.Fatal(http.ListenAndServe(":8080", router))
  }
