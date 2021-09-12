package main

/*
In general TTL and LastUpdated is not included.
This could be include in some new version for a continuously check.
Fore more information see https://oslobysykkel.no/apne-data/sanntid
*/

// Language set default to nb.
type endpoints struct {
	Data struct {
		Nb struct {
			Feeds []struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"feeds"`
		} `json:"nb"`
	} `json:"data"`
}

// We only map the station info and status attributes we are interested in.
type station struct {
	StationID         string `json:"station_id"`
	Name              string `json:"name"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
}

// stations used for both station status and info.
type stations struct {
	Version string `json:"version"`
	Data    struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}
