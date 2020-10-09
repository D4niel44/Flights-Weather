package openweathermap

import (
	"fmt"
	"testing"

	g "myp/Tarea01/app/geo"
)

const testAPIKey = "fc249a737bd3772566de957539505054"

var api = &API{hash: testAPIKey}

// TestNewAPI tests OpenWeatherAPI object creation.
func TestNewAPI(t *testing.T) {
	testHash := "test"
	test := NewAPI(testHash)
	assertNotNil(test, t)
	if test.hash != testHash {
		t.Errorf("expected: %v, got %v", testHash, test.hash)
	}
}

// TestGetWeatherFromCity test GetWeatherFromCity method.
func TestGetWeatherFromCity(t *testing.T) {
	testCases := []struct {
		city  string
		units Units
		lang  Language
	}{
		{"Mexico City", METRIC, ES},
		{"Mexico City", IMPERIAL, ES},
		{"Mexico City", STANDARD, ES},
		{"Mexico City", METRIC, FR},
		{"Mexico City", METRIC, EN},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf(tc.city, tc.units, tc.lang)
		t.Run(testName, func(t *testing.T) {
			response, err := api.GetWeatherFromCity(tc.city, tc.units, tc.lang)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			assertNonNilResponse(response, t)
		})
	}
}

// TestGetWeatherFromCoordinates test GetWeatherFromCoordinates method.
func TestGetWeatherFromCoordinates(t *testing.T) {
	testCases := []struct {
		coordinate g.Coordinate
		units      Units
		lang       Language
	}{
		{g.Coordinate{0, 0}, METRIC, ES},
		{g.Coordinate{90, 180}, IMPERIAL, ES},
		{g.Coordinate{90, -180}, STANDARD, ES},
		{g.Coordinate{-90, 180}, METRIC, FR},
		{g.Coordinate{-90, -180}, METRIC, EN},
		{g.Coordinate{-5.67, 89}, METRIC, EN},
		{g.Coordinate{8.9, -86}, METRIC, EN},
		{g.Coordinate{-89, -76}, METRIC, EN},
		{g.Coordinate{89, 76}, METRIC, EN},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%.2f, %.2f, %v, %v", tc.coordinate.Lat, tc.coordinate.Lon, tc.units, tc.lang)
		t.Run(testName, func(t *testing.T) {
			response, err := api.GetWeatherFromCoordinates(tc.coordinate, tc.units, tc.lang)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			assertNonNilResponse(response, t)
		})
	}
}

func assertNonNilResponse(response *Weather, t *testing.T) {
	assertNotNil(response, t)
	assertNotNil(response.Weather, t)
	for _, v := range response.Weather {
		assertNotNil(v.ID, t)
		assertNotNil(v.Main, t)
		assertNotNil(v.Description, t)
	}
	assertNotNil(response.Main, t)
	assertNotNil(response.Main.Temp, t)
	assertNotNil(response.Main.TempMin, t)
	assertNotNil(response.Main.TempMax, t)
	assertNotNil(response.Main.FeelsLike, t)
	assertNotNil(response.Main.Humidity, t)
}

func assertNotNil(value interface{}, t *testing.T) {
	if value == nil {
		t.Errorf("unexpected nil value %T", value)
	}
}
