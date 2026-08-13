package importer

import "testing"

func validPlace() PlaceData {
    startYear := 1500
    endYear := 1600

    return PlaceData{
        Name:         "Mexico City",
        Type:         "origin",
        Latitude:     19.4326,
        Longitude:    -99.1332,
        Relationship: "origin",
        StartYear:    &startYear,
        EndYear:      &endYear,
    }
}

func TestValidatePlace_Latitude(t *testing.T) {
    tests := []struct {
        name      string
        latitude  float64
        wantError bool
    }{
        {"below minimum", -91, true},
        {"minimum", -90, false},
        {"middle", 0, false},
        {"maximum", 90, false},
        {"above maximum", 91, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            place := validPlace()
            place.Latitude = tt.latitude

            err := ValidatePlace(place)

            if (err != nil) != tt.wantError {
                t.Errorf(
                    "expected error: %v, got error: %v",
                    tt.wantError,
                    err != nil,
                )
            }
        })
    }
}


func TestValidatePlace_Longitude(t *testing.T) {
    tests := []struct {
        name      string
        longitude float64
        wantError bool
    }{
        {"below minimum", -181, true},
        {"minimum", -180, false},
        {"middle", 0, false},
        {"maximum", 180, false},
        {"above maximum", 181, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            place := validPlace()
            place.Longitude = tt.longitude

            err := ValidatePlace(place)

            if (err != nil) != tt.wantError {
                t.Errorf(
                    "expected error: %v, got error: %v",
                    tt.wantError,
                    err != nil,
                )
            }
        })
    }
}


func TestValidatePlace_Years(t *testing.T) {
    tests := []struct {
        name      string
        startYear int
        endYear   int
        wantError bool
    }{
        {"end year after start year", 1500, 1600, false},
        {"same start and end year", 1500, 1500, false},
        {"end year before start year", 1600, 1500, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            place := validPlace()

            place.StartYear = &tt.startYear
            place.EndYear = &tt.endYear

            err := ValidatePlace(place)

            hasError := err != nil

            if hasError != tt.wantError {
                t.Errorf(
                    "expected error: %v, got error: %v",
                    tt.wantError,
                    hasError,
                )
            }
        })
    }
}