package main

// TODO Investigate separating module app from main.
import (
	"fmt"
	"time"

	owm "myp.ciencias.unam.mx/proyecto1/openweathermap"
)

// Runs the app.
func main() {
	panic("not implemented") // TODO
}

// App struct that represents an App for querying cities weather.
type App struct {
	API *owm.API
}

// NewApp Creates A New App
func NewApp(apiKey string) *App {
	return &App{API: owm.NewAPI(apiKey)}
}

// TODO Improve documentation.

// HandleDataSet Process the file in the given path and returns a a list of all flights
// in the dataset, and a map of unique cities by name to its respective city
// structure.
//
// This method requires the input file to be in .csv format specified
// by RFC 4180.
//
// The dataset must contain cloumns with some of the following data of origin a destination cities.
// origin|origen name of the origin city.
// destination|destino name of the destination city.
// origin_latitude|latitud_origen origin city latitude coordinate
// origin_longitude|longitud_origen origin city longitude coordinate
// destination_latitude|latitud_destino destination city latitude coordinate
// destination_longitude|longitud_destino destination city longitude coordinate
// This method will try to return the names in a OpenWeatherMap call ready string.
func (app *App) HandleDataSet(dataset string) (*[]*Flight, *map[string]*City) {
	panic("not implemented") // TODO
}

// QueryWeather Requests the weather for all cities in the map and saves the result in
// the wethaer field of city.
//
// Requires the weather field of city to be nil.
func (app *App) QueryWeather(cities *map[string]*City) {
	queriesCounter := 0
	for cityName, city := range *cities {
		if queriesCounter > 55 {
			time.Sleep(time.Minute)
			queriesCounter = 0
		}
		cityWeather, err := app.API.GetWeatherFromCity(cityName, owm.METRIC, owm.ES)
		if err != nil {
			fmt.Printf("ERROR: %v, \n", err) // TODO Rework this using logger
		} else {
			city.weather = cityWeather
		}
		queriesCounter++
	}
}

// TODO improve documentation.

// PrintWeather Prints the weather to standard output
func PrintWeather(*[]*Flight) {
	panic("not implemented") // TODO
}

// Flight is a struct representing a flight
type Flight struct {
	origin      *City
	destination *City
}

// City is a struct representing a city
type City struct {
	name        *string
	coordinates Coordinates
	weather     *owm.Weather
}

// Coordinates is a struct representing coordinates
type Coordinates struct {
	lat float32
	lon float64
}
