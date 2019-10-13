// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	backend_src "climatehero/backend-src"
	climatehero_errors "climatehero/errors"
	"climatehero/models"
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
			if err := backend_src.AddUser(*params.Info.Email, *params.Info.Pin, *params.Info.Name); err != nil {
				return climatehero_errors.NewError(int(*err.Code), *err.Message)
			}

			return operations.NewAddUserCreated().WithPayload(&models.Name{params.Info.Name})
		})
	}
	if api.CheckItemHandler == nil {
		api.CheckItemHandler = operations.CheckItemHandlerFunc(func(params operations.CheckItemParams) middleware.Responder {
			points, msg, err := backend_src.CheckItem(params.Item.Email, params.Item.Item)
			if err != nil {
				return climatehero_errors.NewError(int(*err.Code), *err.Message)
			}
			points64 := int64(*points)
			return operations.NewCheckItemOK().WithPayload(&models.CheckReturn{&points64, msg})
		})
	}
	if api.GetUserInfoHandler == nil {
		api.GetUserInfoHandler = operations.GetUserInfoHandlerFunc(func(params operations.GetUserInfoParams) middleware.Responder {
			points, checklist, message, err := backend_src.GetUserInfo(params.Email)
			if err != nil {
				return climatehero_errors.NewError(int(*err.Code), *err.Message)
			}

			points64 := int64(*points)
			return operations.NewGetUserInfoOK().WithPayload(&models.UserInfo{checklist, &points64, message})
		})
	}
	if api.LoginHandler == nil {
		api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
			name, err := backend_src.Login(*params.Email.Email, *params.Email.Pin)
			if err != nil {
				return climatehero_errors.NewError(int(*err.Code), *err.Message)
			}

			return operations.NewLoginOK().WithPayload(&models.Name{name})
		})
	}
	if api.UpdateListHandler == nil {
		api.UpdateListHandler = operations.UpdateListHandlerFunc(func(params operations.UpdateListParams) middleware.Responder {
			if err := backend_src.UpdateList(params.Info.Email, params.Info.ListItems); err != nil {
				return climatehero_errors.NewError(int(*err.Code), *err.Message)
			}

			return operations.NewLoginOK()
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
