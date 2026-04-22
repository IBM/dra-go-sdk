//go:build customexamples
// +build customexamples

package powerhaautomationservicev1_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ResourceControllerV2 Examples Tests`, func() {

	var (
		resourceControllerService *resourcecontrollerv2.ResourceControllerV2

		// Configurable values
		resourceGroup        string = "8d445dfd58484a4892207123456" // Replace with your resource group ID
		serviceName          string = "power-dr-automation"         // Catalog service name
		planName             string = "powerha-aix-powervs"         // Mention your plan name for your service to create the provision.
		resourceInstanceName string = "myphasdktest"                // Desired instance name
		targetRegion         string = "us-east"                     // Desired region
		iamAPIKey            string = "<apikey>"                    // Replace with your IBM Cloud API key
	)

	Describe(`Client initialization`, func() {
		It("Successfully construct the service client instance", func() {
			options := &resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.IamAuthenticator{
					ApiKey: iamAPIKey,
				},
			}

			var err error
			resourceControllerService, err = resourcecontrollerv2.NewResourceControllerV2(options)
			Expect(err).To(BeNil())
			Expect(resourceControllerService).ToNot(BeNil())
		})
	})

	Describe(`ResourceControllerV2 request examples`, func() {

		It(`CreateResourceInstance with dynamic plan ID`, func() {
			fmt.Println("\nFetching plan ID...")

			// Step 1: Use global catalog to fetch the plan ID from the given plan name
			catalogClient, err := globalcatalogv1.NewGlobalCatalogV1(
				&globalcatalogv1.GlobalCatalogV1Options{
					Authenticator: &core.IamAuthenticator{ApiKey: iamAPIKey},
				},
			)
			Expect(err).To(BeNil())

			listOpts := catalogClient.NewListCatalogEntriesOptions()
			listOpts.SetQ(fmt.Sprintf("name:%s", serviceName))
			listOpts.SetComplete(true)

			searchResult, _, err := catalogClient.ListCatalogEntries(listOpts)
			Expect(err).To(BeNil())
			Expect(len(searchResult.Resources)).To(BeNumerically(">", 0))

			serviceEntryID := *searchResult.Resources[0].ID

			getChildOpts := catalogClient.NewGetChildObjectsOptions(serviceEntryID, "*")
			getChildOpts.SetComplete(true)

			childResult, _, err := catalogClient.GetChildObjects(getChildOpts)
			Expect(err).To(BeNil())

			// Found the plan id from the user given plan name.
			var resourcePlanID string
			for _, child := range childResult.Resources {
				if child.Name != nil && *child.Name == planName {
					resourcePlanID = *child.ID
					break
				}
			}

			if resourcePlanID == "" {
				log.Fatalf("Plan '%s' not found under service '%s'", planName, serviceName)
			}
			fmt.Printf("Found Plan ID: %s\n", resourcePlanID)

			// Step 2: Create resource instance
			fmt.Println("\nCreateResourceInstance() result:")
			createOpts := resourceControllerService.NewCreateResourceInstanceOptions(
				resourceInstanceName,
				targetRegion,
				resourceGroup,
				resourcePlanID,
			)

			resourceInstance, response, err := resourceControllerService.CreateResourceInstance(createOpts)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resourceInstance).ToNot(BeNil())

			b, _ := json.MarshalIndent(resourceInstance, "", "  ")
			fmt.Println(string(b))

			// Step 3: Fetch details of the created instance
			instanceGUID := *resourceInstance.GUID
			getOpts := resourceControllerService.NewGetResourceInstanceOptions(instanceGUID)

			instanceDetails, response, err := resourceControllerService.GetResourceInstance(getOpts)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDetails).ToNot(BeNil())

			fmt.Println("\nGetResourceInstance() result:")
			b, _ = json.MarshalIndent(instanceDetails, "", "  ")
			fmt.Println(string(b))
		})
		// Example: Deprovision (delete) a previously provisioned resource instance.
		// This prints the instance details that will be used in the delete request.
		It(`DeleteResourceInstance request example`, func() {
			// begin-delete_resource_instance
			instanceGUID := "crn:v1:bluemix:public:power-dr-automation:us-east:a/12345d444c9bd95efca704cf033:a19b6293-9f76-46d8-81b0-7a443db51887::"
			deleteResourceInstanceOptions := resourceControllerService.NewDeleteResourceInstanceOptions(
				instanceGUID,
			)
			deleteResourceInstanceOptions.SetRecursive(false)

			response, err := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_instance
			fmt.Printf("\nDeleteResourceInstance() response status code: %d\n", response.StatusCode)

			time.Sleep(20 * time.Second)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
	})
})
