package models

import "github.com/stripe/stripe-go/v76"

type PaymentIntent struct {
	stripe.PaymentIntent
}

type Customer struct {
	stripe.Customer
}
