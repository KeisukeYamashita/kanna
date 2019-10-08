package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/KeisukeYamashita/kanna/logger"
	"go.uber.org/zap"
)

// Controller ...
type Controller struct {
	Logger     *log.Logger
	OutputType string
	StatusCode int
}

// HealthCheckHandler ...
func (ctr *Controller) HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am healthy"))
	}
}

// NewController ...
func NewController(logger *zap.Logger, verbose bool) *Controller {
	return &Controller{
		Logger: &log.Logger{
			Logger:  logger,
			Verbose: verbose,
		},
	}
}

// MirrorHandler ...
func (ctr *Controller) MirrorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decorder := json.NewDecoder(r.Body)
		var body interface{}
		err := decorder.Decode(&body)
		if err != nil {
			ctr.writeHeader(w, http.StatusInternalServerError)
			w.Write([]byte("cannot get request body\n"))
			return
		}
		b, err := json.MarshalIndent(body, "", " ")
		if err != nil {
			ctr.writeHeader(w, http.StatusInternalServerError)
			w.Write([]byte("cannot parse request body\n"))
			return
		}

		switch ctr.OutputType {
		case "stdout":
			fmt.Println(string(b))
			ctr.writeHeader(w, http.StatusOK)
			w.Write([]byte("Okay\n"))
			return
		case "json":
			if err != nil {
				ctr.writeHeader(w, http.StatusInternalServerError)
				w.Write([]byte("cannot marshal request body\n"))
				return
			}
			date := time.Now().Format("20060102")
			err = ioutil.WriteFile(fmt.Sprintf("dist/kanna-%s.json", date), b, 0644)
			if err != nil {
				ctr.writeHeader(w, http.StatusInternalServerError)
				w.Write([]byte("cannot create request body\n"))
				return
			}
		default:
			ctr.Logger.Fatal("--type should be stdout or json", zap.String("type", ctr.OutputType))
			ctr.writeHeader(w, http.StatusInternalServerError)
			w.Write([]byte("I am mirror"))
			return
		}
	}
}

func (ctr *Controller) writeHeader(w http.ResponseWriter, statusCode int) {
	if 600 > ctr.StatusCode && ctr.StatusCode >= 0 {
		statusCode = ctr.StatusCode
	}
	w.WriteHeader(statusCode)
}
