package main

import (
	"ServerApi/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("create router handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)

}
func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server listening port 8080")
	log.Fatal(server.Serve(listener))
}
