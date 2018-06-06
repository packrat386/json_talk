package compare

import (
	"encoding/json"
	"testing"
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
        },
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
        },
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
        },
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
        },
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

// BENCH OMIT

func BenchmarkNatural(b *testing.B) {
	for n := 0; n < b.N; n++ {
		col := new(NaturalPaymentCollection)
		json.Unmarshal([]byte(data), col)
		for _, p := range col.Payments {
			p.Execute()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		col := new(MapPaymentCollection)
		json.Unmarshal([]byte(data), col)
		for _, p := range col.Payments {
			p.Execute()
		}
	}
}

// END BENCH OMIT
