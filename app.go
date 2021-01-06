package main

import (
	conf "go-rest-mongodb/config"
	"go-rest-mongodb/routers"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var config conf.Config

func init() {
	config.Read()
}

func main() {
	log.SetOutput(os.Stdout)
	logFormatter := new(conf.LogFormatter)
	logFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logFormatter.LevelDesc = []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
	log.SetFormatter(logFormatter)

	r := routers.Routers()
	srv := &http.Server{
		Handler:      r,
		Addr:         config.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
