package models

// Route - route struct composed by an id, name, from point, to point and a geojson string route
type Route struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	From Address `json:"from"`
	To   Address `json:"to"`
	R    string  `json:"r"`
}
