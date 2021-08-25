package tomtom

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type client struct {
	apikey string
}

const routeApiUrl = "https://api.tomtom.com/routing/1/calculateRoute"

// Client - Interface for TomTom api client.
type Client interface {
	GetRoute(from, to RoutePoint) (RoutingApiResponse, error)
}

// GetRoute - Calculates distance and travel time between 2 coordinate points for cars.
func (c *client) GetRoute(from, to RoutePoint) (route RoutingApiResponse, err error) {
	requestUrl := fmt.Sprintf("%s/%f,%f:%f,%f/json",
		routeApiUrl, from.Latitude, from.Longitude, to.Latitude, to.Longitude)
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return route, err
	}
	q := request.URL.Query()
	// Will hardcode query params. Why not for this task ?
	q.Add("maxAlternatives", "0")
	q.Add("traffic", "true")
	q.Add("avoid", "unpavedRoads")
	q.Add("travelMode", "car")
	q.Add("key", c.apikey)
	request.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return route, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return route, err
	}
	if resp.StatusCode != http.StatusOK {
		return route, errors.New(string(body))
	}
	err = json.Unmarshal(body, &route)
	return route, err
}

// NewClient - Returns new TomTom api Client
func NewClient(apiKey string) Client {
	return &client{apikey: apiKey}
}
