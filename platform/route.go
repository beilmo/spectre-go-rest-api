package platform

import "net/http"

// Route stores information to match a request and build URLs.
type Route struct {
	// The name for the route.
	name string

	// HTTP method. e.g.: "GET", "POST", "PUT".
	method string

	// URL path pattern. It accepts a template with zero or more URL variables enclosed by {}. The template must start with a "/". Variables can define an optional regexp pattern to be matched:
	// {name} matches anything until the next slash.
	// {name:pattern} matches the given regexp pattern.
	pattern string

	// Request handler for the route.
	handler func(http.ResponseWriter, *http.Request)
}

// Routes type represents a list of routes.
type Routes = []Route

// Routable interface providing the receivers available routes.
type Routable interface {
	Routes() Routes
}
