// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"climatehero/restapi/operations"
)

//go:generate swagger generate server --target ../../climatehero --name Climatehero --spec ../swagger.yaml

func configureFlags(api *operations.ClimateheroAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ClimateheroAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddUserHandler == nil {
		api.AddUserHandler = operations.AddUserHandlerFunc(func(params operations.AddUserParams) middleware.Responder {
			return middleware.NotImplemented("operation .AddUser has not yet been implemented")
		})
	}
	if api.CheckItemHandler == nil {
		api.CheckItemHandler = operations.CheckItemHandlerFunc(func(params operations.CheckItemParams) middleware.Responder {
			return middleware.NotImplemented("operation .CheckItem has not yet been implemented")
		})
	}
	if api.GetUserInfoHandler == nil {
		api.GetUserInfoHandler = operations.GetUserInfoHandlerFunc(func(params operations.GetUserInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetUserInfo has not yet been implemented")
		})
	}
	if api.LoginHandler == nil {
		api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation .Login has not yet been implemented")
		})
	}
	if api.UpdateListHandler == nil {
		api.UpdateListHandler = operations.UpdateListHandlerFunc(func(params operations.UpdateListParams) middleware.Responder {
			return middleware.NotImplemented("operation .UpdateList has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
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
