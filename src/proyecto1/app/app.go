package main

import owm "myp.ciencias.unam.mx/proyecto1/openweathermap"

// Runs the app.
func main() {
	panic("not implemented") // TODO
}

type flight struct {
	origin      *city
	destination *city
}

type city struct {
	name        string
	coordinates struct {
		lat float32
		lon float32
	}
	weather *owm.Weather
}
