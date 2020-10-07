// TODO package documentation explaining database structure

package openweathermapcities

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"unicode"

	// import the sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	city      = "city"     // city table
	airports  = "airports" // Airports table
	fieldName = "name"     // name field for both cities and airports
	state     = "state"    // state field for cities
	country   = "country"  // to letters identifier for table city, also country name for table airports.
	iata      = "iata"     // iata column name for airports table
)

type OwmCityConverter struct {
	citiesDB *sql.DB
}

// NewOwmCityConverter Tries to create a OwmCityconversor that works with the database in
// the given path.
// Note that the path must be a path to a sqlite database in the file
// sistem with the structure defined in this package documentation.
// Returns an error if the method fails to to load the database or the file.
func NewOwmCityConverter(dbPath string) (*OwmCityConverter, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	OwmCityConverter := OwmCityConverter{citiesDB: db}
	return &OwmCityConverter, nil
}

// Close closes the conection to the underlying database.
func (citiesConverter *OwmCityConverter) Close() {
	citiesConverter.citiesDB.Close()
}

// TODO explore changing coordinates  to its own module to avoid repetition.

// GetCityCoordinates returns the city coordinates related to the given string, if any.
// The string paramater may be both a city name or an airport IATA code.
// Returns an error if this method fails to connect to database.
func (citiesConverter *OwmCityConverter) GetCityCoordinates(name string) (*Coordinates, error) {
	if len(name) == 3 { // if three word length then string may be IATA code.
		coordinate, err := citiesConverter.getCoordinatesByAirportCode("iata", name)
		if err == nil {
			return coordinate, nil
		}
	}
	if len(name) == 4 { // in this case it may be ICAO code.
		coordinate, err := citiesConverter.getCoordinatesByAirportCode("icao", name)
		if err == nil {
			return coordinate, nil
		}
	}
	// If none of the above search by City Name or Airport Name.
	normalizedName := ToAlphaNumeric(name)
	var coordinate *Coordinates
	coordinate, err := citiesConverter.getCoordinatesFromDB("name", "airports", normalizedName)
	if err == nil {
		return coordinate, nil
	}
	coordinate, err = citiesConverter.getCoordinatesFromDB("name", "city", normalizedName)
	if err == nil {
		return coordinate, nil
	}
	// Last try to search by State (only USA states) or country and return first match
	nameUpper := strings.ToUpper(name)
	coordinate, err = citiesConverter.getCoordinatesFromDB("state", "city", nameUpper)
	if err == nil {
		return coordinate, nil
	}
	coordinate, err = citiesConverter.getCoordinatesFromDB("country", "city", normalizedName)
	if err == nil {
		return coordinate, nil
	}
	coordinate, err = citiesConverter.getCoordinatesFromDB("country", "airports", normalizedName)
	if err == nil {
		return coordinate, nil
	}
	// At this point we have nothing we can do to find the string.
	return nil, errors.New("could not find any related record to the string")
}

// gets the coordinate
func (citiesConverter *OwmCityConverter) getCoordinatesByAirportCode(code, value string) (*Coordinates, error) {
	capitalizedValue := strings.ToUpper(value)
	coordinate, err := citiesConverter.getCoordinatesFromDB(code, "airports", capitalizedValue)
	if err != nil {
		return nil, err
	}
	return coordinate, nil
}

// Method for querying lat and coordinates to the database.
// Returns the coordinates of the matching row of the query or an error if it was not found.
func (citiesConverter *OwmCityConverter) getCoordinatesFromDB(column, table, value string) (*Coordinates, error) {
	query := fmt.Sprintf(`SELECT lat, lon FROM %s WHERE %s='%s';`, table, column, value)
	row := citiesConverter.citiesDB.QueryRow(query)
	coordinates := Coordinates{}
	err := row.Scan(&coordinates.lat, &coordinates.lon)
	if err != nil {
		return nil, err
	}
	return &coordinates, nil
}

// ToAlphaNumeric converts the given string to an only alphanumeric  strings without spaces.
func ToAlphaNumeric(ugly string) string {
	ugly = strings.TrimSpace(ugly)
	ugly = strings.ToLower(ugly)
	t := transform.Chain(norm.NFD, runes.Remove(runes.Predicate(func(r rune) bool {
		return !unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Digit, r) || unicode.Is(unicode.Zs, r)
	})), norm.NFC)
	pretty, _, _ := transform.String(t, ugly)
	return pretty
}

type Coordinates struct {
	lat float32
	lon float32
}
