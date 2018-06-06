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
            "to": "nessie",
            "from": "packrat386"
        },
        {
            "id": "12346",
            "amount": "0.02",
            "to": "packrat386",
            "from": "mlarraz"
        }
    ]
}`

// START OMIT
type PaymentCollection struct {
	Payments []Payment `json:"payments"`
}

type Payment struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
	To     string `json:"to"`
	From   string `json:"from"`
}

func main() {
	pmts := new(PaymentCollection)
	err := json.Unmarshal([]byte(data), pmts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", pmts)
}

// END OMIT
