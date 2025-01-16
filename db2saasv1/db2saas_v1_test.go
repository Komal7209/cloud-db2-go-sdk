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

 package db2saasv1_test

 import (
	 "bytes"
	 "context"
	 "encoding/base64"
	 "encoding/json"
	 "fmt"
	 "io"
	 "net/http"
	 "net/http/httptest"
	 "os"
	 "time"
 
	 "github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	 "github.com/IBM/go-sdk-core/v5/core"
	 "github.com/go-openapi/strfmt"
	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
 )
 
 var _ = Describe(`Db2saasV1`, func() {
	 var testServer *httptest.Server
	 Describe(`Service constructor tests`, func() {
		 It(`Instantiate service client`, func() {
			 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				 Authenticator: &core.NoAuthAuthenticator{},
			 })
			 Expect(db2saasService).ToNot(BeNil())
			 Expect(serviceErr).To(BeNil())
		 })
		 It(`Instantiate service client with error: Invalid URL`, func() {
			 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				 URL: "{BAD_URL_STRING",
			 })
			 Expect(db2saasService).To(BeNil())
			 Expect(serviceErr).ToNot(BeNil())
		 })
		 It(`Instantiate service client with error: Invalid Auth`, func() {
			 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				 URL: "https://db2saasv1/api",
				 Authenticator: &core.BasicAuthenticator{
					 Username: "",
					 Password: "",
				 },
			 })
			 Expect(db2saasService).To(BeNil())
			 Expect(serviceErr).ToNot(BeNil())
		 })
	 })
	 Describe(`Service constructor tests using external config`, func() {
		 Context(`Using external config, construct service client instances`, func() {
			 // Map containing environment variables used in testing.
			 var testEnvironment = map[string]string{
				 "DB2SAAS_URL": "https://db2saasv1/api",
				 "DB2SAAS_AUTH_TYPE": "noauth",
			 }
 
			 It(`Create service client using external config successfully`, func() {
				 SetTestEnvironment(testEnvironment)
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				 })
				 Expect(db2saasService).ToNot(BeNil())
				 Expect(serviceErr).To(BeNil())
				 ClearTestEnvironment(testEnvironment)
 
				 clone := db2saasService.Clone()
				 Expect(clone).ToNot(BeNil())
				 Expect(clone.Service != db2saasService.Service).To(BeTrue())
				 Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				 Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			 })
			 It(`Create service client using external config and set url from constructor successfully`, func() {
				 SetTestEnvironment(testEnvironment)
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
					 URL: "https://testService/api",
				 })
				 Expect(db2saasService).ToNot(BeNil())
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				 ClearTestEnvironment(testEnvironment)
 
				 clone := db2saasService.Clone()
				 Expect(clone).ToNot(BeNil())
				 Expect(clone.Service != db2saasService.Service).To(BeTrue())
				 Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				 Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			 })
			 It(`Create service client using external config and set url programatically successfully`, func() {
				 SetTestEnvironment(testEnvironment)
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				 })
				 err := db2saasService.SetServiceURL("https://testService/api")
				 Expect(err).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				 ClearTestEnvironment(testEnvironment)
 
				 clone := db2saasService.Clone()
				 Expect(clone).ToNot(BeNil())
				 Expect(clone.Service != db2saasService.Service).To(BeTrue())
				 Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				 Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			 })
		 })
		 Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			 // Map containing environment variables used in testing.
			 var testEnvironment = map[string]string{
				 "DB2SAAS_URL": "https://db2saasv1/api",
				 "DB2SAAS_AUTH_TYPE": "someOtherAuth",
			 }
 
			 SetTestEnvironment(testEnvironment)
			 db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
			 })
 
			 It(`Instantiate service client with error`, func() {
				 Expect(db2saasService).To(BeNil())
				 Expect(serviceErr).ToNot(BeNil())
				 ClearTestEnvironment(testEnvironment)
			 })
		 })
		 Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			 // Map containing environment variables used in testing.
			 var testEnvironment = map[string]string{
				 "DB2SAAS_AUTH_TYPE":   "NOAuth",
			 }
 
			 SetTestEnvironment(testEnvironment)
			 db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				 URL: "{BAD_URL_STRING",
			 })
 
			 It(`Instantiate service client with error`, func() {
				 Expect(db2saasService).To(BeNil())
				 Expect(serviceErr).ToNot(BeNil())
				 ClearTestEnvironment(testEnvironment)
			 })
		 })
	 })
	 Describe(`Regional endpoint tests`, func() {
		 It(`GetServiceURLForRegion(region string)`, func() {
			 var url string
			 var err error
			 url, err = db2saasv1.GetServiceURLForRegion("INVALID_REGION")
			 Expect(url).To(BeEmpty())
			 Expect(err).ToNot(BeNil())
			 fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		 })
	 })
	 Describe(`Parameterized URL tests`, func() {
		 It(`Format parameterized URL with all default values`, func() {
			 constructedURL, err := db2saasv1.ConstructServiceURL(nil)
			 Expect(constructedURL).To(Equal("https://us-south.db2.saas.ibm.com/dbapi/v4"))
			 Expect(constructedURL).ToNot(BeNil())
			 Expect(err).To(BeNil())
		 })
		 It(`Return an error if a provided variable name is invalid`, func() {
			 var providedUrlVariables = map[string]string{
				 "invalid_variable_name": "value",
			 }
			 constructedURL, err := db2saasv1.ConstructServiceURL(providedUrlVariables)
			 Expect(constructedURL).To(Equal(""))
			 Expect(err).ToNot(BeNil())
		 })
	 })
	 Describe(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) - Operation response error`, func() {
		 getDb2SaasConnectionInfoPath := "/connectioninfo/crn%253Av1%253Astaging%253Apublic%253Adashdb-for-transactions%253Aus-south%253Aa%252Fe7e3e87b512f474381c0684a5ecbba03%253A69db420f-33d5-4953-8bd8-1950abd356f6%253A%253A"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasConnectionInfo with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions)`, func() {
		 getDb2SaasConnectionInfoPath := "/connectioninfo/crn%253Av1%253Astaging%253Apublic%253Adashdb-for-transactions%253Aus-south%253Aa%252Fe7e3e87b512f474381c0684a5ecbba03%253A69db420f-33d5-4953-8bd8-1950abd356f6%253A%253A"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0", "private_serviceName": "us-south-private.db2oc.test.saas.ibm.com:32764", "cloud_service_offering": "dashdb-for-transactions", "vpe_service_crn": "crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud", "db_vpc_endpoint_service": "feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679"}}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasConnectionInfo successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasConnectionInfoWithContext(ctx, getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasConnectionInfoWithContext(ctx, getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0", "private_serviceName": "us-south-private.db2oc.test.saas.ibm.com:32764", "cloud_service_offering": "dashdb-for-transactions", "vpe_service_crn": "crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud", "db_vpc_endpoint_service": "feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679"}}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasConnectionInfo successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasConnectionInfo with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasConnectionInfoOptions model with no property values
				 getDb2SaasConnectionInfoOptionsModelNew := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasConnectionInfo successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				 getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) - Operation response error`, func() {
		 postDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("POST"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke PostDb2SaasAllowlist with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				 postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions)`, func() {
		 postDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasAllowlist successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				 postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.PostDb2SaasAllowlistWithContext(ctx, postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.PostDb2SaasAllowlistWithContext(ctx, postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasAllowlist successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.PostDb2SaasAllowlist(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				 postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke PostDb2SaasAllowlist with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				 postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the PostDb2SaasAllowlistOptions model with no property values
				 postDb2SaasAllowlistOptionsModelNew := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke PostDb2SaasAllowlist successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				 postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				 postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) - Operation response error`, func() {
		 getDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAllowlist with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions)`, func() {
		 getDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAllowlist successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasAllowlistWithContext(ctx, getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasAllowlistWithContext(ctx, getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAllowlist successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasAllowlist(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasAllowlist with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasAllowlistOptions model with no property values
				 getDb2SaasAllowlistOptionsModelNew := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasAllowlist successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				 getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions) - Operation response error`, func() {
		 postDb2SaasUserPath := "/users"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					 Expect(req.Method).To(Equal("POST"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke PostDb2SaasUser with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				 postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				 postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				 postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				 postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				 postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				 postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				 postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions)`, func() {
		 postDb2SaasUserPath := "/users"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasUser successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				 postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				 postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				 postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				 postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				 postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				 postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				 postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.PostDb2SaasUserWithContext(ctx, postDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.PostDb2SaasUserWithContext(ctx, postDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.PostDb2SaasUser(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				 postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				 postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				 postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				 postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				 postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				 postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				 postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke PostDb2SaasUser with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				 postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				 postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				 postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				 postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				 postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				 postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				 postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the PostDb2SaasUserOptions model with no property values
				 postDb2SaasUserOptionsModelNew := new(db2saasv1.PostDb2SaasUserOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke PostDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				 postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				 postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				 postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				 postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				 postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				 postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				 postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions) - Operation response error`, func() {
		 getDb2SaasUserPath := "/users"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasUser with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasUserOptions model
				 getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				 getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions)`, func() {
		 getDb2SaasUserPath := "/users"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"count": 1, "resources": [{"dvRole": "test-role", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "test-formated-ibm-id", "role": "bluadmin", "iamid": "test-iam-id", "permittedActions": ["PermittedActions"], "allClean": false, "password": "nd!@aegr63@989hcRFTcdcs63", "iam": false, "name": "admin", "ibmid": "test-ibm-id", "id": "admin", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "internal", "policy_id": "Default"}}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasUser successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasUserOptions model
				 getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				 getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasUserWithContext(ctx, getDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasUserWithContext(ctx, getDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"count": 1, "resources": [{"dvRole": "test-role", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "test-formated-ibm-id", "role": "bluadmin", "iamid": "test-iam-id", "permittedActions": ["PermittedActions"], "allClean": false, "password": "nd!@aegr63@989hcRFTcdcs63", "iam": false, "name": "admin", "ibmid": "test-ibm-id", "id": "admin", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "internal", "policy_id": "Default"}}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasUser(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasUserOptions model
				 getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				 getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasUser with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasUserOptions model
				 getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				 getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasUserOptions model with no property values
				 getDb2SaasUserOptionsModelNew := new(db2saasv1.GetDb2SaasUserOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasUserOptions model
				 getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				 getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions)`, func() {
		 deleteDb2SaasUserPath := "/users/test-user"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteDb2SaasUserPath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				 }))
			 })
			 It(`Invoke DeleteDb2SaasUser successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the DeleteDb2SaasUserOptions model
				 deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				 deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.DeleteDb2SaasUserWithContext(ctx, deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.DeleteDb2SaasUserWithContext(ctx, deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(deleteDb2SaasUserPath))
					 Expect(req.Method).To(Equal("DELETE"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				 }))
			 })
			 It(`Invoke DeleteDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.DeleteDb2SaasUser(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the DeleteDb2SaasUserOptions model
				 deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				 deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke DeleteDb2SaasUser with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteDb2SaasUserOptions model
				 deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				 deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the DeleteDb2SaasUserOptions model with no property values
				 deleteDb2SaasUserOptionsModelNew := new(db2saasv1.DeleteDb2SaasUserOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke DeleteDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the DeleteDb2SaasUserOptions model
				 deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				 deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				 deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) - Operation response error`, func() {
		 getbyidDb2SaasUserPath := "/users/bluadmin"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetbyidDb2SaasUser with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions)`, func() {
		 getbyidDb2SaasUserPath := "/users/bluadmin"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				 }))
			 })
			 It(`Invoke GetbyidDb2SaasUser successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetbyidDb2SaasUserWithContext(ctx, getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetbyidDb2SaasUserWithContext(ctx, getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				 }))
			 })
			 It(`Invoke GetbyidDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetbyidDb2SaasUser(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetbyidDb2SaasUser with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetbyidDb2SaasUserOptions model with no property values
				 getbyidDb2SaasUserOptionsModelNew := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetbyidDb2SaasUser successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				 getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) - Operation response error`, func() {
		 putDb2SaasAutoscalePath := "/manage/scaling/auto"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("PUT"))
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke PutDb2SaasAutoscale with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				 putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions)`, func() {
		 putDb2SaasAutoscalePath := "/manage/scaling/auto"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("PUT"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				 }))
			 })
			 It(`Invoke PutDb2SaasAutoscale successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				 putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.PutDb2SaasAutoscaleWithContext(ctx, putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.PutDb2SaasAutoscaleWithContext(ctx, putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("PUT"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				 }))
			 })
			 It(`Invoke PutDb2SaasAutoscale successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.PutDb2SaasAutoscale(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				 putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke PutDb2SaasAutoscale with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				 putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the PutDb2SaasAutoscaleOptions model with no property values
				 putDb2SaasAutoscaleOptionsModelNew := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke PutDb2SaasAutoscale successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				 putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				 putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				 putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				 putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) - Operation response error`, func() {
		 getDb2SaasAutoscalePath := "/manage/scaling/auto"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAutoscale with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions)`, func() {
		 getDb2SaasAutoscalePath := "/manage/scaling/auto"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"auto_scaling_allow_plan_limit": false, "auto_scaling_enabled": true, "auto_scaling_max_storage": 21, "auto_scaling_over_time_period": 25, "auto_scaling_pause_limit": 21, "auto_scaling_threshold": 20, "storage_unit": "StorageUnit", "storage_utilization_percentage": 28, "support_auto_scaling": true}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAutoscale successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasAutoscaleWithContext(ctx, getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasAutoscaleWithContext(ctx, getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					 Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"auto_scaling_allow_plan_limit": false, "auto_scaling_enabled": true, "auto_scaling_max_storage": 21, "auto_scaling_over_time_period": 25, "auto_scaling_pause_limit": 21, "auto_scaling_threshold": 20, "storage_unit": "StorageUnit", "storage_utilization_percentage": 28, "support_auto_scaling": true}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasAutoscale successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasAutoscale(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasAutoscale with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasAutoscaleOptions model with no property values
				 getDb2SaasAutoscaleOptionsModelNew := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasAutoscale successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				 getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions) - Operation response error`, func() {
		 postDb2SaasDbConfigurationPath := "/manage/deployments/custom_setting"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					 Expect(req.Method).To(Equal("POST"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke PostDb2SaasDbConfiguration with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 postDb2SaasDbConfigurationOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				 postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				 postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions)`, func() {
		 postDb2SaasDbConfigurationPath := "/manage/deployments/custom_setting"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "status": "Status"}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasDbConfiguration successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 postDb2SaasDbConfigurationOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				 postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				 postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.PostDb2SaasDbConfigurationWithContext(ctx, postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.PostDb2SaasDbConfigurationWithContext(ctx, postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 // For gzip-disabled operation, verify Content-Encoding is not set.
					 Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())
 
					 // If there is a body, then make sure we can read it
					 bodyBuf := new(bytes.Buffer)
					 if req.Header.Get("Content-Encoding") == "gzip" {
						 body, err := core.NewGzipDecompressionReader(req.Body)
						 Expect(err).To(BeNil())
						 _, err = bodyBuf.ReadFrom(body)
						 Expect(err).To(BeNil())
					 } else {
						 _, err := bodyBuf.ReadFrom(req.Body)
						 Expect(err).To(BeNil())
					 }
					 fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "status": "Status"}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasDbConfiguration successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 postDb2SaasDbConfigurationOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				 postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				 postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke PostDb2SaasDbConfiguration with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 postDb2SaasDbConfigurationOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				 postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				 postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the PostDb2SaasDbConfigurationOptions model with no property values
				 postDb2SaasDbConfigurationOptionsModelNew := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke PostDb2SaasDbConfiguration successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				 postDb2SaasDbConfigurationOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				 postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				 postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions) - Operation response error`, func() {
		 getDb2SaasTuneableParamPath := "/manage/tuneable_param"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasTuneableParam with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 getDb2SaasTuneableParamOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions)`, func() {
		 getDb2SaasTuneableParamPath := "/manage/tuneable_param"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"tuneable_param": {"db": {"ACT_SORTMEM_LIMIT": "'NONE', 'range(10, 100)'", "ALT_COLLATE": "'NULL', 'IDENTITY_16BIT'", "APPGROUP_MEM_SZ": "'range(1, 1000000)'", "APPLHEAPSZ": "'AUTOMATIC', 'range(16, 2147483647)'", "APPL_MEMORY": "'AUTOMATIC', 'range(128, 4294967295)'", "APP_CTL_HEAP_SZ": "'range(1, 64000)'", "ARCHRETRYDELAY": "'range(0, 65535)'", "AUTHN_CACHE_DURATION": "'range(1,10000)'", "AUTORESTART": "'ON', 'OFF'", "AUTO_CG_STATS": "'ON', 'OFF'", "AUTO_MAINT": "'ON', 'OFF'", "AUTO_REORG": "'ON', 'OFF'", "AUTO_REVAL": "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'", "AUTO_RUNSTATS": "'ON', 'OFF'", "AUTO_SAMPLING": "'ON', 'OFF'", "AUTO_STATS_VIEWS": "'ON', 'OFF'", "AUTO_STMT_STATS": "'ON', 'OFF'", "AUTO_TBL_MAINT": "'ON', 'OFF'", "AVG_APPLS": "'-'", "CATALOGCACHE_SZ": "'-'", "CHNGPGS_THRESH": "'range(5,99)'", "CUR_COMMIT": "'ON, AVAILABLE, DISABLED'", "DATABASE_MEMORY": "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'", "DBHEAP": "'AUTOMATIC', 'range(32, 2147483647)'", "DB_COLLNAME": "'-'", "DB_MEM_THRESH": "'range(0, 100)'", "DDL_COMPRESSION_DEF": "'YES', 'NO'", "DDL_CONSTRAINT_DEF": "'YES', 'NO'", "DECFLT_ROUNDING": "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'", "DEC_ARITHMETIC": "'-'", "DEC_TO_CHAR_FMT": "'NEW', 'V95'", "DFT_DEGREE": "'-1', 'ANY', 'range(1, 32767)'", "DFT_EXTENT_SZ": "'range(2, 256)'", "DFT_LOADREC_SES": "'range(1, 30000)'", "DFT_MTTB_TYPES": "'-'", "DFT_PREFETCH_SZ": "'range(0, 32767)', 'AUTOMATIC'", "DFT_QUERYOPT": "'range(0, 9)'", "DFT_REFRESH_AGE": "'-'", "DFT_SCHEMAS_DCC": "'YES', 'NO'", "DFT_SQLMATHWARN": "'YES', 'NO'", "DFT_TABLE_ORG": "'COLUMN', 'ROW'", "DLCHKTIME": "'range(1000, 600000)'", "ENABLE_XMLCHAR": "'YES', 'NO'", "EXTENDED_ROW_SZ": "'ENABLE', 'DISABLE'", "GROUPHEAP_RATIO": "'range(1, 99)'", "INDEXREC": "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'", "LARGE_AGGREGATION": "'YES', 'NO'", "LOCKLIST": "'AUTOMATIC', 'range(4, 134217728)'", "LOCKTIMEOUT": "'-1', 'range(0, 32767)'", "LOGINDEXBUILD": "'ON', 'OFF'", "LOG_APPL_INFO": "'YES', 'NO'", "LOG_DDL_STMTS": "'YES', 'NO'", "LOG_DISK_CAP": "'0', '-1', 'range(1, 2147483647)'", "MAXAPPLS": "'range(1, 60000)'", "MAXFILOP": "'range(64, 61440)'", "MAXLOCKS": "'AUTOMATIC', 'range(1, 100)'", "MIN_DEC_DIV_3": "'YES', 'NO'", "MON_ACT_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_DEADLOCK": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LCK_MSG_LVL": "'range(0, 3)'", "MON_LOCKTIMEOUT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LOCKWAIT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LW_THRESH": "'range(1000, 4294967295)'", "MON_OBJ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_PKGLIST_SZ": "'range(0, 1024)'", "MON_REQ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_RTN_DATA": "'NONE', 'BASE'", "MON_RTN_EXECLIST": "'OFF', 'ON'", "MON_UOW_DATA": "'NONE', 'BASE'", "MON_UOW_EXECLIST": "'ON', 'OFF'", "MON_UOW_PKGLIST": "'OFF', 'ON'", "NCHAR_MAPPING": "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'", "NUM_FREQVALUES": "'range(0, 32767)'", "NUM_IOCLEANERS": "'AUTOMATIC', 'range(0, 255)'", "NUM_IOSERVERS": "'AUTOMATIC', 'range(1, 255)'", "NUM_LOG_SPAN": "'range(0, 65535)'", "NUM_QUANTILES": "'range(0, 32767)'", "OPT_BUFFPAGE": "'-'", "OPT_DIRECT_WRKLD": "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'", "OPT_LOCKLIST": "'-'", "OPT_MAXLOCKS": "'-'", "OPT_SORTHEAP": "'-'", "PAGE_AGE_TRGT_GCR": "'range(1, 65535)'", "PAGE_AGE_TRGT_MCR": "'range(1, 65535)'", "PCKCACHESZ": "'AUTOMATIC', '-1', 'range(32, 2147483646)'", "PL_STACK_TRACE": "'NONE', 'ALL', 'UNHANDLED'", "SELF_TUNING_MEM": "'ON', 'OFF'", "SEQDETECT": "'YES', 'NO'", "SHEAPTHRES_SHR": "'AUTOMATIC', 'range(250, 2147483647)'", "SOFTMAX": "'-'", "SORTHEAP": "'AUTOMATIC', 'range(16, 4294967295)'", "SQL_CCFLAGS": "'-'", "STAT_HEAP_SZ": "'AUTOMATIC', 'range(1096, 2147483647)'", "STMTHEAP": "'AUTOMATIC', 'range(128, 2147483647)'", "STMT_CONC": "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'", "STRING_UNITS": "'SYSTEM', 'CODEUNITS32'", "SYSTIME_PERIOD_ADJ": "'NO', 'YES'", "TRACKMOD": "'YES', 'NO'", "UTIL_HEAP_SZ": "'AUTOMATIC', 'range(16, 2147483647)'", "WLM_ADMISSION_CTRL": "'YES', 'NO'", "WLM_AGENT_LOAD_TRGT": "'AUTOMATIC', 'range(1, 65535)'", "WLM_CPU_LIMIT": "'range(0, 100)'", "WLM_CPU_SHARES": "'range(1, 65535)'", "WLM_CPU_SHARE_MODE": "'HARD', 'SOFT'"}, "dbm": {"COMM_BANDWIDTH": "'range(0.1, 100000)', '-1'", "CPUSPEED": "'range(0.0000000001, 1)', '-1'", "DFT_MON_BUFPOOL": "'ON', 'OFF'", "DFT_MON_LOCK": "'ON', 'OFF'", "DFT_MON_SORT": "'ON', 'OFF'", "DFT_MON_STMT": "'ON', 'OFF'", "DFT_MON_TABLE": "'ON', 'OFF'", "DFT_MON_TIMESTAMP": "'ON', 'OFF'", "DFT_MON_UOW": "'ON', 'OFF'", "DIAGLEVEL": "'range(0, 4)'", "FEDERATED_ASYNC": "'range(0, 32767)', '-1', 'ANY'", "INDEXREC": "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'", "INTRA_PARALLEL": "'SYSTEM', 'NO', 'YES'", "KEEPFENCED": "'YES', 'NO'", "MAX_CONNRETRIES": "'range(0, 100)'", "MAX_QUERYDEGREE": "'range(1, 32767)', '-1', 'ANY'", "MON_HEAP_SZ": "'range(0, 2147483647)', 'AUTOMATIC'", "MULTIPARTSIZEMB": "'range(5, 5120)'", "NOTIFYLEVEL": "'range(0, 4)'", "NUM_INITAGENTS": "'range(0, 64000)'", "NUM_INITFENCED": "'range(0, 64000)'", "NUM_POOLAGENTS": "'-1', 'range(0, 64000)'", "RESYNC_INTERVAL": "'range(1, 60000)'", "RQRIOBLK": "'range(4096, 65535)'", "START_STOP_TIME": "'range(1, 1440)'", "UTIL_IMPACT_LIM": "'range(1, 100)'", "WLM_DISPATCHER": "'YES', 'NO'", "WLM_DISP_CONCUR": "'range(1, 32767)', 'COMPUTED'", "WLM_DISP_CPU_SHARES": "'NO', 'YES'", "WLM_DISP_MIN_UTIL": "'range(0, 100)'"}, "registry": {"DB2BIDI": "'YES', 'NO'", "DB2COMPOPT": "'-'", "DB2LOCK_TO_RB": "'STATEMENT'", "DB2STMM": "'NO', 'YES'", "DB2_ALTERNATE_AUTHZ_BEHAVIOUR": "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'", "DB2_ANTIJOIN": "'YES', 'NO', 'EXTEND'", "DB2_ATS_ENABLE": "'YES', 'NO'", "DB2_DEFERRED_PREPARE_SEMANTICS": "'NO', 'YES'", "DB2_EVALUNCOMMITTED": "'NO', 'YES'", "DB2_EXTENDED_OPTIMIZATION": "'-'", "DB2_INDEX_PCTFREE_DEFAULT": "'range(0, 99)'", "DB2_INLIST_TO_NLJN": "'NO', 'YES'", "DB2_MINIMIZE_LISTPREFETCH": "'NO', 'YES'", "DB2_OBJECT_TABLE_ENTRIES": "'range(0, 65532)'", "DB2_OPTPROFILE": "'NO', 'YES'", "DB2_OPTSTATS_LOG": "'-'", "DB2_OPT_MAX_TEMP_SIZE": "'-'", "DB2_PARALLEL_IO": "'-'", "DB2_REDUCED_OPTIMIZATION": "'-'", "DB2_SELECTIVITY": "'YES', 'NO', 'ALL'", "DB2_SKIPDELETED": "'NO', 'YES'", "DB2_SKIPINSERTED": "'NO', 'YES'", "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES": "'NO', 'YES'", "DB2_TRUNCATE_REUSESTORAGE": "'IMPORT', 'LOAD', 'TRUNCATE'", "DB2_USE_ALTERNATE_PAGE_CLEANING": "'ON', 'OFF'", "DB2_VIEW_REOPT_VALUES": "'NO', 'YES'", "DB2_WLM_SETTINGS": "'-'", "DB2_WORKLOAD": "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"}}}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasTuneableParam successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 getDb2SaasTuneableParamOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasTuneableParamWithContext(ctx, getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasTuneableParamWithContext(ctx, getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"tuneable_param": {"db": {"ACT_SORTMEM_LIMIT": "'NONE', 'range(10, 100)'", "ALT_COLLATE": "'NULL', 'IDENTITY_16BIT'", "APPGROUP_MEM_SZ": "'range(1, 1000000)'", "APPLHEAPSZ": "'AUTOMATIC', 'range(16, 2147483647)'", "APPL_MEMORY": "'AUTOMATIC', 'range(128, 4294967295)'", "APP_CTL_HEAP_SZ": "'range(1, 64000)'", "ARCHRETRYDELAY": "'range(0, 65535)'", "AUTHN_CACHE_DURATION": "'range(1,10000)'", "AUTORESTART": "'ON', 'OFF'", "AUTO_CG_STATS": "'ON', 'OFF'", "AUTO_MAINT": "'ON', 'OFF'", "AUTO_REORG": "'ON', 'OFF'", "AUTO_REVAL": "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'", "AUTO_RUNSTATS": "'ON', 'OFF'", "AUTO_SAMPLING": "'ON', 'OFF'", "AUTO_STATS_VIEWS": "'ON', 'OFF'", "AUTO_STMT_STATS": "'ON', 'OFF'", "AUTO_TBL_MAINT": "'ON', 'OFF'", "AVG_APPLS": "'-'", "CATALOGCACHE_SZ": "'-'", "CHNGPGS_THRESH": "'range(5,99)'", "CUR_COMMIT": "'ON, AVAILABLE, DISABLED'", "DATABASE_MEMORY": "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'", "DBHEAP": "'AUTOMATIC', 'range(32, 2147483647)'", "DB_COLLNAME": "'-'", "DB_MEM_THRESH": "'range(0, 100)'", "DDL_COMPRESSION_DEF": "'YES', 'NO'", "DDL_CONSTRAINT_DEF": "'YES', 'NO'", "DECFLT_ROUNDING": "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'", "DEC_ARITHMETIC": "'-'", "DEC_TO_CHAR_FMT": "'NEW', 'V95'", "DFT_DEGREE": "'-1', 'ANY', 'range(1, 32767)'", "DFT_EXTENT_SZ": "'range(2, 256)'", "DFT_LOADREC_SES": "'range(1, 30000)'", "DFT_MTTB_TYPES": "'-'", "DFT_PREFETCH_SZ": "'range(0, 32767)', 'AUTOMATIC'", "DFT_QUERYOPT": "'range(0, 9)'", "DFT_REFRESH_AGE": "'-'", "DFT_SCHEMAS_DCC": "'YES', 'NO'", "DFT_SQLMATHWARN": "'YES', 'NO'", "DFT_TABLE_ORG": "'COLUMN', 'ROW'", "DLCHKTIME": "'range(1000, 600000)'", "ENABLE_XMLCHAR": "'YES', 'NO'", "EXTENDED_ROW_SZ": "'ENABLE', 'DISABLE'", "GROUPHEAP_RATIO": "'range(1, 99)'", "INDEXREC": "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'", "LARGE_AGGREGATION": "'YES', 'NO'", "LOCKLIST": "'AUTOMATIC', 'range(4, 134217728)'", "LOCKTIMEOUT": "'-1', 'range(0, 32767)'", "LOGINDEXBUILD": "'ON', 'OFF'", "LOG_APPL_INFO": "'YES', 'NO'", "LOG_DDL_STMTS": "'YES', 'NO'", "LOG_DISK_CAP": "'0', '-1', 'range(1, 2147483647)'", "MAXAPPLS": "'range(1, 60000)'", "MAXFILOP": "'range(64, 61440)'", "MAXLOCKS": "'AUTOMATIC', 'range(1, 100)'", "MIN_DEC_DIV_3": "'YES', 'NO'", "MON_ACT_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_DEADLOCK": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LCK_MSG_LVL": "'range(0, 3)'", "MON_LOCKTIMEOUT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LOCKWAIT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LW_THRESH": "'range(1000, 4294967295)'", "MON_OBJ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_PKGLIST_SZ": "'range(0, 1024)'", "MON_REQ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_RTN_DATA": "'NONE', 'BASE'", "MON_RTN_EXECLIST": "'OFF', 'ON'", "MON_UOW_DATA": "'NONE', 'BASE'", "MON_UOW_EXECLIST": "'ON', 'OFF'", "MON_UOW_PKGLIST": "'OFF', 'ON'", "NCHAR_MAPPING": "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'", "NUM_FREQVALUES": "'range(0, 32767)'", "NUM_IOCLEANERS": "'AUTOMATIC', 'range(0, 255)'", "NUM_IOSERVERS": "'AUTOMATIC', 'range(1, 255)'", "NUM_LOG_SPAN": "'range(0, 65535)'", "NUM_QUANTILES": "'range(0, 32767)'", "OPT_BUFFPAGE": "'-'", "OPT_DIRECT_WRKLD": "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'", "OPT_LOCKLIST": "'-'", "OPT_MAXLOCKS": "'-'", "OPT_SORTHEAP": "'-'", "PAGE_AGE_TRGT_GCR": "'range(1, 65535)'", "PAGE_AGE_TRGT_MCR": "'range(1, 65535)'", "PCKCACHESZ": "'AUTOMATIC', '-1', 'range(32, 2147483646)'", "PL_STACK_TRACE": "'NONE', 'ALL', 'UNHANDLED'", "SELF_TUNING_MEM": "'ON', 'OFF'", "SEQDETECT": "'YES', 'NO'", "SHEAPTHRES_SHR": "'AUTOMATIC', 'range(250, 2147483647)'", "SOFTMAX": "'-'", "SORTHEAP": "'AUTOMATIC', 'range(16, 4294967295)'", "SQL_CCFLAGS": "'-'", "STAT_HEAP_SZ": "'AUTOMATIC', 'range(1096, 2147483647)'", "STMTHEAP": "'AUTOMATIC', 'range(128, 2147483647)'", "STMT_CONC": "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'", "STRING_UNITS": "'SYSTEM', 'CODEUNITS32'", "SYSTIME_PERIOD_ADJ": "'NO', 'YES'", "TRACKMOD": "'YES', 'NO'", "UTIL_HEAP_SZ": "'AUTOMATIC', 'range(16, 2147483647)'", "WLM_ADMISSION_CTRL": "'YES', 'NO'", "WLM_AGENT_LOAD_TRGT": "'AUTOMATIC', 'range(1, 65535)'", "WLM_CPU_LIMIT": "'range(0, 100)'", "WLM_CPU_SHARES": "'range(1, 65535)'", "WLM_CPU_SHARE_MODE": "'HARD', 'SOFT'"}, "dbm": {"COMM_BANDWIDTH": "'range(0.1, 100000)', '-1'", "CPUSPEED": "'range(0.0000000001, 1)', '-1'", "DFT_MON_BUFPOOL": "'ON', 'OFF'", "DFT_MON_LOCK": "'ON', 'OFF'", "DFT_MON_SORT": "'ON', 'OFF'", "DFT_MON_STMT": "'ON', 'OFF'", "DFT_MON_TABLE": "'ON', 'OFF'", "DFT_MON_TIMESTAMP": "'ON', 'OFF'", "DFT_MON_UOW": "'ON', 'OFF'", "DIAGLEVEL": "'range(0, 4)'", "FEDERATED_ASYNC": "'range(0, 32767)', '-1', 'ANY'", "INDEXREC": "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'", "INTRA_PARALLEL": "'SYSTEM', 'NO', 'YES'", "KEEPFENCED": "'YES', 'NO'", "MAX_CONNRETRIES": "'range(0, 100)'", "MAX_QUERYDEGREE": "'range(1, 32767)', '-1', 'ANY'", "MON_HEAP_SZ": "'range(0, 2147483647)', 'AUTOMATIC'", "MULTIPARTSIZEMB": "'range(5, 5120)'", "NOTIFYLEVEL": "'range(0, 4)'", "NUM_INITAGENTS": "'range(0, 64000)'", "NUM_INITFENCED": "'range(0, 64000)'", "NUM_POOLAGENTS": "'-1', 'range(0, 64000)'", "RESYNC_INTERVAL": "'range(1, 60000)'", "RQRIOBLK": "'range(4096, 65535)'", "START_STOP_TIME": "'range(1, 1440)'", "UTIL_IMPACT_LIM": "'range(1, 100)'", "WLM_DISPATCHER": "'YES', 'NO'", "WLM_DISP_CONCUR": "'range(1, 32767)', 'COMPUTED'", "WLM_DISP_CPU_SHARES": "'NO', 'YES'", "WLM_DISP_MIN_UTIL": "'range(0, 100)'"}, "registry": {"DB2BIDI": "'YES', 'NO'", "DB2COMPOPT": "'-'", "DB2LOCK_TO_RB": "'STATEMENT'", "DB2STMM": "'NO', 'YES'", "DB2_ALTERNATE_AUTHZ_BEHAVIOUR": "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'", "DB2_ANTIJOIN": "'YES', 'NO', 'EXTEND'", "DB2_ATS_ENABLE": "'YES', 'NO'", "DB2_DEFERRED_PREPARE_SEMANTICS": "'NO', 'YES'", "DB2_EVALUNCOMMITTED": "'NO', 'YES'", "DB2_EXTENDED_OPTIMIZATION": "'-'", "DB2_INDEX_PCTFREE_DEFAULT": "'range(0, 99)'", "DB2_INLIST_TO_NLJN": "'NO', 'YES'", "DB2_MINIMIZE_LISTPREFETCH": "'NO', 'YES'", "DB2_OBJECT_TABLE_ENTRIES": "'range(0, 65532)'", "DB2_OPTPROFILE": "'NO', 'YES'", "DB2_OPTSTATS_LOG": "'-'", "DB2_OPT_MAX_TEMP_SIZE": "'-'", "DB2_PARALLEL_IO": "'-'", "DB2_REDUCED_OPTIMIZATION": "'-'", "DB2_SELECTIVITY": "'YES', 'NO', 'ALL'", "DB2_SKIPDELETED": "'NO', 'YES'", "DB2_SKIPINSERTED": "'NO', 'YES'", "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES": "'NO', 'YES'", "DB2_TRUNCATE_REUSESTORAGE": "'IMPORT', 'LOAD', 'TRUNCATE'", "DB2_USE_ALTERNATE_PAGE_CLEANING": "'ON', 'OFF'", "DB2_VIEW_REOPT_VALUES": "'NO', 'YES'", "DB2_WLM_SETTINGS": "'-'", "DB2_WORKLOAD": "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"}}}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasTuneableParam successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 getDb2SaasTuneableParamOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasTuneableParam with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 getDb2SaasTuneableParamOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasTuneableParamOptions model with no property values
				 getDb2SaasTuneableParamOptionsModelNew := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasTuneableParam successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				 getDb2SaasTuneableParamOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions) - Operation response error`, func() {
		 getDb2SaasBackupPath := "/manage/backups"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("GET"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke GetDb2SaasBackup with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				 getDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions)`, func() {
		 getDb2SaasBackupPath := "/manage/backups"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"backups": [{"id": "ID", "type": "Type", "status": "Status", "created_at": "CreatedAt", "size": 4, "duration": 8}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasBackup successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				 getDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.GetDb2SaasBackupWithContext(ctx, getDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.GetDb2SaasBackupWithContext(ctx, getDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("GET"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"backups": [{"id": "ID", "type": "Type", "status": "Status", "created_at": "CreatedAt", "size": 4, "duration": 8}]}`)
				 }))
			 })
			 It(`Invoke GetDb2SaasBackup successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.GetDb2SaasBackup(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				 getDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke GetDb2SaasBackup with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				 getDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the GetDb2SaasBackupOptions model with no property values
				 getDb2SaasBackupOptionsModelNew := new(db2saasv1.GetDb2SaasBackupOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke GetDb2SaasBackup successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				 getDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions) - Operation response error`, func() {
		 postDb2SaasBackupPath := "/manage/backups/backup"
		 Context(`Using mock server endpoint with invalid JSON response`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("POST"))
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprint(res, `} this is not valid json {`)
				 }))
			 })
			 It(`Invoke PostDb2SaasBackup with error: Operation response processing error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				 postDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Expect response parsing to fail since we are receiving a text/plain response
				 result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
 
				 // Enable retries and test again
				 db2saasService.EnableRetries(0, 0)
				 result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions)`, func() {
		 postDb2SaasBackupPath := "/manage/backups/backup"
		 Context(`Using mock server endpoint with timeout`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Sleep a short time to support a timeout test
					 time.Sleep(100 * time.Millisecond)
 
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"task": {"id": "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:0c9c7889-54de-4ecc-8399-09a4d4ff228e:task:51ff2dc7-6cb9-41c0-9345-09e54550fb7b"}}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasBackup successfully with retries`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
				 db2saasService.EnableRetries(0, 0)
 
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				 postDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with a Context to test a timeout error
				 ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc()
				 _, _, operationErr := db2saasService.PostDb2SaasBackupWithContext(ctx, postDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
 
				 // Disable retries and test again
				 db2saasService.DisableRetries()
				 result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
				 // Re-test the timeout error with retries disabled
				 ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				 defer cancelFunc2()
				 _, _, operationErr = db2saasService.PostDb2SaasBackupWithContext(ctx, postDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Verify the contents of the request
					 Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					 Expect(req.Method).To(Equal("POST"))
 
					 Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					 Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					 // Set mock response
					 res.Header().Set("Content-type", "application/json")
					 res.WriteHeader(200)
					 fmt.Fprintf(res, "%s", `{"task": {"id": "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:0c9c7889-54de-4ecc-8399-09a4d4ff228e:task:51ff2dc7-6cb9-41c0-9345-09e54550fb7b"}}`)
				 }))
			 })
			 It(`Invoke PostDb2SaasBackup successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Invoke operation with nil options model (negative test)
				 result, response, operationErr := db2saasService.PostDb2SaasBackup(nil)
				 Expect(operationErr).NotTo(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
 
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				 postDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation with valid options model (positive test)
				 result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
				 Expect(result).ToNot(BeNil())
 
			 })
			 It(`Invoke PostDb2SaasBackup with error: Operation validation and request error`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				 postDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				 // Invoke operation with empty URL (negative test)
				 err := db2saasService.SetServiceURL("")
				 Expect(err).To(BeNil())
				 result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
				 // Construct a second instance of the PostDb2SaasBackupOptions model with no property values
				 postDb2SaasBackupOptionsModelNew := new(db2saasv1.PostDb2SaasBackupOptions)
				 // Invoke operation with invalid model (negative test)
				 result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModelNew)
				 Expect(operationErr).ToNot(BeNil())
				 Expect(response).To(BeNil())
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
		 Context(`Using mock server endpoint with missing response body`, func() {
			 BeforeEach(func() {
				 testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					 defer GinkgoRecover()
 
					 // Set success status code with no respoonse body
					 res.WriteHeader(200)
				 }))
			 })
			 It(`Invoke PostDb2SaasBackup successfully`, func() {
				 db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					 URL:           testServer.URL,
					 Authenticator: &core.NoAuthAuthenticator{},
				 })
				 Expect(serviceErr).To(BeNil())
				 Expect(db2saasService).ToNot(BeNil())
 
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				 postDb2SaasBackupOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
 
				 // Invoke operation
				 result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				 Expect(operationErr).To(BeNil())
				 Expect(response).ToNot(BeNil())
 
				 // Verify a nil result
				 Expect(result).To(BeNil())
			 })
			 AfterEach(func() {
				 testServer.Close()
			 })
		 })
	 })
	 Describe(`Model constructor tests`, func() {
		 Context(`Using a service client instance`, func() {
			 db2saasService, _ := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				 URL:           "http://db2saasv1modelgenerator.com",
				 Authenticator: &core.NoAuthAuthenticator{},
			 })
			 It(`Invoke NewCreateUserAuthentication successfully`, func() {
				 method := "internal"
				 policyID := "Default"
				 _model, err := db2saasService.NewCreateUserAuthentication(method, policyID)
				 Expect(_model).ToNot(BeNil())
				 Expect(err).To(BeNil())
			 })
			 It(`Invoke NewDeleteDb2SaasUserOptions successfully`, func() {
				 // Construct an instance of the DeleteDb2SaasUserOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 id := "test-user"
				 deleteDb2SaasUserOptionsModel := db2saasService.NewDeleteDb2SaasUserOptions(xDeploymentID, id)
				 deleteDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 deleteDb2SaasUserOptionsModel.SetID("test-user")
				 deleteDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(deleteDb2SaasUserOptionsModel).ToNot(BeNil())
				 Expect(deleteDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(deleteDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				 Expect(deleteDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasAllowlistOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasAllowlistOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getDb2SaasAllowlistOptionsModel := db2saasService.NewGetDb2SaasAllowlistOptions(xDeploymentID)
				 getDb2SaasAllowlistOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasAllowlistOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasAllowlistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getDb2SaasAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasAutoscaleOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasAutoscaleOptions model
				 xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				 getDb2SaasAutoscaleOptionsModel := db2saasService.NewGetDb2SaasAutoscaleOptions(xDbProfile)
				 getDb2SaasAutoscaleOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 getDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasAutoscaleOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				 Expect(getDb2SaasAutoscaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasBackupOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasBackupOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getDb2SaasBackupOptionsModel := db2saasService.NewGetDb2SaasBackupOptions(xDeploymentID)
				 getDb2SaasBackupOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasBackupOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasBackupOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getDb2SaasBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasConnectionInfoOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasConnectionInfoOptions model
				 deploymentID := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A"
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getDb2SaasConnectionInfoOptionsModel := db2saasService.NewGetDb2SaasConnectionInfoOptions(deploymentID, xDeploymentID)
				 getDb2SaasConnectionInfoOptionsModel.SetDeploymentID("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				 getDb2SaasConnectionInfoOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasConnectionInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasConnectionInfoOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasConnectionInfoOptionsModel.DeploymentID).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")))
				 Expect(getDb2SaasConnectionInfoOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getDb2SaasConnectionInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasTuneableParamOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasTuneableParamOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getDb2SaasTuneableParamOptionsModel := db2saasService.NewGetDb2SaasTuneableParamOptions(xDeploymentID)
				 getDb2SaasTuneableParamOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasTuneableParamOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasTuneableParamOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasTuneableParamOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getDb2SaasTuneableParamOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetDb2SaasUserOptions successfully`, func() {
				 // Construct an instance of the GetDb2SaasUserOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getDb2SaasUserOptionsModel := db2saasService.NewGetDb2SaasUserOptions(xDeploymentID)
				 getDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getDb2SaasUserOptionsModel).ToNot(BeNil())
				 Expect(getDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewGetbyidDb2SaasUserOptions successfully`, func() {
				 // Construct an instance of the GetbyidDb2SaasUserOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 getbyidDb2SaasUserOptionsModel := db2saasService.NewGetbyidDb2SaasUserOptions(xDeploymentID)
				 getbyidDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 getbyidDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(getbyidDb2SaasUserOptionsModel).ToNot(BeNil())
				 Expect(getbyidDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(getbyidDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewIpAddress successfully`, func() {
				 address := "127.0.0.1"
				 description := "A sample IP address"
				 _model, err := db2saasService.NewIpAddress(address, description)
				 Expect(_model).ToNot(BeNil())
				 Expect(err).To(BeNil())
			 })
			 It(`Invoke NewPostDb2SaasAllowlistOptions successfully`, func() {
				 // Construct an instance of the IpAddress model
				 ipAddressModel := new(db2saasv1.IpAddress)
				 Expect(ipAddressModel).ToNot(BeNil())
				 ipAddressModel.Address = core.StringPtr("127.0.0.1")
				 ipAddressModel.Description = core.StringPtr("A sample IP address")
				 Expect(ipAddressModel.Address).To(Equal(core.StringPtr("127.0.0.1")))
				 Expect(ipAddressModel.Description).To(Equal(core.StringPtr("A sample IP address")))
 
				 // Construct an instance of the PostDb2SaasAllowlistOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 postDb2SaasAllowlistOptionsIpAddresses := []db2saasv1.IpAddress{}
				 postDb2SaasAllowlistOptionsModel := db2saasService.NewPostDb2SaasAllowlistOptions(xDeploymentID, postDb2SaasAllowlistOptionsIpAddresses)
				 postDb2SaasAllowlistOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasAllowlistOptionsModel.SetIpAddresses([]db2saasv1.IpAddress{*ipAddressModel})
				 postDb2SaasAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(postDb2SaasAllowlistOptionsModel).ToNot(BeNil())
				 Expect(postDb2SaasAllowlistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(postDb2SaasAllowlistOptionsModel.IpAddresses).To(Equal([]db2saasv1.IpAddress{*ipAddressModel}))
				 Expect(postDb2SaasAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewPostDb2SaasBackupOptions successfully`, func() {
				 // Construct an instance of the PostDb2SaasBackupOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 postDb2SaasBackupOptionsModel := db2saasService.NewPostDb2SaasBackupOptions(xDeploymentID)
				 postDb2SaasBackupOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(postDb2SaasBackupOptionsModel).ToNot(BeNil())
				 Expect(postDb2SaasBackupOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(postDb2SaasBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewPostDb2SaasDbConfigurationOptions successfully`, func() {
				 // Construct an instance of the CreateCustomSettingsRegistry model
				 createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				 Expect(createCustomSettingsRegistryModel).ToNot(BeNil())
				 createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("ALL")
				 createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				 Expect(createCustomSettingsRegistryModel.DB2SELECTIVITY).To(Equal(core.StringPtr("ALL")))
				 Expect(createCustomSettingsRegistryModel.DB2ANTIJOIN).To(Equal(core.StringPtr("EXTEND")))
 
				 // Construct an instance of the CreateCustomSettingsDb model
				 createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				 Expect(createCustomSettingsDbModel).ToNot(BeNil())
				 createCustomSettingsDbModel.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
				 Expect(createCustomSettingsDbModel.SORTHEAP).To(Equal(core.StringPtr("50000 AUTOMATIC")))
 
				 // Construct an instance of the PostDb2SaasDbConfigurationOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 postDb2SaasDbConfigurationOptionsModel := db2saasService.NewPostDb2SaasDbConfigurationOptions(xDeploymentID)
				 postDb2SaasDbConfigurationOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasDbConfigurationOptionsModel.SetRegistry(createCustomSettingsRegistryModel)
				 postDb2SaasDbConfigurationOptionsModel.SetDb(createCustomSettingsDbModel)
				 postDb2SaasDbConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(postDb2SaasDbConfigurationOptionsModel).ToNot(BeNil())
				 Expect(postDb2SaasDbConfigurationOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(postDb2SaasDbConfigurationOptionsModel.Registry).To(Equal(createCustomSettingsRegistryModel))
				 Expect(postDb2SaasDbConfigurationOptionsModel.Db).To(Equal(createCustomSettingsDbModel))
				 Expect(postDb2SaasDbConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewPostDb2SaasUserOptions successfully`, func() {
				 // Construct an instance of the CreateUserAuthentication model
				 createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				 Expect(createUserAuthenticationModel).ToNot(BeNil())
				 createUserAuthenticationModel.Method = core.StringPtr("internal")
				 createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
				 Expect(createUserAuthenticationModel.Method).To(Equal(core.StringPtr("internal")))
				 Expect(createUserAuthenticationModel.PolicyID).To(Equal(core.StringPtr("Default")))
 
				 // Construct an instance of the PostDb2SaasUserOptions model
				 xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				 postDb2SaasUserOptionsID := "test-user"
				 postDb2SaasUserOptionsIam := false
				 postDb2SaasUserOptionsIbmid := "test-ibm-id"
				 postDb2SaasUserOptionsName := "test_user"
				 postDb2SaasUserOptionsPassword := "dEkMc43@gfAPl!867^dSbu"
				 postDb2SaasUserOptionsRole := "bluuser"
				 postDb2SaasUserOptionsEmail := "test_user@mycompany.com"
				 postDb2SaasUserOptionsLocked := "no"
				 var postDb2SaasUserOptionsAuthentication *db2saasv1.CreateUserAuthentication = nil
				 postDb2SaasUserOptionsModel := db2saasService.NewPostDb2SaasUserOptions(xDeploymentID, postDb2SaasUserOptionsID, postDb2SaasUserOptionsIam, postDb2SaasUserOptionsIbmid, postDb2SaasUserOptionsName, postDb2SaasUserOptionsPassword, postDb2SaasUserOptionsRole, postDb2SaasUserOptionsEmail, postDb2SaasUserOptionsLocked, postDb2SaasUserOptionsAuthentication)
				 postDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				 postDb2SaasUserOptionsModel.SetID("test-user")
				 postDb2SaasUserOptionsModel.SetIam(false)
				 postDb2SaasUserOptionsModel.SetIbmid("test-ibm-id")
				 postDb2SaasUserOptionsModel.SetName("test_user")
				 postDb2SaasUserOptionsModel.SetPassword("dEkMc43@gfAPl!867^dSbu")
				 postDb2SaasUserOptionsModel.SetRole("bluuser")
				 postDb2SaasUserOptionsModel.SetEmail("test_user@mycompany.com")
				 postDb2SaasUserOptionsModel.SetLocked("no")
				 postDb2SaasUserOptionsModel.SetAuthentication(createUserAuthenticationModel)
				 postDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(postDb2SaasUserOptionsModel).ToNot(BeNil())
				 Expect(postDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				 Expect(postDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				 Expect(postDb2SaasUserOptionsModel.Iam).To(Equal(core.BoolPtr(false)))
				 Expect(postDb2SaasUserOptionsModel.Ibmid).To(Equal(core.StringPtr("test-ibm-id")))
				 Expect(postDb2SaasUserOptionsModel.Name).To(Equal(core.StringPtr("test_user")))
				 Expect(postDb2SaasUserOptionsModel.Password).To(Equal(core.StringPtr("dEkMc43@gfAPl!867^dSbu")))
				 Expect(postDb2SaasUserOptionsModel.Role).To(Equal(core.StringPtr("bluuser")))
				 Expect(postDb2SaasUserOptionsModel.Email).To(Equal(core.StringPtr("test_user@mycompany.com")))
				 Expect(postDb2SaasUserOptionsModel.Locked).To(Equal(core.StringPtr("no")))
				 Expect(postDb2SaasUserOptionsModel.Authentication).To(Equal(createUserAuthenticationModel))
				 Expect(postDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
			 It(`Invoke NewPutDb2SaasAutoscaleOptions successfully`, func() {
				 // Construct an instance of the PutDb2SaasAutoscaleOptions model
				 xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				 putDb2SaasAutoscaleOptionsModel := db2saasService.NewPutDb2SaasAutoscaleOptions(xDbProfile)
				 putDb2SaasAutoscaleOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				 putDb2SaasAutoscaleOptionsModel.SetAutoScalingEnabled("true")
				 putDb2SaasAutoscaleOptionsModel.SetAutoScalingThreshold(int64(90))
				 putDb2SaasAutoscaleOptionsModel.SetAutoScalingOverTimePeriod(float64(5))
				 putDb2SaasAutoscaleOptionsModel.SetAutoScalingPauseLimit(int64(70))
				 putDb2SaasAutoscaleOptionsModel.SetAutoScalingAllowPlanLimit("YES")
				 putDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				 Expect(putDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				 Expect(putDb2SaasAutoscaleOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				 Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled).To(Equal(core.StringPtr("true")))
				 Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold).To(Equal(core.Int64Ptr(int64(90))))
				 Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod).To(Equal(core.Float64Ptr(float64(5))))
				 Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit).To(Equal(core.Int64Ptr(int64(70))))
				 Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit).To(Equal(core.StringPtr("YES")))
				 Expect(putDb2SaasAutoscaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			 })
		 })
	 })
	 Describe(`Model unmarshaling tests`, func() {
		 It(`Invoke UnmarshalCreateCustomSettingsDb successfully`, func() {
			 // Construct an instance of the model.
			 model := new(db2saasv1.CreateCustomSettingsDb)
			 model.SORTHEAP = core.StringPtr("50000 AUTOMATIC")
 
			 b, err := json.Marshal(model)
			 Expect(err).To(BeNil())
 
			 var raw map[string]json.RawMessage
			 err = json.Unmarshal(b, &raw)
			 Expect(err).To(BeNil())
 
			 var result *db2saasv1.CreateCustomSettingsDb
			 err = db2saasv1.UnmarshalCreateCustomSettingsDb(raw, &result)
			 Expect(err).To(BeNil())
			 Expect(result).ToNot(BeNil())
			 Expect(result).To(Equal(model))
		 })
		 It(`Invoke UnmarshalCreateCustomSettingsRegistry successfully`, func() {
			 // Construct an instance of the model.
			 model := new(db2saasv1.CreateCustomSettingsRegistry)
			 model.DB2SELECTIVITY = core.StringPtr("ALL")
			 model.DB2ANTIJOIN = core.StringPtr("EXTEND")
 
			 b, err := json.Marshal(model)
			 Expect(err).To(BeNil())
 
			 var raw map[string]json.RawMessage
			 err = json.Unmarshal(b, &raw)
			 Expect(err).To(BeNil())
 
			 var result *db2saasv1.CreateCustomSettingsRegistry
			 err = db2saasv1.UnmarshalCreateCustomSettingsRegistry(raw, &result)
			 Expect(err).To(BeNil())
			 Expect(result).ToNot(BeNil())
			 Expect(result).To(Equal(model))
		 })
		 It(`Invoke UnmarshalCreateUserAuthentication successfully`, func() {
			 // Construct an instance of the model.
			 model := new(db2saasv1.CreateUserAuthentication)
			 model.Method = core.StringPtr("internal")
			 model.PolicyID = core.StringPtr("Default")
 
			 b, err := json.Marshal(model)
			 Expect(err).To(BeNil())
 
			 var raw map[string]json.RawMessage
			 err = json.Unmarshal(b, &raw)
			 Expect(err).To(BeNil())
 
			 var result *db2saasv1.CreateUserAuthentication
			 err = db2saasv1.UnmarshalCreateUserAuthentication(raw, &result)
			 Expect(err).To(BeNil())
			 Expect(result).ToNot(BeNil())
			 Expect(result).To(Equal(model))
		 })
		 It(`Invoke UnmarshalIpAddress successfully`, func() {
			 // Construct an instance of the model.
			 model := new(db2saasv1.IpAddress)
			 model.Address = core.StringPtr("127.0.0.1")
			 model.Description = core.StringPtr("A sample IP address")
 
			 b, err := json.Marshal(model)
			 Expect(err).To(BeNil())
 
			 var raw map[string]json.RawMessage
			 err = json.Unmarshal(b, &raw)
			 Expect(err).To(BeNil())
 
			 var result *db2saasv1.IpAddress
			 err = db2saasv1.UnmarshalIpAddress(raw, &result)
			 Expect(err).To(BeNil())
			 Expect(result).ToNot(BeNil())
			 Expect(result).To(Equal(model))
		 })
	 })
	 Describe(`Utility function tests`, func() {
		 It(`Invoke CreateMockByteArray() successfully`, func() {
			 mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			 Expect(mockByteArray).ToNot(BeNil())
		 })
		 It(`Invoke CreateMockUUID() successfully`, func() {
			 mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			 Expect(mockUUID).ToNot(BeNil())
		 })
		 It(`Invoke CreateMockReader() successfully`, func() {
			 mockReader := CreateMockReader("This is a test.")
			 Expect(mockReader).ToNot(BeNil())
		 })
		 It(`Invoke CreateMockDate() successfully`, func() {
			 mockDate := CreateMockDate("2019-01-01")
			 Expect(mockDate).ToNot(BeNil())
		 })
		 It(`Invoke CreateMockDateTime() successfully`, func() {
			 mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			 Expect(mockDateTime).ToNot(BeNil())
		 })
	 })
 })
 
 //
 // Utility functions used by the generated test code
 //
 
 func CreateMockByteArray(encodedString string) *[]byte {
	 ba, err := base64.StdEncoding.DecodeString(encodedString)
	 if err != nil {
		 panic(err)
	 }
	 return &ba
 }
 
 func CreateMockUUID(mockData string) *strfmt.UUID {
	 uuid := strfmt.UUID(mockData)
	 return &uuid
 }
 
 func CreateMockReader(mockData string) io.ReadCloser {
	 return io.NopCloser(bytes.NewReader([]byte(mockData)))
 }
 
 func CreateMockDate(mockData string) *strfmt.Date {
	 d, err := core.ParseDate(mockData)
	 if err != nil {
		 return nil
	 }
	 return &d
 }
 
 func CreateMockDateTime(mockData string) *strfmt.DateTime {
	 d, err := core.ParseDateTime(mockData)
	 if err != nil {
		 return nil
	 }
	 return &d
 }
 
 func SetTestEnvironment(testEnvironment map[string]string) {
	 for key, value := range testEnvironment {
		 os.Setenv(key, value)
	 }
 }
 
 func ClearTestEnvironment(testEnvironment map[string]string) {
	 for key := range testEnvironment {
		 os.Unsetenv(key)
	 }
 }
 