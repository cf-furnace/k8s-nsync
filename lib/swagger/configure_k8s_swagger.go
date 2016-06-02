package swagger

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cf-furnace/k8s-nsync/lib/swagger/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.K8sSwaggerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

// ConfigureAPI configures the Nsync API server
func ConfigureAPI(api *operations.K8sSwaggerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CancelTaskHandler = operations.CancelTaskHandlerFunc(func(params operations.CancelTaskParams) middleware.Responder {
		return middleware.NotImplemented("operation .CancelTask has not yet been implemented")
	})
	api.DesireAppHandler = operations.DesireAppHandlerFunc(func(params operations.DesireAppParams) middleware.Responder {
		return middleware.NotImplemented("operation .DesireApp has not yet been implemented")
	})
	api.DesireTaskHandler = operations.DesireTaskHandlerFunc(func(params operations.DesireTaskParams) middleware.Responder {
		return middleware.NotImplemented("operation .DesireTask has not yet been implemented")
	})
	api.KillIndexHandler = operations.KillIndexHandlerFunc(func(params operations.KillIndexParams) middleware.Responder {
		return middleware.NotImplemented("operation .KillIndex has not yet been implemented")
	})
	api.StopAppHandler = operations.StopAppHandlerFunc(func(params operations.StopAppParams) middleware.Responder {
		return middleware.NotImplemented("operation .StopApp has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
