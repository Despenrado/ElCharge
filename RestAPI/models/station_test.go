package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStation(t *testing.T) {
	testCase := []struct {
		name    string
		station *Station
		isValid bool
	}{
		{
			name: "default",
			station: &Station{
				Description: "testText",
				StationName: "station name",
				OwnerID:     "1234567890123456789012345",
				Latitude:    1.12,
				Longitude:   12.2,
			},
			isValid: true,
		},
		{
			name: "invalid description",
			station: &Station{
				Description: "",
				StationName: "station name",
				OwnerID:     "1234567890123456789012345",
				Latitude:    15.12,
				Longitude:   12.2,
			},
			isValid: false,
		},
		{
			name: "invalid name",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "s",
				OwnerID:     "1234567890123456789012345",
				Latitude:    15.12,
				Longitude:   12.2,
			},
			isValid: false,
		},
		{
			name: "invalid ownerID",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "station name",
				OwnerID:     "1234567890",
				Latitude:    -15.12,
				Longitude:   15.2,
			},
			isValid: false,
		},
		{
			name: "OK",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "station name",
				OwnerID:     "1234567890123456789012345",
				Latitude:    -15.12,
				Longitude:   -15.2,
			},
			isValid: true,
		},
		{
			name: "invalid Latitude",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "station name",
				OwnerID:     "1234567890123456789012345",
				Latitude:    -105.12,
				Longitude:   -15.2,
			},
			isValid: false,
		},
		{
			name: "invalid Longitude",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "station name",
				OwnerID:     "1234567890123456789012345",
				Latitude:    -10.12,
				Longitude:   -1500.2,
			},
			isValid: false,
		},
	}
	for _, item := range testCase {
		t.Run(item.name, func(t *testing.T) {
			if item.isValid {
				assert.NoError(t, item.station.Validate())
			} else {
				assert.Error(t, item.station.Validate())
			}
		})
	}
}

func TestBeforeCreateStation(t *testing.T) {
	station := &Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	assert.NoError(t, station.BeforeCreate())
	assert.NotNil(t, station.UpdateAt)
	assert.NotNil(t, station.CreateAt)
}
