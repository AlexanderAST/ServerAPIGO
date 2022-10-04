package main

import (
	"ServerApi/internal/config"
	"ServerApi/internal/user"
	"ServerApi/internal/user/db"
	"ServerApi/pkg/client/mongodb"
	"ServerApi/pkg/logging"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()
	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)
	user1 := user.User{
		ID:           "",
		Email:        "mymail@example.com",
		Username:     "Zolafarre",
		PasswordHash: "143311",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	logger.Info(user1ID)
	user2 := user.User{
		ID:           "",
		Email:        "koding00@bk.ru",
		Username:     "Alexander Astakhov",
		PasswordHash: "1355842428437525",
	}
	user2ID, err := storage.Create(context.Background(), user2)
	if err != nil {
		panic(err)
	}
	logger.Info(user2ID)

	logger.Info("create router handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}
func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")
	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("create unix socket ")

		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server listening port %s", socketPath)

	} else {
		logger.Info("listener socket ")

		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
