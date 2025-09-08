package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCoordinates_Valid(t *testing.T) {
	err := ValidateCoordinates("50.5", "109.9")
	assert.NoError(t, err)
}

func TestValidateCoordinates_MissingLatitude(t *testing.T) {
	err := ValidateCoordinates("", "109.9")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "latitude is required")
}

func TestValidateCoordinates_InvalidLatitudeFormat(t *testing.T) {
	err := ValidateCoordinates("abc", "109.9")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid latitude format")
}

func TestValidateCoordinates_OutOfRange(t *testing.T) {
	err := ValidateCoordinates("200", "109.9")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "latitude must be between -90 and 90")
}

func TestGetCoordinates(t *testing.T) {
	lat, lon := GetCoordinates("50.5", "109.9")
	assert.Equal(t, 50.5, lat)
	assert.Equal(t, 109.9, lon)
}
