// TODO package documentation explaining database structure

package openweathermapCities

import (
	"database/sql"

	// import the sqlite driver
	_ "github.com/mattn/go-sqlite3"
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

// GetCityID returns the city ID related to the given string, if any.
// The string paramater may be both a city name or an airport IATA code.
func (citiesConverter *OwmCityConverter) GetCityID(name int) {
	panic("not implemented") // TODO
}

func (citiesConverter *OwmCityConverter) makeQuery(column, match string) {
	panic("not implemented") // TODO
}
