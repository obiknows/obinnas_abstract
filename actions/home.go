package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// AboutHandler is a default handler to serve up
// an about page.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}

// WorkHandler is a default handler to serve up
// a work page.
func WorkHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("work.html"))
}

// ContactHandler is a default handler to serve up
// a contact page.
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact.html"))
}

// RoutesHandler is a default handler to serve up
// a home page.
func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("routes.html"))
}
