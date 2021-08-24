package tomtom

import (
	"time"
)

type RoutePoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RoutingApiResponse struct {
	FormatVersion string `json:"formatVersion"`
	Routes        []struct {
		Summary struct {
			LengthInMeters        uint64       `json:"lengthInMeters"`
			TravelTimeInSeconds   uint64       `json:"travelTimeInSeconds"`
			TrafficDelayInSeconds uint64       `json:"trafficDelayInSeconds"`
			TrafficLengthInMeters uint64       `json:"trafficLengthInMeters"`
			DepartureTime         time.Time `json:"departureTime"`
			ArrivalTime           time.Time `json:"arrivalTime"`
		} `json:"summary"`
		Legs []struct {
			Summary struct {
				LengthInMeters        uint64       `json:"lengthInMeters"`
				TravelTimeInSeconds   uint64       `json:"travelTimeInSeconds"`
				TrafficDelayInSeconds uint64       `json:"trafficDelayInSeconds"`
				TrafficLengthInMeters uint64       `json:"trafficLengthInMeters"`
				DepartureTime         time.Time `json:"departureTime"`
				ArrivalTime           time.Time `json:"arrivalTime"`
			} `json:"summary"`
		} `json:"legs"`
		Sections []struct {
			StartPointIndex int    `json:"startPointIndex"`
			EndPointIndex   int    `json:"endPointIndex"`
			SectionType     string `json:"sectionType"`
			TravelMode      string `json:"travelMode"`
		} `json:"sections"`
	} `json:"routes"`
}
