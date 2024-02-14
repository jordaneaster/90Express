// structs package
package structs

import (
	"database/sql"
)

type GeometryType string

type Geometry struct {
	Type            GeometryType           `json:"type"`
	BoundingBox     []float64              `json:"bbox,omitempty"`
	Point           []float64              `json:"coordinates,omitempty"`
	MultiPoint      [][]float64            `json:"coordinates,omitempty"`
	LineString      [][]float64            `json:"coordinates,omitempty"`
	MultiLineString [][][]float64          `json:"coordinates,omitempty"`
	Polygon         [][][]float64          `json:"coordinates,omitempty"`
	MultiPolygon    [][][][]float64        `json:"coordinates,omitempty"`
	Geometries      []*Geometry            `json:"geometries,omitempty"`
	CRS             map[string]interface{} `json:"crs,omitempty"`
}

type Feature struct {
	ID         interface{}            `json:"id,omitempty"`
	Type       string                 `json:"type"`
	BBox       [4]float64             `json:"bbox,omitempty"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	CRS        map[string]interface{} `json:"crs,omitempty"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	FileName string    `json:"filename"`
	CRS      CRSObject `json:"crs"`
	Features []Feature `json:"features"`
}

type CRSObject struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}

type User struct {
	ID       int
	UserID   string
	Password string
}

type AuthDB struct {
	DB *sql.DB
}

func NewAuthDB(db *sql.DB) *AuthDB {
	return &AuthDB{DB: db}
}

type Profile struct {
	UserID   int
	Username string
	Email    string
	FullName string
	Company  string
}
