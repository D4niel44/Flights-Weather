package main

import (
	"testing"

	g "myp.ciencias.unam.mx/geo"
)

const (
	testAPIKey      = "fc249a737bd3772566de957539505054"
	testDBPath      = "../../../bin/db/cities.db"
	testDatasetPath = "../../../testDatasets/"
)

var app = NewApp(testAPIKey, testDBPath)

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

// TODO Improve tests to check no ban when more than 60 cities are requested.
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
