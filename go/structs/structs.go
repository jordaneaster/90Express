package main

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type GeometryType string

type Geometry struct {
	Type            GeometryType `json:"type"`
	BoundingBox     []float64    `json:"bbox,omitempty"`
	Point           []float64
	MultiPoint      [][]float64
	LineString      [][]float64
	MultiLineString [][][]float64
	Polygon         [][][]float64
	MultiPolygon    [][][][]float64
	Geometries      []*Geometry
	CRS             map[string]interface{} `json:"crs,omitempty"`
}

type Feature struct {
	ID         interface{}            `json:"id,omitempty"`
	Type       string                 `json:"type"`
	BBox       geojson.BBox           `json:"bbox,omitempty"`
	Geometry   orb.Geometry           `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	CRS        map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

type FeatureCollection struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Name        string      `json:"Name"`
			Description interface{} `json:"description"`
		} `json:"properties"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

type Polygon struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
		Geometry struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}
