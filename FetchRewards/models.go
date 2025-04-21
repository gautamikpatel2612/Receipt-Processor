// Declares that this file belongs to the "main" package. Required for running the
// app as an executable.
package main

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Receipt and Item structs
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}
