package middlewares

import "github.com/labstack/echo/v4"

type GoMiddleware struct {
	// another stuff , may be needed by middlewares
}

// CORS will handle the CORS middlewares
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Cursor")
		return next(c)
	}
}
