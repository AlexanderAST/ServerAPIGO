package main

import (
	"ServerApi/internal/user"
	"ServerApi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	logger.Info("create router handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)

}
func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("server listening port 8080")
	logger.Fatal(server.Serve(listener))
}
