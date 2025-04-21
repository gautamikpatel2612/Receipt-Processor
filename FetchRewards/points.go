package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculatePoints takes a Receipt and returns the calculated points based on specific rules.
func CalculatePoints(r Receipt) int {
	points := 0 // Start with 0 points

	// 1. Alphanumeric characters in retailer name
	for _, c := range r.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			points++
		}
	}

	// 2. Round dollar amount
	if strings.HasSuffix(r.Total, ".00") {
		points += 50
	}

	// 3. Total is multiple of 0.25
	totalFloat, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		log.Printf("Error parsing total as float: %v", err)
	} else if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points per 2 items
	points += (len(r.Items) / 2) * 5

	// 5. Item descriptions length multiple of 3
	for _, item := range r.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("Error parsing item price: %v", err)
			} else {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// 6. Day of purchase is odd
	date, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		log.Printf("Error parsing purchase date: %v", err)
	} else if date.Day()%2 == 1 {
		points += 6
	}

	// 7. Time of purchase is between 2:00pm and 4:00pm (exclusive)
	purchaseTime, err := time.Parse("15:04", r.PurchaseTime)
	if err != nil {
		log.Printf("Error parsing purchase time: %v", err)
	} else {
		hour := purchaseTime.Hour()
		minute := purchaseTime.Minute()
		if (hour == 14 && minute >= 0) || (hour == 15 && minute < 60) {
			points += 10
		}
	}

	return points
}

// Rule 1: +1 point for every alphanumeric in retailer name
// Rule 2: +50 points if total is a round dollar (no cents)
// Rule 3: +25 points if total is multiple of 0.25
// Rule 4: +5 points for every two items
// Rule 5: +ceil(item price * 0.2) if trimmed item description length divisible by 3
// Rule 6: +6 points if purchase day is odd
// Rule 7: +10 points if time is between 2pm and 4pm (exclusive)
