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

// CreateUserParams defines model for CreateUserParams.
type CreateUserParams struct {
	Bio          *string `json:"bio,omitempty"`
	CaptchaToken string  `json:"captcha_token"`
	Email        *string `json:"email,omitempty"`
	Fullname     *string `json:"fullname,omitempty"`
	Institution  *string `json:"institution,omitempty"`
	Password     string  `json:"password"`
	Username     string  `json:"username"`
	Website      *string `json:"website,omitempty"`
}

// Electrode defines model for Electrode.
type Electrode struct {
	ElectrodeId string  `json:"electrode_id"`
	Value       float32 `json:"value"`
}

// Frame defines model for Frame.
type Frame struct {
	Duration     float32               `json:"duration"`
	Electrodes   []Electrode           `json:"electrodes"`
	Magnets      *[]IndexedMagnet      `json:"magnets,omitempty"`
	Rank         float32               `json:"rank"`
	Temperatures *[]IndexedTemperature `json:"temperatures,omitempty"`
}

// FullDevice defines model for FullDevice.
type FullDevice struct {
	Electrodes []string `json:"electrodes"`
	Id         float32  `json:"id"`
	Name       string   `json:"name"`
}

// FullProtocol defines model for FullProtocol.
type FullProtocol struct {
	AuthorList    *[]RankedAuthor `json:"author_list,omitempty"`
	Description   *string         `json:"description,omitempty"`
	DeviceId      float32         `json:"device_id"`
	FrameCount    float32         `json:"frame_count"`
	Frames        []Frame         `json:"frames"`
	Id            float32         `json:"id"`
	Name          string          `json:"name"`
	Public        bool            `json:"public"`
	TotalDuration float32         `json:"total_duration"`
}

// IndexedMagnet defines model for IndexedMagnet.
type IndexedMagnet struct {
	Index float32 `json:"index"`
	Value bool    `json:"value"`
}

// IndexedTemperature defines model for IndexedTemperature.
type IndexedTemperature struct {
	Index float32 `json:"index"`
	Value float32 `json:"value"`
}

// LoginParams defines model for LoginParams.
type LoginParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginToken defines model for LoginToken.
type LoginToken struct {
	ExpirationDate *string `json:"expiration_date,omitempty"`
	Token          *string `json:"token,omitempty"`
}

// PublicUser defines model for PublicUser.
type PublicUser struct {
	Bio         *string `json:"bio,omitempty"`
	Email       *string `json:"email,omitempty"`
	Fullname    *string `json:"fullname,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Institution *string `json:"institution,omitempty"`
	Username    *string `json:"username,omitempty"`
	Website     *string `json:"website,omitempty"`
}

// RankedAuthor defines model for RankedAuthor.
type RankedAuthor struct {
	Author string  `json:"author"`
	Rank   float32 `json:"rank"`
}

// ShortProtocol defines model for ShortProtocol.
type ShortProtocol struct {
	AuthorList    []RankedAuthor `json:"author_list"`
	AuthorRank    float32        `json:"author_rank"`
	Description   string         `json:"description"`
	DeviceId      float32        `json:"device_id"`
	FrameCount    float32        `json:"frame_count"`
	Id            float32        `json:"id"`
	MaskFrame     []Frame        `json:"mask_frame"`
	Name          string         `json:"name"`
	Public        bool           `json:"public"`
	TotalDuration float32        `json:"total_duration"`
}

// ShortProtocolsList defines model for ShortProtocolsList.
type ShortProtocolsList struct {
	Protocols []ShortProtocol `json:"protocols"`
}

// UploadProtocolParams defines model for UploadProtocolParams.
type UploadProtocolParams struct {
	AuthorList  []RankedAuthor `json:"author_list"`
	Description *string        `json:"description,omitempty"`
	DeviceId    float32        `json:"device_id"`
	Frames      []Frame        `json:"frames"`
	Name        string         `json:"name"`
	Public      bool           `json:"public"`
}

// User defines model for User.
type User struct {
	Email    *string       `json:"email,omitempty"`
	Id       *int64        `json:"id,omitempty"`
	Tokens   *[]LoginToken `json:"tokens,omitempty"`
	Username *string       `json:"username,omitempty"`
}

// UploadProtocolJSONBody defines parameters for UploadProtocol.
type UploadProtocolJSONBody = UploadProtocolParams

// OverwriteProtocolJSONBody defines parameters for OverwriteProtocol.
type OverwriteProtocolJSONBody = UploadProtocolParams

// GetPublicProtocolsListParams defines parameters for GetPublicProtocolsList.
type GetPublicProtocolsListParams struct {
	// The number of items to skip before starting to collect the result set
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// The numbers of items to return
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody = CreateUserParams

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody = LoginParams

// UploadProtocolJSONRequestBody defines body for UploadProtocol for application/json ContentType.
type UploadProtocolJSONRequestBody = UploadProtocolJSONBody

// OverwriteProtocolJSONRequestBody defines body for OverwriteProtocol for application/json ContentType.
type OverwriteProtocolJSONRequestBody = OverwriteProtocolJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUserJSONBody

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody = LoginUserJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a particular dive infos by its ID
	// (GET /hardware/device/all)
	GetDevices(ctx echo.Context) error
	// Upload a protocol
	// (POST /protocol)
	UploadProtocol(ctx echo.Context) error
	// Get token bearer protocols list
	// (GET /protocol/me)
	GetSelfProtocolList(ctx echo.Context) error
	// delete a particular entire protocol by its ID
	// (DELETE /protocol/{protocolID})
	DeleteProtocol(ctx echo.Context, protocolID int) error
	// Get a particular protocol by its ID
	// (GET /protocol/{protocolID})
	GetProtocol(ctx echo.Context, protocolID int) error
	// update a particular entire protocol by its ID
	// (PUT /protocol/{protocolID})
	OverwriteProtocol(ctx echo.Context, protocolID int) error
	// Get public protocols list
	// (GET /public/protocol/all)
	GetPublicProtocolsList(ctx echo.Context, params GetPublicProtocolsListParams) error
	// Get a particular protocol by its ID
	// (GET /public/protocol/{protocolID})
	GetPublicProtocol(ctx echo.Context, protocolID int) error
	// Serve a json file representing this swaggerfile
	// (GET /swagger.json)
	ServeSwaggerFile(ctx echo.Context) error
	// Create user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Logs user into the system
	// (POST /user/login)
	LoginUser(ctx echo.Context) error
	// Get user infos of token bearer
	// (GET /user/me)
	GetSelfUser(ctx echo.Context) error
	// get a user
	// (GET /user/{username})
	GetUser(ctx echo.Context, username string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDevices converts echo context to params.
func (w *ServerInterfaceWrapper) GetDevices(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDevices(ctx)
	return err
}

// UploadProtocol converts echo context to params.
func (w *ServerInterfaceWrapper) UploadProtocol(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UploadProtocol(ctx)
	return err
}

// GetSelfProtocolList converts echo context to params.
func (w *ServerInterfaceWrapper) GetSelfProtocolList(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSelfProtocolList(ctx)
	return err
}

// DeleteProtocol converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteProtocol(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "protocolID" -------------
	var protocolID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "protocolID", runtime.ParamLocationPath, ctx.Param("protocolID"), &protocolID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter protocolID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteProtocol(ctx, protocolID)
	return err
}

// GetProtocol converts echo context to params.
func (w *ServerInterfaceWrapper) GetProtocol(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "protocolID" -------------
	var protocolID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "protocolID", runtime.ParamLocationPath, ctx.Param("protocolID"), &protocolID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter protocolID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProtocol(ctx, protocolID)
	return err
}

// OverwriteProtocol converts echo context to params.
func (w *ServerInterfaceWrapper) OverwriteProtocol(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "protocolID" -------------
	var protocolID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "protocolID", runtime.ParamLocationPath, ctx.Param("protocolID"), &protocolID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter protocolID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.OverwriteProtocol(ctx, protocolID)
	return err
}

// GetPublicProtocolsList converts echo context to params.
func (w *ServerInterfaceWrapper) GetPublicProtocolsList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPublicProtocolsListParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPublicProtocolsList(ctx, params)
	return err
}

// GetPublicProtocol converts echo context to params.
func (w *ServerInterfaceWrapper) GetPublicProtocol(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "protocolID" -------------
	var protocolID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "protocolID", runtime.ParamLocationPath, ctx.Param("protocolID"), &protocolID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter protocolID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPublicProtocol(ctx, protocolID)
	return err
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

// GetSelfUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetSelfUser(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSelfUser(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUser(ctx, username)
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

	router.GET(baseURL+"/hardware/device/all", wrapper.GetDevices)
	router.POST(baseURL+"/protocol", wrapper.UploadProtocol)
	router.GET(baseURL+"/protocol/me", wrapper.GetSelfProtocolList)
	router.DELETE(baseURL+"/protocol/:protocolID", wrapper.DeleteProtocol)
	router.GET(baseURL+"/protocol/:protocolID", wrapper.GetProtocol)
	router.PUT(baseURL+"/protocol/:protocolID", wrapper.OverwriteProtocol)
	router.GET(baseURL+"/public/protocol/all", wrapper.GetPublicProtocolsList)
	router.GET(baseURL+"/public/protocol/:protocolID", wrapper.GetPublicProtocol)
	router.GET(baseURL+"/swagger.json", wrapper.ServeSwaggerFile)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.POST(baseURL+"/user/login", wrapper.LoginUser)
	router.GET(baseURL+"/user/me", wrapper.GetSelfUser)
	router.GET(baseURL+"/user/:username", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"k+sHu9vdfwMAAP//GI8MLn4qAAA=",
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
