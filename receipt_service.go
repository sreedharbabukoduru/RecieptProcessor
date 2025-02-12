package services

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"receipt-processor/internal/models"
)

type ReceiptService struct {
	receipts map[string]*models.Receipt
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{
		receipts: make(map[string]*models.Receipt),
	}
}

func (s *ReceiptService) ProcessReceipt(receipt *models.Receipt) (string, error) {
	if receipt == nil {
		return "", errors.New("receipt cannot be nil")
	}

	id := generateID()
	s.receipts[id] = receipt
	return id, nil
}

func (s *ReceiptService) GetPoints(id string) (int, error) {
	receipt, exists := s.receipts[id]
	if !exists {
		return 0, errors.New("receipt not found")
	}

	points := 0
	points += len(strings.ReplaceAll(receipt.Retailer, " ", "")) // Alphanumeric characters in retailer name

	if isRoundDollar(receipt.Total) {
		points += 50
	}
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}
	points += (len(receipt.Items) / 2) * 5 // 5 points for every two items

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	if total, _ := strconv.ParseFloat(receipt.Total, 64); total > 10.00 {
		points += 5
	}

	if day := extractDay(receipt.PurchaseDate); day%2 != 0 {
		points += 6
	}

	if time := extractTime(receipt.PurchaseTime); time >= 14 && time < 16 {
		points += 10
	}

	return points, nil
}

func generateID() string {
	// Implementation for generating a unique ID
}

func isRoundDollar(total string) bool {
	// Implementation to check if total is a round dollar amount
}

func isMultipleOfQuarter(total string) bool {
	// Implementation to check if total is a multiple of 0.25
}

func extractDay(date string) int {
	// Implementation to extract the day from the purchase date
}

func extractTime(time string) int {
	// Implementation to extract the hour from the purchase time
}