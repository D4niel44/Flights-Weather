module myp/Tarea01

go 1.15

replace myp/Tarea01/app/geo => ./app/geo

replace myp/Tarea01/app/openweathermap => ./app/openweathermap

replace myp/Tarea01/app/openweathermapcities => ./app/openweathermapcities

require (
	github.com/mattn/go-sqlite3 v1.14.4
	golang.org/x/text v0.3.3
)
