package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// RoutesHandler is a default handler to serve up
// a home page.
func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("routes.html"))
}
