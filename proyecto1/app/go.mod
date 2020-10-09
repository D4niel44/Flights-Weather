module myp.ciencias.unam.mx/proyecto1/app

go 1.15

replace myp.ciencias.unam.mx/proyecto1/openweathermap => ../openweathermap

replace myp.ciencias.unam.mx/geo => ../geo

replace myp.ciencias.unam.mx/proyecto1/openweathermapcities => ../openweathermapcities

require (
	myp.ciencias.unam.mx/geo v0.0.0-00010101000000-000000000000
	myp.ciencias.unam.mx/proyecto1/openweathermap v0.0.0-00010101000000-000000000000
	myp.ciencias.unam.mx/proyecto1/openweathermapcities v0.0.0-00010101000000-000000000000
)
