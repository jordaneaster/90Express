package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

var filename string

func readKML(filename string) (*KML, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening KML file: %w", err) // contains filename
	}
	defer f.Close() // reading, ignoring error is acceptable
	var kml KML
	if err := xml.NewDecoder(f).Decode(&kml); err != nil {
		return nil, fmt.Errorf("decoding XML from %q as KML: %w", filename, err)
	}
	return &kml, nil
}

func writeKML(filename string, kml *KML) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating KML file: %w", err) // contains filename
	}
	defer f.Close() // double close is OK for *os.File
	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	if err := enc.Encode(kml); err != nil {
		return fmt.Errorf("encoding KML to %q: %w", filename, err)
	}
	return nil
}

func main() {
	kml, err := readKML("maps/MA_NY.kml")
	if err != nil {
		return // contains context
	}

	coordinates := strings.Fields(kml.Folder.Document.Placemark.LineString.Coordinates)

	for coord := range coordinates {
		fmt.Println(coord)
		//   coordinates = append(coordinates, coord)

		kml.Folder.Document.Placemark.LineString.Coordinates = strings.Join(coordinates, "\n")
		if err := writeKML(filename, kml); err != nil {
			log.Printf("Warning: failed to update %q: %s", filename, err)
		}
	}
}
