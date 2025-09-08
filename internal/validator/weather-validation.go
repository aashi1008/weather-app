package validator

import (
	"errors"
	"strconv"
)

func ValidateCoordinates(latitude, longitude string) error {
	if latitude == "" {
		return errors.New("latitude is required")
	}
	if longitude == "" {
		return errors.New("longitude is required")
	}

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return errors.New("invalid latitude format")
	}

	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return errors.New("invalid longitude format")
	}

	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90")
	}

	if lon < -180 || lon > 180 {
		return errors.New("longitude must be between -180 and 180")
	}

	return nil
}

func GetCoordinates(latitude, longitude string) (float64, float64) {
	lat, _ := strconv.ParseFloat(latitude, 64)
	lon, _ := strconv.ParseFloat(longitude, 64)
	return lat, lon
}
