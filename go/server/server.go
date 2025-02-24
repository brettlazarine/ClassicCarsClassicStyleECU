package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func RunServer() error {
	const PORT = "8080"
	mux := http.NewServeMux()

	// Root, verify communication
	mux.HandleFunc("GET /", handler_verify)

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	log.Printf("Server running on port %s", PORT)
	log.Fatal(srv.ListenAndServe())

	defer srv.Close()

	return http.ListenAndServe(":"+PORT, nil)
}

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

// This is boilerplate, may require modification for my calls
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling JSON: %v", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

func handler_verify(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Connected to CCCS server!")
}
