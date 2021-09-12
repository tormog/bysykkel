package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// main prints a table of different stations with corresponding available spots and bicycles.
// URL is the endpoints feed giving new URLs to new endpoint services for status and information.
func main() {
	d := NewDefaultStationDownloader()
	stations, err := d.GetStationSummary("https://gbfs.urbansharing.com/oslobysykkel.no/gbfs.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Stations\tAvailable bikes\tDocks Available\t")
	for _, station := range stations.Data.Stations {
		fmt.Fprintf(w, "%s\t%d\t%d\t\n", station.Name, station.NumBikesAvailable, station.NumDocksAvailable)
	}
	w.Flush()
}
