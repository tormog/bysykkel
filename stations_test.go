package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockGetEndpoints(url string) (map[string]string, error) {
	return map[string]string{
		"station_information": "https://example.com/something",
		"station_status":      "https://example.com/someother",
	}, nil
}

func mockGetStationStatus(url string) (*stations, error) {
	stationsStatus := &stations{}

	stationsStatus.Data.Stations = append(stationsStatus.Data.Stations, station{
		StationID:         "627",
		NumBikesAvailable: 1,
		NumDocksAvailable: 10,
	})

	return stationsStatus, nil
}

func mockGetStationInformation(url string) (*stations, error) {
	stationsInformation := &stations{}

	stationsInformation.Data.Stations = append(stationsInformation.Data.Stations, station{
		StationID: "627",
		Name:      "Skøyen Stasjon",
	})

	return stationsInformation, nil
}

func TestGetStationSummary(t *testing.T) {
	d := NewStationDownloader(mockGetEndpoints, mockGetStationInformation, mockGetStationStatus)
	stations, err := d.GetStationSummary("https://example.com")
	if err != nil {
		t.Error(err)
	}

	if len(stations.Data.Stations) == 0 {
		t.Error("List of stations is empty")
	}

	for _, val := range stations.Data.Stations {
		if val.StationID != "627" {
			t.Errorf("Station ID %s not correct", val.StationID)
		}
	}
}

func TestGetEndpoints(t *testing.T) {
	b, err := ioutil.ReadFile("test/data/endpoints.json")
	if err != nil {
		t.Error(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(b)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}))
	endpoints, err := GetEndpoints(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if len(endpoints) != 3 {
		t.Error("not the correct amount endpoint information returned")
	}

	if _, ok := endpoints["station_information"]; !ok {
		t.Error("expected endpoint not found")
	}
}

func TestGetStationInformation(t *testing.T) {
	b, err := ioutil.ReadFile("test/data/station_information.json")
	if err != nil {
		t.Error(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(b)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}))
	station, err := GetStationInformation(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if len(station.Data.Stations) != 3 {
		t.Error("not the correct amount of station information returned")
	}
}

func TestGetStationStatus(t *testing.T) {
	b, err := ioutil.ReadFile("test/data/station_status.json")
	if err != nil {
		t.Error(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(b)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}))
	station, err := GetStationStatus(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if len(station.Data.Stations) != 3 {
		t.Error("not the correct amount of station status returned")
	}
}
