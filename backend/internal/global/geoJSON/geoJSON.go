package geoJSON

import (
	errors2 "adsb-api/internal/global/errorMsg"
	"adsb-api/internal/global/models"
	"errors"
)

// GeoJson FeatureCollection for a Point type

type FeatureCollectionPoint struct {
	Type     string         `json:"type"`
	Features []featurePoint `json:"features"`
}

type featurePoint struct {
	Type       string                    `json:"type"`
	Geometry   geometryPoint             `json:"geometry"`
	Properties aircraftCurrentProperties `json:"properties"`
}

type aircraftCurrentProperties struct {
	Icao         string `json:"icao"`
	Callsign     string `json:"callsign"`
	Altitude     int    `json:"altitude"`
	Speed        int    `json:"speed"`
	Track        int    `json:"track"`
	VerticalRate int    `json:"vspeed"`
	Timestamp    string `json:"timestamp"`
}

type geometryPoint struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}

// GeoJson FeatureCollection for a LineString type

type FeatureCollectionLineString struct {
	Type     string              `json:"type"`
	Features []featureLineString `json:"features"`
}

type featureLineString struct {
	Type       string                 `json:"type"`
	Properties aircraftHistProperties `json:"properties"`
	Geometry   geometryLineString     `json:"geometry"`
}

type aircraftHistProperties struct {
	Icao string `json:"icao"`
}

type geometryLineString struct {
	Coordinates [][]float32 `json:"coordinates"`
	Type        string      `json:"type"`
}

func ConvertCurrentModelToGeoJson(aircraft []models.AircraftCurrentModel) (FeatureCollectionPoint, error) {
	var features []featurePoint
	for _, ac := range aircraft {
		var feature featurePoint
		feature.Type = "Feature"
		properties := aircraftCurrentProperties{
			Icao:         ac.Icao,
			Callsign:     ac.Callsign,
			Altitude:     ac.Altitude,
			Speed:        ac.Speed,
			Track:        ac.Track,
			VerticalRate: ac.VerticalRate,
			Timestamp:    ac.Timestamp,
		}
		feature.Properties = properties
		feature.Geometry.Type = "Point"
		feature.Geometry.Coordinates = append(feature.Geometry.Coordinates, ac.Longitude, ac.Latitude)
		features = append(features, feature)
	}

	var featureCollection FeatureCollectionPoint
	featureCollection.Features = features
	featureCollection.Type = "FeatureCollection"
	return featureCollection, nil
}

func ConvertHistoryModelToGeoJson(aircraft []models.AircraftHistoryModel) (FeatureCollectionLineString, error) {
	if len(aircraft) < 2 {
		return FeatureCollectionLineString{}, errors.New(errors2.ErrorGeoJsonTooFewCoordinates)
	}

	var coordinates [][]float32
	for _, ac := range aircraft {
		point := []float32{ac.Latitude, ac.Longitude}
		coordinates = append(coordinates, point)
	}

	var features []featureLineString
	var feature featureLineString
	feature.Type = "Feature"
	feature.Properties.Icao = aircraft[0].Icao
	feature.Geometry.Coordinates = coordinates
	feature.Geometry.Type = "LineString"
	features = append(features, feature)

	var featureCollection FeatureCollectionLineString
	featureCollection.Features = features
	featureCollection.Type = "FeatureCollection"
	return featureCollection, nil
}
