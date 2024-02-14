package fileprocessor

import (
	"encoding/json"
	"fmt"
	"os"
	"scanner"

	"github.com/paulmach/orb/geojson"
)

// FeatureCollection represents a collection of features
type FeatureCollection struct {
	FileName string                  `json:"fileName"`
	Features map[string][][2]float64 `json:"features"`
}

// ProcessFiles reads and processes the files, returning the processed feature collections
func ProcessFiles() ([]*FeatureCollection, error) {
	kvs := map[string]string{
		"A1.json": "go/data/A1.json",
		"A2.json": "go/data/A2.json",
		"A3.json": "go/data/A3.json",
		"A4.json": "go/data/A4.json",
		"A5.json": "go/data/A5.json",
		"A6.json": "go/data/A6.json",
		"A7.json": "go/data/A7.json",
	}

	// Open the regions file
	regionsFile, err := os.Open("go/data/polygons.json")
	if err != nil {
		return nil, fmt.Errorf("error opening regions file: %v", err)
	}
	defer regionsFile.Close()

	// Read the regions data
	var regions geojson.FeatureCollection
	err = json.NewDecoder(regionsFile).Decode(&regions)
	if err != nil {
		return nil, fmt.Errorf("error decoding regions data: %v", err)
	}

	// Create a slice to hold processed feature collections
	var processedCollections []*FeatureCollection

	// Iterate over each file in kvs
	for fileName, filePath := range kvs {
		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		// Read the file content
		bytevalue, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		// Unmarshal the feature collection
		fc, err := geojson.UnmarshalFeatureCollection(bytevalue)
		if err != nil {
			fmt.Println("Error unmarshalling feature collection:", err)
			continue
		}

		// Check each feature collection against the regions
		for _, feature := range fc.Features {
			result := scanner.PointsInsidePolygons(&regions, feature.Point())
			if len(result) > 0 {
				// If the point is inside one or more polygons, convert the result to the desired format
				// and add it to processed collections
				fc := &FeatureCollection{
					FileName: fileName,
					Features: make(map[string][][2]float64),
				}
				for k, v := range result {
					var convertedPoints [][2]float64
					for _, p := range v {
						convertedPoints = append(convertedPoints, [2]float64{p[0], p[1]})
					}
					fc.Features[k] = convertedPoints
				}
				processedCollections = append(processedCollections, fc)
			}
		}
	}
	return processedCollections, nil
}
