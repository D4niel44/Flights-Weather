package main

import (
	"testing"
)

const (
	testAPIKey = "fc249a737bd3772566de957539505054"
)

// TODO Improve tests to check no ban when more than 60 cities are requested.
// TestWueryTest runs simple test to check QueryTest is working
func TestQuerySet(t *testing.T) {
	cityName := "Mexico City"
	city := City{
		name: &cityName,
		coordinates: Coordinates{
			lat: 19,
			lon: -99,
		},
	}
	app := NewApp(testAPIKey)
	cities := make(map[string]*City)
	cities[cityName] = &city
	t.Run("Simple Test", func(t *testing.T) {
		app.QueryWeather(&cities)
		if cities[cityName].weather == nil {
			t.Error("Expected non empty weather")
		}
	})
}
