package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/KeisukeYamashita/kanna/controller"
	"go.uber.org/zap"
)

func main() {
	var outputType string
	var statusCode int
	var verbose bool
	flag.StringVar(&outputType, "type", "stdout", "Stdout the request")
	flag.IntVar(&statusCode, "statusCode", -1, "HTTP Status code for response")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	ctr := controller.NewController(logger, verbose)
	ctr.OutputType = outputType
	ctr.StatusCode = statusCode

	ctr.Logger.Info("server is ready to serve", zap.Int("port", 8080), zap.String("type", outputType))
	http.HandleFunc("/healthz", ctr.HealthCheckHandler())
	http.HandleFunc("/mirror", ctr.MirrorHandler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		ctr.Logger.Fatal("cannot listen port and server", zap.Int("port", 8080), zap.String("type", outputType))
	}
}
