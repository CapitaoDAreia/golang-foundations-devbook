package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		ret := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := ret.Area()

		if expectedArea != receivedArea {
			t.Fatal("Unexpected Result.")
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Unexpected Result.")
		}
	})
}
