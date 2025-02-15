package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	g "myp/Tarea01/app/geo"
	owm "myp/Tarea01/app/openweathermap"
)

// Flight is a struct representing a flight
type Flight struct {
	origin      *City
	destination *City
}

// City is a struct representing a city
type City struct {
	name       string
	coordinate *g.Coordinate
	weather    *owm.Weather
}

var defaultCity = City{
	name:       "Mexico City",
	coordinate: &g.Coordinate{Lat: 19.4363, Lon: -99.072098},
}

// App struct that represents an App for querying cities weather.
type App struct {
	API                 *owm.API
	DB                  *owm.OwmCityConverter
	QueriesCounter      int
	MaxQueriesPerMinute int
}

// NewApp Creates A New App
func NewApp(apiKey, dbPath string, maxQueriesPerMinute int) *App {
	db, _ := owm.NewOwmCityConverter(dbPath)
	return &App{API: owm.NewAPI(apiKey), DB: db, QueriesCounter: 0, MaxQueriesPerMinute: maxQueriesPerMinute}
}

// Close closes the underlying database connection, neccessary to run this app.
func (app *App) Close() {
	app.DB.Close()
}

// HandleDataSet Process the file in the given path and returns a a list of all flights
// in the dataset, and a map of unique cities by name to its respective city
// structure.
//
// This method requires the input file to be in .csv format specified
// by RFC 4180.
//
// The dataset must contain cloumns with some of the following data of origin a destination cities.
// origin|origen name of the origin city.
// destination/destino name of the destination city.
// origin_latitude origin city latitude coordinate
// origin_longitude origin city longitude coordinate
// destination_latitude destination city latitude coordinate
// destination_longitude destination city longitude coordinate
// This method will try to make all its cities ready to call owm API by coordinates
// This method will panic if it cannot open datasets and or database.
func (app *App) HandleDataSet(dataset string) (*[]*Flight, map[string]*City) {
	fileReader, err := os.Open(dataset)
	if err != nil {
		panic("error reading dataset")
	}
	defer fileReader.Close()

	// read headers
	csvReader := csv.NewReader(fileReader)
	headersRow, err := csvReader.Read()
	if err != nil {
		panic("error reading csv")
	}

	// index to fields we want
	var index *Index
	if headersRow[0] == "destino" { // We are in dataset 1
		index = &Index{-1, 0, -1, -1, -1, -1}
	} else {
		index = &Index{0, 1, 2, 3, 4, 5}
	}

	var flights []*Flight
	cities := make(map[string]*City)
	if index.origin == -1 || index.destination == -1 {
		cities[owm.ToAlphaNumeric(defaultCity.name)] = &defaultCity
	}
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic("error while reading csv")
	}
	fc := FlightCreator{db: app.DB, index: index, flights: flights, cities: cities}
	for _, row := range rows {
		fc.addFlight(row)
	}
	return &fc.flights, cities
}

// FlightCreator is an auxiliar type for creating flights
type FlightCreator struct {
	db      *owm.OwmCityConverter
	index   *Index
	flights []*Flight
	cities  map[string]*City
}

// Index is a struct representing indexes to csv columns
type Index struct {
	origin         int
	destination    int
	originLat      int
	originLon      int
	destinationLat int
	detinationLon  int
}

func (fc *FlightCreator) addFlight(row []string) {
	if fc.index.origin == -1 && fc.index.destination == -1 {
		fmt.Print("Not enough information about flight.")
		return
	}
	newFlight := Flight{}
	origin := fc.createCity(row, fc.index.origin, fc.index.originLat, fc.index.originLon)
	origin = addToCities(fc.cities, origin)
	newFlight.origin = origin
	destination := fc.createCity(row, fc.index.destination, fc.index.destinationLat, fc.index.detinationLon)
	destination = addToCities(fc.cities, destination)
	newFlight.destination = destination
	resultingFlights := append(fc.flights, &newFlight)
	fc.flights = resultingFlights
}

func (fc *FlightCreator) createCity(row []string, cityI, latI, lonI int) *City {
	var city City
	if cityI == -1 {
		city = defaultCity
	} else {
		cityName := strings.TrimSpace(row[cityI])
		city = City{name: cityName}
		coordinate, err := fc.db.GetCityCoordinates(cityName)
		if err != nil && latI != -1 && lonI != -1 {
			lat, _ := strconv.ParseFloat(row[latI], 64)
			lon, _ := strconv.ParseFloat(row[lonI], 64)
			coordinate, _ = g.NewCoordinate(lat, lon)
		}
		city.coordinate = coordinate
	}
	return &city
}

func addToCities(cities map[string]*City, city *City) *City {
	normalizedCityName := owm.ToAlphaNumeric(city.name)
	previousCity, found := cities[normalizedCityName]
	if found {
		if previousCity.coordinate == nil && city.coordinate != nil {
			previousCity.coordinate = city.coordinate
		}
		return previousCity
	}
	cities[normalizedCityName] = city
	return city
}

// QueryWeather Tries to request the weather for all cities in the map and saves the result in
// the wethaer field of city.
//
// Requires the weather field of city to be nil.
// This method will panic if encounters a network error while trying to retrieve weather condition
// from the database (for example when there is no connection to internet), this does not includes
// when a weather query fails.
func (app *App) QueryWeather(cities map[string]*City) {
	for _, city := range cities {
		if app.QueriesCounter > app.MaxQueriesPerMinute {
			time.Sleep(time.Minute)
			app.QueriesCounter = 0
		}
		app.QueriesCounter++
		var cityWeather *owm.Weather
		var err error
		if city.coordinate != nil {
			cityWeather, err = app.API.GetWeatherFromCoordinates(city.coordinate, owm.METRIC, owm.ES)
		} else if city.name != "" {
			cityWeather, err = app.API.GetWeatherFromCity(city.name, owm.METRIC, owm.ES)
		} else {
			err = errors.New("insuficcient data about city")
		}
		if err != nil {
			fmt.Printf("ERROR: city not supported %s \n", city.name)
		} else {
			city.weather = cityWeather
		}
	}
}

// PrintWeather Prints the weather to standard output
// It prints information about max and min temperature,
// humidity an a description as well for each fligth.
func PrintWeather(flights *[]*Flight) {
	fmt.Println()
	for _, flight := range *flights {
		fmt.Println("#############################")
		printCityWeather("Origen", flight.origin)
		fmt.Println("-----------------------------")
		printCityWeather("Destino", flight.destination)
	}
	fmt.Println("#############################")
	fmt.Println()
}

// Prints the weatherr for an specific city
func printCityWeather(header string, city *City) {
	fmt.Printf("%s: %s \n", header, city.name)
	if city.weather == nil {
		fmt.Println("Clima no Encontrado")
	} else {
		fmt.Printf("Descripcion: %s \n", city.weather.Weather[0].Description)
		fmt.Printf("Temperatura Mínima: %.1f °C\n", city.weather.Main.TempMin)
		fmt.Printf("Temperatura Máxima: %.1f °C\n", city.weather.Main.TempMax)
		fmt.Printf("Humedad: %.0f%% \n", city.weather.Main.Humidity)
	}
}
