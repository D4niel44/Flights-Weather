package openweathermapcities

import (
	"testing"
)

// TODO remaining tests
func TestGetCityCoordinates(t *testing.T) {
	testCases := []struct {
		testName string
		query    string
	}{
		{"City Name", "LaPas"},
		{"IATA", "MEX"},
		{"ICAO", "PAVD"},
		{"Airport Name", "gambellaairport"},
		{"State", "NY"},
		{"Country Code", "US"},
		{"Country", "brazil"},
	}
	// Assumes NewOwmCityConverter works
	cv, err := NewOwmCityConverter("../../../bin/db/cities.db")
	defer cv.Close()
	if err != nil {
		t.Errorf("Error while conecting to DB. \n %v", err)
	}
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			coordinate, err := cv.GetCityCoordinates(testCase.query)
			if err != nil {
				t.Errorf("failed to retrieve coordinates for %s", testCase.query)
			}
			if coordinate == nil {
				t.Errorf("Unexpected nil return.")
			}
		})
	}
}
