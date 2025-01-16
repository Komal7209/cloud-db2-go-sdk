/**
 * (C) Copyright IBM Corp. 2025.
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
	if getDb2SaasConnectionInfoOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasConnectionInfoOptions.XDeploymentID))
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

// PostDb2SaasAllowlist : Allow listing of new IPs
func (db2saas *Db2saasV1) PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) (result *SuccessPostAllowedlistIPs, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasAllowlistWithContext(context.Background(), postDb2SaasAllowlistOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasAllowlistWithContext is an alternate form of the PostDb2SaasAllowlist method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasAllowlistWithContext(ctx context.Context, postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) (result *SuccessPostAllowedlistIPs, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasAllowlistOptions, "postDb2SaasAllowlistOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasAllowlistOptions, "postDb2SaasAllowlistOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/dbsettings/whitelistips`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if postDb2SaasAllowlistOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasAllowlistOptions.XDeploymentID))
	}

	body := make(map[string]interface{})
	if postDb2SaasAllowlistOptions.IpAddresses != nil {
		body["ip_addresses"] = postDb2SaasAllowlistOptions.IpAddresses
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_allowlist", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessPostAllowedlistIPs)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasAllowlist : Get allowed list of IPs
func (db2saas *Db2saasV1) GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) (result *SuccessGetAllowlistIPs, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasAllowlistWithContext(context.Background(), getDb2SaasAllowlistOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasAllowlistWithContext is an alternate form of the GetDb2SaasAllowlist method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasAllowlistWithContext(ctx context.Context, getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) (result *SuccessGetAllowlistIPs, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasAllowlistOptions, "getDb2SaasAllowlistOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasAllowlistOptions, "getDb2SaasAllowlistOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/dbsettings/whitelistips`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasAllowlistOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasAllowlistOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_allowlist", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetAllowlistIPs)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PostDb2SaasUser : Create new user ( available only for platform users)
func (db2saas *Db2saasV1) PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions) (result *SuccessUserResponse, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasUserWithContext(context.Background(), postDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasUserWithContext is an alternate form of the PostDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasUserWithContext(ctx context.Context, postDb2SaasUserOptions *PostDb2SaasUserOptions) (result *SuccessUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasUserOptions, "postDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasUserOptions, "postDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if postDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasUserOptions.XDeploymentID))
	}

	body := make(map[string]interface{})
	if postDb2SaasUserOptions.ID != nil {
		body["id"] = postDb2SaasUserOptions.ID
	}
	if postDb2SaasUserOptions.Iam != nil {
		body["iam"] = postDb2SaasUserOptions.Iam
	}
	if postDb2SaasUserOptions.Ibmid != nil {
		body["ibmid"] = postDb2SaasUserOptions.Ibmid
	}
	if postDb2SaasUserOptions.Name != nil {
		body["name"] = postDb2SaasUserOptions.Name
	}
	if postDb2SaasUserOptions.Password != nil {
		body["password"] = postDb2SaasUserOptions.Password
	}
	if postDb2SaasUserOptions.Role != nil {
		body["role"] = postDb2SaasUserOptions.Role
	}
	if postDb2SaasUserOptions.Email != nil {
		body["email"] = postDb2SaasUserOptions.Email
	}
	if postDb2SaasUserOptions.Locked != nil {
		body["locked"] = postDb2SaasUserOptions.Locked
	}
	if postDb2SaasUserOptions.Authentication != nil {
		body["authentication"] = postDb2SaasUserOptions.Authentication
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessUserResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasUser : Get the list of Users
func (db2saas *Db2saasV1) GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions) (result *SuccessGetUserInfo, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasUserWithContext(context.Background(), getDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasUserWithContext is an alternate form of the GetDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasUserWithContext(ctx context.Context, getDb2SaasUserOptions *GetDb2SaasUserOptions) (result *SuccessGetUserInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasUserOptions, "getDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasUserOptions, "getDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetUserInfo)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDb2SaasUser : Delete a user (only platform admin)
func (db2saas *Db2saasV1) DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions) (result map[string]interface{}, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.DeleteDb2SaasUserWithContext(context.Background(), deleteDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDb2SaasUserWithContext is an alternate form of the DeleteDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) DeleteDb2SaasUserWithContext(ctx context.Context, deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions) (result map[string]interface{}, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDb2SaasUserOptions, "deleteDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDb2SaasUserOptions, "deleteDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteDb2SaasUserOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "DeleteDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*deleteDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = db2saas.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetbyidDb2SaasUser : Get specific user by Id
func (db2saas *Db2saasV1) GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) (result *SuccessGetUserByID, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetbyidDb2SaasUserWithContext(context.Background(), getbyidDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetbyidDb2SaasUserWithContext is an alternate form of the GetbyidDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) GetbyidDb2SaasUserWithContext(ctx context.Context, getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) (result *SuccessGetUserByID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getbyidDb2SaasUserOptions, "getbyidDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getbyidDb2SaasUserOptions, "getbyidDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users/bluadmin`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getbyidDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetbyidDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getbyidDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getbyidDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getbyid_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetUserByID)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PutDb2SaasAutoscale : Update auto scaling configuration
func (db2saas *Db2saasV1) PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) (result *SuccessUpdateAutoScale, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PutDb2SaasAutoscaleWithContext(context.Background(), putDb2SaasAutoscaleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PutDb2SaasAutoscaleWithContext is an alternate form of the PutDb2SaasAutoscale method which supports a Context parameter
func (db2saas *Db2saasV1) PutDb2SaasAutoscaleWithContext(ctx context.Context, putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) (result *SuccessUpdateAutoScale, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putDb2SaasAutoscaleOptions, "putDb2SaasAutoscaleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(putDb2SaasAutoscaleOptions, "putDb2SaasAutoscaleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/scaling/auto`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range putDb2SaasAutoscaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PutDb2SaasAutoscale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if putDb2SaasAutoscaleOptions.XDbProfile != nil {
		builder.AddHeader("x-db-profile", fmt.Sprint(*putDb2SaasAutoscaleOptions.XDbProfile))
	}

	body := make(map[string]interface{})
	if putDb2SaasAutoscaleOptions.AutoScalingEnabled != nil {
		body["auto_scaling_enabled"] = putDb2SaasAutoscaleOptions.AutoScalingEnabled
	}
	if putDb2SaasAutoscaleOptions.AutoScalingThreshold != nil {
		body["auto_scaling_threshold"] = putDb2SaasAutoscaleOptions.AutoScalingThreshold
	}
	if putDb2SaasAutoscaleOptions.AutoScalingOverTimePeriod != nil {
		body["auto_scaling_over_time_period"] = putDb2SaasAutoscaleOptions.AutoScalingOverTimePeriod
	}
	if putDb2SaasAutoscaleOptions.AutoScalingPauseLimit != nil {
		body["auto_scaling_pause_limit"] = putDb2SaasAutoscaleOptions.AutoScalingPauseLimit
	}
	if putDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit != nil {
		body["auto_scaling_allow_plan_limit"] = putDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "put_db2_saas_autoscale", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessUpdateAutoScale)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasAutoscale : Get auto scaling info
func (db2saas *Db2saasV1) GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) (result *SuccessAutoScaling, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasAutoscaleWithContext(context.Background(), getDb2SaasAutoscaleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasAutoscaleWithContext is an alternate form of the GetDb2SaasAutoscale method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasAutoscaleWithContext(ctx context.Context, getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) (result *SuccessAutoScaling, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasAutoscaleOptions, "getDb2SaasAutoscaleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasAutoscaleOptions, "getDb2SaasAutoscaleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/scaling/auto`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasAutoscaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasAutoscale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasAutoscaleOptions.XDbProfile != nil {
		builder.AddHeader("x-db-profile", fmt.Sprint(*getDb2SaasAutoscaleOptions.XDbProfile))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_autoscale", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessAutoScaling)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PostDb2SaasDbConfiguration : Set database and database manager configuration
func (db2saas *Db2saasV1) PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions) (result *SuccessPostCustomSettings, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasDbConfigurationWithContext(context.Background(), postDb2SaasDbConfigurationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasDbConfigurationWithContext is an alternate form of the PostDb2SaasDbConfiguration method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasDbConfigurationWithContext(ctx context.Context, postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions) (result *SuccessPostCustomSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasDbConfigurationOptions, "postDb2SaasDbConfigurationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasDbConfigurationOptions, "postDb2SaasDbConfigurationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/deployments/custom_setting`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasDbConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasDbConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if postDb2SaasDbConfigurationOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasDbConfigurationOptions.XDeploymentID))
	}

	body := make(map[string]interface{})
	if postDb2SaasDbConfigurationOptions.Registry != nil {
		body["registry"] = postDb2SaasDbConfigurationOptions.Registry
	}
	if postDb2SaasDbConfigurationOptions.Db != nil {
		body["db"] = postDb2SaasDbConfigurationOptions.Db
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_db_configuration", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessPostCustomSettings)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasTuneableParam : Retrieves the values of tunable parameters of the DB2 instance
func (db2saas *Db2saasV1) GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions) (result *SuccessTuneableParams, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasTuneableParamWithContext(context.Background(), getDb2SaasTuneableParamOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasTuneableParamWithContext is an alternate form of the GetDb2SaasTuneableParam method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasTuneableParamWithContext(ctx context.Context, getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions) (result *SuccessTuneableParams, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasTuneableParamOptions, "getDb2SaasTuneableParamOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasTuneableParamOptions, "getDb2SaasTuneableParamOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/tuneable_param`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasTuneableParamOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasTuneableParam")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasTuneableParamOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasTuneableParamOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_tuneable_param", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessTuneableParams)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasBackup : Get Db2 instance backup information
func (db2saas *Db2saasV1) GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions) (result *SuccessGetBackups, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasBackupWithContext(context.Background(), getDb2SaasBackupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasBackupWithContext is an alternate form of the GetDb2SaasBackup method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasBackupWithContext(ctx context.Context, getDb2SaasBackupOptions *GetDb2SaasBackupOptions) (result *SuccessGetBackups, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasBackupOptions, "getDb2SaasBackupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasBackupOptions, "getDb2SaasBackupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/backups`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasBackupOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasBackupOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_backup", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetBackups)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PostDb2SaasBackup : Create backup of an instance
func (db2saas *Db2saasV1) PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions) (result *SuccessGetBackup, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasBackupWithContext(context.Background(), postDb2SaasBackupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasBackupWithContext is an alternate form of the PostDb2SaasBackup method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasBackupWithContext(ctx context.Context, postDb2SaasBackupOptions *PostDb2SaasBackupOptions) (result *SuccessGetBackup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasBackupOptions, "postDb2SaasBackupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasBackupOptions, "postDb2SaasBackupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/backups/backup`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if postDb2SaasBackupOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasBackupOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_backup", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetBackup)
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

// Backup : Info of backup.
type Backup struct {
	// CRN of the db2 instance.
	ID *string `json:"id" validate:"required"`

	// Defines the type of execution of backup.
	Type *string `json:"type" validate:"required"`

	// Status of the backup.
	Status *string `json:"status" validate:"required"`

	// Timestamp of the backup created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// Size of the backup or data set.
	Size *int64 `json:"size" validate:"required"`

	// The duration of the backup operation in seconds.
	Duration *int64 `json:"duration" validate:"required"`
}

// UnmarshalBackup unmarshals an instance of Backup from the specified map of raw messages.
func UnmarshalBackup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		err = core.SDKErrorf(err, "", "size-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		err = core.SDKErrorf(err, "", "duration-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCustomSettingsDb : Container for general database settings.
type CreateCustomSettingsDb struct {
	// Defines the amount of memory (in kb) that DB2 can use for sorting operations during query execution. It has in
	// format AUTOMATIC, range(16, 4294967295).
	SORTHEAP *string `json:"SORTHEAP,omitempty"`
}

// UnmarshalCreateCustomSettingsDb unmarshals an instance of CreateCustomSettingsDb from the specified map of raw messages.
func UnmarshalCreateCustomSettingsDb(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCustomSettingsDb)
	err = core.UnmarshalPrimitive(m, "SORTHEAP", &obj.SORTHEAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "SORTHEAP-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCustomSettingsRegistry : registry for db2 related configuration settings/configurations.
type CreateCustomSettingsRegistry struct {
	// Determines the selectivity of a query.
	DB2SELECTIVITY *string `json:"DB2_SELECTIVITY,omitempty"`

	// Configures how DB2 handles anti-joins.
	DB2ANTIJOIN *string `json:"DB2_ANTIJOIN,omitempty"`
}

// Constants associated with the CreateCustomSettingsRegistry.DB2SELECTIVITY property.
// Determines the selectivity of a query.
const (
	CreateCustomSettingsRegistry_DB2SELECTIVITY_All = "ALL"
	CreateCustomSettingsRegistry_DB2SELECTIVITY_No = "NO"
	CreateCustomSettingsRegistry_DB2SELECTIVITY_Yes = "YES"
)

// Constants associated with the CreateCustomSettingsRegistry.DB2ANTIJOIN property.
// Configures how DB2 handles anti-joins.
const (
	CreateCustomSettingsRegistry_DB2ANTIJOIN_Extend = "EXTEND"
	CreateCustomSettingsRegistry_DB2ANTIJOIN_No = "NO"
	CreateCustomSettingsRegistry_DB2ANTIJOIN_Yes = "YES"
)

// UnmarshalCreateCustomSettingsRegistry unmarshals an instance of CreateCustomSettingsRegistry from the specified map of raw messages.
func UnmarshalCreateCustomSettingsRegistry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCustomSettingsRegistry)
	err = core.UnmarshalPrimitive(m, "DB2_SELECTIVITY", &obj.DB2SELECTIVITY)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_SELECTIVITY-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_ANTIJOIN", &obj.DB2ANTIJOIN)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_ANTIJOIN-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateUserAuthentication : CreateUserAuthentication struct
type CreateUserAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Authentication policy ID.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// NewCreateUserAuthentication : Instantiate CreateUserAuthentication (Generic Model Constructor)
func (*Db2saasV1) NewCreateUserAuthentication(method string, policyID string) (_model *CreateUserAuthentication, err error) {
	_model = &CreateUserAuthentication{
		Method: core.StringPtr(method),
		PolicyID: core.StringPtr(policyID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalCreateUserAuthentication unmarshals an instance of CreateUserAuthentication from the specified map of raw messages.
func UnmarshalCreateUserAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateUserAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDb2SaasUserOptions : The DeleteDb2SaasUser options.
type DeleteDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// id of the user.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDb2SaasUserOptions : Instantiate DeleteDb2SaasUserOptions
func (*Db2saasV1) NewDeleteDb2SaasUserOptions(xDeploymentID string, id string) *DeleteDb2SaasUserOptions {
	return &DeleteDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		ID: core.StringPtr(id),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *DeleteDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *DeleteDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetID : Allow user to set ID
func (_options *DeleteDb2SaasUserOptions) SetID(id string) *DeleteDb2SaasUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDb2SaasUserOptions) SetHeaders(param map[string]string) *DeleteDb2SaasUserOptions {
	options.Headers = param
	return options
}

// GetDb2SaasAllowlistOptions : The GetDb2SaasAllowlist options.
type GetDb2SaasAllowlistOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasAllowlistOptions : Instantiate GetDb2SaasAllowlistOptions
func (*Db2saasV1) NewGetDb2SaasAllowlistOptions(xDeploymentID string) *GetDb2SaasAllowlistOptions {
	return &GetDb2SaasAllowlistOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasAllowlistOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasAllowlistOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasAllowlistOptions) SetHeaders(param map[string]string) *GetDb2SaasAllowlistOptions {
	options.Headers = param
	return options
}

// GetDb2SaasAutoscaleOptions : The GetDb2SaasAutoscale options.
type GetDb2SaasAutoscaleOptions struct {
	// Encoded CRN deployment id.
	XDbProfile *string `json:"x-db-profile" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasAutoscaleOptions : Instantiate GetDb2SaasAutoscaleOptions
func (*Db2saasV1) NewGetDb2SaasAutoscaleOptions(xDbProfile string) *GetDb2SaasAutoscaleOptions {
	return &GetDb2SaasAutoscaleOptions{
		XDbProfile: core.StringPtr(xDbProfile),
	}
}

// SetXDbProfile : Allow user to set XDbProfile
func (_options *GetDb2SaasAutoscaleOptions) SetXDbProfile(xDbProfile string) *GetDb2SaasAutoscaleOptions {
	_options.XDbProfile = core.StringPtr(xDbProfile)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasAutoscaleOptions) SetHeaders(param map[string]string) *GetDb2SaasAutoscaleOptions {
	options.Headers = param
	return options
}

// GetDb2SaasBackupOptions : The GetDb2SaasBackup options.
type GetDb2SaasBackupOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasBackupOptions : Instantiate GetDb2SaasBackupOptions
func (*Db2saasV1) NewGetDb2SaasBackupOptions(xDeploymentID string) *GetDb2SaasBackupOptions {
	return &GetDb2SaasBackupOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasBackupOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasBackupOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasBackupOptions) SetHeaders(param map[string]string) *GetDb2SaasBackupOptions {
	options.Headers = param
	return options
}

// GetDb2SaasConnectionInfoOptions : The GetDb2SaasConnectionInfo options.
type GetDb2SaasConnectionInfoOptions struct {
	// Encoded CRN deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasConnectionInfoOptions : Instantiate GetDb2SaasConnectionInfoOptions
func (*Db2saasV1) NewGetDb2SaasConnectionInfoOptions(deploymentID string, xDeploymentID string) *GetDb2SaasConnectionInfoOptions {
	return &GetDb2SaasConnectionInfoOptions{
		DeploymentID: core.StringPtr(deploymentID),
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *GetDb2SaasConnectionInfoOptions) SetDeploymentID(deploymentID string) *GetDb2SaasConnectionInfoOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasConnectionInfoOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasConnectionInfoOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasConnectionInfoOptions) SetHeaders(param map[string]string) *GetDb2SaasConnectionInfoOptions {
	options.Headers = param
	return options
}

// GetDb2SaasTuneableParamOptions : The GetDb2SaasTuneableParam options.
type GetDb2SaasTuneableParamOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasTuneableParamOptions : Instantiate GetDb2SaasTuneableParamOptions
func (*Db2saasV1) NewGetDb2SaasTuneableParamOptions(xDeploymentID string) *GetDb2SaasTuneableParamOptions {
	return &GetDb2SaasTuneableParamOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasTuneableParamOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasTuneableParamOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasTuneableParamOptions) SetHeaders(param map[string]string) *GetDb2SaasTuneableParamOptions {
	options.Headers = param
	return options
}

// GetDb2SaasUserOptions : The GetDb2SaasUser options.
type GetDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasUserOptions : Instantiate GetDb2SaasUserOptions
func (*Db2saasV1) NewGetDb2SaasUserOptions(xDeploymentID string) *GetDb2SaasUserOptions {
	return &GetDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasUserOptions) SetHeaders(param map[string]string) *GetDb2SaasUserOptions {
	options.Headers = param
	return options
}

// GetbyidDb2SaasUserOptions : The GetbyidDb2SaasUser options.
type GetbyidDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetbyidDb2SaasUserOptions : Instantiate GetbyidDb2SaasUserOptions
func (*Db2saasV1) NewGetbyidDb2SaasUserOptions(xDeploymentID string) *GetbyidDb2SaasUserOptions {
	return &GetbyidDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetbyidDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *GetbyidDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetbyidDb2SaasUserOptions) SetHeaders(param map[string]string) *GetbyidDb2SaasUserOptions {
	options.Headers = param
	return options
}

// IpAddress : Details of an IP address.
type IpAddress struct {
	// The IP address, in IPv4/ipv6 format.
	Address *string `json:"address" validate:"required"`

	// Description of the IP address.
	Description *string `json:"description" validate:"required"`
}

// NewIpAddress : Instantiate IpAddress (Generic Model Constructor)
func (*Db2saasV1) NewIpAddress(address string, description string) (_model *IpAddress, err error) {
	_model = &IpAddress{
		Address: core.StringPtr(address),
		Description: core.StringPtr(description),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalIpAddress unmarshals an instance of IpAddress from the specified map of raw messages.
func UnmarshalIpAddress(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpAddress)
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		err = core.SDKErrorf(err, "", "address-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostDb2SaasAllowlistOptions : The PostDb2SaasAllowlist options.
type PostDb2SaasAllowlistOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// List of IP addresses.
	IpAddresses []IpAddress `json:"ip_addresses" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPostDb2SaasAllowlistOptions : Instantiate PostDb2SaasAllowlistOptions
func (*Db2saasV1) NewPostDb2SaasAllowlistOptions(xDeploymentID string, ipAddresses []IpAddress) *PostDb2SaasAllowlistOptions {
	return &PostDb2SaasAllowlistOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		IpAddresses: ipAddresses,
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasAllowlistOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasAllowlistOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetIpAddresses : Allow user to set IpAddresses
func (_options *PostDb2SaasAllowlistOptions) SetIpAddresses(ipAddresses []IpAddress) *PostDb2SaasAllowlistOptions {
	_options.IpAddresses = ipAddresses
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasAllowlistOptions) SetHeaders(param map[string]string) *PostDb2SaasAllowlistOptions {
	options.Headers = param
	return options
}

// PostDb2SaasBackupOptions : The PostDb2SaasBackup options.
type PostDb2SaasBackupOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPostDb2SaasBackupOptions : Instantiate PostDb2SaasBackupOptions
func (*Db2saasV1) NewPostDb2SaasBackupOptions(xDeploymentID string) *PostDb2SaasBackupOptions {
	return &PostDb2SaasBackupOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasBackupOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasBackupOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasBackupOptions) SetHeaders(param map[string]string) *PostDb2SaasBackupOptions {
	options.Headers = param
	return options
}

// PostDb2SaasDbConfigurationOptions : The PostDb2SaasDbConfiguration options.
type PostDb2SaasDbConfigurationOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// registry for db2 related configuration settings/configurations.
	Registry *CreateCustomSettingsRegistry `json:"registry,omitempty"`

	// Container for general database settings.
	Db *CreateCustomSettingsDb `json:"db,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPostDb2SaasDbConfigurationOptions : Instantiate PostDb2SaasDbConfigurationOptions
func (*Db2saasV1) NewPostDb2SaasDbConfigurationOptions(xDeploymentID string) *PostDb2SaasDbConfigurationOptions {
	return &PostDb2SaasDbConfigurationOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasDbConfigurationOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasDbConfigurationOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetRegistry : Allow user to set Registry
func (_options *PostDb2SaasDbConfigurationOptions) SetRegistry(registry *CreateCustomSettingsRegistry) *PostDb2SaasDbConfigurationOptions {
	_options.Registry = registry
	return _options
}

// SetDb : Allow user to set Db
func (_options *PostDb2SaasDbConfigurationOptions) SetDb(db *CreateCustomSettingsDb) *PostDb2SaasDbConfigurationOptions {
	_options.Db = db
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasDbConfigurationOptions) SetHeaders(param map[string]string) *PostDb2SaasDbConfigurationOptions {
	options.Headers = param
	return options
}

// PostDb2SaasUserOptions : The PostDb2SaasUser options.
type PostDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// The id of the User.
	ID *string `json:"id" validate:"required"`

	// Indicates if IAM is enabled.
	Iam *bool `json:"iam" validate:"required"`

	// IBM ID of the User.
	Ibmid *string `json:"ibmid" validate:"required"`

	// The name of the User.
	Name *string `json:"name" validate:"required"`

	// Password of the User.
	Password *string `json:"password" validate:"required"`

	// Role of the User.
	Role *string `json:"role" validate:"required"`

	// Email of the User.
	Email *string `json:"email" validate:"required"`

	// Indicates if the account is locked.
	Locked *string `json:"locked" validate:"required"`

	Authentication *CreateUserAuthentication `json:"authentication" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the PostDb2SaasUserOptions.Role property.
// Role of the User.
const (
	PostDb2SaasUserOptions_Role_Bluadmin = "bluadmin"
	PostDb2SaasUserOptions_Role_Bluuser = "bluuser"
)

// Constants associated with the PostDb2SaasUserOptions.Locked property.
// Indicates if the account is locked.
const (
	PostDb2SaasUserOptions_Locked_No = "no"
	PostDb2SaasUserOptions_Locked_Yes = "yes"
)

// NewPostDb2SaasUserOptions : Instantiate PostDb2SaasUserOptions
func (*Db2saasV1) NewPostDb2SaasUserOptions(xDeploymentID string, id string, iam bool, ibmid string, name string, password string, role string, email string, locked string, authentication *CreateUserAuthentication) *PostDb2SaasUserOptions {
	return &PostDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		ID: core.StringPtr(id),
		Iam: core.BoolPtr(iam),
		Ibmid: core.StringPtr(ibmid),
		Name: core.StringPtr(name),
		Password: core.StringPtr(password),
		Role: core.StringPtr(role),
		Email: core.StringPtr(email),
		Locked: core.StringPtr(locked),
		Authentication: authentication,
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetID : Allow user to set ID
func (_options *PostDb2SaasUserOptions) SetID(id string) *PostDb2SaasUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIam : Allow user to set Iam
func (_options *PostDb2SaasUserOptions) SetIam(iam bool) *PostDb2SaasUserOptions {
	_options.Iam = core.BoolPtr(iam)
	return _options
}

// SetIbmid : Allow user to set Ibmid
func (_options *PostDb2SaasUserOptions) SetIbmid(ibmid string) *PostDb2SaasUserOptions {
	_options.Ibmid = core.StringPtr(ibmid)
	return _options
}

// SetName : Allow user to set Name
func (_options *PostDb2SaasUserOptions) SetName(name string) *PostDb2SaasUserOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPassword : Allow user to set Password
func (_options *PostDb2SaasUserOptions) SetPassword(password string) *PostDb2SaasUserOptions {
	_options.Password = core.StringPtr(password)
	return _options
}

// SetRole : Allow user to set Role
func (_options *PostDb2SaasUserOptions) SetRole(role string) *PostDb2SaasUserOptions {
	_options.Role = core.StringPtr(role)
	return _options
}

// SetEmail : Allow user to set Email
func (_options *PostDb2SaasUserOptions) SetEmail(email string) *PostDb2SaasUserOptions {
	_options.Email = core.StringPtr(email)
	return _options
}

// SetLocked : Allow user to set Locked
func (_options *PostDb2SaasUserOptions) SetLocked(locked string) *PostDb2SaasUserOptions {
	_options.Locked = core.StringPtr(locked)
	return _options
}

// SetAuthentication : Allow user to set Authentication
func (_options *PostDb2SaasUserOptions) SetAuthentication(authentication *CreateUserAuthentication) *PostDb2SaasUserOptions {
	_options.Authentication = authentication
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasUserOptions) SetHeaders(param map[string]string) *PostDb2SaasUserOptions {
	options.Headers = param
	return options
}

// PutDb2SaasAutoscaleOptions : The PutDb2SaasAutoscale options.
type PutDb2SaasAutoscaleOptions struct {
	// Encoded CRN deployment id.
	XDbProfile *string `json:"x-db-profile" validate:"required"`

	// Indicates if automatic scaling is enabled or not.
	AutoScalingEnabled *string `json:"auto_scaling_enabled,omitempty"`

	// Specifies the resource utilization level that triggers an auto-scaling.
	AutoScalingThreshold *int64 `json:"auto_scaling_threshold,omitempty"`

	// Defines the time period over which auto-scaling adjustments are monitored and applied.
	AutoScalingOverTimePeriod *float64 `json:"auto_scaling_over_time_period,omitempty"`

	// Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
	AutoScalingPauseLimit *int64 `json:"auto_scaling_pause_limit,omitempty"`

	// Indicates the maximum number of scaling actions that are allowed within a specified time period.
	AutoScalingAllowPlanLimit *string `json:"auto_scaling_allow_plan_limit,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the PutDb2SaasAutoscaleOptions.AutoScalingEnabled property.
// Indicates if automatic scaling is enabled or not.
const (
	PutDb2SaasAutoscaleOptions_AutoScalingEnabled_False = "false"
	PutDb2SaasAutoscaleOptions_AutoScalingEnabled_True = "true"
)

// Constants associated with the PutDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit property.
// Indicates the maximum number of scaling actions that are allowed within a specified time period.
const (
	PutDb2SaasAutoscaleOptions_AutoScalingAllowPlanLimit_No = "NO"
	PutDb2SaasAutoscaleOptions_AutoScalingAllowPlanLimit_Yes = "YES"
)

// NewPutDb2SaasAutoscaleOptions : Instantiate PutDb2SaasAutoscaleOptions
func (*Db2saasV1) NewPutDb2SaasAutoscaleOptions(xDbProfile string) *PutDb2SaasAutoscaleOptions {
	return &PutDb2SaasAutoscaleOptions{
		XDbProfile: core.StringPtr(xDbProfile),
	}
}

// SetXDbProfile : Allow user to set XDbProfile
func (_options *PutDb2SaasAutoscaleOptions) SetXDbProfile(xDbProfile string) *PutDb2SaasAutoscaleOptions {
	_options.XDbProfile = core.StringPtr(xDbProfile)
	return _options
}

// SetAutoScalingEnabled : Allow user to set AutoScalingEnabled
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingEnabled(autoScalingEnabled string) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingEnabled = core.StringPtr(autoScalingEnabled)
	return _options
}

// SetAutoScalingThreshold : Allow user to set AutoScalingThreshold
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingThreshold(autoScalingThreshold int64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingThreshold = core.Int64Ptr(autoScalingThreshold)
	return _options
}

// SetAutoScalingOverTimePeriod : Allow user to set AutoScalingOverTimePeriod
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingOverTimePeriod(autoScalingOverTimePeriod float64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingOverTimePeriod = core.Float64Ptr(autoScalingOverTimePeriod)
	return _options
}

// SetAutoScalingPauseLimit : Allow user to set AutoScalingPauseLimit
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingPauseLimit(autoScalingPauseLimit int64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingPauseLimit = core.Int64Ptr(autoScalingPauseLimit)
	return _options
}

// SetAutoScalingAllowPlanLimit : Allow user to set AutoScalingAllowPlanLimit
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingAllowPlanLimit(autoScalingAllowPlanLimit string) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingAllowPlanLimit = core.StringPtr(autoScalingAllowPlanLimit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PutDb2SaasAutoscaleOptions) SetHeaders(param map[string]string) *PutDb2SaasAutoscaleOptions {
	options.Headers = param
	return options
}

// SuccessAutoScaling : The details of the autoscale.
type SuccessAutoScaling struct {
	// Indicates the maximum number of scaling actions that are allowed within a specified time period.
	AutoScalingAllowPlanLimit *bool `json:"auto_scaling_allow_plan_limit" validate:"required"`

	// Indicates if automatic scaling is enabled or not.
	AutoScalingEnabled *bool `json:"auto_scaling_enabled" validate:"required"`

	// The maximum limit for automatically increasing storage capacity to handle growing data needs.
	AutoScalingMaxStorage *int64 `json:"auto_scaling_max_storage" validate:"required"`

	// Defines the time period over which auto-scaling adjustments are monitored and applied.
	AutoScalingOverTimePeriod *int64 `json:"auto_scaling_over_time_period" validate:"required"`

	// Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
	AutoScalingPauseLimit *int64 `json:"auto_scaling_pause_limit" validate:"required"`

	// Specifies the resource utilization level that triggers an auto-scaling.
	AutoScalingThreshold *int64 `json:"auto_scaling_threshold" validate:"required"`

	// Specifies the unit of measurement for storage capacity.
	StorageUnit *string `json:"storage_unit" validate:"required"`

	// Represents the percentage of total storage capacity currently in use.
	StorageUtilizationPercentage *int64 `json:"storage_utilization_percentage" validate:"required"`

	// Indicates whether a system or service can automatically adjust resources based on demand.
	SupportAutoScaling *bool `json:"support_auto_scaling" validate:"required"`
}

// UnmarshalSuccessAutoScaling unmarshals an instance of SuccessAutoScaling from the specified map of raw messages.
func UnmarshalSuccessAutoScaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessAutoScaling)
	err = core.UnmarshalPrimitive(m, "auto_scaling_allow_plan_limit", &obj.AutoScalingAllowPlanLimit)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_allow_plan_limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_enabled", &obj.AutoScalingEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_max_storage", &obj.AutoScalingMaxStorage)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_max_storage-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_over_time_period", &obj.AutoScalingOverTimePeriod)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_over_time_period-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_pause_limit", &obj.AutoScalingPauseLimit)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_pause_limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_threshold", &obj.AutoScalingThreshold)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_threshold-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_unit", &obj.StorageUnit)
	if err != nil {
		err = core.SDKErrorf(err, "", "storage_unit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_utilization_percentage", &obj.StorageUtilizationPercentage)
	if err != nil {
		err = core.SDKErrorf(err, "", "storage_utilization_percentage-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "support_auto_scaling", &obj.SupportAutoScaling)
	if err != nil {
		err = core.SDKErrorf(err, "", "support_auto_scaling-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

	SslPort *string `json:"sslPort,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	PrivateServiceName *string `json:"private_serviceName,omitempty"`

	CloudServiceOffering *string `json:"cloud_service_offering,omitempty"`

	VpeServiceCrn *string `json:"vpe_service_crn,omitempty"`

	DbVpcEndpointService *string `json:"db_vpc_endpoint_service,omitempty"`
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
	err = core.UnmarshalPrimitive(m, "private_serviceName", &obj.PrivateServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "private_serviceName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cloud_service_offering", &obj.CloudServiceOffering)
	if err != nil {
		err = core.SDKErrorf(err, "", "cloud_service_offering-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "vpe_service_crn", &obj.VpeServiceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "vpe_service_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "db_vpc_endpoint_service", &obj.DbVpcEndpointService)
	if err != nil {
		err = core.SDKErrorf(err, "", "db_vpc_endpoint_service-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfoPublic : SuccessConnectionInfoPublic struct
type SuccessConnectionInfoPublic struct {
	Hostname *string `json:"hostname,omitempty"`

	DatabaseName *string `json:"databaseName,omitempty"`

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

// SuccessGetAllowlistIPs : Success response of get allowlist IPs.
type SuccessGetAllowlistIPs struct {
	// List of IP addresses.
	IpAddresses []IpAddress `json:"ip_addresses" validate:"required"`
}

// UnmarshalSuccessGetAllowlistIPs unmarshals an instance of SuccessGetAllowlistIPs from the specified map of raw messages.
func UnmarshalSuccessGetAllowlistIPs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetAllowlistIPs)
	err = core.UnmarshalModel(m, "ip_addresses", &obj.IpAddresses, UnmarshalIpAddress)
	if err != nil {
		err = core.SDKErrorf(err, "", "ip_addresses-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetBackup : Success response of post backup.
type SuccessGetBackup struct {
	Task *SuccessGetBackupTask `json:"task" validate:"required"`
}

// UnmarshalSuccessGetBackup unmarshals an instance of SuccessGetBackup from the specified map of raw messages.
func UnmarshalSuccessGetBackup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetBackup)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalSuccessGetBackupTask)
	if err != nil {
		err = core.SDKErrorf(err, "", "task-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetBackupTask : SuccessGetBackupTask struct
type SuccessGetBackupTask struct {
	// CRN of the instance.
	ID *string `json:"id,omitempty"`
}

// UnmarshalSuccessGetBackupTask unmarshals an instance of SuccessGetBackupTask from the specified map of raw messages.
func UnmarshalSuccessGetBackupTask(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetBackupTask)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetBackups : The details of the backups.
type SuccessGetBackups struct {
	Backups []Backup `json:"backups" validate:"required"`
}

// UnmarshalSuccessGetBackups unmarshals an instance of SuccessGetBackups from the specified map of raw messages.
func UnmarshalSuccessGetBackups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetBackups)
	err = core.UnmarshalModel(m, "backups", &obj.Backups, UnmarshalBackup)
	if err != nil {
		err = core.SDKErrorf(err, "", "backups-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserByID : The details of the users.
type SuccessGetUserByID struct {
	// User's DV role.
	DvRole *string `json:"dvRole" validate:"required"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata" validate:"required"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid" validate:"required"`

	// Role assigned to the user.
	Role *string `json:"role" validate:"required"`

	// IAM ID for the user.
	Iamid *string `json:"iamid" validate:"required"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions" validate:"required"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean" validate:"required"`

	// User's password.
	Password *string `json:"password" validate:"required"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam" validate:"required"`

	// The display name of the user.
	Name *string `json:"name" validate:"required"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid" validate:"required"`

	// Unique identifier for the user.
	ID *string `json:"id" validate:"required"`

	// Account lock status for the user.
	Locked *string `json:"locked" validate:"required"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg" validate:"required"`

	// Email address of the user.
	Email *string `json:"email" validate:"required"`

	// Authentication details for the user.
	Authentication *SuccessGetUserByIDAuthentication `json:"authentication" validate:"required"`
}

// Constants associated with the SuccessGetUserByID.Role property.
// Role assigned to the user.
const (
	SuccessGetUserByID_Role_Bluadmin = "bluadmin"
	SuccessGetUserByID_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessGetUserByID.Locked property.
// Account lock status for the user.
const (
	SuccessGetUserByID_Locked_No = "no"
	SuccessGetUserByID_Locked_Yes = "yes"
)

// UnmarshalSuccessGetUserByID unmarshals an instance of SuccessGetUserByID from the specified map of raw messages.
func UnmarshalSuccessGetUserByID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserByID)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessGetUserByIDAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserByIDAuthentication : Authentication details for the user.
type SuccessGetUserByIDAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessGetUserByIDAuthentication unmarshals an instance of SuccessGetUserByIDAuthentication from the specified map of raw messages.
func UnmarshalSuccessGetUserByIDAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserByIDAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfo : Success response of get user.
type SuccessGetUserInfo struct {
	// The total number of resources.
	Count *int64 `json:"count" validate:"required"`

	// A list of user resource.
	Resources []SuccessGetUserInfoResourcesItem `json:"resources" validate:"required"`
}

// UnmarshalSuccessGetUserInfo unmarshals an instance of SuccessGetUserInfo from the specified map of raw messages.
func UnmarshalSuccessGetUserInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfo)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		err = core.SDKErrorf(err, "", "count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalSuccessGetUserInfoResourcesItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfoResourcesItem : SuccessGetUserInfoResourcesItem struct
type SuccessGetUserInfoResourcesItem struct {
	// User's DV role.
	DvRole *string `json:"dvRole,omitempty"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid,omitempty"`

	// Role assigned to the user.
	Role *string `json:"role,omitempty"`

	// IAM ID for the user.
	Iamid *string `json:"iamid,omitempty"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions,omitempty"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean,omitempty"`

	// User's password.
	Password *string `json:"password,omitempty"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam,omitempty"`

	// The display name of the user.
	Name *string `json:"name,omitempty"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid,omitempty"`

	// Unique identifier for the user.
	ID *string `json:"id,omitempty"`

	// Account lock status for the user.
	Locked *string `json:"locked,omitempty"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg,omitempty"`

	// Email address of the user.
	Email *string `json:"email,omitempty"`

	// Authentication details for the user.
	Authentication *SuccessGetUserInfoResourcesItemAuthentication `json:"authentication,omitempty"`
}

// Constants associated with the SuccessGetUserInfoResourcesItem.Role property.
// Role assigned to the user.
const (
	SuccessGetUserInfoResourcesItem_Role_Bluadmin = "bluadmin"
	SuccessGetUserInfoResourcesItem_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessGetUserInfoResourcesItem.Locked property.
// Account lock status for the user.
const (
	SuccessGetUserInfoResourcesItem_Locked_No = "no"
	SuccessGetUserInfoResourcesItem_Locked_Yes = "yes"
)

// UnmarshalSuccessGetUserInfoResourcesItem unmarshals an instance of SuccessGetUserInfoResourcesItem from the specified map of raw messages.
func UnmarshalSuccessGetUserInfoResourcesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfoResourcesItem)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessGetUserInfoResourcesItemAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfoResourcesItemAuthentication : Authentication details for the user.
type SuccessGetUserInfoResourcesItemAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessGetUserInfoResourcesItemAuthentication unmarshals an instance of SuccessGetUserInfoResourcesItemAuthentication from the specified map of raw messages.
func UnmarshalSuccessGetUserInfoResourcesItemAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfoResourcesItemAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessPostAllowedlistIPs : Success response of post allowlist IPs.
type SuccessPostAllowedlistIPs struct {
	// status of the post allowlist IPs request.
	Status *string `json:"status" validate:"required"`
}

// UnmarshalSuccessPostAllowedlistIPs unmarshals an instance of SuccessPostAllowedlistIPs from the specified map of raw messages.
func UnmarshalSuccessPostAllowedlistIPs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessPostAllowedlistIPs)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessPostCustomSettings : The details of created custom settings of db2.
type SuccessPostCustomSettings struct {
	// Describes the operation done.
	Description *string `json:"description" validate:"required"`

	// CRN of the db2 instance.
	ID *string `json:"id" validate:"required"`

	// Defines the status of the instance.
	Status *string `json:"status" validate:"required"`
}

// UnmarshalSuccessPostCustomSettings unmarshals an instance of SuccessPostCustomSettings from the specified map of raw messages.
func UnmarshalSuccessPostCustomSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessPostCustomSettings)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessTuneableParams : Response of tuneable params of the Db2 instance.
type SuccessTuneableParams struct {
	TuneableParam *SuccessTuneableParamsTuneableParam `json:"tuneable_param,omitempty"`
}

// UnmarshalSuccessTuneableParams unmarshals an instance of SuccessTuneableParams from the specified map of raw messages.
func UnmarshalSuccessTuneableParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessTuneableParams)
	err = core.UnmarshalModel(m, "tuneable_param", &obj.TuneableParam, UnmarshalSuccessTuneableParamsTuneableParam)
	if err != nil {
		err = core.SDKErrorf(err, "", "tuneable_param-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessTuneableParamsTuneableParam : SuccessTuneableParamsTuneableParam struct
type SuccessTuneableParamsTuneableParam struct {
	// Tunable parameters related to the Db2 database instance.
	Db *SuccessTuneableParamsTuneableParamDb `json:"db,omitempty"`

	// Tunable parameters related to the Db2 instance manager (dbm).
	Dbm *SuccessTuneableParamsTuneableParamDbm `json:"dbm,omitempty"`

	// Tunable parameters related to the Db2 registry.
	Registry *SuccessTuneableParamsTuneableParamRegistry `json:"registry,omitempty"`
}

// UnmarshalSuccessTuneableParamsTuneableParam unmarshals an instance of SuccessTuneableParamsTuneableParam from the specified map of raw messages.
func UnmarshalSuccessTuneableParamsTuneableParam(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessTuneableParamsTuneableParam)
	err = core.UnmarshalModel(m, "db", &obj.Db, UnmarshalSuccessTuneableParamsTuneableParamDb)
	if err != nil {
		err = core.SDKErrorf(err, "", "db-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "dbm", &obj.Dbm, UnmarshalSuccessTuneableParamsTuneableParamDbm)
	if err != nil {
		err = core.SDKErrorf(err, "", "dbm-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "registry", &obj.Registry, UnmarshalSuccessTuneableParamsTuneableParamRegistry)
	if err != nil {
		err = core.SDKErrorf(err, "", "registry-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessTuneableParamsTuneableParamDb : Tunable parameters related to the Db2 database instance.
type SuccessTuneableParamsTuneableParamDb struct {
	ACTSORTMEMLIMIT *string `json:"ACT_SORTMEM_LIMIT,omitempty"`

	ALTCOLLATE *string `json:"ALT_COLLATE,omitempty"`

	APPGROUPMEMSZ *string `json:"APPGROUP_MEM_SZ,omitempty"`

	APPLHEAPSZ *string `json:"APPLHEAPSZ,omitempty"`

	APPLMEMORY *string `json:"APPL_MEMORY,omitempty"`

	APPCTLHEAPSZ *string `json:"APP_CTL_HEAP_SZ,omitempty"`

	ARCHRETRYDELAY *string `json:"ARCHRETRYDELAY,omitempty"`

	AUTHNCACHEDURATION *string `json:"AUTHN_CACHE_DURATION,omitempty"`

	AUTORESTART *string `json:"AUTORESTART,omitempty"`

	AUTOCGSTATS *string `json:"AUTO_CG_STATS,omitempty"`

	AUTOMAINT *string `json:"AUTO_MAINT,omitempty"`

	AUTOREORG *string `json:"AUTO_REORG,omitempty"`

	AUTOREVAL *string `json:"AUTO_REVAL,omitempty"`

	AUTORUNSTATS *string `json:"AUTO_RUNSTATS,omitempty"`

	AUTOSAMPLING *string `json:"AUTO_SAMPLING,omitempty"`

	AUTOSTATSVIEWS *string `json:"AUTO_STATS_VIEWS,omitempty"`

	AUTOSTMTSTATS *string `json:"AUTO_STMT_STATS,omitempty"`

	AUTOTBLMAINT *string `json:"AUTO_TBL_MAINT,omitempty"`

	AVGAPPLS *string `json:"AVG_APPLS,omitempty"`

	CATALOGCACHESZ *string `json:"CATALOGCACHE_SZ,omitempty"`

	CHNGPGSTHRESH *string `json:"CHNGPGS_THRESH,omitempty"`

	CURCOMMIT *string `json:"CUR_COMMIT,omitempty"`

	DATABASEMEMORY *string `json:"DATABASE_MEMORY,omitempty"`

	DBHEAP *string `json:"DBHEAP,omitempty"`

	DBCOLLNAME *string `json:"DB_COLLNAME,omitempty"`

	DBMEMTHRESH *string `json:"DB_MEM_THRESH,omitempty"`

	DDLCOMPRESSIONDEF *string `json:"DDL_COMPRESSION_DEF,omitempty"`

	DDLCONSTRAINTDEF *string `json:"DDL_CONSTRAINT_DEF,omitempty"`

	DECFLTROUNDING *string `json:"DECFLT_ROUNDING,omitempty"`

	DECARITHMETIC *string `json:"DEC_ARITHMETIC,omitempty"`

	DECTOCHARFMT *string `json:"DEC_TO_CHAR_FMT,omitempty"`

	DFTDEGREE *string `json:"DFT_DEGREE,omitempty"`

	DFTEXTENTSZ *string `json:"DFT_EXTENT_SZ,omitempty"`

	DFTLOADRECSES *string `json:"DFT_LOADREC_SES,omitempty"`

	DFTMTTBTYPES *string `json:"DFT_MTTB_TYPES,omitempty"`

	DFTPREFETCHSZ *string `json:"DFT_PREFETCH_SZ,omitempty"`

	DFTQUERYOPT *string `json:"DFT_QUERYOPT,omitempty"`

	DFTREFRESHAGE *string `json:"DFT_REFRESH_AGE,omitempty"`

	DFTSCHEMASDCC *string `json:"DFT_SCHEMAS_DCC,omitempty"`

	DFTSQLMATHWARN *string `json:"DFT_SQLMATHWARN,omitempty"`

	DFTTABLEORG *string `json:"DFT_TABLE_ORG,omitempty"`

	DLCHKTIME *string `json:"DLCHKTIME,omitempty"`

	ENABLEXMLCHAR *string `json:"ENABLE_XMLCHAR,omitempty"`

	EXTENDEDROWSZ *string `json:"EXTENDED_ROW_SZ,omitempty"`

	GROUPHEAPRATIO *string `json:"GROUPHEAP_RATIO,omitempty"`

	INDEXREC *string `json:"INDEXREC,omitempty"`

	LARGEAGGREGATION *string `json:"LARGE_AGGREGATION,omitempty"`

	LOCKLIST *string `json:"LOCKLIST,omitempty"`

	LOCKTIMEOUT *string `json:"LOCKTIMEOUT,omitempty"`

	LOGINDEXBUILD *string `json:"LOGINDEXBUILD,omitempty"`

	LOGAPPLINFO *string `json:"LOG_APPL_INFO,omitempty"`

	LOGDDLSTMTS *string `json:"LOG_DDL_STMTS,omitempty"`

	LOGDISKCAP *string `json:"LOG_DISK_CAP,omitempty"`

	MAXAPPLS *string `json:"MAXAPPLS,omitempty"`

	MAXFILOP *string `json:"MAXFILOP,omitempty"`

	MAXLOCKS *string `json:"MAXLOCKS,omitempty"`

	MINDECDIV3 *string `json:"MIN_DEC_DIV_3,omitempty"`

	MONACTMETRICS *string `json:"MON_ACT_METRICS,omitempty"`

	MONDEADLOCK *string `json:"MON_DEADLOCK,omitempty"`

	MONLCKMSGLVL *string `json:"MON_LCK_MSG_LVL,omitempty"`

	MONLOCKTIMEOUT *string `json:"MON_LOCKTIMEOUT,omitempty"`

	MONLOCKWAIT *string `json:"MON_LOCKWAIT,omitempty"`

	MONLWTHRESH *string `json:"MON_LW_THRESH,omitempty"`

	MONOBJMETRICS *string `json:"MON_OBJ_METRICS,omitempty"`

	MONPKGLISTSZ *string `json:"MON_PKGLIST_SZ,omitempty"`

	MONREQMETRICS *string `json:"MON_REQ_METRICS,omitempty"`

	MONRTNDATA *string `json:"MON_RTN_DATA,omitempty"`

	MONRTNEXECLIST *string `json:"MON_RTN_EXECLIST,omitempty"`

	MONUOWDATA *string `json:"MON_UOW_DATA,omitempty"`

	MONUOWEXECLIST *string `json:"MON_UOW_EXECLIST,omitempty"`

	MONUOWPKGLIST *string `json:"MON_UOW_PKGLIST,omitempty"`

	NCHARMAPPING *string `json:"NCHAR_MAPPING,omitempty"`

	NUMFREQVALUES *string `json:"NUM_FREQVALUES,omitempty"`

	NUMIOCLEANERS *string `json:"NUM_IOCLEANERS,omitempty"`

	NUMIOSERVERS *string `json:"NUM_IOSERVERS,omitempty"`

	NUMLOGSPAN *string `json:"NUM_LOG_SPAN,omitempty"`

	NUMQUANTILES *string `json:"NUM_QUANTILES,omitempty"`

	OPTBUFFPAGE *string `json:"OPT_BUFFPAGE,omitempty"`

	OPTDIRECTWRKLD *string `json:"OPT_DIRECT_WRKLD,omitempty"`

	OPTLOCKLIST *string `json:"OPT_LOCKLIST,omitempty"`

	OPTMAXLOCKS *string `json:"OPT_MAXLOCKS,omitempty"`

	OPTSORTHEAP *string `json:"OPT_SORTHEAP,omitempty"`

	PAGEAGETRGTGCR *string `json:"PAGE_AGE_TRGT_GCR,omitempty"`

	PAGEAGETRGTMCR *string `json:"PAGE_AGE_TRGT_MCR,omitempty"`

	PCKCACHESZ *string `json:"PCKCACHESZ,omitempty"`

	PLSTACKTRACE *string `json:"PL_STACK_TRACE,omitempty"`

	SELFTUNINGMEM *string `json:"SELF_TUNING_MEM,omitempty"`

	SEQDETECT *string `json:"SEQDETECT,omitempty"`

	SHEAPTHRESSHR *string `json:"SHEAPTHRES_SHR,omitempty"`

	SOFTMAX *string `json:"SOFTMAX,omitempty"`

	SORTHEAP *string `json:"SORTHEAP,omitempty"`

	SQLCCFLAGS *string `json:"SQL_CCFLAGS,omitempty"`

	STATHEAPSZ *string `json:"STAT_HEAP_SZ,omitempty"`

	STMTHEAP *string `json:"STMTHEAP,omitempty"`

	STMTCONC *string `json:"STMT_CONC,omitempty"`

	STRINGUNITS *string `json:"STRING_UNITS,omitempty"`

	SYSTIMEPERIODADJ *string `json:"SYSTIME_PERIOD_ADJ,omitempty"`

	TRACKMOD *string `json:"TRACKMOD,omitempty"`

	UTILHEAPSZ *string `json:"UTIL_HEAP_SZ,omitempty"`

	WLMADMISSIONCTRL *string `json:"WLM_ADMISSION_CTRL,omitempty"`

	WLMAGENTLOADTRGT *string `json:"WLM_AGENT_LOAD_TRGT,omitempty"`

	WLMCPULIMIT *string `json:"WLM_CPU_LIMIT,omitempty"`

	WLMCPUSHARES *string `json:"WLM_CPU_SHARES,omitempty"`

	WLMCPUSHAREMODE *string `json:"WLM_CPU_SHARE_MODE,omitempty"`
}

// UnmarshalSuccessTuneableParamsTuneableParamDb unmarshals an instance of SuccessTuneableParamsTuneableParamDb from the specified map of raw messages.
func UnmarshalSuccessTuneableParamsTuneableParamDb(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessTuneableParamsTuneableParamDb)
	err = core.UnmarshalPrimitive(m, "ACT_SORTMEM_LIMIT", &obj.ACTSORTMEMLIMIT)
	if err != nil {
		err = core.SDKErrorf(err, "", "ACT_SORTMEM_LIMIT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ALT_COLLATE", &obj.ALTCOLLATE)
	if err != nil {
		err = core.SDKErrorf(err, "", "ALT_COLLATE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "APPGROUP_MEM_SZ", &obj.APPGROUPMEMSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "APPGROUP_MEM_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "APPLHEAPSZ", &obj.APPLHEAPSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "APPLHEAPSZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "APPL_MEMORY", &obj.APPLMEMORY)
	if err != nil {
		err = core.SDKErrorf(err, "", "APPL_MEMORY-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "APP_CTL_HEAP_SZ", &obj.APPCTLHEAPSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "APP_CTL_HEAP_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ARCHRETRYDELAY", &obj.ARCHRETRYDELAY)
	if err != nil {
		err = core.SDKErrorf(err, "", "ARCHRETRYDELAY-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTHN_CACHE_DURATION", &obj.AUTHNCACHEDURATION)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTHN_CACHE_DURATION-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTORESTART", &obj.AUTORESTART)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTORESTART-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_CG_STATS", &obj.AUTOCGSTATS)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_CG_STATS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_MAINT", &obj.AUTOMAINT)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_MAINT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_REORG", &obj.AUTOREORG)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_REORG-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_REVAL", &obj.AUTOREVAL)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_REVAL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_RUNSTATS", &obj.AUTORUNSTATS)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_RUNSTATS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_SAMPLING", &obj.AUTOSAMPLING)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_SAMPLING-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_STATS_VIEWS", &obj.AUTOSTATSVIEWS)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_STATS_VIEWS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_STMT_STATS", &obj.AUTOSTMTSTATS)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_STMT_STATS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTO_TBL_MAINT", &obj.AUTOTBLMAINT)
	if err != nil {
		err = core.SDKErrorf(err, "", "AUTO_TBL_MAINT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "AVG_APPLS", &obj.AVGAPPLS)
	if err != nil {
		err = core.SDKErrorf(err, "", "AVG_APPLS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "CATALOGCACHE_SZ", &obj.CATALOGCACHESZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "CATALOGCACHE_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "CHNGPGS_THRESH", &obj.CHNGPGSTHRESH)
	if err != nil {
		err = core.SDKErrorf(err, "", "CHNGPGS_THRESH-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "CUR_COMMIT", &obj.CURCOMMIT)
	if err != nil {
		err = core.SDKErrorf(err, "", "CUR_COMMIT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DATABASE_MEMORY", &obj.DATABASEMEMORY)
	if err != nil {
		err = core.SDKErrorf(err, "", "DATABASE_MEMORY-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DBHEAP", &obj.DBHEAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "DBHEAP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB_COLLNAME", &obj.DBCOLLNAME)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB_COLLNAME-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB_MEM_THRESH", &obj.DBMEMTHRESH)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB_MEM_THRESH-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DDL_COMPRESSION_DEF", &obj.DDLCOMPRESSIONDEF)
	if err != nil {
		err = core.SDKErrorf(err, "", "DDL_COMPRESSION_DEF-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DDL_CONSTRAINT_DEF", &obj.DDLCONSTRAINTDEF)
	if err != nil {
		err = core.SDKErrorf(err, "", "DDL_CONSTRAINT_DEF-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DECFLT_ROUNDING", &obj.DECFLTROUNDING)
	if err != nil {
		err = core.SDKErrorf(err, "", "DECFLT_ROUNDING-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DEC_ARITHMETIC", &obj.DECARITHMETIC)
	if err != nil {
		err = core.SDKErrorf(err, "", "DEC_ARITHMETIC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DEC_TO_CHAR_FMT", &obj.DECTOCHARFMT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DEC_TO_CHAR_FMT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_DEGREE", &obj.DFTDEGREE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_DEGREE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_EXTENT_SZ", &obj.DFTEXTENTSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_EXTENT_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_LOADREC_SES", &obj.DFTLOADRECSES)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_LOADREC_SES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MTTB_TYPES", &obj.DFTMTTBTYPES)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MTTB_TYPES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_PREFETCH_SZ", &obj.DFTPREFETCHSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_PREFETCH_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_QUERYOPT", &obj.DFTQUERYOPT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_QUERYOPT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_REFRESH_AGE", &obj.DFTREFRESHAGE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_REFRESH_AGE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_SCHEMAS_DCC", &obj.DFTSCHEMASDCC)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_SCHEMAS_DCC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_SQLMATHWARN", &obj.DFTSQLMATHWARN)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_SQLMATHWARN-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_TABLE_ORG", &obj.DFTTABLEORG)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_TABLE_ORG-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DLCHKTIME", &obj.DLCHKTIME)
	if err != nil {
		err = core.SDKErrorf(err, "", "DLCHKTIME-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ENABLE_XMLCHAR", &obj.ENABLEXMLCHAR)
	if err != nil {
		err = core.SDKErrorf(err, "", "ENABLE_XMLCHAR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "EXTENDED_ROW_SZ", &obj.EXTENDEDROWSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "EXTENDED_ROW_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "GROUPHEAP_RATIO", &obj.GROUPHEAPRATIO)
	if err != nil {
		err = core.SDKErrorf(err, "", "GROUPHEAP_RATIO-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "INDEXREC", &obj.INDEXREC)
	if err != nil {
		err = core.SDKErrorf(err, "", "INDEXREC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LARGE_AGGREGATION", &obj.LARGEAGGREGATION)
	if err != nil {
		err = core.SDKErrorf(err, "", "LARGE_AGGREGATION-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOCKLIST", &obj.LOCKLIST)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOCKLIST-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOCKTIMEOUT", &obj.LOCKTIMEOUT)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOCKTIMEOUT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOGINDEXBUILD", &obj.LOGINDEXBUILD)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOGINDEXBUILD-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOG_APPL_INFO", &obj.LOGAPPLINFO)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOG_APPL_INFO-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOG_DDL_STMTS", &obj.LOGDDLSTMTS)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOG_DDL_STMTS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "LOG_DISK_CAP", &obj.LOGDISKCAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "LOG_DISK_CAP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MAXAPPLS", &obj.MAXAPPLS)
	if err != nil {
		err = core.SDKErrorf(err, "", "MAXAPPLS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MAXFILOP", &obj.MAXFILOP)
	if err != nil {
		err = core.SDKErrorf(err, "", "MAXFILOP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MAXLOCKS", &obj.MAXLOCKS)
	if err != nil {
		err = core.SDKErrorf(err, "", "MAXLOCKS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MIN_DEC_DIV_3", &obj.MINDECDIV3)
	if err != nil {
		err = core.SDKErrorf(err, "", "MIN_DEC_DIV_3-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_ACT_METRICS", &obj.MONACTMETRICS)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_ACT_METRICS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_DEADLOCK", &obj.MONDEADLOCK)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_DEADLOCK-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_LCK_MSG_LVL", &obj.MONLCKMSGLVL)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_LCK_MSG_LVL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_LOCKTIMEOUT", &obj.MONLOCKTIMEOUT)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_LOCKTIMEOUT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_LOCKWAIT", &obj.MONLOCKWAIT)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_LOCKWAIT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_LW_THRESH", &obj.MONLWTHRESH)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_LW_THRESH-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_OBJ_METRICS", &obj.MONOBJMETRICS)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_OBJ_METRICS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_PKGLIST_SZ", &obj.MONPKGLISTSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_PKGLIST_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_REQ_METRICS", &obj.MONREQMETRICS)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_REQ_METRICS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_RTN_DATA", &obj.MONRTNDATA)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_RTN_DATA-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_RTN_EXECLIST", &obj.MONRTNEXECLIST)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_RTN_EXECLIST-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_UOW_DATA", &obj.MONUOWDATA)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_UOW_DATA-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_UOW_EXECLIST", &obj.MONUOWEXECLIST)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_UOW_EXECLIST-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_UOW_PKGLIST", &obj.MONUOWPKGLIST)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_UOW_PKGLIST-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NCHAR_MAPPING", &obj.NCHARMAPPING)
	if err != nil {
		err = core.SDKErrorf(err, "", "NCHAR_MAPPING-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_FREQVALUES", &obj.NUMFREQVALUES)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_FREQVALUES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_IOCLEANERS", &obj.NUMIOCLEANERS)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_IOCLEANERS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_IOSERVERS", &obj.NUMIOSERVERS)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_IOSERVERS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_LOG_SPAN", &obj.NUMLOGSPAN)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_LOG_SPAN-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_QUANTILES", &obj.NUMQUANTILES)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_QUANTILES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "OPT_BUFFPAGE", &obj.OPTBUFFPAGE)
	if err != nil {
		err = core.SDKErrorf(err, "", "OPT_BUFFPAGE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "OPT_DIRECT_WRKLD", &obj.OPTDIRECTWRKLD)
	if err != nil {
		err = core.SDKErrorf(err, "", "OPT_DIRECT_WRKLD-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "OPT_LOCKLIST", &obj.OPTLOCKLIST)
	if err != nil {
		err = core.SDKErrorf(err, "", "OPT_LOCKLIST-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "OPT_MAXLOCKS", &obj.OPTMAXLOCKS)
	if err != nil {
		err = core.SDKErrorf(err, "", "OPT_MAXLOCKS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "OPT_SORTHEAP", &obj.OPTSORTHEAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "OPT_SORTHEAP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "PAGE_AGE_TRGT_GCR", &obj.PAGEAGETRGTGCR)
	if err != nil {
		err = core.SDKErrorf(err, "", "PAGE_AGE_TRGT_GCR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "PAGE_AGE_TRGT_MCR", &obj.PAGEAGETRGTMCR)
	if err != nil {
		err = core.SDKErrorf(err, "", "PAGE_AGE_TRGT_MCR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "PCKCACHESZ", &obj.PCKCACHESZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "PCKCACHESZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "PL_STACK_TRACE", &obj.PLSTACKTRACE)
	if err != nil {
		err = core.SDKErrorf(err, "", "PL_STACK_TRACE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SELF_TUNING_MEM", &obj.SELFTUNINGMEM)
	if err != nil {
		err = core.SDKErrorf(err, "", "SELF_TUNING_MEM-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SEQDETECT", &obj.SEQDETECT)
	if err != nil {
		err = core.SDKErrorf(err, "", "SEQDETECT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SHEAPTHRES_SHR", &obj.SHEAPTHRESSHR)
	if err != nil {
		err = core.SDKErrorf(err, "", "SHEAPTHRES_SHR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SOFTMAX", &obj.SOFTMAX)
	if err != nil {
		err = core.SDKErrorf(err, "", "SOFTMAX-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SORTHEAP", &obj.SORTHEAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "SORTHEAP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SQL_CCFLAGS", &obj.SQLCCFLAGS)
	if err != nil {
		err = core.SDKErrorf(err, "", "SQL_CCFLAGS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "STAT_HEAP_SZ", &obj.STATHEAPSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "STAT_HEAP_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "STMTHEAP", &obj.STMTHEAP)
	if err != nil {
		err = core.SDKErrorf(err, "", "STMTHEAP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "STMT_CONC", &obj.STMTCONC)
	if err != nil {
		err = core.SDKErrorf(err, "", "STMT_CONC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "STRING_UNITS", &obj.STRINGUNITS)
	if err != nil {
		err = core.SDKErrorf(err, "", "STRING_UNITS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "SYSTIME_PERIOD_ADJ", &obj.SYSTIMEPERIODADJ)
	if err != nil {
		err = core.SDKErrorf(err, "", "SYSTIME_PERIOD_ADJ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "TRACKMOD", &obj.TRACKMOD)
	if err != nil {
		err = core.SDKErrorf(err, "", "TRACKMOD-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "UTIL_HEAP_SZ", &obj.UTILHEAPSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "UTIL_HEAP_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_ADMISSION_CTRL", &obj.WLMADMISSIONCTRL)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_ADMISSION_CTRL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_AGENT_LOAD_TRGT", &obj.WLMAGENTLOADTRGT)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_AGENT_LOAD_TRGT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_CPU_LIMIT", &obj.WLMCPULIMIT)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_CPU_LIMIT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_CPU_SHARES", &obj.WLMCPUSHARES)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_CPU_SHARES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_CPU_SHARE_MODE", &obj.WLMCPUSHAREMODE)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_CPU_SHARE_MODE-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessTuneableParamsTuneableParamDbm : Tunable parameters related to the Db2 instance manager (dbm).
type SuccessTuneableParamsTuneableParamDbm struct {
	COMMBANDWIDTH *string `json:"COMM_BANDWIDTH,omitempty"`

	CPUSPEED *string `json:"CPUSPEED,omitempty"`

	DFTMONBUFPOOL *string `json:"DFT_MON_BUFPOOL,omitempty"`

	DFTMONLOCK *string `json:"DFT_MON_LOCK,omitempty"`

	DFTMONSORT *string `json:"DFT_MON_SORT,omitempty"`

	DFTMONSTMT *string `json:"DFT_MON_STMT,omitempty"`

	DFTMONTABLE *string `json:"DFT_MON_TABLE,omitempty"`

	DFTMONTIMESTAMP *string `json:"DFT_MON_TIMESTAMP,omitempty"`

	DFTMONUOW *string `json:"DFT_MON_UOW,omitempty"`

	DIAGLEVEL *string `json:"DIAGLEVEL,omitempty"`

	FEDERATEDASYNC *string `json:"FEDERATED_ASYNC,omitempty"`

	INDEXREC *string `json:"INDEXREC,omitempty"`

	INTRAPARALLEL *string `json:"INTRA_PARALLEL,omitempty"`

	KEEPFENCED *string `json:"KEEPFENCED,omitempty"`

	MAXCONNRETRIES *string `json:"MAX_CONNRETRIES,omitempty"`

	MAXQUERYDEGREE *string `json:"MAX_QUERYDEGREE,omitempty"`

	MONHEAPSZ *string `json:"MON_HEAP_SZ,omitempty"`

	MULTIPARTSIZEMB *string `json:"MULTIPARTSIZEMB,omitempty"`

	NOTIFYLEVEL *string `json:"NOTIFYLEVEL,omitempty"`

	NUMINITAGENTS *string `json:"NUM_INITAGENTS,omitempty"`

	NUMINITFENCED *string `json:"NUM_INITFENCED,omitempty"`

	NUMPOOLAGENTS *string `json:"NUM_POOLAGENTS,omitempty"`

	RESYNCINTERVAL *string `json:"RESYNC_INTERVAL,omitempty"`

	RQRIOBLK *string `json:"RQRIOBLK,omitempty"`

	STARTSTOPTIME *string `json:"START_STOP_TIME,omitempty"`

	UTILIMPACTLIM *string `json:"UTIL_IMPACT_LIM,omitempty"`

	WLMDISPATCHER *string `json:"WLM_DISPATCHER,omitempty"`

	WLMDISPCONCUR *string `json:"WLM_DISP_CONCUR,omitempty"`

	WLMDISPCPUSHARES *string `json:"WLM_DISP_CPU_SHARES,omitempty"`

	WLMDISPMINUTIL *string `json:"WLM_DISP_MIN_UTIL,omitempty"`
}

// UnmarshalSuccessTuneableParamsTuneableParamDbm unmarshals an instance of SuccessTuneableParamsTuneableParamDbm from the specified map of raw messages.
func UnmarshalSuccessTuneableParamsTuneableParamDbm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessTuneableParamsTuneableParamDbm)
	err = core.UnmarshalPrimitive(m, "COMM_BANDWIDTH", &obj.COMMBANDWIDTH)
	if err != nil {
		err = core.SDKErrorf(err, "", "COMM_BANDWIDTH-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "CPUSPEED", &obj.CPUSPEED)
	if err != nil {
		err = core.SDKErrorf(err, "", "CPUSPEED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_BUFPOOL", &obj.DFTMONBUFPOOL)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_BUFPOOL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_LOCK", &obj.DFTMONLOCK)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_LOCK-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_SORT", &obj.DFTMONSORT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_SORT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_STMT", &obj.DFTMONSTMT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_STMT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_TABLE", &obj.DFTMONTABLE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_TABLE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_TIMESTAMP", &obj.DFTMONTIMESTAMP)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_TIMESTAMP-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DFT_MON_UOW", &obj.DFTMONUOW)
	if err != nil {
		err = core.SDKErrorf(err, "", "DFT_MON_UOW-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DIAGLEVEL", &obj.DIAGLEVEL)
	if err != nil {
		err = core.SDKErrorf(err, "", "DIAGLEVEL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "FEDERATED_ASYNC", &obj.FEDERATEDASYNC)
	if err != nil {
		err = core.SDKErrorf(err, "", "FEDERATED_ASYNC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "INDEXREC", &obj.INDEXREC)
	if err != nil {
		err = core.SDKErrorf(err, "", "INDEXREC-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "INTRA_PARALLEL", &obj.INTRAPARALLEL)
	if err != nil {
		err = core.SDKErrorf(err, "", "INTRA_PARALLEL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "KEEPFENCED", &obj.KEEPFENCED)
	if err != nil {
		err = core.SDKErrorf(err, "", "KEEPFENCED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MAX_CONNRETRIES", &obj.MAXCONNRETRIES)
	if err != nil {
		err = core.SDKErrorf(err, "", "MAX_CONNRETRIES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MAX_QUERYDEGREE", &obj.MAXQUERYDEGREE)
	if err != nil {
		err = core.SDKErrorf(err, "", "MAX_QUERYDEGREE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MON_HEAP_SZ", &obj.MONHEAPSZ)
	if err != nil {
		err = core.SDKErrorf(err, "", "MON_HEAP_SZ-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "MULTIPARTSIZEMB", &obj.MULTIPARTSIZEMB)
	if err != nil {
		err = core.SDKErrorf(err, "", "MULTIPARTSIZEMB-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NOTIFYLEVEL", &obj.NOTIFYLEVEL)
	if err != nil {
		err = core.SDKErrorf(err, "", "NOTIFYLEVEL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_INITAGENTS", &obj.NUMINITAGENTS)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_INITAGENTS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_INITFENCED", &obj.NUMINITFENCED)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_INITFENCED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "NUM_POOLAGENTS", &obj.NUMPOOLAGENTS)
	if err != nil {
		err = core.SDKErrorf(err, "", "NUM_POOLAGENTS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "RESYNC_INTERVAL", &obj.RESYNCINTERVAL)
	if err != nil {
		err = core.SDKErrorf(err, "", "RESYNC_INTERVAL-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "RQRIOBLK", &obj.RQRIOBLK)
	if err != nil {
		err = core.SDKErrorf(err, "", "RQRIOBLK-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "START_STOP_TIME", &obj.STARTSTOPTIME)
	if err != nil {
		err = core.SDKErrorf(err, "", "START_STOP_TIME-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "UTIL_IMPACT_LIM", &obj.UTILIMPACTLIM)
	if err != nil {
		err = core.SDKErrorf(err, "", "UTIL_IMPACT_LIM-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_DISPATCHER", &obj.WLMDISPATCHER)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_DISPATCHER-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_DISP_CONCUR", &obj.WLMDISPCONCUR)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_DISP_CONCUR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_DISP_CPU_SHARES", &obj.WLMDISPCPUSHARES)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_DISP_CPU_SHARES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "WLM_DISP_MIN_UTIL", &obj.WLMDISPMINUTIL)
	if err != nil {
		err = core.SDKErrorf(err, "", "WLM_DISP_MIN_UTIL-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessTuneableParamsTuneableParamRegistry : Tunable parameters related to the Db2 registry.
type SuccessTuneableParamsTuneableParamRegistry struct {
	DB2BIDI *string `json:"DB2BIDI,omitempty"`

	DB2COMPOPT *string `json:"DB2COMPOPT,omitempty"`

	DB2LOCKTORB *string `json:"DB2LOCK_TO_RB,omitempty"`

	DB2STMM *string `json:"DB2STMM,omitempty"`

	DB2ALTERNATEAUTHZBEHAVIOUR *string `json:"DB2_ALTERNATE_AUTHZ_BEHAVIOUR,omitempty"`

	DB2ANTIJOIN *string `json:"DB2_ANTIJOIN,omitempty"`

	DB2ATSENABLE *string `json:"DB2_ATS_ENABLE,omitempty"`

	DB2DEFERREDPREPARESEMANTICS *string `json:"DB2_DEFERRED_PREPARE_SEMANTICS,omitempty"`

	DB2EVALUNCOMMITTED *string `json:"DB2_EVALUNCOMMITTED,omitempty"`

	DB2EXTENDEDOPTIMIZATION *string `json:"DB2_EXTENDED_OPTIMIZATION,omitempty"`

	DB2INDEXPCTFREEDEFAULT *string `json:"DB2_INDEX_PCTFREE_DEFAULT,omitempty"`

	DB2INLISTTONLJN *string `json:"DB2_INLIST_TO_NLJN,omitempty"`

	DB2MINIMIZELISTPREFETCH *string `json:"DB2_MINIMIZE_LISTPREFETCH,omitempty"`

	DB2OBJECTTABLEENTRIES *string `json:"DB2_OBJECT_TABLE_ENTRIES,omitempty"`

	DB2OPTPROFILE *string `json:"DB2_OPTPROFILE,omitempty"`

	DB2OPTSTATSLOG *string `json:"DB2_OPTSTATS_LOG,omitempty"`

	DB2OPTMAXTEMPSIZE *string `json:"DB2_OPT_MAX_TEMP_SIZE,omitempty"`

	DB2PARALLELIO *string `json:"DB2_PARALLEL_IO,omitempty"`

	DB2REDUCEDOPTIMIZATION *string `json:"DB2_REDUCED_OPTIMIZATION,omitempty"`

	DB2SELECTIVITY *string `json:"DB2_SELECTIVITY,omitempty"`

	DB2SKIPDELETED *string `json:"DB2_SKIPDELETED,omitempty"`

	DB2SKIPINSERTED *string `json:"DB2_SKIPINSERTED,omitempty"`

	DB2SYNCRELEASELOCKATTRIBUTES *string `json:"DB2_SYNC_RELEASE_LOCK_ATTRIBUTES,omitempty"`

	DB2TRUNCATEREUSESTORAGE *string `json:"DB2_TRUNCATE_REUSESTORAGE,omitempty"`

	DB2USEALTERNATEPAGECLEANING *string `json:"DB2_USE_ALTERNATE_PAGE_CLEANING,omitempty"`

	DB2VIEWREOPTVALUES *string `json:"DB2_VIEW_REOPT_VALUES,omitempty"`

	DB2WLMSETTINGS *string `json:"DB2_WLM_SETTINGS,omitempty"`

	DB2WORKLOAD *string `json:"DB2_WORKLOAD,omitempty"`
}

// UnmarshalSuccessTuneableParamsTuneableParamRegistry unmarshals an instance of SuccessTuneableParamsTuneableParamRegistry from the specified map of raw messages.
func UnmarshalSuccessTuneableParamsTuneableParamRegistry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessTuneableParamsTuneableParamRegistry)
	err = core.UnmarshalPrimitive(m, "DB2BIDI", &obj.DB2BIDI)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2BIDI-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2COMPOPT", &obj.DB2COMPOPT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2COMPOPT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2LOCK_TO_RB", &obj.DB2LOCKTORB)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2LOCK_TO_RB-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2STMM", &obj.DB2STMM)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2STMM-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_ALTERNATE_AUTHZ_BEHAVIOUR", &obj.DB2ALTERNATEAUTHZBEHAVIOUR)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_ALTERNATE_AUTHZ_BEHAVIOUR-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_ANTIJOIN", &obj.DB2ANTIJOIN)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_ANTIJOIN-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_ATS_ENABLE", &obj.DB2ATSENABLE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_ATS_ENABLE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_DEFERRED_PREPARE_SEMANTICS", &obj.DB2DEFERREDPREPARESEMANTICS)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_DEFERRED_PREPARE_SEMANTICS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_EVALUNCOMMITTED", &obj.DB2EVALUNCOMMITTED)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_EVALUNCOMMITTED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_EXTENDED_OPTIMIZATION", &obj.DB2EXTENDEDOPTIMIZATION)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_EXTENDED_OPTIMIZATION-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_INDEX_PCTFREE_DEFAULT", &obj.DB2INDEXPCTFREEDEFAULT)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_INDEX_PCTFREE_DEFAULT-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_INLIST_TO_NLJN", &obj.DB2INLISTTONLJN)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_INLIST_TO_NLJN-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_MINIMIZE_LISTPREFETCH", &obj.DB2MINIMIZELISTPREFETCH)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_MINIMIZE_LISTPREFETCH-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_OBJECT_TABLE_ENTRIES", &obj.DB2OBJECTTABLEENTRIES)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_OBJECT_TABLE_ENTRIES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_OPTPROFILE", &obj.DB2OPTPROFILE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_OPTPROFILE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_OPTSTATS_LOG", &obj.DB2OPTSTATSLOG)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_OPTSTATS_LOG-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_OPT_MAX_TEMP_SIZE", &obj.DB2OPTMAXTEMPSIZE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_OPT_MAX_TEMP_SIZE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_PARALLEL_IO", &obj.DB2PARALLELIO)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_PARALLEL_IO-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_REDUCED_OPTIMIZATION", &obj.DB2REDUCEDOPTIMIZATION)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_REDUCED_OPTIMIZATION-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_SELECTIVITY", &obj.DB2SELECTIVITY)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_SELECTIVITY-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_SKIPDELETED", &obj.DB2SKIPDELETED)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_SKIPDELETED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_SKIPINSERTED", &obj.DB2SKIPINSERTED)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_SKIPINSERTED-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES", &obj.DB2SYNCRELEASELOCKATTRIBUTES)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_TRUNCATE_REUSESTORAGE", &obj.DB2TRUNCATEREUSESTORAGE)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_TRUNCATE_REUSESTORAGE-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_USE_ALTERNATE_PAGE_CLEANING", &obj.DB2USEALTERNATEPAGECLEANING)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_USE_ALTERNATE_PAGE_CLEANING-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_VIEW_REOPT_VALUES", &obj.DB2VIEWREOPTVALUES)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_VIEW_REOPT_VALUES-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_WLM_SETTINGS", &obj.DB2WLMSETTINGS)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_WLM_SETTINGS-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "DB2_WORKLOAD", &obj.DB2WORKLOAD)
	if err != nil {
		err = core.SDKErrorf(err, "", "DB2_WORKLOAD-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUpdateAutoScale : Response of successful updation of scaling configurations.
type SuccessUpdateAutoScale struct {
	// Indicates the message of the updation.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalSuccessUpdateAutoScale unmarshals an instance of SuccessUpdateAutoScale from the specified map of raw messages.
func UnmarshalSuccessUpdateAutoScale(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUpdateAutoScale)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUserResponse : The details of the users.
type SuccessUserResponse struct {
	// User's DV role.
	DvRole *string `json:"dvRole" validate:"required"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata" validate:"required"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid" validate:"required"`

	// Role assigned to the user.
	Role *string `json:"role" validate:"required"`

	// IAM ID for the user.
	Iamid *string `json:"iamid" validate:"required"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions" validate:"required"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean" validate:"required"`

	// User's password.
	Password *string `json:"password" validate:"required"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam" validate:"required"`

	// The display name of the user.
	Name *string `json:"name" validate:"required"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid" validate:"required"`

	// Unique identifier for the user.
	ID *string `json:"id" validate:"required"`

	// Account lock status for the user.
	Locked *string `json:"locked" validate:"required"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg" validate:"required"`

	// Email address of the user.
	Email *string `json:"email" validate:"required"`

	// Authentication details for the user.
	Authentication *SuccessUserResponseAuthentication `json:"authentication" validate:"required"`
}

// Constants associated with the SuccessUserResponse.Role property.
// Role assigned to the user.
const (
	SuccessUserResponse_Role_Bluadmin = "bluadmin"
	SuccessUserResponse_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessUserResponse.Locked property.
// Account lock status for the user.
const (
	SuccessUserResponse_Locked_No = "no"
	SuccessUserResponse_Locked_Yes = "yes"
)

// UnmarshalSuccessUserResponse unmarshals an instance of SuccessUserResponse from the specified map of raw messages.
func UnmarshalSuccessUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUserResponse)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessUserResponseAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUserResponseAuthentication : Authentication details for the user.
type SuccessUserResponseAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessUserResponseAuthentication unmarshals an instance of SuccessUserResponseAuthentication from the specified map of raw messages.
func UnmarshalSuccessUserResponseAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUserResponseAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
