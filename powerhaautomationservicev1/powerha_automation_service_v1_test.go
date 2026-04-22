/**
 * (C) Copyright IBM Corp. 2026.
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

package powerhaautomationservicev1_test

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

	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PowerhaAutomationServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(powerhaAutomationServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(powerhaAutomationServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
				URL: "https://powerhaautomationservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(powerhaAutomationServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POWERHA_AUTOMATION_SERVICE_URL":       "https://powerhaautomationservicev1/api",
				"POWERHA_AUTOMATION_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{})
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := powerhaAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != powerhaAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(powerhaAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(powerhaAutomationServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := powerhaAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != powerhaAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(powerhaAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(powerhaAutomationServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{})
				err := powerhaAutomationServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := powerhaAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != powerhaAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(powerhaAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(powerhaAutomationServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POWERHA_AUTOMATION_SERVICE_URL":       "https://powerhaautomationservicev1/api",
				"POWERHA_AUTOMATION_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(powerhaAutomationServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"POWERHA_AUTOMATION_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(powerhaAutomationServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = powerhaautomationservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions) - Operation response error`, func() {
		getAPIKeyPath := "/powerha_automation/v1/api_key/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAPIKey with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(powerhaautomationservicev1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getAPIKeyOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions)`, func() {
		getAPIKeyPath := "/powerha_automation/v1/api_key/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Success", "id": "9676767890"}`)
				}))
			})
			It(`Invoke GetAPIKey successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(powerhaautomationservicev1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getAPIKeyOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetAPIKeyWithContext(ctx, getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetAPIKeyWithContext(ctx, getAPIKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAPIKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Success", "id": "9676767890"}`)
				}))
			})
			It(`Invoke GetAPIKey successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(powerhaautomationservicev1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getAPIKeyOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAPIKey with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(powerhaautomationservicev1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getAPIKeyOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAPIKeyOptions model with no property values
				getAPIKeyOptionsModelNew := new(powerhaautomationservicev1.GetAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModelNew)
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
			It(`Invoke GetAPIKey successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetAPIKeyOptions model
				getAPIKeyOptionsModel := new(powerhaautomationservicev1.GetAPIKeyOptions)
				getAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getAPIKeyOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptionsModel)
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
	Describe(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions) - Operation response error`, func() {
		createAPIKeyPath := "/powerha_automation/v1/api_key/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAPIKey with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {
		createAPIKeyPath := "/powerha_automation/v1/api_key/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"status": "Success", "id": "9676767890"}`)
				}))
			})
			It(`Invoke CreateAPIKey successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.CreateAPIKeyWithContext(ctx, createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.CreateAPIKeyWithContext(ctx, createAPIKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAPIKeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"status": "Success", "id": "9676767890"}`)
				}))
			})
			It(`Invoke CreateAPIKey successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.CreateAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAPIKey with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAPIKeyOptions model with no property values
				createAPIKeyOptionsModelNew := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateAPIKey successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateAPIKeyOptions model
				createAPIKeyOptionsModel := new(powerhaautomationservicev1.CreateAPIKeyOptions)
				createAPIKeyOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createAPIKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptionsModel)
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
	Describe(`GetClusterNode(getClusterNodeOptions *GetClusterNodeOptions) - Operation response error`, func() {
		getClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterNodePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetClusterNode with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetClusterNodeOptions model
				getClusterNodeOptionsModel := new(powerhaautomationservicev1.GetClusterNodeOptions)
				getClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetClusterNode(getClusterNodeOptions *GetClusterNodeOptions)`, func() {
		getClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterNodePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke GetClusterNode successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetClusterNodeOptions model
				getClusterNodeOptionsModel := new(powerhaautomationservicev1.GetClusterNodeOptions)
				getClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetClusterNodeWithContext(ctx, getClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetClusterNodeWithContext(ctx, getClusterNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getClusterNodePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke GetClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetClusterNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClusterNodeOptions model
				getClusterNodeOptionsModel := new(powerhaautomationservicev1.GetClusterNodeOptions)
				getClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetClusterNode with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetClusterNodeOptions model
				getClusterNodeOptionsModel := new(powerhaautomationservicev1.GetClusterNodeOptions)
				getClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClusterNodeOptions model with no property values
				getClusterNodeOptionsModelNew := new(powerhaautomationservicev1.GetClusterNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModelNew)
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
			It(`Invoke GetClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetClusterNodeOptions model
				getClusterNodeOptionsModel := new(powerhaautomationservicev1.GetClusterNodeOptions)
				getClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptionsModel)
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
	Describe(`CreateClusterNode(createClusterNodeOptions *CreateClusterNodeOptions) - Operation response error`, func() {
		createClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClusterNodePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateClusterNode with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateClusterNodeOptions model
				createClusterNodeOptionsModel := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				createClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.SecondaryClusterNodes = []string{"ede4c36e-1234-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateClusterNode(createClusterNodeOptions *CreateClusterNodeOptions)`, func() {
		createClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClusterNodePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke CreateClusterNode successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreateClusterNodeOptions model
				createClusterNodeOptionsModel := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				createClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.SecondaryClusterNodes = []string{"ede4c36e-1234-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.CreateClusterNodeWithContext(ctx, createClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.CreateClusterNodeWithContext(ctx, createClusterNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createClusterNodePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke CreateClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.CreateClusterNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateClusterNodeOptions model
				createClusterNodeOptionsModel := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				createClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.SecondaryClusterNodes = []string{"ede4c36e-1234-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateClusterNode with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateClusterNodeOptions model
				createClusterNodeOptionsModel := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				createClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.SecondaryClusterNodes = []string{"ede4c36e-1234-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateClusterNodeOptions model with no property values
				createClusterNodeOptionsModelNew := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateClusterNodeOptions model
				createClusterNodeOptionsModel := new(powerhaautomationservicev1.CreateClusterNodeOptions)
				createClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.SecondaryClusterNodes = []string{"ede4c36e-1234-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptionsModel)
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
	Describe(`DeleteClusterNode(deleteClusterNodeOptions *DeleteClusterNodeOptions) - Operation response error`, func() {
		deleteClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClusterNodePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["vm_id"]).To(Equal([]string{"r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteClusterNode with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteClusterNodeOptions model
				deleteClusterNodeOptionsModel := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				deleteClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.VMID = core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				deleteClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteClusterNode(deleteClusterNodeOptions *DeleteClusterNodeOptions)`, func() {
		deleteClusterNodePath := "/powerha_automation/v1/cluster_nodes/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClusterNodePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["vm_id"]).To(Equal([]string{"r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke DeleteClusterNode successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the DeleteClusterNodeOptions model
				deleteClusterNodeOptionsModel := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				deleteClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.VMID = core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				deleteClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.DeleteClusterNodeWithContext(ctx, deleteClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.DeleteClusterNodeWithContext(ctx, deleteClusterNodeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteClusterNodePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["vm_id"]).To(Equal([]string{"r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "cluster-response-01", "primary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}], "secondary_node_details": [{"agent_status": "running", "cores": 8.0, "ip_addresses": ["IPAddresses"], "memory": 64.0, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-9b7c2d11", "vm_name": "pha-node-primary-1", "vm_status": "ACTIVE", "workspace_id": "workspace-primary-001"}]}`)
				}))
			})
			It(`Invoke DeleteClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.DeleteClusterNode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteClusterNodeOptions model
				deleteClusterNodeOptionsModel := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				deleteClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.VMID = core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				deleteClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteClusterNode with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteClusterNodeOptions model
				deleteClusterNodeOptionsModel := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				deleteClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.VMID = core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				deleteClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteClusterNodeOptions model with no property values
				deleteClusterNodeOptionsModelNew := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModelNew)
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
			It(`Invoke DeleteClusterNode successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteClusterNodeOptions model
				deleteClusterNodeOptionsModel := new(powerhaautomationservicev1.DeleteClusterNodeOptions)
				deleteClusterNodeOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.VMID = core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				deleteClusterNodeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptionsModel)
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
	Describe(`GetPowervsWorkspace(getPowervsWorkspaceOptions *GetPowervsWorkspaceOptions) - Operation response error`, func() {
		getPowervsWorkspacePath := "/powerha_automation/v1/powervs_workspaces/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPowervsWorkspace with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspaceOptions model
				getPowervsWorkspaceOptionsModel := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				getPowervsWorkspaceOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.LocationID = core.StringPtr("us-south")
				getPowervsWorkspaceOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPowervsWorkspaceOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPowervsWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPowervsWorkspace(getPowervsWorkspaceOptions *GetPowervsWorkspaceOptions)`, func() {
		getPowervsWorkspacePath := "/powerha_automation/v1/powervs_workspaces/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": [{"id": "ws-001", "name": "primary-workspace"}]}`)
				}))
			})
			It(`Invoke GetPowervsWorkspace successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetPowervsWorkspaceOptions model
				getPowervsWorkspaceOptionsModel := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				getPowervsWorkspaceOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.LocationID = core.StringPtr("us-south")
				getPowervsWorkspaceOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPowervsWorkspaceOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPowervsWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetPowervsWorkspaceWithContext(ctx, getPowervsWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetPowervsWorkspaceWithContext(ctx, getPowervsWorkspaceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": [{"id": "ws-001", "name": "primary-workspace"}]}`)
				}))
			})
			It(`Invoke GetPowervsWorkspace successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetPowervsWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPowervsWorkspaceOptions model
				getPowervsWorkspaceOptionsModel := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				getPowervsWorkspaceOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.LocationID = core.StringPtr("us-south")
				getPowervsWorkspaceOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPowervsWorkspaceOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPowervsWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPowervsWorkspace with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspaceOptions model
				getPowervsWorkspaceOptionsModel := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				getPowervsWorkspaceOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.LocationID = core.StringPtr("us-south")
				getPowervsWorkspaceOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPowervsWorkspaceOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPowervsWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPowervsWorkspaceOptions model with no property values
				getPowervsWorkspaceOptionsModelNew := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModelNew)
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
			It(`Invoke GetPowervsWorkspace successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspaceOptions model
				getPowervsWorkspaceOptionsModel := new(powerhaautomationservicev1.GetPowervsWorkspaceOptions)
				getPowervsWorkspaceOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.LocationID = core.StringPtr("us-south")
				getPowervsWorkspaceOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPowervsWorkspaceOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPowervsWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptionsModel)
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
	Describe(`GetPhaLastOperation(getPhaLastOperationOptions *GetPhaLastOperationOptions) - Operation response error`, func() {
		getPhaLastOperationPath := "/powerha_automation/v1/last_operation/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPhaLastOperation with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaLastOperationOptions model
				getPhaLastOperationOptionsModel := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				getPhaLastOperationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaLastOperationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPhaLastOperation(getPhaLastOperationOptions *GetPhaLastOperationOptions)`, func() {
		getPhaLastOperationPath := "/powerha_automation/v1/last_operation/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment_name": "pha-deployment-prod-01", "provision_id": "8eefautr-4c02-0009-0086-8bd4d8cf61b6", "resource_group": "ResourceGroup", "status": "ACTIVE"}`)
				}))
			})
			It(`Invoke GetPhaLastOperation successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetPhaLastOperationOptions model
				getPhaLastOperationOptionsModel := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				getPhaLastOperationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaLastOperationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetPhaLastOperationWithContext(ctx, getPhaLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetPhaLastOperationWithContext(ctx, getPhaLastOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPhaLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment_name": "pha-deployment-prod-01", "provision_id": "8eefautr-4c02-0009-0086-8bd4d8cf61b6", "resource_group": "ResourceGroup", "status": "ACTIVE"}`)
				}))
			})
			It(`Invoke GetPhaLastOperation successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetPhaLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPhaLastOperationOptions model
				getPhaLastOperationOptionsModel := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				getPhaLastOperationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaLastOperationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPhaLastOperation with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaLastOperationOptions model
				getPhaLastOperationOptionsModel := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				getPhaLastOperationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaLastOperationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPhaLastOperationOptions model with no property values
				getPhaLastOperationOptionsModelNew := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModelNew)
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
			It(`Invoke GetPhaLastOperation successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaLastOperationOptions model
				getPhaLastOperationOptionsModel := new(powerhaautomationservicev1.GetPhaLastOperationOptions)
				getPhaLastOperationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaLastOperationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptionsModel)
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
	Describe(`GetPhaDeployment(getPhaDeploymentOptions *GetPhaDeploymentOptions) - Operation response error`, func() {
		getPhaDeploymentPath := "/powerha_automation/v1/pha_deployment/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaDeploymentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPhaDeployment with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaDeploymentOptions model
				getPhaDeploymentOptionsModel := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				getPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPhaDeployment(getPhaDeploymentOptions *GetPhaDeploymentOptions)`, func() {
		getPhaDeploymentPath := "/powerha_automation/v1/pha_deployment/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaDeploymentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cloud_account_id": "adfadfdsafsdfdsf", "connectivity_type": "private", "creation_time": "2026-01-10T08:15:30Z", "custom_network": ["CustomNetwork"], "deprovision_time": "2026-01-20T12:45:00Z", "guid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "is_duplicate": false, "plan_id": "powerha-standard", "plan_name": "PowerHA Standard", "powerha_cluster_name": "pha-cluster-prod", "powerha_cluster_type": "standard", "powerha_level": "7.2.1", "primary_cluster_nodes_details": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "primary_location": "us-south", "primary_region_name": "Dallas", "primary_workspace": "ws-primary-001", "primary_workspace_name": "primary-ws-01", "provision_end_time": "2026-01-10T08:30:00Z", "id": "prov-9f8a7b6c", "provision_start_time": "2026-01-10T08:16:00Z", "provision_status": "SUCCEEDED", "region_id": "us-south", "resource_group": "rg-pha-prod", "resource_group_crn": "crn:v1:bluemix:public:resource-group:us-south:a/123456::rg:abcd1234", "resource_instance": "resource-instance-01", "secondary_cluster_nodes": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "secondary_location": "us-east", "secondary_workspace": "ws-secondary-001", "service_description": "PowerHA disaster recovery deployment", "service_id": "powerha", "service_name": "IBM PowerHA", "standby_region_name": "Washington", "standby_workspace_name": "standby-ws-01", "user_tags": "env:prod,team:dr"}`)
				}))
			})
			It(`Invoke GetPhaDeployment successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetPhaDeploymentOptions model
				getPhaDeploymentOptionsModel := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				getPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetPhaDeploymentWithContext(ctx, getPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetPhaDeploymentWithContext(ctx, getPhaDeploymentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPhaDeploymentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cloud_account_id": "adfadfdsafsdfdsf", "connectivity_type": "private", "creation_time": "2026-01-10T08:15:30Z", "custom_network": ["CustomNetwork"], "deprovision_time": "2026-01-20T12:45:00Z", "guid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "is_duplicate": false, "plan_id": "powerha-standard", "plan_name": "PowerHA Standard", "powerha_cluster_name": "pha-cluster-prod", "powerha_cluster_type": "standard", "powerha_level": "7.2.1", "primary_cluster_nodes_details": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "primary_location": "us-south", "primary_region_name": "Dallas", "primary_workspace": "ws-primary-001", "primary_workspace_name": "primary-ws-01", "provision_end_time": "2026-01-10T08:30:00Z", "id": "prov-9f8a7b6c", "provision_start_time": "2026-01-10T08:16:00Z", "provision_status": "SUCCEEDED", "region_id": "us-south", "resource_group": "rg-pha-prod", "resource_group_crn": "crn:v1:bluemix:public:resource-group:us-south:a/123456::rg:abcd1234", "resource_instance": "resource-instance-01", "secondary_cluster_nodes": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "secondary_location": "us-east", "secondary_workspace": "ws-secondary-001", "service_description": "PowerHA disaster recovery deployment", "service_id": "powerha", "service_name": "IBM PowerHA", "standby_region_name": "Washington", "standby_workspace_name": "standby-ws-01", "user_tags": "env:prod,team:dr"}`)
				}))
			})
			It(`Invoke GetPhaDeployment successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetPhaDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPhaDeploymentOptions model
				getPhaDeploymentOptionsModel := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				getPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPhaDeployment with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaDeploymentOptions model
				getPhaDeploymentOptionsModel := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				getPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPhaDeploymentOptions model with no property values
				getPhaDeploymentOptionsModelNew := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModelNew)
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
			It(`Invoke GetPhaDeployment successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaDeploymentOptions model
				getPhaDeploymentOptionsModel := new(powerhaautomationservicev1.GetPhaDeploymentOptions)
				getPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptionsModel)
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
	Describe(`CreatePhaDeployment(createPhaDeploymentOptions *CreatePhaDeploymentOptions) - Operation response error`, func() {
		createPhaDeploymentPath := "/powerha_automation/v1/pha_deployment/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPhaDeploymentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePhaDeployment with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreatePhaDeploymentOptions model
				createPhaDeploymentOptionsModel := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				createPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.LocationID = core.StringPtr("loc-us-south-01")
				createPhaDeploymentOptionsModel.PrimaryWorkspace = core.StringPtr("workspace-primary")
				createPhaDeploymentOptionsModel.APIKey = core.StringPtr("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.ClusterType = core.StringPtr("standard")
				createPhaDeploymentOptionsModel.ConfigureType = core.StringPtr("automatic")
				createPhaDeploymentOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createPhaDeploymentOptionsModel.StandbyClusterNodes = []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}
				createPhaDeploymentOptionsModel.PrimaryLocation = core.StringPtr("us-south")
				createPhaDeploymentOptionsModel.SecondaryLocation = core.StringPtr("us-east")
				createPhaDeploymentOptionsModel.SecondaryWorkspace = core.StringPtr("workspace-secondary")
				createPhaDeploymentOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePhaDeployment(createPhaDeploymentOptions *CreatePhaDeploymentOptions)`, func() {
		createPhaDeploymentPath := "/powerha_automation/v1/pha_deployment/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPhaDeploymentPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"cloud_account_id": "adfadfdsafsdfdsf", "connectivity_type": "private", "creation_time": "2026-01-10T08:15:30Z", "custom_network": ["CustomNetwork"], "deprovision_time": "2026-01-20T12:45:00Z", "guid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "is_duplicate": false, "plan_id": "powerha-standard", "plan_name": "PowerHA Standard", "powerha_cluster_name": "pha-cluster-prod", "powerha_cluster_type": "standard", "powerha_level": "7.2.1", "primary_cluster_nodes_details": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "primary_location": "us-south", "primary_region_name": "Dallas", "primary_workspace": "ws-primary-001", "primary_workspace_name": "primary-ws-01", "provision_end_time": "2026-01-10T08:30:00Z", "id": "prov-9f8a7b6c", "provision_start_time": "2026-01-10T08:16:00Z", "provision_status": "SUCCEEDED", "region_id": "us-south", "resource_group": "rg-pha-prod", "resource_group_crn": "crn:v1:bluemix:public:resource-group:us-south:a/123456::rg:abcd1234", "resource_instance": "resource-instance-01", "secondary_cluster_nodes": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "secondary_location": "us-east", "secondary_workspace": "ws-secondary-001", "service_description": "PowerHA disaster recovery deployment", "service_id": "powerha", "service_name": "IBM PowerHA", "standby_region_name": "Washington", "standby_workspace_name": "standby-ws-01", "user_tags": "env:prod,team:dr"}`)
				}))
			})
			It(`Invoke CreatePhaDeployment successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreatePhaDeploymentOptions model
				createPhaDeploymentOptionsModel := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				createPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.LocationID = core.StringPtr("loc-us-south-01")
				createPhaDeploymentOptionsModel.PrimaryWorkspace = core.StringPtr("workspace-primary")
				createPhaDeploymentOptionsModel.APIKey = core.StringPtr("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.ClusterType = core.StringPtr("standard")
				createPhaDeploymentOptionsModel.ConfigureType = core.StringPtr("automatic")
				createPhaDeploymentOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createPhaDeploymentOptionsModel.StandbyClusterNodes = []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}
				createPhaDeploymentOptionsModel.PrimaryLocation = core.StringPtr("us-south")
				createPhaDeploymentOptionsModel.SecondaryLocation = core.StringPtr("us-east")
				createPhaDeploymentOptionsModel.SecondaryWorkspace = core.StringPtr("workspace-secondary")
				createPhaDeploymentOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.CreatePhaDeploymentWithContext(ctx, createPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.CreatePhaDeploymentWithContext(ctx, createPhaDeploymentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPhaDeploymentPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"cloud_account_id": "adfadfdsafsdfdsf", "connectivity_type": "private", "creation_time": "2026-01-10T08:15:30Z", "custom_network": ["CustomNetwork"], "deprovision_time": "2026-01-20T12:45:00Z", "guid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "is_duplicate": false, "plan_id": "powerha-standard", "plan_name": "PowerHA Standard", "powerha_cluster_name": "pha-cluster-prod", "powerha_cluster_type": "standard", "powerha_level": "7.2.1", "primary_cluster_nodes_details": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "primary_location": "us-south", "primary_region_name": "Dallas", "primary_workspace": "ws-primary-001", "primary_workspace_name": "primary-ws-01", "provision_end_time": "2026-01-10T08:30:00Z", "id": "prov-9f8a7b6c", "provision_start_time": "2026-01-10T08:16:00Z", "provision_status": "SUCCEEDED", "region_id": "us-south", "resource_group": "rg-pha-prod", "resource_group_crn": "crn:v1:bluemix:public:resource-group:us-south:a/123456::rg:abcd1234", "resource_instance": "resource-instance-01", "secondary_cluster_nodes": [{"agent_status": "RUNNING", "cores": 8.0, "ip_address": "10.0.2.45", "memory": 32, "pha_level": "7.2.1", "region": "us-south", "vm_id": "vm-3c91af27", "vm_name": "pha-node-01", "vm_status": "ACTIVE", "workspace_id": "workspace-pha-prod"}], "secondary_location": "us-east", "secondary_workspace": "ws-secondary-001", "service_description": "PowerHA disaster recovery deployment", "service_id": "powerha", "service_name": "IBM PowerHA", "standby_region_name": "Washington", "standby_workspace_name": "standby-ws-01", "user_tags": "env:prod,team:dr"}`)
				}))
			})
			It(`Invoke CreatePhaDeployment successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.CreatePhaDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreatePhaDeploymentOptions model
				createPhaDeploymentOptionsModel := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				createPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.LocationID = core.StringPtr("loc-us-south-01")
				createPhaDeploymentOptionsModel.PrimaryWorkspace = core.StringPtr("workspace-primary")
				createPhaDeploymentOptionsModel.APIKey = core.StringPtr("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.ClusterType = core.StringPtr("standard")
				createPhaDeploymentOptionsModel.ConfigureType = core.StringPtr("automatic")
				createPhaDeploymentOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createPhaDeploymentOptionsModel.StandbyClusterNodes = []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}
				createPhaDeploymentOptionsModel.PrimaryLocation = core.StringPtr("us-south")
				createPhaDeploymentOptionsModel.SecondaryLocation = core.StringPtr("us-east")
				createPhaDeploymentOptionsModel.SecondaryWorkspace = core.StringPtr("workspace-secondary")
				createPhaDeploymentOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePhaDeployment with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreatePhaDeploymentOptions model
				createPhaDeploymentOptionsModel := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				createPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.LocationID = core.StringPtr("loc-us-south-01")
				createPhaDeploymentOptionsModel.PrimaryWorkspace = core.StringPtr("workspace-primary")
				createPhaDeploymentOptionsModel.APIKey = core.StringPtr("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.ClusterType = core.StringPtr("standard")
				createPhaDeploymentOptionsModel.ConfigureType = core.StringPtr("automatic")
				createPhaDeploymentOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createPhaDeploymentOptionsModel.StandbyClusterNodes = []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}
				createPhaDeploymentOptionsModel.PrimaryLocation = core.StringPtr("us-south")
				createPhaDeploymentOptionsModel.SecondaryLocation = core.StringPtr("us-east")
				createPhaDeploymentOptionsModel.SecondaryWorkspace = core.StringPtr("workspace-secondary")
				createPhaDeploymentOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePhaDeploymentOptions model with no property values
				createPhaDeploymentOptionsModelNew := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreatePhaDeployment successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreatePhaDeploymentOptions model
				createPhaDeploymentOptionsModel := new(powerhaautomationservicev1.CreatePhaDeploymentOptions)
				createPhaDeploymentOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.LocationID = core.StringPtr("loc-us-south-01")
				createPhaDeploymentOptionsModel.PrimaryWorkspace = core.StringPtr("workspace-primary")
				createPhaDeploymentOptionsModel.APIKey = core.StringPtr("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.ClusterType = core.StringPtr("standard")
				createPhaDeploymentOptionsModel.ConfigureType = core.StringPtr("automatic")
				createPhaDeploymentOptionsModel.PrimaryClusterNodes = []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createPhaDeploymentOptionsModel.StandbyClusterNodes = []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}
				createPhaDeploymentOptionsModel.PrimaryLocation = core.StringPtr("us-south")
				createPhaDeploymentOptionsModel.SecondaryLocation = core.StringPtr("us-east")
				createPhaDeploymentOptionsModel.SecondaryWorkspace = core.StringPtr("workspace-secondary")
				createPhaDeploymentOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				createPhaDeploymentOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				createPhaDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptionsModel)
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
	Describe(`GetSupportedLocation(getSupportedLocationOptions *GetSupportedLocationOptions) - Operation response error`, func() {
		getSupportedLocationPath := "/powerha_automation/v1/supported_locations/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedLocationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSupportedLocation with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetSupportedLocationOptions model
				getSupportedLocationOptionsModel := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				getSupportedLocationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getSupportedLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSupportedLocation(getSupportedLocationOptions *GetSupportedLocationOptions)`, func() {
		getSupportedLocationPath := "/powerha_automation/v1/supported_locations/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedLocationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"id": "loc-us-south-01", "name": "Dallas (us-south)"}]}`)
				}))
			})
			It(`Invoke GetSupportedLocation successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetSupportedLocationOptions model
				getSupportedLocationOptionsModel := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				getSupportedLocationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getSupportedLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetSupportedLocationWithContext(ctx, getSupportedLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetSupportedLocationWithContext(ctx, getSupportedLocationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSupportedLocationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": [{"id": "loc-us-south-01", "name": "Dallas (us-south)"}]}`)
				}))
			})
			It(`Invoke GetSupportedLocation successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetSupportedLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSupportedLocationOptions model
				getSupportedLocationOptionsModel := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				getSupportedLocationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getSupportedLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSupportedLocation with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetSupportedLocationOptions model
				getSupportedLocationOptionsModel := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				getSupportedLocationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getSupportedLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSupportedLocationOptions model with no property values
				getSupportedLocationOptionsModelNew := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModelNew)
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
			It(`Invoke GetSupportedLocation successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetSupportedLocationOptions model
				getSupportedLocationOptionsModel := new(powerhaautomationservicev1.GetSupportedLocationOptions)
				getSupportedLocationOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getSupportedLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptionsModel)
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
	Describe(`GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptions *GetPhaAgentFileDownloadJobStatusOptions) - Operation response error`, func() {
		getPhaAgentFileDownloadJobStatusPath := "/powerha_automation/v1/pha_agent/download/8eefautr-4c02-0009-0086-8bd4d8cf61b6/jobs/4235r23r5vdfdf-2323"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaAgentFileDownloadJobStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPhaAgentFileDownloadJobStatus with error: Operation response processing error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				getPhaAgentFileDownloadJobStatusOptionsModel := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID = core.StringPtr("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				powerhaAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptions *GetPhaAgentFileDownloadJobStatusOptions)`, func() {
		getPhaAgentFileDownloadJobStatusPath := "/powerha_automation/v1/pha_agent/download/8eefautr-4c02-0009-0086-8bd4d8cf61b6/jobs/4235r23r5vdfdf-2323"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPhaAgentFileDownloadJobStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bytes_downloaded": 52428800, "creation_at": "2026-01-08T11:00:00.000Z", "file_name": "power_agent", "job_id": "job-98765", "last_updated_at": "2026-01-08T12:15:00.000Z", "service_instance_id": "service-12345", "status": "running", "total_bytes": 104857600, "vm_id": "vm-12345"}`)
				}))
			})
			It(`Invoke GetPhaAgentFileDownloadJobStatus successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				getPhaAgentFileDownloadJobStatusOptionsModel := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID = core.StringPtr("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatusWithContext(ctx, getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatusWithContext(ctx, getPhaAgentFileDownloadJobStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPhaAgentFileDownloadJobStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bytes_downloaded": 52428800, "creation_at": "2026-01-08T11:00:00.000Z", "file_name": "power_agent", "job_id": "job-98765", "last_updated_at": "2026-01-08T12:15:00.000Z", "service_instance_id": "service-12345", "status": "running", "total_bytes": 104857600, "vm_id": "vm-12345"}`)
				}))
			})
			It(`Invoke GetPhaAgentFileDownloadJobStatus successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				getPhaAgentFileDownloadJobStatusOptionsModel := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID = core.StringPtr("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPhaAgentFileDownloadJobStatus with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				getPhaAgentFileDownloadJobStatusOptionsModel := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID = core.StringPtr("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPhaAgentFileDownloadJobStatusOptions model with no property values
				getPhaAgentFileDownloadJobStatusOptionsModelNew := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModelNew)
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
			It(`Invoke GetPhaAgentFileDownloadJobStatus successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				getPhaAgentFileDownloadJobStatusOptionsModel := new(powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions)
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID = core.StringPtr("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptionsModel)
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
	Describe(`DownloadPhaAgentFile(downloadPhaAgentFileOptions *DownloadPhaAgentFileOptions)`, func() {
		downloadPhaAgentFilePath := "/powerha_automation/v1/pha_agent/download/8eefautr-4c02-0009-0086-8bd4d8cf61b6"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(downloadPhaAgentFilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DownloadPhaAgentFile successfully with retries`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())
				powerhaAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the DownloadPhaAgentFileOptions model
				downloadPhaAgentFileOptionsModel := new(powerhaautomationservicev1.DownloadPhaAgentFileOptions)
				downloadPhaAgentFileOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				downloadPhaAgentFileOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				downloadPhaAgentFileOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				downloadPhaAgentFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := powerhaAutomationServiceService.DownloadPhaAgentFileWithContext(ctx, downloadPhaAgentFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				powerhaAutomationServiceService.DisableRetries()
				result, response, operationErr := powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = powerhaAutomationServiceService.DownloadPhaAgentFileWithContext(ctx, downloadPhaAgentFileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(downloadPhaAgentFilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "en-US")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "abcdef")))
					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DownloadPhaAgentFile successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := powerhaAutomationServiceService.DownloadPhaAgentFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DownloadPhaAgentFileOptions model
				downloadPhaAgentFileOptionsModel := new(powerhaautomationservicev1.DownloadPhaAgentFileOptions)
				downloadPhaAgentFileOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				downloadPhaAgentFileOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				downloadPhaAgentFileOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				downloadPhaAgentFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DownloadPhaAgentFile with error: Operation validation and request error`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the DownloadPhaAgentFileOptions model
				downloadPhaAgentFileOptionsModel := new(powerhaautomationservicev1.DownloadPhaAgentFileOptions)
				downloadPhaAgentFileOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				downloadPhaAgentFileOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				downloadPhaAgentFileOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				downloadPhaAgentFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := powerhaAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DownloadPhaAgentFileOptions model with no property values
				downloadPhaAgentFileOptionsModelNew := new(powerhaautomationservicev1.DownloadPhaAgentFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptionsModelNew)
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
			It(`Invoke DownloadPhaAgentFile successfully`, func() {
				powerhaAutomationServiceService, serviceErr := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(powerhaAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the DownloadPhaAgentFileOptions model
				downloadPhaAgentFileOptionsModel := new(powerhaautomationservicev1.DownloadPhaAgentFileOptions)
				downloadPhaAgentFileOptionsModel.PhaInstanceID = core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				downloadPhaAgentFileOptionsModel.AcceptLanguage = core.StringPtr("en-US")
				downloadPhaAgentFileOptionsModel.IfNoneMatch = core.StringPtr("abcdef")
				downloadPhaAgentFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			powerhaAutomationServiceService, _ := powerhaautomationservicev1.NewPowerhaAutomationServiceV1(&powerhaautomationservicev1.PowerhaAutomationServiceV1Options{
				URL:           "http://powerhaautomationservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateAPIKeyOptions successfully`, func() {
				// Construct an instance of the CreateAPIKeyOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				createAPIKeyOptionsModel := powerhaAutomationServiceService.NewCreateAPIKeyOptions(phaInstanceID)
				createAPIKeyOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createAPIKeyOptionsModel.SetAPIKey("adfadfdsafsdfdsf")
				createAPIKeyOptionsModel.SetAcceptLanguage("en-US")
				createAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAPIKeyOptionsModel).ToNot(BeNil())
				Expect(createAPIKeyOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(createAPIKeyOptionsModel.APIKey).To(Equal(core.StringPtr("adfadfdsafsdfdsf")))
				Expect(createAPIKeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(createAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateClusterNodeOptions successfully`, func() {
				// Construct an instance of the CreateClusterNodeOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				createClusterNodeOptionsPrimaryClusterNodes := []string{"ede4c36e-002c-48da-992e-6039d230c478"}
				createClusterNodeOptionsModel := powerhaAutomationServiceService.NewCreateClusterNodeOptions(phaInstanceID, createClusterNodeOptionsPrimaryClusterNodes)
				createClusterNodeOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createClusterNodeOptionsModel.SetPrimaryClusterNodes([]string{"ede4c36e-002c-48da-992e-6039d230c478"})
				createClusterNodeOptionsModel.SetSecondaryClusterNodes([]string{"ede4c36e-1234-48da-992e-6039d230c478"})
				createClusterNodeOptionsModel.SetAcceptLanguage("en-US")
				createClusterNodeOptionsModel.SetIfNoneMatch("abcdef")
				createClusterNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createClusterNodeOptionsModel).ToNot(BeNil())
				Expect(createClusterNodeOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(createClusterNodeOptionsModel.PrimaryClusterNodes).To(Equal([]string{"ede4c36e-002c-48da-992e-6039d230c478"}))
				Expect(createClusterNodeOptionsModel.SecondaryClusterNodes).To(Equal([]string{"ede4c36e-1234-48da-992e-6039d230c478"}))
				Expect(createClusterNodeOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(createClusterNodeOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(createClusterNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePhaDeploymentOptions successfully`, func() {
				// Construct an instance of the CreatePhaDeploymentOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				createPhaDeploymentOptionsLocationID := "loc-us-south-01"
				createPhaDeploymentOptionsPrimaryWorkspace := "workspace-primary"
				createPhaDeploymentOptionsModel := powerhaAutomationServiceService.NewCreatePhaDeploymentOptions(phaInstanceID, createPhaDeploymentOptionsLocationID, createPhaDeploymentOptionsPrimaryWorkspace)
				createPhaDeploymentOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				createPhaDeploymentOptionsModel.SetLocationID("loc-us-south-01")
				createPhaDeploymentOptionsModel.SetPrimaryWorkspace("workspace-primary")
				createPhaDeploymentOptionsModel.SetAPIKey("123635364646fghrtfhbfdhb")
				createPhaDeploymentOptionsModel.SetClusterType("standard")
				createPhaDeploymentOptionsModel.SetConfigureType("automatic")
				createPhaDeploymentOptionsModel.SetPrimaryClusterNodes([]string{"ede4c36e-002c-48da-992e-6039d230c478"})
				createPhaDeploymentOptionsModel.SetStandbyClusterNodes([]string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"})
				createPhaDeploymentOptionsModel.SetPrimaryLocation("us-south")
				createPhaDeploymentOptionsModel.SetSecondaryLocation("us-east")
				createPhaDeploymentOptionsModel.SetSecondaryWorkspace("workspace-secondary")
				createPhaDeploymentOptionsModel.SetAcceptLanguage("en-US")
				createPhaDeploymentOptionsModel.SetIfNoneMatch("abcdef")
				createPhaDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPhaDeploymentOptionsModel).ToNot(BeNil())
				Expect(createPhaDeploymentOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(createPhaDeploymentOptionsModel.LocationID).To(Equal(core.StringPtr("loc-us-south-01")))
				Expect(createPhaDeploymentOptionsModel.PrimaryWorkspace).To(Equal(core.StringPtr("workspace-primary")))
				Expect(createPhaDeploymentOptionsModel.APIKey).To(Equal(core.StringPtr("123635364646fghrtfhbfdhb")))
				Expect(createPhaDeploymentOptionsModel.ClusterType).To(Equal(core.StringPtr("standard")))
				Expect(createPhaDeploymentOptionsModel.ConfigureType).To(Equal(core.StringPtr("automatic")))
				Expect(createPhaDeploymentOptionsModel.PrimaryClusterNodes).To(Equal([]string{"ede4c36e-002c-48da-992e-6039d230c478"}))
				Expect(createPhaDeploymentOptionsModel.StandbyClusterNodes).To(Equal([]string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"}))
				Expect(createPhaDeploymentOptionsModel.PrimaryLocation).To(Equal(core.StringPtr("us-south")))
				Expect(createPhaDeploymentOptionsModel.SecondaryLocation).To(Equal(core.StringPtr("us-east")))
				Expect(createPhaDeploymentOptionsModel.SecondaryWorkspace).To(Equal(core.StringPtr("workspace-secondary")))
				Expect(createPhaDeploymentOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(createPhaDeploymentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(createPhaDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteClusterNodeOptions successfully`, func() {
				// Construct an instance of the DeleteClusterNodeOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				vmID := "r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"
				deleteClusterNodeOptionsModel := powerhaAutomationServiceService.NewDeleteClusterNodeOptions(phaInstanceID, vmID)
				deleteClusterNodeOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				deleteClusterNodeOptionsModel.SetVMID("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")
				deleteClusterNodeOptionsModel.SetIfNoneMatch("abcdef")
				deleteClusterNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteClusterNodeOptionsModel).ToNot(BeNil())
				Expect(deleteClusterNodeOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(deleteClusterNodeOptionsModel.VMID).To(Equal(core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2")))
				Expect(deleteClusterNodeOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(deleteClusterNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDownloadPhaAgentFileOptions successfully`, func() {
				// Construct an instance of the DownloadPhaAgentFileOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				downloadPhaAgentFileOptionsModel := powerhaAutomationServiceService.NewDownloadPhaAgentFileOptions(phaInstanceID)
				downloadPhaAgentFileOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				downloadPhaAgentFileOptionsModel.SetAcceptLanguage("en-US")
				downloadPhaAgentFileOptionsModel.SetIfNoneMatch("abcdef")
				downloadPhaAgentFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(downloadPhaAgentFileOptionsModel).ToNot(BeNil())
				Expect(downloadPhaAgentFileOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(downloadPhaAgentFileOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(downloadPhaAgentFileOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(downloadPhaAgentFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAPIKeyOptions successfully`, func() {
				// Construct an instance of the GetAPIKeyOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				getAPIKeyOptionsModel := powerhaAutomationServiceService.NewGetAPIKeyOptions(phaInstanceID)
				getAPIKeyOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getAPIKeyOptionsModel.SetAcceptLanguage("en-US")
				getAPIKeyOptionsModel.SetIfNoneMatch("abcdef")
				getAPIKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAPIKeyOptionsModel).ToNot(BeNil())
				Expect(getAPIKeyOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getAPIKeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(getAPIKeyOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getAPIKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClusterNodeOptions successfully`, func() {
				// Construct an instance of the GetClusterNodeOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				getClusterNodeOptionsModel := powerhaAutomationServiceService.NewGetClusterNodeOptions(phaInstanceID)
				getClusterNodeOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getClusterNodeOptionsModel.SetIfNoneMatch("abcdef")
				getClusterNodeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClusterNodeOptionsModel).ToNot(BeNil())
				Expect(getClusterNodeOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getClusterNodeOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getClusterNodeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPhaAgentFileDownloadJobStatusOptions successfully`, func() {
				// Construct an instance of the GetPhaAgentFileDownloadJobStatusOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				phaJobID := "4235r23r5vdfdf-2323"
				getPhaAgentFileDownloadJobStatusOptionsModel := powerhaAutomationServiceService.NewGetPhaAgentFileDownloadJobStatusOptions(phaInstanceID, phaJobID)
				getPhaAgentFileDownloadJobStatusOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaAgentFileDownloadJobStatusOptionsModel.SetPhaJobID("4235r23r5vdfdf-2323")
				getPhaAgentFileDownloadJobStatusOptionsModel.SetAcceptLanguage("en-US")
				getPhaAgentFileDownloadJobStatusOptionsModel.SetIfNoneMatch("abcdef")
				getPhaAgentFileDownloadJobStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel).ToNot(BeNil())
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel.PhaJobID).To(Equal(core.StringPtr("4235r23r5vdfdf-2323")))
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getPhaAgentFileDownloadJobStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPhaDeploymentOptions successfully`, func() {
				// Construct an instance of the GetPhaDeploymentOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				getPhaDeploymentOptionsModel := powerhaAutomationServiceService.NewGetPhaDeploymentOptions(phaInstanceID)
				getPhaDeploymentOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaDeploymentOptionsModel.SetIfNoneMatch("abcdef")
				getPhaDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPhaDeploymentOptionsModel).ToNot(BeNil())
				Expect(getPhaDeploymentOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getPhaDeploymentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getPhaDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPhaLastOperationOptions successfully`, func() {
				// Construct an instance of the GetPhaLastOperationOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				getPhaLastOperationOptionsModel := powerhaAutomationServiceService.NewGetPhaLastOperationOptions(phaInstanceID)
				getPhaLastOperationOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPhaLastOperationOptionsModel.SetAcceptLanguage("en-US")
				getPhaLastOperationOptionsModel.SetIfNoneMatch("abcdef")
				getPhaLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPhaLastOperationOptionsModel).ToNot(BeNil())
				Expect(getPhaLastOperationOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getPhaLastOperationOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(getPhaLastOperationOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getPhaLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPowervsWorkspaceOptions successfully`, func() {
				// Construct an instance of the GetPowervsWorkspaceOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				locationID := "us-south"
				getPowervsWorkspaceOptionsModel := powerhaAutomationServiceService.NewGetPowervsWorkspaceOptions(phaInstanceID, locationID)
				getPowervsWorkspaceOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getPowervsWorkspaceOptionsModel.SetLocationID("us-south")
				getPowervsWorkspaceOptionsModel.SetAcceptLanguage("en-US")
				getPowervsWorkspaceOptionsModel.SetIfNoneMatch("abcdef")
				getPowervsWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPowervsWorkspaceOptionsModel).ToNot(BeNil())
				Expect(getPowervsWorkspaceOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getPowervsWorkspaceOptionsModel.LocationID).To(Equal(core.StringPtr("us-south")))
				Expect(getPowervsWorkspaceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("en-US")))
				Expect(getPowervsWorkspaceOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getPowervsWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSupportedLocationOptions successfully`, func() {
				// Construct an instance of the GetSupportedLocationOptions model
				phaInstanceID := "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
				getSupportedLocationOptionsModel := powerhaAutomationServiceService.NewGetSupportedLocationOptions(phaInstanceID)
				getSupportedLocationOptionsModel.SetPhaInstanceID("8eefautr-4c02-0009-0086-8bd4d8cf61b6")
				getSupportedLocationOptionsModel.SetIfNoneMatch("abcdef")
				getSupportedLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSupportedLocationOptionsModel).ToNot(BeNil())
				Expect(getSupportedLocationOptionsModel.PhaInstanceID).To(Equal(core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6")))
				Expect(getSupportedLocationOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("abcdef")))
				Expect(getSupportedLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
