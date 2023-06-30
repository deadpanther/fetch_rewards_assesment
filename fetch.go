package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptsMap map[string]int

var receipts ReceiptsMap

func main() {
	receipts = make(ReceiptsMap)
	router := gin.Default()
	router.POST("/receipts/process", processReceipts)
	router.Run(":5000")
}

func processReceipts(c *gin.Context) {
	var receipt Receipt
	err := c.ShouldBindJSON(&receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if receipt.Items == nil || len(receipt.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt cannot be made without Items"})
		return
	}
	if receipt.PurchaseTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt should have a purchase time"})
		return
	}
	if receipt.PurchaseDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt should have a purchase date"})
		return
	}
	if receipt.Total == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt should have a total"})
		return
	}
	if receipt.Retailer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt should have a retailer"})
		return
	}
	if len(receipt.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt should have at least one item"})
		return
	}
	receiptID := uuid.New().String()
	points := calculatePoints(receipt)
	receipts[receiptID] = points

	//pass back the receiptId and points
	c.JSON(http.StatusOK, gin.H{"receiptId": receiptID, "points": points})
}

func calculatePoints(receipt Receipt) int {
	points := 0
	points += rule1(receipt.Retailer)
	points += rule2(receipt.Total)
	points += rule3(receipt.Total)
	points += rule4(receipt.Items)
	points += rule5(receipt.Items)
	points += rule6(receipt.PurchaseDate)
	points += rule7(receipt.PurchaseTime)
	return points
}

func rule1(retailer string) int {
	// 1 point for each alphanumeric character in the retailer name
	count := 0
	for _, ch := range retailer {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			count++
		}
	}
	return count
}

func rule2(total string) int {
	// 50 points if the total is a round dollar amount with no cents.
	points := 0
	parsedTotal, err := strconv.ParseFloat(total, 64)
	if err == nil && parsedTotal == float64(int(parsedTotal)) {
		points += 50
	}
	return points
}

func rule3(total string) int {
	// 25 points if the total is a multiple of 0.25
	points := 0
	parsedTotal, err := strconv.ParseFloat(total, 64)
	if err == nil && math.Mod(parsedTotal*100, 25) == 0 {
		points += 25
	}
	return points
}

func rule4(items []Item) int {
	// 5 points for every two items on the receipt
	points := 0
	if len(items) > 0 {
		points += len(items) / 2 * 5
	}
	return points
}

func rule5(items []Item) int {
	//If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	points := 0
	for _, item := range items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}
	return points
}

func rule6(purchaseDate string) int {
	// 6 points if the day of the month is odd
	points := 0
	day, err := strconv.Atoi(strings.Split(purchaseDate, "-")[2])
	if err == nil && day%2 != 0 {
		points += 6
	}
	return points
}

func rule7(purchaseTime string) int {
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	points := 0
	hour, err := strconv.Atoi(strings.Split(purchaseTime, ":")[0])
	if err == nil && hour >= 14 && hour < 16 {
		points += 10
	}
	return points
}
