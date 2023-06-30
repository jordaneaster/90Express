package main

import "encoding/xml"

type KML struct {
	XMLName xml.Name `xml:"kml"`
	Text    string   `xml:",chardata"`
	XMLNS   string   `xml:"xmlns,attr"`
	GX      string   `xml:"gx,attr"`
	KML     string   `xml:"kml,attr"`
	Atom    string   `xml:"atom,attr"`
	Folder  struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name"`
		Open     string `xml:"open"`
		Document struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name"`
			Placemark struct {
				Text  string `xml:",chardata"`
				Style struct {
					Text      string `xml:",chardata"`
					LineStyle struct {
						Text  string `xml:",chardata"`
						Color string `xml:"color"`
						Width string `xml:"width"`
					} `xml:"LineStyle"`
					PolyStyle struct {
						Text    string `xml:",chardata"`
						Color   string `xml:"color"`
						Fill    string `xml:"fill"`
						Outline string `xml:"outline"`
					} `xml:"PolyStyle"`
				} `xml:"Style"`
				LineString struct {
					Text        string `xml:",chardata"`
					Tessellate  string `xml:"tessellate"`
					Coordinates string `xml:"coordinates"`
				} `xml:"LineString"`
			} `xml:"Placemark"`
		} `xml:"Document"`
	} `xml:"Folder"`
}
