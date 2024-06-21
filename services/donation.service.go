package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mrdiamond096/go-challenge/cipher"
	"github.com/Mrdiamond096/go-challenge/models"
)

type DonationProcessor struct {
	*models.DonationProcessor // Embedding the imported struct
}

func NewDonationProcessor() *DonationProcessor {
	return &DonationProcessor{
		DonationProcessor: &models.DonationProcessor{}, // Initialize the embedded struct
	}
}

func OpenFileCsv(filePath string) *cipher.Rot128Reader {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	rotReader, err := cipher.NewRot128Reader(file)
	if err != nil {
		log.Fatalf("Failed to create Rot128 reader: %v", err)
	}
	return rotReader
}

func (dp *DonationProcessor) ProcessDonation(data string) error {
	fmt.Println("🚀  data:", data)
	parts := strings.Split(data, ",")
	amount, err := models.ParseAmount(parts[1])
	if err != nil {
		fmt.Println("🚀  response: failed", err)
	}

	// donation := models.Donation{
	// 	DonorName: parts[0],
	// 	Amount:    amount,
	// }
	chargeRequest := models.ChargeRequest{
		Amount:   amount,
		Currency: "THB",
		Card:     parts[2], // Example token; replace with your token
		Capture:  true,
	}

	response, err := Charge(chargeRequest)
	fmt.Println("🚀  response:", response)
	if err != nil {
		dp.FaultyDonation += chargeRequest.Amount
		return err
	}
	dp.SuccessfullyDonated += chargeRequest.Amount
	dp.Donations = append(dp.Donations, chargeRequest)
	return nil
}

func (dp *DonationProcessor) Summarize() string {
	return fmt.Sprintf(`
Total Received: THB %.2f
Successfully Donated: THB %.2f
Faulty Donation: THB %.2f
Average per Person: THB %.2f`,
		dp.TotalReceived, dp.SuccessfullyDonated, dp.FaultyDonation,
		dp.SuccessfullyDonated/float64(len(dp.Donations)))
}
