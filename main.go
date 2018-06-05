package main

import (
	"log"
	"net/http"

	"github.com/gorkemsari/golang-rest-api-httprouter/route"
	"github.com/gorkemsari/golang-rest-api-httprouter/repository"
)

func main() {

	repository.InitDB("mysql", "username:password@tcp(ipaddress:port)/dbname")
	router := route.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}