// Package openweathermap provides methods for easily querying OpenWeatherMap API.
package openweathermap

// API represents an OpenWatherApi client to make climate and weather requests.
type API struct {
	hash string
}

// NewAPI creates a new API instance with the given key.
func NewAPI(hash string) *API {
	return &API{hash: hash}
}

// GetWeatherFromCity gets the weather for a given city.
// Requires the city to be city recognized by OpenWeatherMap.
// Note this method does not validate the city before making the request,
// so calling this method with an invalid  city will make a request to the API.
// Units and Language are optional parameters, Language defaults to SPANISH and
// units defaults to METRIC.
func (api *API) GetWeatherFromCity(city string, units Units, lang Language) (*Weather, error) {
	panic("not implemented") // TODO
}

// GetWeatherFromCoordinates gets the weather for the given coordinates.
// Requires -90 < lat < 90 and -180 < lon < 180.
func (api *API) GetWeatherFromCoordinates(lat, lon int, units Units, lang Language) (*Weather, error) {
	panic("not implemented") // TODO
}

// Units is an enum for temperature units in an Api response.
type Units string

const (
	// METRIC for temperature values in Celsius degrees.
	METRIC Units = "metric"
	// IMPERIAL for temperature values in Fahrenheit degrees.
	IMPERIAL Units = "imperial"
	// STANDARD for temperature values in Kelvin degrees.
	STANDARD Units = "standard"
)

// Language is an enum for supported response languages when querying the API.
// For now only requests in english, spanish and french are supported.
type Language string

const (
	// EN for query responses in english.
	EN Language = "en"
	// ES for query responses in spanish.
	ES Language = "es"
	// FR for query response in french.
	FR Language = "fr"
)

// Weather represents a json response from OpenWeatherAPI.
// For simplicity only some fields of the OpenWeatherMap reponse are provided.
type Weather struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"name"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Humidity  float32 `json:"humidity"`
	} `json:"main"`
}

const baseURL = "api.openweathermap.org/data/2.5/weather"
