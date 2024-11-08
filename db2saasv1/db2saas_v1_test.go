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

package db2saasv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
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
		getDb2SaasConnectionInfoPath := "/connectioninfo/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Id"]).ToNot(BeNil())
					Expect(req.Header["Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.ID = core.StringPtr("testString")
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
		getDb2SaasConnectionInfoPath := "/connectioninfo/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Id"]).ToNot(BeNil())
					Expect(req.Header["Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "Hostname", "databaseName": "DatabaseName", "host_ros": "HostRos", "certificateBase64": "CertificateBase64", "sslPort": "SslPort", "ssl": false, "databaseVersion": "DatabaseVersion"}, "private": {"hostname": "Hostname", "databaseName": "DatabaseName", "host_ros": "HostRos", "certificateBase64": "CertificateBase64", "sslPort": "SslPort", "ssl": false, "databaseVersion": "DatabaseVersion"}}`)
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.ID = core.StringPtr("testString")
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

					Expect(req.Header["Id"]).ToNot(BeNil())
					Expect(req.Header["Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "Hostname", "databaseName": "DatabaseName", "host_ros": "HostRos", "certificateBase64": "CertificateBase64", "sslPort": "SslPort", "ssl": false, "databaseVersion": "DatabaseVersion"}, "private": {"hostname": "Hostname", "databaseName": "DatabaseName", "host_ros": "HostRos", "certificateBase64": "CertificateBase64", "sslPort": "SslPort", "ssl": false, "databaseVersion": "DatabaseVersion"}}`)
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.ID = core.StringPtr("testString")
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.ID = core.StringPtr("testString")
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.ID = core.StringPtr("testString")
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			db2saasService, _ := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				URL:           "http://db2saasv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetDb2SaasConnectionInfoOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				deploymentID := "testString"
				id := "testString"
				getDb2SaasConnectionInfoOptionsModel := db2saasService.NewGetDb2SaasConnectionInfoOptions(deploymentID, id)
				getDb2SaasConnectionInfoOptionsModel.SetDeploymentID("testString")
				getDb2SaasConnectionInfoOptionsModel.SetID("testString")
				getDb2SaasConnectionInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasConnectionInfoOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasConnectionInfoOptionsModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasConnectionInfoOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasConnectionInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
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
