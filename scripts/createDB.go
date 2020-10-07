package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type JsonCity struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	Coordinates struct {
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
	} `json:"coord"`
}

type JsonAirport struct {
	City    string `json:"city"`
	Country string `json:"country"`
	IATA    string `json:"iata"`
	ICAO    string `json:"icao"`
	Lat     string `json:"latitude"`
	Lon     string `json:"longitude"`
	Name    string `json:"name"`
}

const (
	citiesDataSource   = "data_sources/city.list.json"
	airportsDataSource = "data_sources/airports.json"
)

// TODO Refactor this.
// make this look nicer
func main() {
	// Open File for cities
	citiesFile, err := os.Open(citiesDataSource)
	if err != nil {
		panic("Error opening cities json")
	}
	defer citiesFile.Close()

	// Decode cities file into memory
	var cities []JsonCity
	err = json.NewDecoder(citiesFile).Decode(&cities)
	if err != nil {
		panic("Error at decoding cities file.")
	}

	// crete db directory
	baseDir := "../bin/db/"
	dbPath := baseDir + "cities.db"
	os.MkdirAll(baseDir, 0777)
	os.Create(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic("error opening db")
	}
	defer db.Close()

	// Create table for cities
	citiesTable := `CREATE TABLE city (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		country TEXT NOT NULL,
		lat REAL NOT NULL,
		lon REAL NOT NULL
	);
	`
	statement, err := db.Prepare(citiesTable)
	if err != nil {
		panic("Error in query preparing")
	}
	statement.Exec()

	transaction, _ := db.Begin()
	for _, city := range cities {
		insertCity := fmt.Sprintf(`INSERT INTO city(id, name, country, lat, lon) VALUES (%.0f, "%s", "%s", %.2f, %.2f);`,
			city.ID, toAlphaNumeric(city.Name), toAlphaNumeric(city.Country), city.Coordinates.Lat, city.Coordinates.Lon)
		_, err := transaction.Exec(insertCity)
		if err != nil {
			transaction.Rollback()
			panic(fmt.Sprintf("Error while inserting city. \n %s \n %v", insertCity, err))
		}
	}
	transaction.Commit()

	airportsFile, err := os.Open(airportsDataSource)
	if err != nil {
		panic("Error opening cities json")
	}
	defer airportsFile.Close()

	// Decode airports file into memory
	var airports map[string]JsonAirport
	err = json.NewDecoder(airportsFile).Decode(&airports)
	if err != nil {
		panic(fmt.Sprintf("Error at decoding airports file. \n %v", err))
	}

	// Create table for Airports
	airportsTable := `CREATE TABLE airports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		city INTEGER,
		country TEXT NOT NULL,
		iata TEXT NOT NULL,
		icao TEXT NOT NULL,
		lat REAL NOT NULL,
		lon REAL NOT NULL,
		name TEXT NOT NULL
	);
	`
	statement, err = db.Prepare(airportsTable)
	if err != nil {
		panic(fmt.Sprintf("Error in query preparing \n %s \n %v", airportsTable, err))
	}
	statement.Exec()

	transaction, _ = db.Begin()
	for _, airport := range airports {
		cityName := toAlphaNumeric(airport.City)
		row := transaction.QueryRow(fmt.Sprintf("SELECT id FROM city WHERE name='%s';", cityName))
		var cityID string
		err = row.Scan(&cityID)
		if err != nil {
			fmt.Printf("City not found for %s. Setting value to null \n", cityName)
			cityID = "null" // If no equivalent openweathermap city then set the value to null in the table
		}
		insertAirport := fmt.Sprintf(`INSERT INTO airports (city, country, iata, icao, lat, lon, name) VALUES (%s, "%s", "%s", "%s", %s, %s, "%s");`,
			cityID,
			toAlphaNumeric(airport.Country),
			airport.IATA,
			airport.ICAO,
			airport.Lat,
			airport.Lon,
			toAlphaNumeric(airport.Name))
		_, err := transaction.Exec(insertAirport)
		if err != nil {
			transaction.Rollback()
			panic(fmt.Sprintf("Error while inserting airport. \n %s \n %v", insertAirport, err))
		}
	}
	transaction.Commit()

}

func toAlphaNumeric(ugly string) string {
	ugly = strings.TrimSpace(ugly)
	ugly = strings.ToLower(ugly)
	t := transform.Chain(norm.NFD, runes.Remove(runes.Predicate(func(r rune) bool {
		return !unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Digit, r) || unicode.Is(unicode.Zs, r)
	})), norm.NFC)
	pretty, _, _ := transform.String(t, ugly)
	return pretty
}
