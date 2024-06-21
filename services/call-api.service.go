package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Mrdiamond096/go-challenge/models"
)

const chargeAPIURL = "https://docs.opn.ooo/charges-api"

func Charge(donation models.ChargeRequest) (*ChargeResponse, error) {
	secretKey := "skey_test_no1t4tnemucod0e51mo"
	data, err := json.Marshal(donation)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", chargeAPIURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(secretKey, "")

	resp, err := client.Do(req)
	// resp, err := http.Post(chargeAPIURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChargeResponse struct {
	Success bool
	Message string
}
