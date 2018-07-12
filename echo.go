package session

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// DefaultContextKey DefaultContextKey
	DefaultContextKey = "_CONTEXT_SESSION_KEY"
)

type (
	// MiddlewareConfig MiddlewareConfig
	MiddlewareConfig struct {
		Skipper middleware.Skipper
		Manager *Manager
	}
)

// Middleware Middleware for echo only
func Middleware(config MiddlewareConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}

	if config.Manager == nil {
		panic("config.Manager must be set")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			sess, err := config.Manager.SessionStart(c.Response(), c.Request())
			if err != nil {
				return err
			}

			c.Set(DefaultContextKey, sess)

			err = next(c)
			if err != nil {
				return err
			}

			sess.SessionRelease(c.Response())

			return nil
		}
	}
}

// Session Session Store from context
func Session(c echo.Context) Store {
	return c.Get(DefaultContextKey).(Store)
}
