# gostripe

This project demonstrates how to integrate Stripe payments and webhooks in a Golang backend.

## Setup
1. Replace `stripeSecretKey` and `stripeWebhookSecret` in `services/stripe.go` with your Stripe keys.
2. Run the server:
   ```bash
   go run main.go