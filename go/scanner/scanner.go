package scanner

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func PointsInsidePolygons(fc *geojson.FeatureCollection, point orb.Point) map[string][]orb.Point {
	containedPoints := make(map[string][]orb.Point)

	for _, feature := range fc.Features {
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			for _, polygon := range multiPoly {
				if planar.PolygonContains(polygon, point) {
					name := feature.Properties["name"].(string)
					containedPoints[name] = append(containedPoints[name], point)
				}
			}
		} else {
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if planar.PolygonContains(polygon, point) {
					name := feature.Properties["name"].(string)
					containedPoints[name] = append(containedPoints[name], point)
				}
			}
		}
	}
	return containedPoints
}
