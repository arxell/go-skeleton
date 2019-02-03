package main

import (
	"compress/gzip"
	"fmt"
	"go-skeleton/src/config"
	"go-skeleton/src/helpers"
	"go-skeleton/src/server"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/google/logger"
)

func main() {
	fmt.Println(helpers.Add(1, 2))
	logger.Info("Start!")

	logger.Info("Create Config!")
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logger.Info("Create AppServer!")
	appserver, err := server.NewServer(c)
	if err != nil {
		panic(err)
	}

	logger.Info("Create Gzip handler!")
	handler, err := gziphandler.NewGzipLevelHandler(gzip.BestSpeed)
	if err != nil {
		panic(err)
	}

	logger.Info("Create server!")
	serv := http.Server{
		Addr:    c.Hostport,
		Handler: handler(appserver),
	}

	logger.Info("Start server!")
	err = serv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
