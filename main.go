package main

import (
	"github.com/guaychou/flusher/cmd/server"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		TimestampFormat:  "Mon, 02 Jan 2006 15:04:05",
	})

}

func main() {
	flusher := server.Init()
	flusher.Run()
}
