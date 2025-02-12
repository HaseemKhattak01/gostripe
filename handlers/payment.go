package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HaseemKhattak01/gostripe/services"
)

func HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req struct {
		Amount int64  `json:"amount"`
		Email  string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := services.CreateCustomer(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientSecret, err := services.CreatePaymentIntent(req.Amount, customer.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return PaymentIntent client_secret
	response := struct {
		ClientSecret string `json:"client_secret"`
	}{
		ClientSecret: clientSecret,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	err := services.HandleStripeWebhook(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
