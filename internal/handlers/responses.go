package handlers

type ErrorJson struct {
	Error string `json:"error"`
}

type RouteCreatedJson struct {
	RouteId uint64 `json:"route_id"`
}