// Package geopoint provides simple coordinate abstraction.
package geopoint

import (
	"errors"
	"math"
)

// GeoPoint is an interface for representing geographical coordinates.
// Provides methods for retrieving latitude and longitude of an object.
type GeoPoint interface {
	Latitude() float64
	Longitude() float64
}

// Coordinate is a simple coordinates type which implements geopoint interface.
type Coordinate struct {
	Lat float64
	Lon float64
}

// NewCoordinate Returns a new coordinate.
// returns an error if |latitude| > 90 or |longitude| > 180.
func NewCoordinate(latitude, longitude float64) (*Coordinate, error) {
	if math.Abs(latitude) > 90 || math.Abs(longitude) > 180 {
		return nil, errors.New("invalid argument value for coordinate")
	}
	return &Coordinate{Lat: latitude, Lon: longitude}, nil
}

// Latitude retrieves the latitude property of the coordinate.
func (c Coordinate) Latitude() float64 {
	return c.Lat
}

// Longitude retrieves the longitude property of the coordinate.
func (c Coordinate) Longitude() float64 {
	return c.Lon
}
