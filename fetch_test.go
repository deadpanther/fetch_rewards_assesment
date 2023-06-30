package main

import (
	"testing"
)

func TestRule1(t *testing.T) {
	retailer := "ABC Store"
	expectedPoints := 8

	points := rule1(retailer)
	if points != expectedPoints {
		t.Errorf("Rule 1 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule2(t *testing.T) {
	total := "100.00"
	expectedPoints := 50

	points := rule2(total)
	if points != expectedPoints {
		t.Errorf("Rule 2 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule3(t *testing.T) {
	total := "25.00"
	expectedPoints := 25

	points := rule3(total)
	if points != expectedPoints {
		t.Errorf("Rule 3 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule4(t *testing.T) {
	items := []Item{
		{ShortDescription: "Item 1", Price: "10.00"},
		{ShortDescription: "Item 2", Price: "20.00"},
		{ShortDescription: "Item 3", Price: "30.00"},
	}
	expectedPoints := 5

	points := rule4(items)
	if points != expectedPoints {
		t.Errorf("Rule 4 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule5(t *testing.T) {
	items := []Item{
		{ShortDescription: "Item 1", Price: "10.00"},
		{ShortDescription: "Item 2", Price: "20.00"},
		{ShortDescription: "Item 3", Price: "30.00"},
		{ShortDescription: "Item 4", Price: "40.00"},
	}
	expectedPoints := 20

	points := rule5(items)
	if points != expectedPoints {
		t.Errorf("Rule 5 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule6(t *testing.T) {
	purchaseDate := "2023-06-28"
	expectedPoints := 0

	points := rule6(purchaseDate)
	if points != expectedPoints {
		t.Errorf("Rule 6 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestRule7(t *testing.T) {
	purchaseTime := "15:30"
	expectedPoints := 10

	points := rule7(purchaseTime)
	if points != expectedPoints {
		t.Errorf("Rule 7 failed. Expected points: %d, got: %d", expectedPoints, points)
	}
}
