//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the powerhaautomationservicev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PowerhaAutomationServiceV1 Integration Tests`, func() {
	const externalConfigFile = "../powerha_automation_service_v1.env"

	var (
		err                             error
		powerhaAutomationServiceService *powerhaautomationservicev1.PowerhaAutomationServiceV1
		serviceURL                      string
		config                          map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(powerhaautomationservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			powerhaAutomationServiceServiceOptions := &powerhaautomationservicev1.PowerhaAutomationServiceV1Options{}

			powerhaAutomationServiceService, err = powerhaautomationservicev1.NewPowerhaAutomationServiceV1UsingExternalConfig(powerhaAutomationServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(powerhaAutomationServiceService).ToNot(BeNil())
			Expect(powerhaAutomationServiceService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			powerhaAutomationServiceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetAPIKey - Get the apikey for the specified PowerHA service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions)`, func() {
			getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				AcceptLanguage: core.StringPtr("en-US"),
				IfNoneMatch:    core.StringPtr("abcdef"),
			}

			apiKeyResponse, response, err := powerhaAutomationServiceService.GetAPIKey(getAPIKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateAPIKey - Update the apikey for the specified PowerHA service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {
			createAPIKeyOptions := &powerhaautomationservicev1.CreateAPIKeyOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				APIKey:         core.StringPtr("adfadfdsafsdfdsf"),
				AcceptLanguage: core.StringPtr("en-US"),
			}

			apiKeyResponse, response, err := powerhaAutomationServiceService.CreateAPIKey(createAPIKeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`GetClusterNode - Retrieves the cluster node details for the specified PowerHA instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetClusterNode(getClusterNodeOptions *GetClusterNodeOptions)`, func() {
			getClusterNodeOptions := &powerhaautomationservicev1.GetClusterNodeOptions{
				PhaInstanceID: core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				IfNoneMatch:   core.StringPtr("abcdef"),
			}

			clusterNodeResponse, response, err := powerhaAutomationServiceService.GetClusterNode(getClusterNodeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateClusterNode - Add a new cluster node`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateClusterNode(createClusterNodeOptions *CreateClusterNodeOptions)`, func() {
			createClusterNodeOptions := &powerhaautomationservicev1.CreateClusterNodeOptions{
				PhaInstanceID:         core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				PrimaryClusterNodes:   []string{"ede4c36e-002c-48da-992e-6039d230c478"},
				SecondaryClusterNodes: []string{"ede4c36e-1234-48da-992e-6039d230c478"},
				AcceptLanguage:        core.StringPtr("en-US"),
				IfNoneMatch:           core.StringPtr("abcdef"),
			}

			clusterNodeResponse, response, err := powerhaAutomationServiceService.CreateClusterNode(createClusterNodeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
	})

	Describe(`GetPowervsWorkspace - List of PowerVS workspaces based on location`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPowervsWorkspace(getPowervsWorkspaceOptions *GetPowervsWorkspaceOptions)`, func() {
			getPowervsWorkspaceOptions := &powerhaautomationservicev1.GetPowervsWorkspaceOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				LocationID:     core.StringPtr("us-south"),
				AcceptLanguage: core.StringPtr("en-US"),
				IfNoneMatch:    core.StringPtr("abcdef"),
			}

			phaWorkspacesRegionResponse, response, err := powerhaAutomationServiceService.GetPowervsWorkspace(getPowervsWorkspaceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaWorkspacesRegionResponse).ToNot(BeNil())
		})
	})

	Describe(`GetPhaLastOperation - View details of the last operation performed on the PowerHA instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPhaLastOperation(getPhaLastOperationOptions *GetPhaLastOperationOptions)`, func() {
			getPhaLastOperationOptions := &powerhaautomationservicev1.GetPhaLastOperationOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				AcceptLanguage: core.StringPtr("en-US"),
				IfNoneMatch:    core.StringPtr("abcdef"),
			}

			serviceInstancePhaStatus, response, err := powerhaAutomationServiceService.GetPhaLastOperation(getPhaLastOperationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstancePhaStatus).ToNot(BeNil())
		})
	})

	Describe(`GetPhaDeployment - Retrieves deployment details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPhaDeployment(getPhaDeploymentOptions *GetPhaDeploymentOptions)`, func() {
			getPhaDeploymentOptions := &powerhaautomationservicev1.GetPhaDeploymentOptions{
				PhaInstanceID: core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				IfNoneMatch:   core.StringPtr("abcdef"),
			}

			phaDeploymentResponse, response, err := powerhaAutomationServiceService.GetPhaDeployment(getPhaDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaDeploymentResponse).ToNot(BeNil())
		})
	})

	Describe(`CreatePhaDeployment - Create a new deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePhaDeployment(createPhaDeploymentOptions *CreatePhaDeploymentOptions)`, func() {
			createPhaDeploymentOptions := &powerhaautomationservicev1.CreatePhaDeploymentOptions{
				PhaInstanceID:       core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				LocationID:          core.StringPtr("loc-us-south-01"),
				PrimaryWorkspace:    core.StringPtr("workspace-primary"),
				APIKey:              core.StringPtr("123635364646fghrtfhbfdhb"),
				ClusterType:         core.StringPtr("standard"),
				ConfigureType:       core.StringPtr("automatic"),
				PrimaryClusterNodes: []string{"ede4c36e-002c-48da-992e-6039d230c478"},
				StandbyClusterNodes: []string{"843a8e1f-05bb-4164-8c73-de39e016c2b4"},
				PrimaryLocation:     core.StringPtr("us-south"),
				SecondaryLocation:   core.StringPtr("us-east"),
				SecondaryWorkspace:  core.StringPtr("workspace-secondary"),
				AcceptLanguage:      core.StringPtr("en-US"),
				IfNoneMatch:         core.StringPtr("abcdef"),
			}

			phaDeploymentResponse, response, err := powerhaAutomationServiceService.CreatePhaDeployment(createPhaDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(phaDeploymentResponse).ToNot(BeNil())
		})
	})

	Describe(`GetSupportedLocation - Get PowerVS locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSupportedLocation(getSupportedLocationOptions *GetSupportedLocationOptions)`, func() {
			getSupportedLocationOptions := &powerhaautomationservicev1.GetSupportedLocationOptions{
				PhaInstanceID: core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				IfNoneMatch:   core.StringPtr("abcdef"),
			}

			phaSupportedLocationsResponse, response, err := powerhaAutomationServiceService.GetSupportedLocation(getSupportedLocationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaSupportedLocationsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetPhaAgentFileDownloadJobStatus - Get the Job status of the downloaded powerHA agent file`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptions *GetPhaAgentFileDownloadJobStatusOptions)`, func() {
			getPhaAgentFileDownloadJobStatusOptions := &powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				PhaJobID:       core.StringPtr("4235r23r5vdfdf-2323"),
				AcceptLanguage: core.StringPtr("en-US"),
				IfNoneMatch:    core.StringPtr("abcdef"),
			}

			phaAgentJobStatusResponse, response, err := powerhaAutomationServiceService.GetPhaAgentFileDownloadJobStatus(getPhaAgentFileDownloadJobStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(phaAgentJobStatusResponse).ToNot(BeNil())
		})
	})

	Describe(`DownloadPhaAgentFile - Downloads PowerHA Agent file`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DownloadPhaAgentFile(downloadPhaAgentFileOptions *DownloadPhaAgentFileOptions)`, func() {
			downloadPhaAgentFileOptions := &powerhaautomationservicev1.DownloadPhaAgentFileOptions{
				PhaInstanceID:  core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				AcceptLanguage: core.StringPtr("en-US"),
				IfNoneMatch:    core.StringPtr("abcdef"),
			}

			result, response, err := powerhaAutomationServiceService.DownloadPhaAgentFile(downloadPhaAgentFileOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeleteClusterNode - Delete a cluster node`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteClusterNode(deleteClusterNodeOptions *DeleteClusterNodeOptions)`, func() {
			deleteClusterNodeOptions := &powerhaautomationservicev1.DeleteClusterNodeOptions{
				PhaInstanceID: core.StringPtr("8eefautr-4c02-0009-0086-8bd4d8cf61b6"),
				VMID:          core.StringPtr("r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"),
				IfNoneMatch:   core.StringPtr("abcdef"),
			}

			clusterNodeResponse, response, err := powerhaAutomationServiceService.DeleteClusterNode(deleteClusterNodeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterNodeResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
