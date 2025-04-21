package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

var receiptStore = make(map[string]Receipt)
var pointsStore = make(map[string]int)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		return
	}

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Printf("[ERROR] Failed to decode receipt JSON: %v", err)
		return
	}

	// Validate receipt
	if err := validateReceipt(receipt); err != nil {
		http.Error(w, fmt.Sprintf("Invalid receipt: %v", err), http.StatusBadRequest)
		log.Printf("[ERROR] Receipt validation failed: %v", err)
		return
	}

	// Generate ID and store the receipt
	id := uuid.New().String()
	receiptStore[id] = receipt
	log.Printf("[INFO] Stored receipt with ID: %s | Data: %+v", id, receipt)

	// Calculate points and store them
	points := CalculatePoints(receipt)
	pointsStore[id] = points

	// Respond with the ID
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"id": id}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("[ERROR] Failed to write JSON response: %v", err)
	}
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	if id == "" {
		http.Error(w, "Missing receipt ID", http.StatusBadRequest)
		log.Println("[ERROR] Missing receipt ID in path")
		return
	}

	points, exists := pointsStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		log.Printf("[WARN] Receipt ID not found: %s", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]int{"points": points}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("[ERROR] Failed to encode points response: %v", err)
	}
}

func validateReceipt(r Receipt) error {
	log.Printf("[INFO] Validating receipt: %+v", r)

	if strings.TrimSpace(r.Retailer) == "" {
		log.Println("[ERROR] Retailer is empty")
		return fmt.Errorf("retailer is required")
	}
	if strings.TrimSpace(r.PurchaseDate) == "" {
		log.Println("[ERROR] PurchaseDate is empty")
		return fmt.Errorf("purchase date is required")
	}
	if strings.TrimSpace(r.PurchaseTime) == "" {
		log.Println("[ERROR] PurchaseTime is empty")
		return fmt.Errorf("purchase time is required")
	}
	if strings.TrimSpace(r.Total) == "" {
		log.Println("[ERROR] Total is empty")
		return fmt.Errorf("total amount is required")
	}
	if len(r.Items) == 0 {
		log.Println("[ERROR] Items list is empty")
		return fmt.Errorf("at least one item is required")
	}

	// Check valid date format
	if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil {
		log.Printf("[ERROR] Invalid date format: %v", err)
		return fmt.Errorf("invalid purchase date format (expected YYYY-MM-DD)")
	}

	// Check valid time format
	if _, err := time.Parse("15:04", r.PurchaseTime); err != nil {
		log.Printf("[ERROR] Invalid time format: %v", err)
		return fmt.Errorf("invalid purchase time format (expected HH:MM in 24-hour format)")
	}

	for i, item := range r.Items {
		if strings.TrimSpace(item.ShortDescription) == "" {
			log.Printf("[ERROR] Item %d missing short description", i+1)
			return fmt.Errorf("item %d missing short description", i+1)
		}
		if strings.TrimSpace(item.Price) == "" {
			log.Printf("[ERROR] Item %d missing price", i+1)
			return fmt.Errorf("item %d missing price", i+1)
		}
	}

	log.Println("[INFO] Receipt validated successfully")
	return nil
}
