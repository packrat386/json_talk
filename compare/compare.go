package compare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type NaturalPaymentCollection struct {
	Payments []*NaturalPayment `json"payments"`
}

type NaturalPayment struct {
	paymentData
	Detail
}

func (p *NaturalPayment) UnmarshalJSON(data []byte) error {
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

func (p *NaturalPayment) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(p.Detail)
	if err != nil {
		return make([]byte, 0), err
	}

	p.paymentData.Detail = data

	return json.Marshal(p.paymentData)
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

type achDetail struct {
	RoutingNumber string `json:"routing_number"`
	AccountNumber string `json:"account_number"`
}

func (a *achDetail) Execute() {
	fmt.Fprintf(ioutil.Discard, "Executing ach payment\nrouting_number: %s\naccount_number: %s\n", a.RoutingNumber, a.AccountNumber)
}

type cardDetail struct {
	CardNumber string `json:"card_number"`
	Expiration string `json:"expiration"`
}

func (c *cardDetail) Execute() {
	fmt.Fprintf(ioutil.Discard, "Executing cc payment\ncard_number: %s\nexpiration: %s\n", c.CardNumber, c.Expiration)
}

type MapPaymentCollection struct {
	Payments []MapPayment `json"payments"`
}

type MapPayment struct {
	ID             string                 `json:"id"`
	Amount         string                 `json:"amount"`
	Type           string                 `json:"type"`
	PaymentDetails map[string]interface{} `json:"payment_details"`
}

func (p MapPayment) Execute() {
	if p.Type == "ach" {
		routingNumber, ok := p.PaymentDetails["routing_number"].(string)
		if !ok {
			panic("routing number not a string")
		}

		accountNumber, ok := p.PaymentDetails["account_number"].(string)
		if !ok {
			panic("account number not a string")
		}

		fmt.Fprintf(
			ioutil.Discard,
			"Executing ach payment\nrouting_number: %s\naccount_number: %s\n",
			routingNumber,
			accountNumber,
		)
	} else if p.Type == "credit_card" {
		cardNumber, ok := p.PaymentDetails["card_number"].(string)
		if !ok {
			panic("routing number not a string")
		}

		expiration, ok := p.PaymentDetails["expiration"].(string)
		if !ok {
			panic("account number not a string")
		}

		fmt.Fprintf(
			ioutil.Discard,
			"Executing credit card payment\ncard_payment: %s\nexpiration: %s\n",
			cardNumber,
			expiration,
		)
	} else {
		panic("unrecognized type")
	}
}
