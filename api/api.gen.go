// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for AWSRegionEnum.
const (
	AfSouth1     AWSRegionEnum = "af-south-1"
	ApEast1      AWSRegionEnum = "ap-east-1"
	ApNortheast1 AWSRegionEnum = "ap-northeast-1"
	ApNortheast2 AWSRegionEnum = "ap-northeast-2"
	ApNortheast3 AWSRegionEnum = "ap-northeast-3"
	ApSouth1     AWSRegionEnum = "ap-south-1"
	ApSouth2     AWSRegionEnum = "ap-south-2"
	ApSoutheast1 AWSRegionEnum = "ap-southeast-1"
	ApSoutheast2 AWSRegionEnum = "ap-southeast-2"
	ApSoutheast3 AWSRegionEnum = "ap-southeast-3"
	ApSoutheast4 AWSRegionEnum = "ap-southeast-4"
	CaCentral1   AWSRegionEnum = "ca-central-1"
	CaWest1      AWSRegionEnum = "ca-west-1"
	EuCentral1   AWSRegionEnum = "eu-central-1"
	EuCentral2   AWSRegionEnum = "eu-central-2"
	EuNorth1     AWSRegionEnum = "eu-north-1"
	EuSouth1     AWSRegionEnum = "eu-south-1"
	EuSouth2     AWSRegionEnum = "eu-south-2"
	EuWest1      AWSRegionEnum = "eu-west-1"
	EuWest2      AWSRegionEnum = "eu-west-2"
	EuWest3      AWSRegionEnum = "eu-west-3"
	IlCentral1   AWSRegionEnum = "il-central-1"
	MeCentral1   AWSRegionEnum = "me-central-1"
	MeSouth1     AWSRegionEnum = "me-south-1"
	SaEast1      AWSRegionEnum = "sa-east-1"
	UsEast1      AWSRegionEnum = "us-east-1"
	UsEast2      AWSRegionEnum = "us-east-2"
	UsGovEast1   AWSRegionEnum = "us-gov-east-1"
	UsGovWest1   AWSRegionEnum = "us-gov-west-1"
	UsWest1      AWSRegionEnum = "us-west-1"
	UsWest2      AWSRegionEnum = "us-west-2"
)

// Defines values for AwsRDSEngineEnum.
const (
	AuroraMysql      AwsRDSEngineEnum = "aurora_mysql"
	AuroraPostgresql AwsRDSEngineEnum = "aurora_postgresql"
	Mysql            AwsRDSEngineEnum = "mysql"
	Postgresql       AwsRDSEngineEnum = "postgresql"
)

// Defines values for OsEnum.
const (
	Linux   OsEnum = "linux"
	Windows OsEnum = "windows"
)

// AWSRegion defines model for AWSRegion.
type AWSRegion struct {
	Region AWSRegionEnum `json:"region" validate:"oneof=ap-south-1 ap-south-2 ap-northeast-1 ap-northeast-2 ap-northeast-3 ap-southeast-1 ap-southeast-2 ap-southeast-3 ap-southeast-4 ap-east-1 us-east-1 us-east-2 us-west-1 us-west-2 ca-central-1 ca-west-1 us-gov-east-1 us-gov-west-1 sa-east-1 eu-west-1 eu-west-2 eu-west-3 eu-central-1 eu-central-2 eu-north-1 eu-south-1 eu-south-2 il-central-1 af-south-1 me-central-1 me-south-1"`
}

// AWSRegionEnum defines model for AWSRegionEnum.
type AWSRegionEnum string

// AwsRDSEngine defines model for awsRDSEngine.
type AwsRDSEngine struct {
	Engine AwsRDSEngineEnum `json:"engine" validate:"oneof=aurora_mysql aurora_postgresql mysql postgresql"`
}

// AwsRDSEngineEnum defines model for awsRDSEngineEnum.
type AwsRDSEngineEnum string

// Os defines model for os.
type Os struct {
	Os OsEnum `json:"os" validate:"oneof=linux windows"`
}

// OsEnum defines model for osEnum.
type OsEnum string

// GetAwsEc2InstancesParams defines parameters for GetAwsEc2Instances.
type GetAwsEc2InstancesParams struct {
	// Region Filter instance by region
	Region *AWSRegion `form:"region,omitempty" json:"region,omitempty"`

	// Os Filter instance by name
	Os *Os `form:"os,omitempty" json:"os,omitempty"`

	// Instancetype Filter instance by type
	Instancetype *string `form:"instancetype,omitempty" json:"instancetype,omitempty"`
}

// GetAwsRdsInstancesParams defines parameters for GetAwsRdsInstances.
type GetAwsRdsInstancesParams struct {
	// Region Filter instance by region
	Region *AWSRegion `form:"region,omitempty" json:"region,omitempty"`

	// Engine Filter instance by engine
	Engine *AwsRDSEngine `form:"engine,omitempty" json:"engine,omitempty"`

	// Instancetype Filter instance by type
	Instancetype *string `form:"instancetype,omitempty" json:"instancetype,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get EC2 instances cost with optional filtering
	// (GET /aws/ec2/instances)
	GetAwsEc2Instances(c *gin.Context, params GetAwsEc2InstancesParams)
	// Get EC2 instance cost by SKU
	// (GET /aws/ec2/instances/{instance_sku})
	GetAwsEc2InstancesInstanceSku(c *gin.Context, instanceSku string)
	// Get AWS RDS instances cost with optional filtering
	// (GET /aws/rds/instances)
	GetAwsRdsInstances(c *gin.Context, params GetAwsRdsInstancesParams)
	// Get RDS instance cost by SKU
	// (GET /aws/rds/instances/{instance_sku})
	GetAwsRdsInstancesInstanceSku(c *gin.Context, instanceSku string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAwsEc2Instances operation middleware
func (siw *ServerInterfaceWrapper) GetAwsEc2Instances(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAwsEc2InstancesParams

	// ------------- Optional query parameter "region" -------------

	err = runtime.BindQueryParameter("form", true, false, "region", c.Request.URL.Query(), &params.Region)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter region: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "os" -------------

	err = runtime.BindQueryParameter("form", true, false, "os", c.Request.URL.Query(), &params.Os)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter os: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "instancetype" -------------

	err = runtime.BindQueryParameter("form", true, false, "instancetype", c.Request.URL.Query(), &params.Instancetype)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter instancetype: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAwsEc2Instances(c, params)
}

// GetAwsEc2InstancesInstanceSku operation middleware
func (siw *ServerInterfaceWrapper) GetAwsEc2InstancesInstanceSku(c *gin.Context) {

	var err error

	// ------------- Path parameter "instance_sku" -------------
	var instanceSku string

	err = runtime.BindStyledParameterWithOptions("simple", "instance_sku", c.Param("instance_sku"), &instanceSku, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter instance_sku: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAwsEc2InstancesInstanceSku(c, instanceSku)
}

// GetAwsRdsInstances operation middleware
func (siw *ServerInterfaceWrapper) GetAwsRdsInstances(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAwsRdsInstancesParams

	// ------------- Optional query parameter "region" -------------

	err = runtime.BindQueryParameter("form", true, false, "region", c.Request.URL.Query(), &params.Region)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter region: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "engine" -------------

	err = runtime.BindQueryParameter("form", true, false, "engine", c.Request.URL.Query(), &params.Engine)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter engine: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "instancetype" -------------

	err = runtime.BindQueryParameter("form", true, false, "instancetype", c.Request.URL.Query(), &params.Instancetype)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter instancetype: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAwsRdsInstances(c, params)
}

// GetAwsRdsInstancesInstanceSku operation middleware
func (siw *ServerInterfaceWrapper) GetAwsRdsInstancesInstanceSku(c *gin.Context) {

	var err error

	// ------------- Path parameter "instance_sku" -------------
	var instanceSku string

	err = runtime.BindStyledParameterWithOptions("simple", "instance_sku", c.Param("instance_sku"), &instanceSku, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter instance_sku: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAwsRdsInstancesInstanceSku(c, instanceSku)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/aws/ec2/instances", wrapper.GetAwsEc2Instances)
	router.GET(options.BaseURL+"/aws/ec2/instances/:instance_sku", wrapper.GetAwsEc2InstancesInstanceSku)
	router.GET(options.BaseURL+"/aws/rds/instances", wrapper.GetAwsRdsInstances)
	router.GET(options.BaseURL+"/aws/rds/instances/:instance_sku", wrapper.GetAwsRdsInstancesInstanceSku)
}

type GetAwsEc2InstancesRequestObject struct {
	Params GetAwsEc2InstancesParams
}

type GetAwsEc2InstancesResponseObject interface {
	VisitGetAwsEc2InstancesResponse(w http.ResponseWriter) error
}

type GetAwsEc2Instances200JSONResponse []struct {
	Id             *string  `json:"id,omitempty"`
	Instancefamily *string  `json:"instancefamily,omitempty"`
	Instancetype   *string  `json:"instancetype,omitempty"`
	Memory         *float32 `json:"memory,omitempty"`
	Ondemandprice  *float32 `json:"ondemandprice,omitempty"`
	Vcpu           *float32 `json:"vcpu,omitempty"`
}

func (response GetAwsEc2Instances200JSONResponse) VisitGetAwsEc2InstancesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAwsEc2Instances400Response struct {
}

func (response GetAwsEc2Instances400Response) VisitGetAwsEc2InstancesResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetAwsEc2InstancesInstanceSkuRequestObject struct {
	InstanceSku string `json:"instance_sku"`
}

type GetAwsEc2InstancesInstanceSkuResponseObject interface {
	VisitGetAwsEc2InstancesInstanceSkuResponse(w http.ResponseWriter) error
}

type GetAwsEc2InstancesInstanceSku200JSONResponse struct {
	Clockspeed                        *string  `json:"clockspeed,omitempty"`
	DedicatedebsThroughput            *string  `gorm:"column:dedicatedebsthroughput" json:"dedicatedebs_throughput,omitempty"`
	Ecu                               *float32 `json:"ecu,omitempty"`
	GpuMemory                         *float32 `gorm:"column:gpumemory" json:"gpu_memory,omitempty"`
	Id                                *string  `json:"id,omitempty"`
	InstanceFamily                    *string  `gorm:"column:instancefamily" json:"instance_family,omitempty"`
	InstanceType                      *string  `gorm:"column:instancetype" json:"instance_type,omitempty"`
	Memory                            *float32 `json:"memory,omitempty"`
	NetworkperFormance                *string  `gorm:"column:networkperformance" json:"networkper_formance,omitempty"`
	OndemandPrice                     *float32 `gorm:"column:ondemandprice" json:"ondemand_price,omitempty"`
	OneYearReservedConvertiblePrice   *float32 `json:"one_year_reserved_convertible_price,omitempty"`
	OneYearReservedStandardPrice      *float32 `json:"one_year_reserved_standard_price,omitempty"`
	Os                                *string  `gorm:"column:operatingsystem" json:"os,omitempty"`
	PhysicalProcessor                 *string  `gorm:"column:physicalprocessor" json:"physical_processor,omitempty"`
	ProcessorFeatures                 *string  `gorm:"column:processorfeatures" json:"processor_features,omitempty"`
	Region                            *string  `gorm:"column:regioncode" json:"region,omitempty"`
	Storage                           *string  `json:"storage,omitempty"`
	ThreeYearReservedConvertiblePrice *float32 `json:"three_year_reserved_convertible_price,omitempty"`
	ThreeYearReservedStandardPrice    *float32 `json:"three_year_reserved_standard_price,omitempty"`
	Vcpu                              *float32 `json:"vcpu,omitempty"`
}

func (response GetAwsEc2InstancesInstanceSku200JSONResponse) VisitGetAwsEc2InstancesInstanceSkuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAwsEc2InstancesInstanceSku404Response struct {
}

func (response GetAwsEc2InstancesInstanceSku404Response) VisitGetAwsEc2InstancesInstanceSkuResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetAwsRdsInstancesRequestObject struct {
	Params GetAwsRdsInstancesParams
}

type GetAwsRdsInstancesResponseObject interface {
	VisitGetAwsRdsInstancesResponse(w http.ResponseWriter) error
}

type GetAwsRdsInstances200JSONResponse []struct {
	DatabaseEngine *string  `gorm:"column:databaseengine" json:"database_engine,omitempty"`
	Id             *string  `json:"id,omitempty"`
	InstanceFamily *string  `gorm:"column:instancefamily" json:"instance_family,omitempty"`
	InstanceType   *string  `gorm:"column:instancetype" json:"instance_type,omitempty"`
	Memory         *float32 `json:"memory,omitempty"`
	Ondemandprice  *float32 `gorm:"column:ondemandprice" json:"ondemandprice,omitempty"`
	Regioncode     *string  `json:"regioncode,omitempty"`
	Vcpu           *float32 `json:"vcpu,omitempty"`
}

func (response GetAwsRdsInstances200JSONResponse) VisitGetAwsRdsInstancesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAwsRdsInstances400Response struct {
}

func (response GetAwsRdsInstances400Response) VisitGetAwsRdsInstancesResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetAwsRdsInstancesInstanceSkuRequestObject struct {
	InstanceSku string `json:"instance_sku"`
}

type GetAwsRdsInstancesInstanceSkuResponseObject interface {
	VisitGetAwsRdsInstancesInstanceSkuResponse(w http.ResponseWriter) error
}

type GetAwsRdsInstancesInstanceSku200JSONResponse struct {
	Clockspeed                     *string  `json:"clockspeed,omitempty"`
	DatabaseEngine                 *string  `gorm:"column:databaseengine" json:"database_engine,omitempty"`
	DedicatedEbsThroughput         *string  `gorm:"column:dedicatedebsthroughput" json:"dedicated_ebs_throughput,omitempty"`
	Id                             *string  `json:"id,omitempty"`
	InstanceFamily                 *string  `gorm:"column:instancefamily" json:"instance_family,omitempty"`
	InstanceType                   *string  `gorm:"column:instancetype" json:"instance_type,omitempty"`
	Memory                         *float32 `json:"memory,omitempty"`
	NetworkperFormance             *string  `gorm:"column:networkperformance" json:"networkper_formance,omitempty"`
	OndemandPrice                  *float32 `gorm:"column:ondemandprice" json:"ondemand_price,omitempty"`
	OneYearReservedStandardPrice   *float32 `json:"one_year_reserved_standard_price,omitempty"`
	PhysicalProcessor              *string  `gorm:"column:physicalprocessor" json:"physical_processor,omitempty"`
	Region                         *string  `gorm:"column:regioncode" json:"region,omitempty"`
	Storage                        *string  `json:"storage,omitempty"`
	ThreeYearReservedStandardPrice *float32 `json:"three_year_reserved_standard_price,omitempty"`
	Vcpu                           *float32 `json:"vcpu,omitempty"`
}

func (response GetAwsRdsInstancesInstanceSku200JSONResponse) VisitGetAwsRdsInstancesInstanceSkuResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAwsRdsInstancesInstanceSku404Response struct {
}

func (response GetAwsRdsInstancesInstanceSku404Response) VisitGetAwsRdsInstancesInstanceSkuResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get EC2 instances cost with optional filtering
	// (GET /aws/ec2/instances)
	GetAwsEc2Instances(ctx context.Context, request GetAwsEc2InstancesRequestObject) (GetAwsEc2InstancesResponseObject, error)
	// Get EC2 instance cost by SKU
	// (GET /aws/ec2/instances/{instance_sku})
	GetAwsEc2InstancesInstanceSku(ctx context.Context, request GetAwsEc2InstancesInstanceSkuRequestObject) (GetAwsEc2InstancesInstanceSkuResponseObject, error)
	// Get AWS RDS instances cost with optional filtering
	// (GET /aws/rds/instances)
	GetAwsRdsInstances(ctx context.Context, request GetAwsRdsInstancesRequestObject) (GetAwsRdsInstancesResponseObject, error)
	// Get RDS instance cost by SKU
	// (GET /aws/rds/instances/{instance_sku})
	GetAwsRdsInstancesInstanceSku(ctx context.Context, request GetAwsRdsInstancesInstanceSkuRequestObject) (GetAwsRdsInstancesInstanceSkuResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetAwsEc2Instances operation middleware
func (sh *strictHandler) GetAwsEc2Instances(ctx *gin.Context, params GetAwsEc2InstancesParams) {
	var request GetAwsEc2InstancesRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAwsEc2Instances(ctx, request.(GetAwsEc2InstancesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAwsEc2Instances")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAwsEc2InstancesResponseObject); ok {
		if err := validResponse.VisitGetAwsEc2InstancesResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetAwsEc2InstancesInstanceSku operation middleware
func (sh *strictHandler) GetAwsEc2InstancesInstanceSku(ctx *gin.Context, instanceSku string) {
	var request GetAwsEc2InstancesInstanceSkuRequestObject

	request.InstanceSku = instanceSku

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAwsEc2InstancesInstanceSku(ctx, request.(GetAwsEc2InstancesInstanceSkuRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAwsEc2InstancesInstanceSku")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAwsEc2InstancesInstanceSkuResponseObject); ok {
		if err := validResponse.VisitGetAwsEc2InstancesInstanceSkuResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetAwsRdsInstances operation middleware
func (sh *strictHandler) GetAwsRdsInstances(ctx *gin.Context, params GetAwsRdsInstancesParams) {
	var request GetAwsRdsInstancesRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAwsRdsInstances(ctx, request.(GetAwsRdsInstancesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAwsRdsInstances")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAwsRdsInstancesResponseObject); ok {
		if err := validResponse.VisitGetAwsRdsInstancesResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetAwsRdsInstancesInstanceSku operation middleware
func (sh *strictHandler) GetAwsRdsInstancesInstanceSku(ctx *gin.Context, instanceSku string) {
	var request GetAwsRdsInstancesInstanceSkuRequestObject

	request.InstanceSku = instanceSku

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAwsRdsInstancesInstanceSku(ctx, request.(GetAwsRdsInstancesInstanceSkuRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAwsRdsInstancesInstanceSku")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetAwsRdsInstancesInstanceSkuResponseObject); ok {
		if err := validResponse.VisitGetAwsRdsInstancesInstanceSkuResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYX2/bOBL/KgTvHv0ndfIk4B5ybXrI7S5QxF30IQgEmhrLbCSS5Z+43kDffUFSoqhY",
	"SZwaLZrdfTI5HP6GnN9wZqx7TEUtBQduNM7usaYbqIkfnn9aXkHJBHcTqYQEZRj4JRXl/1awxhn+17yH",
	"mbcY8whwwW2Nm2aCFXyxTEGBs+sO42aCzU4CzrBYfQZqcDPBw43ZPQb/e42JnGphzWb6Bk/6ySJMuFBm",
	"A0SbbrUX7GmcJvvTLb1g8VCwt+UsCOJ2q/fHizDeQi/3YyenZEqBG0Uqv0TJQK0UdwM4N4/rmvRrYHt5",
	"N14k49MwTk0l01bTe6Zb610cJ06NVQMQsk4Uaxis1RDXen61UYyXeIK/TgWRbEpFASXwKXw1ikwNKX1o",
	"3ZGKFcT4gOAg1v/pOUc942jI93D6YPUUDZkeThfD6QPlMxQZRpHfOFqgyG0cLVDKK4qsogGnaMAoinyi",
	"yGYcLeLoFKU8phOv03KIegb74QKl7KGeO5QyhxLe3HMlW331bnnBS8ZhPwlAlD+VBFKM0TzQwozlgb29",
	"aSqwSiiS1zv9pXLRGKZSaFMqCLJuLREeH42JWbRnFAV5YtDdVuh95wXZU44TetRdQo+6qtVOHFQxbr/i",
	"Cd4yXoitPv7qHhB1cI07mAZqFTO7pTtzuNh/gShQ59Zs3GzlZ++FqonBGf7/p48ue3ltnLWrOB5sY4wM",
	"wIyvhdtfgKaKSeOrDT7/cInWQqGacFIyXiJaCVsgKrTRE2QUobdOajUpYYIIL5BQJeHsDyfVoO4YBY0Y",
	"NwJJJZzr9MxZZ6Zy5t96tLdCG/SbswA1cIPOP1ziCb4DpcMh3sxOZife5RI4kQxn+HR2MnNpTxKz8U5w",
	"UT8Hupgzrg3hNLimBOOpl6CIu9FlgTP8PzDnW31BF5dR1QEpUoMBpXF2/dAL71llQKEOGq12qK2lzm84",
	"w18sqB2eYE68l+NiiKyDq7YLvQNseyvjloU+2KrQB5rzoTJurlNrVXrDw7hvmhv3pLQUXAdiFicn7ocK",
	"boB7joiUFaOepflnHVqdHo8ZqEfeNCtGjE3isdakZtXuSZWwMKJQQy1UupfbegXKhyEvoCa8kIpRGNW4",
	"o9KOLDQjaaQVEKXILrzEISG/Mm2QWEdG3J6z4L2h4iX3+QMxLq3xtrSta+Lu4EIeXbxdRBDtXzDaMrNB",
	"wgOQCq099YGvyciDmt93w1zf2uYF76sbLG/tc0/t4wbQ8pff3Y3NBgaH7mLQvfn9EHRnwmniNsrC9wzJ",
	"YSTSStBbLQHGI7KAwiFBASudm40Sttw4mvZ0nykRpVA1zjAVla15lqImoI4+oHY0Mktp88dC+0W2S2lb",
	"HF89nn6I+WMv8UUmHzzrJjUw/o6/Cd5DNE+nAA5mK9StBJWvXanl9FjrPWIEbJJckz+SbF5kY5i4Ajzk",
	"OyAqV+CqNRQ5FfzOhfWqgvzxBLe/z/muIKp4apM+0kdtnuGl3mkDoVOTm51mlFS5VIKC1kIdaaQD7PG8",
	"mW6Wr4EYq+DYu0TAiBf6zu5v/hHQAcRphjJghCLleJUzGwXfFgFjOw+IgYMr434l7KqIr12hDJ6NlcFW",
	"iwuD1sLy4plaGErhaufKTl/5VKEPbiWvCv1aW8n2r+C47bh4mO3Bf9dX31oWxJAV0ZD3f7qPqdQtGkTv",
	"/O1K5jNd85GFLMl5Y179GRvy809LdPVu+eKmfJCaXtaUp5nq25vy9NCvuin/3k889uf5j2r7/+nD/yp9",
	"+AG91A9qfX+GnvTVdJZpbnzQWSZfb32OTb/bXt+41Ocv22Zgq6r2+6zO5nMi2cx/ep06zBkVNW5umj8D",
	"AAD//9evP6dLHAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
