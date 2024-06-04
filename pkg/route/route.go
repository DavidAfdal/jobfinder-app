package route

import "github.com/labstack/echo/v4"

type Route struct {
	Methode string
	Path string
	Handler echo.HandlerFunc
}
