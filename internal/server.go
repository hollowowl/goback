package internal

import (
	"arduino-playground.xyz/goback/config"
	"encoding/json"
	"net/http"
	"strings"
	"log"
)

type Server struct {
	storage *Storage
	config *config.ServerConfig
}

func NewServer(config *config.Config) (*Server, error) {
	storage, err := NewStorage(config.Database)
	if (err != nil) {
		return nil, err
	}
	server := &Server{}
	server.config = &config.Server
	server.storage = storage
	return server, nil
}

func (server *Server) Run() {
	http.Handle("/processData", processDataHandler(server))
	http.ListenAndServe(server.config.Addr, nil)
}


func processDataHandler(server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Palun, kasuta POST või söö sitta", http.StatusMethodNotAllowed)
			return
		}
		contentType := strings.ToLower(r.Header.Get("Content-Type"))
		if contentType != "" && contentType != "application/json" {
			http.Error(w, "Palun, kasuta Content-Type 'application/json' või söö sitta", http.StatusUnsupportedMediaType)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		var data Data
		//TODO ??? store request before decoding - how to get board name from header?
		if err := dec.Decode(&data); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
			return
		}
		dataID, err := server.storage.RegisterIncomingData(&data)
		if err != nil {
			log.Println("registerIncomingData error: ", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
			return
		}
		decision, err := server.makeDecision(dataID, &data)
		if err != nil {
			log.Println("makeDecision error: ", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
			return
		}
		encoded, err := json.Marshal(decision)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(encoded)
	})
}

func (server *Server) makeDecision(dataID *DataID, data *Data) (*Decision, error) {
	decision := Decision{}
	if data.SoilHumidity < 25 {
		decision.Action = TurnPumpForSec
		decision.Value = 2
	}
	if decision.Action != "" {
		if err := server.storage.RegisterDecision(dataID, &decision); err != nil {
			return nil, err
		}
	}
	return &decision, nil
}
