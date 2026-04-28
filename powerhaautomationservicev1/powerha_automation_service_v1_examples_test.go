//go:build examples

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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the PowerhaAutomation Service service.
//
// The following configuration properties are assumed to be defined:
// POWERHA_AUTOMATION_SERVICE_URL=<service base url>
// POWERHA_AUTOMATION_SERVICE_AUTH_TYPE=iam
// POWERHA_AUTOMATION_SERVICE_APIKEY=<IAM apikey>
// POWERHA_AUTOMATION_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PowerhaAutomationServiceV1 Examples Tests`, func() {

	const externalConfigFile = "../powerha_automation_service_v1.env"

	var (
		powerhaAutomationServiceService *powerhaautomationservicev1.PowerhaAutomationServiceV1
		config                          map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(powerhaautomationservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			powerhaAutomationServiceServiceOptions := &powerhaautomationservicev1.PowerhaAutomationServiceV1Options{}

			powerhaAutomationServiceService, err = powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(powerhaAutomationServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(powerhaAutomationServiceService).ToNot(BeNil())
		})
	})

	Describe(`PowerhaAutomationServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKey request example`, func() {
			fmt.Println("\nGetAPIKey() result:")
			// begin-get_api_key

			getAPIKeyOptions := powerhaAutomationServiceService.NewGetAPIKeyOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			getAPIKeyOptions.SetAcceptLanguage("en-US")
			getAPIKeyOptions.SetIfNoneMatch("abcdef")

			apiKeyResponse, response, err := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyResponse).ToNot(BeNil())
		})
		It(`CreateAPIKey request example`, func() {
			fmt.Println("\nCreateAPIKey() result:")
			// begin-create_api_key

			createAPIKeyOptions := powerhaAutomationServiceService.NewCreateAPIKeyOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			createAPIKeyOptions.SetAcceptLanguage("en-US")

			apiKeyResponse, response, err := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKeyResponse).ToNot(BeNil())
		})
		It(`GetClusterNode request example`, func() {
			fmt.Println("\nGetClusterNode() result:")
			// begin-get_cluster_node

			getClusterNodeOptions := powerhaAutomationServiceService.NewGetClusterNodeOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			getClusterNodeOptions.SetIfNoneMatch("abcdef")

			clusterNodeResponse, response, err := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterNodeResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_cluster_node

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
		It(`CreateClusterNode request example`, func() {
			fmt.Println("\nCreateClusterNode() result:")
			// begin-create_cluster_node

			createClusterNodeOptions := powerhaAutomationServiceService.NewCreateClusterNodeOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
				[]string{"ede4c36e-002c-48da-992e-6039d230c478"},
			)
			createClusterNodeOptions.SetAcceptLanguage("en-US")
			createClusterNodeOptions.SetIfNoneMatch("abcdef")

			clusterNodeResponse, response, err := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterNodeResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_cluster_node

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
		It(`GetPowervsWorkspace request example`, func() {
			fmt.Println("\nGetPowervsWorkspace() result:")
			// begin-get_powervs_workspace

			getPowervsWorkspaceOptions := powerhaAutomationServiceService.NewGetPowervsWorkspaceOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
				"us-south",
			)
			getPowervsWorkspaceOptions.SetAcceptLanguage("en-US")
			getPowervsWorkspaceOptions.SetIfNoneMatch("abcdef")

			phaWorkspacesRegionResponse, response, err := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(phaWorkspacesRegionResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_powervs_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaWorkspacesRegionResponse).ToNot(BeNil())
		})
		It(`GetPhaLastOperation request example`, func() {
			fmt.Println("\nGetPhaLastOperation() result:")
			// begin-get_pha_last_operation

			getPhaLastOperationOptions := powerhaAutomationServiceService.NewGetPhaLastOperationOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			getPhaLastOperationOptions.SetAcceptLanguage("en-US")
			getPhaLastOperationOptions.SetIfNoneMatch("abcdef")

			serviceInstancePhaStatus, response, err := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceInstancePhaStatus, "", "  ")
			fmt.Println(string(b))

			// end-get_pha_last_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstancePhaStatus).ToNot(BeNil())
		})
		It(`GetPhaDeployment request example`, func() {
			fmt.Println("\nGetPhaDeployment() result:")
			// begin-get_pha_deployment

			getPhaDeploymentOptions := powerhaAutomationServiceService.NewGetPhaDeploymentOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			getPhaDeploymentOptions.SetIfNoneMatch("abcdef")

			phaDeploymentResponse, response, err := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(phaDeploymentResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_pha_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaDeploymentResponse).ToNot(BeNil())
		})
		It(`CreatePhaDeployment request example`, func() {
			fmt.Println("\nCreatePhaDeployment() result:")
			// begin-create_pha_deployment

			createPhaDeploymentOptions := powerhaAutomationServiceService.NewCreatePhaDeploymentOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
				"loc-us-south-01",
				"workspace-primary",
			)
			createPhaDeploymentOptions.SetAcceptLanguage("en-US")
			createPhaDeploymentOptions.SetIfNoneMatch("abcdef")

			phaDeploymentResponse, response, err := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(phaDeploymentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_pha_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(phaDeploymentResponse).ToNot(BeNil())
		})
		It(`GetSupportedLocation request example`, func() {
			fmt.Println("\nGetSupportedLocation() result:")
			// begin-get_supported_location

			getSupportedLocationOptions := powerhaAutomationServiceService.NewGetSupportedLocationOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			getSupportedLocationOptions.SetIfNoneMatch("abcdef")

			phaSupportedLocationsResponse, response, err := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(phaSupportedLocationsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_supported_location

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaSupportedLocationsResponse).ToNot(BeNil())
		})
		It(`GetPhaAgentFileDownloadJobStatus request example`, func() {
			fmt.Println("\nGetPhaAgentFileDownloadJobStatus() result:")
			// begin-get_pha_agent_file_download_job_status

			getPhaAgentFileDownloadJobStatusOptions := powerhaAutomationServiceService.NewGetPhaAgentFileDownloadJobStatusOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
				"4235r23r5vdfdf-2323",
			)
			getPhaAgentFileDownloadJobStatusOptions.SetAcceptLanguage("en-US")
			getPhaAgentFileDownloadJobStatusOptions.SetIfNoneMatch("abcdef")

			phaAgentJobStatusResponse, response, err := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(phaAgentJobStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_pha_agent_file_download_job_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaAgentJobStatusResponse).ToNot(BeNil())
		})
		It(`DownloadPhaAgentFile request example`, func() {
			fmt.Println("\nDownloadPhaAgentFile() result:")
			// begin-download_pha_agent_file

			downloadPhaAgentFileOptions := powerhaAutomationServiceService.NewDownloadPhaAgentFileOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
			)
			downloadPhaAgentFileOptions.SetAcceptLanguage("en-US")
			downloadPhaAgentFileOptions.SetIfNoneMatch("abcdef")

			result, response, err := powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil {
					panic(err)
				}
			}

			// end-download_pha_agent_file

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`DeleteClusterNode request example`, func() {
			fmt.Println("\nDeleteClusterNode() result:")
			// begin-delete_cluster_node

			deleteClusterNodeOptions := powerhaAutomationServiceService.NewDeleteClusterNodeOptions(
				"8eefautr-4c02-0009-0086-8bd4d8cf61b6",
				"r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2",
			)
			deleteClusterNodeOptions.SetIfNoneMatch("abcdef")

			clusterNodeResponse, response, err := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterNodeResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_cluster_node

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
	})
})
