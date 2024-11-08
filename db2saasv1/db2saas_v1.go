/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

// Package db2saasv1 : Operations and models for the Db2saasV1 service
package db2saasv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/cloud-db2-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// Db2saasV1 : Manage lifecycle of your Db2 on Cloud resources using the  APIs.
//
// API Version: 1.0.0
type Db2saasV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.db2.saas.ibm.com/dbapi/v4"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "db2saas"

const ParameterizedServiceURL = "https://{region}.db2.saas.ibm.com/dbapi/v4"

var defaultUrlVariables = map[string]string{
	"region": "us-south",
}

// Db2saasV1Options : Service options
type Db2saasV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDb2saasV1UsingExternalConfig : constructs an instance of Db2saasV1 with passed in options and external configuration.
func NewDb2saasV1UsingExternalConfig(options *Db2saasV1Options) (db2saas *Db2saasV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	db2saas, err = NewDb2saasV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = db2saas.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = db2saas.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDb2saasV1 : constructs an instance of Db2saasV1 with passed in options.
func NewDb2saasV1(options *Db2saasV1Options) (service *Db2saasV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &Db2saasV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "db2saas" suitable for processing requests.
func (db2saas *Db2saasV1) Clone() *Db2saasV1 {
	if core.IsNil(db2saas) {
		return nil
	}
	clone := *db2saas
	clone.Service = db2saas.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (db2saas *Db2saasV1) SetServiceURL(url string) error {
	err := db2saas.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (db2saas *Db2saasV1) GetServiceURL() string {
	return db2saas.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (db2saas *Db2saasV1) SetDefaultHeaders(headers http.Header) {
	db2saas.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (db2saas *Db2saasV1) SetEnableGzipCompression(enableGzip bool) {
	db2saas.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (db2saas *Db2saasV1) GetEnableGzipCompression() bool {
	return db2saas.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (db2saas *Db2saasV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	db2saas.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (db2saas *Db2saasV1) DisableRetries() {
	db2saas.Service.DisableRetries()
}

// GetDb2SaasConnectionInfo : Get Db2 connection information
func (db2saas *Db2saasV1) GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) (result *SuccessConnectionInfo, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasConnectionInfoWithContext(context.Background(), getDb2SaasConnectionInfoOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasConnectionInfoWithContext is an alternate form of the GetDb2SaasConnectionInfo method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasConnectionInfoWithContext(ctx context.Context, getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) (result *SuccessConnectionInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasConnectionInfoOptions, "getDb2SaasConnectionInfoOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasConnectionInfoOptions, "getDb2SaasConnectionInfoOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *getDb2SaasConnectionInfoOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/connectioninfo/{deployment_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasConnectionInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasConnectionInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasConnectionInfoOptions.ID != nil {
		builder.AddHeader("id", fmt.Sprint(*getDb2SaasConnectionInfoOptions.ID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_connection_info", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessConnectionInfo)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// GetDb2SaasConnectionInfoOptions : The GetDb2SaasConnectionInfo options.
type GetDb2SaasConnectionInfoOptions struct {
	// Encoded CRN deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// CRN deployment id.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasConnectionInfoOptions : Instantiate GetDb2SaasConnectionInfoOptions
func (*Db2saasV1) NewGetDb2SaasConnectionInfoOptions(deploymentID string, id string) *GetDb2SaasConnectionInfoOptions {
	return &GetDb2SaasConnectionInfoOptions{
		DeploymentID: core.StringPtr(deploymentID),
		ID: core.StringPtr(id),
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *GetDb2SaasConnectionInfoOptions) SetDeploymentID(deploymentID string) *GetDb2SaasConnectionInfoOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetID : Allow user to set ID
func (_options *GetDb2SaasConnectionInfoOptions) SetID(id string) *GetDb2SaasConnectionInfoOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasConnectionInfoOptions) SetHeaders(param map[string]string) *GetDb2SaasConnectionInfoOptions {
	options.Headers = param
	return options
}

// SuccessConnectionInfo : Responds with JSON of the connection information for the Db2 SaaS Instance.
type SuccessConnectionInfo struct {
	Public *SuccessConnectionInfoPublic `json:"public,omitempty"`

	Private *SuccessConnectionInfoPrivate `json:"private,omitempty"`
}

// UnmarshalSuccessConnectionInfo unmarshals an instance of SuccessConnectionInfo from the specified map of raw messages.
func UnmarshalSuccessConnectionInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfo)
	err = core.UnmarshalModel(m, "public", &obj.Public, UnmarshalSuccessConnectionInfoPublic)
	if err != nil {
		err = core.SDKErrorf(err, "", "public-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "private", &obj.Private, UnmarshalSuccessConnectionInfoPrivate)
	if err != nil {
		err = core.SDKErrorf(err, "", "private-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfoPrivate : SuccessConnectionInfoPrivate struct
type SuccessConnectionInfoPrivate struct {
	Hostname *string `json:"hostname,omitempty"`

	DatabaseName *string `json:"databaseName,omitempty"`

	HostRos *string `json:"host_ros,omitempty"`

	CertificateBase64 *string `json:"certificateBase64,omitempty"`

	SslPort *string `json:"sslPort,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	DatabaseVersion *string `json:"databaseVersion,omitempty"`
}

// UnmarshalSuccessConnectionInfoPrivate unmarshals an instance of SuccessConnectionInfoPrivate from the specified map of raw messages.
func UnmarshalSuccessConnectionInfoPrivate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfoPrivate)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseName", &obj.DatabaseName)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_ros", &obj.HostRos)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_ros-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "certificateBase64", &obj.CertificateBase64)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificateBase64-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sslPort", &obj.SslPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "sslPort-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		err = core.SDKErrorf(err, "", "ssl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseVersion", &obj.DatabaseVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseVersion-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfoPublic : SuccessConnectionInfoPublic struct
type SuccessConnectionInfoPublic struct {
	Hostname *string `json:"hostname,omitempty"`

	DatabaseName *string `json:"databaseName,omitempty"`

	HostRos *string `json:"host_ros,omitempty"`

	CertificateBase64 *string `json:"certificateBase64,omitempty"`

	SslPort *string `json:"sslPort,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	DatabaseVersion *string `json:"databaseVersion,omitempty"`
}

// UnmarshalSuccessConnectionInfoPublic unmarshals an instance of SuccessConnectionInfoPublic from the specified map of raw messages.
func UnmarshalSuccessConnectionInfoPublic(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfoPublic)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseName", &obj.DatabaseName)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_ros", &obj.HostRos)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_ros-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "certificateBase64", &obj.CertificateBase64)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificateBase64-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sslPort", &obj.SslPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "sslPort-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		err = core.SDKErrorf(err, "", "ssl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseVersion", &obj.DatabaseVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseVersion-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
