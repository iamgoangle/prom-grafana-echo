package main

import (
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// urlSkipper ignores metrics route on some middleware
func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	return false
}

func main() {
	e := echo.New()
	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", urlSkipper)
	p.Use(e)

	e.Logger.Fatal(e.Start(":9009"))
}
