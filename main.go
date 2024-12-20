package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, "fail", "Invalid JSON format")
		return
	}

	if requestData.Message == "" {
		respondJSON(w, http.StatusBadRequest, "fail", "Invalid JSON message")
		return
	}

	fmt.Printf("Message received: %s\n", requestData.Message)
	respondJSON(w, http.StatusOK, "success", "Data successfully received")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	respondJSON(w, http.StatusOK, "success", "Send a POST request with JSON data to test!")
}

func respondJSON(w http.ResponseWriter, statusCode int, status, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ResponseData{
		Status:  status,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/get", handleGet)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
