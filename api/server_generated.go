// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение всех баннеров c фильтрацией по фиче и/или тегу
	// (GET /banner)
	GetBanner(ctx echo.Context, params GetBannerParams) error
	// Создание нового баннера
	// (POST /banner)
	PostBanner(ctx echo.Context, params PostBannerParams) error
	// Удаление баннера по идентификатору
	// (DELETE /banner/{id})
	DeleteBannerId(ctx echo.Context, id int, params DeleteBannerIdParams) error
	// Обновление содержимого баннера
	// (PATCH /banner/{id})
	PatchBannerId(ctx echo.Context, id int, params PatchBannerIdParams) error
	// Получение баннера для пользователя
	// (GET /user_banner)
	GetUserBanner(ctx echo.Context, params GetUserBannerParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetBanner converts echo context to params.
func (w *ServerInterfaceWrapper) GetBanner(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBannerParams
	// ------------- Optional query parameter "feature_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "feature_id", ctx.QueryParams(), &params.FeatureId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter feature_id: %s", err))
	}

	// ------------- Optional query parameter "tag_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag_id", ctx.QueryParams(), &params.TagId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tag_id: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for token, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
		}

		params.Token = &Token
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBanner(ctx, params)
	return err
}

// PostBanner converts echo context to params.
func (w *ServerInterfaceWrapper) PostBanner(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostBannerParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for token, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
		}

		params.Token = &Token
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostBanner(ctx, params)
	return err
}

// DeleteBannerId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteBannerId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteBannerIdParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for token, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
		}

		params.Token = &Token
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteBannerId(ctx, id, params)
	return err
}

// PatchBannerId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchBannerId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchBannerIdParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for token, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
		}

		params.Token = &Token
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchBannerId(ctx, id, params)
	return err
}

// GetUserBanner converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserBanner(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserBannerParams
	// ------------- Required query parameter "tag_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "tag_id", ctx.QueryParams(), &params.TagId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tag_id: %s", err))
	}

	// ------------- Required query parameter "feature_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "feature_id", ctx.QueryParams(), &params.FeatureId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter feature_id: %s", err))
	}

	// ------------- Optional query parameter "use_last_revision" -------------

	err = runtime.BindQueryParameter("form", true, false, "use_last_revision", ctx.QueryParams(), &params.UseLastRevision)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter use_last_revision: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Optional header parameter "token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("token")]; found {
		var Token string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for token, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "token", valueList[0], &Token, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
		}

		params.Token = &Token
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserBanner(ctx, params)
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

	router.GET(baseURL+"/banner", wrapper.GetBanner)
	router.POST(baseURL+"/banner", wrapper.PostBanner)
	router.DELETE(baseURL+"/banner/:id", wrapper.DeleteBannerId)
	router.PATCH(baseURL+"/banner/:id", wrapper.PatchBannerId)
	router.GET(baseURL+"/user_banner", wrapper.GetUserBanner)

}