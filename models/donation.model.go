package models

import "strconv"

// DonationProcessor handles the processing of donations
type DonationProcessor struct {
	TotalReceived       float64
	SuccessfullyDonated float64
	FaultyDonation      float64
	Donations           []ChargeRequest
}

type Donation struct {
	DonorName string
	Amount    float64
}

type ChargeRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description,omitempty"`
	Card        string  `json:"card,omitempty"`
	Customer    string  `json:"customer,omitempty"`
	Capture     bool    `json:"capture,omitempty"`
}

type ChargeResponse struct {
	Status      string `json:"status"`
	FailureCode string `json:"failure_code,omitempty"`
	Description string `json:"description,omitempty"`
}

func ParseAmount(amountStr string) (float64, error) {
	return strconv.ParseFloat(amountStr, 64)
}
