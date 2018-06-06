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

type PaymentCollection struct {
	Payments []*Payment `json"payments"`
}

// STRUCTS OMIT

type Payment struct {
	paymentData
	Detail
}

type paymentData struct {
	ID     string          `json:"id"`
	Amount string          `json:"amount"`
	Type   string          `json:"type"`
	Detail json.RawMessage `json:"payment_details"`
}

type Detail interface {
	Execute()
}

// END STRUCTS OMIT

// UNMARSHAL OMIT

func (p *Payment) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &p.paymentData)
	if err != nil {
		return err
	}

	var detail Detail
	if p.Type == "ach" {
		detail = new(achDetail)
	} else if p.Type == "credit_card" {
		detail = new(cardDetail)
	} else {
		return fmt.Errorf("unrecognized type")
	}

	err = json.Unmarshal(p.paymentData.Detail, detail)
	if err != nil {
		return err
	}
	p.Detail = detail
	return nil
}

// END UNMARSHAL OMIT

// MARSHAL OMIT

func (p *Payment) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(p.Detail)
	if err != nil {
		return make([]byte, 0), err
	}

	p.paymentData.Detail = data

	return json.Marshal(p.paymentData)
}

// END MARSHAL OMIT

type achDetail struct {
	RoutingNumber string `json:"routing_number"`
	AccountNumber string `json:"account_number"`
}

func (a *achDetail) Execute() {
	fmt.Printf("Executing ach payment\nrouting_number: %s\naccount_number: %s\n", a.RoutingNumber, a.AccountNumber)
}

type cardDetail struct {
	CardNumber string `json:"card_number"`
	Expiration string `json:"expiration"`
}

func (c *cardDetail) Execute() {
	fmt.Printf("Executing cc payment\ncard_number: %s\nexpiration: %s\n", c.CardNumber, c.Expiration)
}

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

	output, err := json.Marshal(pmts)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

// END MAIN OMIT
