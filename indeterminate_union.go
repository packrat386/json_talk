package main

import (
	"encoding/json"
	"fmt"
)

var data = `{
    "payments": [
        {
            "id": "12345",
            "amount": "3.50",
            "type": "ach",
            "payment_details": {
                "routing_number": "123456789",
                "account_number": "987654321"
            }
        },
        {
            "id": "12345",
            "amount": "3.50",
            "type": "credit_card",
            "payment_details": {
                "card_number": "1111222233334444",
                "expiration": "0618"
            }
        }
    ]
}`

// STRUCTS OMIT
type PaymentCollection struct {
	Payments []Payment `json"payments"`
}

type Payment struct {
	ID             string `json:"id"`
	Amount         string `json:"amount"`
	Type           string `json:"type"`
	PaymentDetails Detail `json:"payment_details"`
}

type Detail struct {
	RoutingNumber string `json:"routing_number"`
	AccountNumber string `json:"account_number"`
	CardNumber    string `json:"card_number"`
	Expiration    string `json:"expiration"`
}

// END STRUCTS OMIT

// MAIN OMIT
func main() {
	pmts := new(PaymentCollection)
	err := json.Unmarshal([]byte(data), pmts)
	if err != nil {
		panic(err)
	}

	for _, p := range pmts.Payments {
		p.Execute()
	}
}

// END MAIN OMIT

// USAGE OMIT

func (p Payment) Execute() {
	if p.Type == "ach" {
		fmt.Printf(
			"Executing ach payment\nrouting_number: %s\naccount_number: %s\n",
			p.PaymentDetails.RoutingNumber,
			p.PaymentDetails.AccountNumber,
		)
	} else if p.Type == "credit_card" {
		fmt.Printf(
			"Executing credit card payment\ncard_number: %s\nexpiration: %s\n",
			p.PaymentDetails.CardNumber,
			p.PaymentDetails.Expiration,
		)
	} else {
		panic("unrecognized type")
	}
}

// END USAGE OMIT
