package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/HaseemKhattak01/gostripe/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/webhook"
)

const (
	stripeSecretKey     = "sk_test_51QnEUnKvqOfHz2kK63G5UjwWxRP1GByBsdKibgKqudCk7PGQtXa1GIFVJQMrjkEq7YsI1izWi9igLFxW4qEufUZj006ijq97DO"
	stripeWebhookSecret = "whsec_AMDPHqMp2x3pY9aBVBZwjawl1VFcdB65"
)

func init() {
	stripe.Key = stripeSecretKey
}

func CreatePaymentIntent(amount int64) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String("usd"),
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
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	// Verify webhook signature
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), stripeWebhookSecret)
	if err != nil {
		return fmt.Errorf("error verifying webhook signature: %v", err)
	}

	// Handle the event
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
