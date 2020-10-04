// Package openweathermap provides methods for easily querying OpenWeatherMap API.
package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
func (api *API) GetWeatherFromCity(city string, units Units, lang Language) (*Weather, error) {
	if city == "" {
		return nil, errors.New("city must be a nonempty string")
	}
	urlBuilder := api.baseURLBuilder(units, lang)
	urlBuilder.addParameter("q", city)
	return makeQuery(urlBuilder.makeURL())
}

// GetWeatherFromCoordinates gets the weather for the given coordinates.
// Requires -90 < lat < 90 and -180 < lon < 180.
func (api *API) GetWeatherFromCoordinates(lat, lon float32, units Units, lang Language) (*Weather, error) {
	if lat < -90 || lat > 90 || lon < -180 || lon > 180 {
		return nil, errors.New("wrong coordinates")
	}
	urlBuilder := api.baseURLBuilder(units, lang)
	urlBuilder.addParameter("lat", fmt.Sprintf("%.2f", lat))
	urlBuilder.addParameter("lon", fmt.Sprintf("%.2f", lon))
	return makeQuery(urlBuilder.makeURL())
}

const baseURL = "https://api.openweathermap.org/data/2.5/weather?"

// Creates an urlBuilder with common parameters for both requests
func (api *API) baseURLBuilder(units Units, lang Language) *getRequestURLBuilder {
	urlBuilder := newGetRequestURLBuilder(baseURL)
	urlBuilder.addParameter("appid", api.hash)
	urlBuilder.addParameter("units", string(units))
	urlBuilder.addParameter("lang", string(lang))
	return urlBuilder
}

// Queries the url and returns and struct containing the response
func makeQuery(url string) (*Weather, error) {
	// make query
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode response
	weather := &Weather{}
	jsonErr := json.NewDecoder(resp.Body).Decode(weather)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return weather, nil
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

type getRequestURLBuilder struct {
	base   string
	params map[string]string
}

func newGetRequestURLBuilder(baseURL string) *getRequestURLBuilder {
	return &getRequestURLBuilder{base: baseURL, params: make(map[string]string)}
}

func (r *getRequestURLBuilder) addParameter(key, value string) {
	r.params[key] = value
}

func (r *getRequestURLBuilder) makeURL() string {
	url := r.base
	i := 1
	length := len(r.params)
	for key, value := range r.params {
		url += fmt.Sprintf("%s=%s", key, value)
		if i != length {
			url += "&"
		}
		i++
	}
	return url
}
