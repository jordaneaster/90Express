package main

import (
	"fmt"
	"os"
	"scanner"

	"github.com/paulmach/orb/geojson"
)

var z []byte

func main() {
	// need to make these directories global environment variables -> then secrets in ansible
	zone1, err := os.Open("go/data/A1.json")
	zone2, err := os.Open("go/data/A2.json")
	zone3, err := os.Open("go/data/A3.json")
	zone4, err := os.Open("go/data/A4.json")
	zone5, err := os.Open("go/data/A5.json")
	zone6, err := os.Open("go/data/A6.json")
	zone7, err := os.Open("go/data/A7.json")
	regions, err := os.Open("go/data/polygons.json")

	kvs := map[string]os.File{zone1.Name(): *zone1, zone2.Name(): *zone2, zone3.Name(): *zone3, zone4.Name(): *zone4, zone5.Name(): *zone5, zone6.Name(): *zone6, zone7.Name(): *zone7, regions.Name(): *regions}

	for _, file := range kvs {
		os.Open(file.Name())
		defer file.Close()

		bytevalue, _ := os.ReadFile(file.Name())
		fc, _ := geojson.UnmarshalFeatureCollection(bytevalue)
		// fmt.Printf("Zone: %s ", fc.Features)
		fmt.Print(scanner.IsPointInsidePolygon(fc, fc.Features[0].Point()))
	}
	if err != nil {
		fmt.Print(err)
	}

}
