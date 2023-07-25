// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Lists all users
	// (GET /user)
	GetUser(ctx echo.Context) error
	// Creates a user
	// (POST /user)
	PostUser(ctx echo.Context) error
	// Deletes a user by ID
	// (DELETE /user/{id})
	DeleteUserId(ctx echo.Context, id uuid.UUID) error
	// Retries a user by ID
	// (GET /user/{id})
	GetUserId(ctx echo.Context, id uuid.UUID) error
	// Updates a user
	// (PUT /user/{id})
	PutUserId(ctx echo.Context, id uuid.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUser(ctx)
	return err
}

// PostUser converts echo context to params.
func (w *ServerInterfaceWrapper) PostUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUser(ctx)
	return err
}

// DeleteUserId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUserId(ctx, id)
	return err
}

// GetUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserId(ctx, id)
	return err
}

// PutUserId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUserId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/user", wrapper.GetUser)
	router.POST(baseURL+"/user", wrapper.PostUser)
	router.DELETE(baseURL+"/user/:id", wrapper.DeleteUserId)
	router.GET(baseURL+"/user/:id", wrapper.GetUserId)
	router.PUT(baseURL+"/user/:id", wrapper.PutUserId)

}