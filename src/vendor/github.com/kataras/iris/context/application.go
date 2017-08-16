package context

import (
	"io"
	"net/http"

	"github.com/kataras/golog"
)

// Application is the context's owner.
// This interface contains the functions that can be used with safety inside a Handler
// by `context.Application()`.
type Application interface {
	// ConfigurationReadOnly returns all the available configuration values can be used on a request.
	ConfigurationReadOnly() ConfigurationReadOnly

	// Logger returns the golog logger instance(pointer) that is being used inside the "app".
	Logger() *golog.Logger

	// View executes and write the result of a template file to the writer.
	//
	// Use context.View to render templates to the client instead.
	// Returns an error on failure, otherwise nil.
	View(writer io.Writer, filename string, layout string, bindingData interface{}) error

	// ServeHTTPC is the internal router, it's visible because it can be used for advanced use cases,
	// i.e: routing within a foreign context.
	//
	// It is ready to use after Build state.
	ServeHTTPC(ctx Context)

	// ServeHTTP is the main router handler which calls the .Serve and acquires a new context from the pool.
	//
	// It is ready to use after Build state.
	ServeHTTP(w http.ResponseWriter, r *http.Request)

	// FireErrorCode executes an error http status code handler
	// based on the context's status code.
	//
	// If a handler is not already registered,
	// then it creates & registers a new trivial handler on the-fly.
	FireErrorCode(ctx Context)
}
