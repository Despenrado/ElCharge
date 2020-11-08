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
				Location:    "156.12 1235.2",
			},
			isValid: true,
		},
		{
			name: "invalid description",
			station: &Station{
				Description: "",
				StationName: "station name",
				Location:    "156.12 1235.2",
			},
			isValid: false,
		},
		{
			name: "invalid name",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "s",
				Location:    "156.12 1235.2",
			},
			isValid: false,
		},
		{
			name: "invalid location",
			station: &Station{
				Description: "asdasdasdasdas",
				StationName: "station name",
				Location:    "15",
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
		Location:    "156.12 1235.2",
	}
	assert.NoError(t, station.BeforeCreate())
	assert.NotNil(t, station.UpdateAt)
	assert.NotNil(t, station.CreateAt)
}
