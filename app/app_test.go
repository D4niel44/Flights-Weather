package main

import (
	"testing"

	g "myp/Tarea01/app/geo"
)

const (
	testAPIKey      = "fc249a737bd3772566de957539505054"
	testDBPath      = "../bin/db/cities.db"
	testDatasetPath = "../testDatasets/"
)

var app = NewApp(testAPIKey, testDBPath, 40)

func TestDataset1(t *testing.T) {
	testDataset("dataset1.csv", t)
}

func TestDataset2(t *testing.T) {
	testDataset("dataset2.csv", t)
}

func testDataset(dataset string, t *testing.T) {
	flights, cities := app.HandleDataSet(testDatasetPath + dataset)
	app.QueryWeather(cities)
	PrintWeather(flights)
}

// TestWueryTest runs simple test to check QueryTest is working
func TestQuerySet(t *testing.T) {
	cityName := "Mexico City"
	city := City{
		name: cityName,
		coordinate: &g.Coordinate{
			Lat: 19,
			Lon: -99,
		},
	}
	cities := make(map[string]*City)
	cities[cityName] = &city
	t.Run("Simple Test", func(t *testing.T) {
		app.QueryWeather(cities)
		if cities[cityName].weather == nil {
			t.Error("Expected non empty weather")
		}
	})
}
