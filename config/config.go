package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StripeSecretKey     string
	StripeWebhookSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
	if StripeSecretKey == "" {
		log.Fatal("STRIPE_SECRET_KEY must be set in the environment")
	}

	StripeWebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
	if StripeWebhookSecret == "" {
		log.Fatal("STRIPE_WEBHOOK_SECRET must be set in the environment")
	}
}
