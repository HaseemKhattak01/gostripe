package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaseemKhattak01/gostripe/handlers"
)

func main() {
	// Define routes
	http.HandleFunc("/create-payment-intent", handlers.HandleCreatePaymentIntent)
	http.HandleFunc("/stripe-webhook", handlers.HandleStripeWebhook)

	// Start server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
