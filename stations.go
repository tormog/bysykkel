package main

import (
	"encoding/json"
	"fmt"
)

var headers = map[string]string{
	"Accept":            "application/json",
	"Client-Identifier": "my-client-test",
}

// Use this setup for mocking
type Endpoints func(url string) (map[string]string, error)
type StationInformation func(url string) (*stations, error)
type StationStatus func(url string) (*stations, error)

type StationDownloader struct {
	getEndpoints          Endpoints
	getStationInformation StationInformation
	getStationStatus      StationStatus
}

// Used for finelizing mock setup
func NewStationDownloader(endpoints Endpoints, stationInfo StationInformation, stationStatus StationStatus) *StationDownloader {
	return &StationDownloader{
		getEndpoints:          endpoints,
		getStationInformation: stationInfo,
		getStationStatus:      stationStatus,
	}
}

// Used by external
func NewDefaultStationDownloader() *StationDownloader {
	return &StationDownloader{
		getEndpoints:          GetEndpoints,
		getStationInformation: GetStationInformation,
		getStationStatus:      GetStationStatus,
	}
}

/*
We reuse one type, stations, throughout the code.
This is possible duo to the actual data required for processing.
Golang will exempt all other data under request parsing.
*/

// GetEndpoints retrieves all relevant endpoints from bysykkel
func GetEndpoints(url string) (map[string]string, error) {
	mapOfEndpoints := map[string]string{}
	body, code, err := GetRequest(headers, url)
	if err != nil {
		return nil, err
	}
	if code == 200 {
		endpoints := &endpoints{}
		err := json.Unmarshal(body, &endpoints)
		if err != nil {
			return nil, err
		}
		for _, data := range endpoints.Data.Nb.Feeds {
			mapOfEndpoints[data.Name] = data.URL
		}
		return mapOfEndpoints, nil
	} else {
		return nil, fmt.Errorf("expected other status code than %d", code)
	}
}

/*
GetStationStatus and GetStationInformation is basically the same code.
We separate them to duo the mock setup.
*/

// GetStation retrieves station status.
func GetStationStatus(url string) (*stations, error) {
	stations := &stations{}
	body, code, err := GetRequest(headers, url)
	if err != nil {
		return nil, err
	}
	if code == 200 {
		err := json.Unmarshal(body, &stations)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("expected other status code than %d", code)
	}
	return stations, nil
}

// GetStation retrieves station information.
func GetStationInformation(url string) (*stations, error) {
	stations := &stations{}
	body, code, err := GetRequest(headers, url)
	if err != nil {
		return nil, err
	}
	if code == 200 {
		err := json.Unmarshal(body, &stations)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("expected other status code than %d", code)
	}
	return stations, nil
}

// GetStationsSummary will summarize both station info and status.
func (d *StationDownloader) GetStationSummary(endpoints string) (*stations, error) {
	mapOfEndpoints, err := d.getEndpoints(endpoints)
	if err != nil {
		return nil, err
	}

	stationInformation := &stations{}
	stationStatus := &stations{}
	stationSummary := &stations{}

	for key, elem := range mapOfEndpoints {
		if key == "station_information" {
			stationInformation, err = d.getStationInformation(elem)
			if err != nil {
				return nil, err
			}
		} else if key == "station_status" {
			stationStatus, err = d.getStationStatus(elem)
			if err != nil {
				return nil, err
			}
		}
	}

	// Setup for simple index lookup
	stationMap := map[string]string{}
	for _, elem := range stationInformation.Data.Stations {
		stationMap[elem.StationID] = elem.Name
	}

	for _, elem := range stationStatus.Data.Stations {
		stationSummary.Data.Stations = append(stationSummary.Data.Stations, station{
			StationID:         elem.StationID,
			Name:              stationMap[elem.StationID],
			NumBikesAvailable: elem.NumBikesAvailable,
			NumDocksAvailable: elem.NumDocksAvailable,
		},
		)
	}

	return stationSummary, nil
}
