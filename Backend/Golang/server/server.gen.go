// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// LoginParams defines model for LoginParams.
type LoginParams struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// LoginToken defines model for LoginToken.
type LoginToken struct {
	ExpirationDate *string `json:"expiration_date,omitempty"`
	Token          *string `json:"token,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// User defines model for User.
type User struct {
	Email    *string `json:"email,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`

	// User Status
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody = User

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody = LoginParams

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody = User

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUserJSONBody

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody = LoginUserJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Serve a json file representing this swaggerfile
	// (GET /swagger.json)
	ServeSwaggerFile(ctx echo.Context) error
	// Create user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Logs user into the system
	// (GET /user/login)
	LoginUser(ctx echo.Context) error
	// Logs out current logged in user session
	// (GET /user/logout)
	LogoutUser(ctx echo.Context) error
	// Delete user
	// (DELETE /user/{username})
	DeleteUser(ctx echo.Context, username string) error
	// Get user by user name
	// (GET /user/{username})
	GetUserByName(ctx echo.Context, username string) error
	// Update user
	// (PUT /user/{username})
	UpdateUser(ctx echo.Context, username string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ServeSwaggerFile converts echo context to params.
func (w *ServerInterfaceWrapper) ServeSwaggerFile(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ServeSwaggerFile(ctx)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// LoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) LoginUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LoginUser(ctx)
	return err
}

// LogoutUser converts echo context to params.
func (w *ServerInterfaceWrapper) LogoutUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LogoutUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUser(ctx, username)
	return err
}

// GetUserByName converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserByName(ctx, username)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateUser(ctx, username)
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

	router.GET(baseURL+"/swagger.json", wrapper.ServeSwaggerFile)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.GET(baseURL+"/user/login", wrapper.LoginUser)
	router.GET(baseURL+"/user/logout", wrapper.LogoutUser)
	router.DELETE(baseURL+"/user/:username", wrapper.DeleteUser)
	router.GET(baseURL+"/user/:username", wrapper.GetUserByName)
	router.PUT(baseURL+"/user/:username", wrapper.UpdateUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RXbW/bthP/KgT//5ey5ad0hV61adciQLEUTTxsK4KBls4SO4nUyKMdN/B3H46UbMsP",
	"bZIlA4a9EkWRd7+7+92D7niqq1orUGh5cscN2ForC/7lnTYzmWWgfjRGG9rJwKZG1ii14gl/7bAAhTIV",
	"tMGkmmtTNWvLKmmtVDnThkm1EKXM+DriUyUcFtrIr5A9mdh1xG1aQCU87A86l+qjMKLyr7XRNRiUwaZa",
	"WLvUJqM13IqqLoEnfGrB0BcecVzVtGPRSJUTYmfBKFFB9wYWQJcOL6w3O3r2BVLkEb+tSrochHh5/pjH",
	"ea3/AHUIE25rabzNv2cC93SPBsNRbzDpjcbXw5fJ6CyZjPtnw+Fvx9BjK397/ZevZ5/ghxdYvrn+Ncsu",
	"fl5OL+vzi3/UdH//0OhKyLKr64su1Cu/3091dQyj7IZyOIh44AtPuFT4YrK9JBVCTgiiEzQYjsaTs1OO",
	"uEKBzh7SlWxhzcdoB0gXx3h0FMczeZgSAlJnJK6uKDGCf89BGDCUXvTmM4buzPz2VlmBWAcZlHp0NNUK",
	"RYo7MeKilgiiemWXIs/B9KUmc7qeuQJgucTCzVimU1eBwpDGAhkpsUkch+8U3Pi9cJn8IGY2fitziaI8",
	"l7rU+epjKZBcSQjBVPZyfgVmIdMW6yPESPR+bg6w5gRrj7DXHy9Yj13WoGg17g94xBdgbLBr2B/0h0Oy",
	"V9egRC15wsf9QX/MiVhYeGfHrWe+WO0TMAfvQCK898JF5l1kFnAVTr6TJfCoW4JHg8Eh46xLU7B27kq2",
	"kRaKoKsqYVatXCYYKWdzWQIzUBuwVFdVzrCQljUA50Etitzy5DNvdvkNCYxdm6ja4iGQNwYEkhoX2Nq1",
	"LXxtiGzgTwcWz3W2ahkFyosUdV02tT5ufRWqOa3+b2DOE/6/eNun4qbWx9O2nOyKuO0tl8seRbHnTAkq",
	"1Rlkf1dmyLEHSTjIhuCOzPuKbdKX/CINIUTjYL0X/gzmwpX4zB57CuvuwcmGLQ1XWr751y3Z4pK64sl8",
	"8T3zGSm1OzscsfJTUMmEN5b55soWMiQAleB401i+F9rRYPi0qMMkcc/QRA32NJCSatkkADqmZwM8Phzf",
	"ujH+oHMbGC4VaoYFMLuyCNW3I64dfivk2uEm5sez48Fk9EC1Q5Y6Y0AhK3WeQ8akCvAtWF/tT8O+a2O+",
	"DgBKCHNaF8o1VdpUKKZVuWIzYJlWwGYr75quzv5BBX3rhTam18RKQDCE5lANMMLCsBDIFEBG7PT6vAzi",
	"I6WV70882h0X/HKfrdEO8/ZnkJu9MEyO9aiLMJtvEoNZR/RuqTY5MUcpjWyuncr2whU8caJ2RMeJ8x48",
	"a85XPwULH+nAOWBaQNZnUxsADNlcG4ZgqZX22TN6tun+/8rSHz2KF+Pvl6C9v9KH0KkZiX34d4fhzzfk",
	"+i3b3gOGKjBbhWcTyQPe1Q6fOOWndSbul/Jdtj5vtv/XZ7cQFSYUg1tpyfq2yYUWh9oAXz9he2oUnpiV",
	"PJfNouWFM+XOr5DShbZIcZ5ro/Qy/JrtnkriuNSpKOlc8nLwcsApyo2SfbCXLULLxIwaZoOp88N5s/4r",
	"AAD//2Y78htJEgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
