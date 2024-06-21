package routes

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Mrdiamond096/go-challenge/services"
)

func StartServer() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go-tamboon <path-to-encrypted-csv>")
	}

	filePath := os.Args[1]
	rotReader := services.OpenFileCsv(filePath)

	scanner := bufio.NewScanner(rotReader)
	donationProcessor := services.NewDonationProcessor()

	for scanner.Scan() {
		line := scanner.Text()
		if err := donationProcessor.ProcessDonation(line); err != nil {
			log.Printf("Failed to process donation: %v", err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	summary := donationProcessor.Summarize()
	fmt.Println(summary)
}
