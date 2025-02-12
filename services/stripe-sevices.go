package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	config "github.com/HaseemKhattak01/gostripe/congfig"
	"github.com/HaseemKhattak01/gostripe/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/webhook"
)

func CreateCustomer(email string) (*models.Customer, error) {
	stripe.Key = config.StripeSecretKey

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	c, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	return &models.Customer{Customer: *c}, nil
}

func CreatePaymentIntent(amount int64, customerID string) (string, error) {
	stripe.Key = config.StripeSecretKey
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String("usd"),
		Customer: stripe.String(customerID),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}
	return pi.ClientSecret, nil
}

func HandleStripeWebhook(w http.ResponseWriter, r *http.Request) error {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	stripeWebhookSecret := config.StripeWebhookSecret
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), stripeWebhookSecret)
	if err != nil {
		return fmt.Errorf("error verifying webhook signature: %v", err)
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent models.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			return fmt.Errorf("error parsing payment_intent.succeeded: %v", err)
		}
		fmt.Println("PaymentIntent was successful!")
		// Add your business logic here
	default:
		fmt.Printf("Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
