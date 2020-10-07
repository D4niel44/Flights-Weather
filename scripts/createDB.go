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

const (
	citiesDataSource   = "data_sources/city.list.json"
	airportsDataSource = "data_sources/airports.json"
)

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
		lon REAL NOT NULL,
		lat REAL NOT NULL
	);
	`
	statement, err := db.Prepare(citiesTable)
	if err != nil {
		panic("Error in query preparing")
	}
	statement.Exec()

	for _, city := range cities {
		insertCity := fmt.Sprintf(`INSERT INTO city(id, name, country, lon, lat) VALUES (%.0f, "%s", "%s", %.2f, %.2f);`,
			city.ID, toAlphaNumeric(city.Name), toAlphaNumeric(city.Country), city.Coordinates.Lat, city.Coordinates.Lon)
		statement, err := db.Prepare(insertCity)
		if err != nil {
			panic(fmt.Sprintf("Error while inserting city. \n %s \n %v", insertCity, err))
		}
		statement.Exec()
	}
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
