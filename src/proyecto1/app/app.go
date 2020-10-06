package main

import owm "myp.ciencias.unam.mx/proyecto1/openweathermap"

// Runs the app.
func main() {
	panic("not implemented") // TODO
}

// TODO Improve documentation.
// Process the file in the given path and returns a a list of all flights
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
func handleDataSet(dataset string) (*[]flight, *map[string]*city) {
	panic("not implemented") // TODO
}

// Requests the weather for all cities in the map and saves the result in
// the wethaer field of city.
//
// Requires the weather field of city to be nil.
func queryWeather(*map[string]*city) {
	panic("not implemented") // TODO
}

// TODO improve documentation.
// Prints the weather.
func printWeather(*[]flight) {
	panic("not implemented") // TODO
}

type flight struct {
	origin      *city
	destination *city
}

type city struct {
	name        *string
	coordinates struct {
		lat float32
		lon float32
	}
	weather *owm.Weather
}
