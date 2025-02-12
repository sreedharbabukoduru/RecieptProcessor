package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
)

var (
	receipts = make(map[string]models.Receipt)
	mu       sync.Mutex
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	id := uuid.New().String()
	receipts[id] = receipt

	points := services.CalculatePoints(receipt)
	response := map[string]string{"id": id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/receipts/"):len("/receipts/")+36]

	mu.Lock()
	defer mu.Unlock()

	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := services.CalculatePoints(receipt)
	response := map[string]int{"points": points}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}