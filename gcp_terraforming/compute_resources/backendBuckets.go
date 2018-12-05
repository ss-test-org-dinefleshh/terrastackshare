// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"
)

var backendBucketsIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,
}

var backendBucketsAllowEmptyValues = map[string]bool{}

var backendBucketsAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type BackendBucketsGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on backendBucketsList and create for each TerraformResource
func (BackendBucketsGenerator) createResources(backendBucketsList *compute.BackendBucketsListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := backendBucketsList.Pages(ctx, func(page *compute.BackendBucketList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_backend_bucket",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
					"region":  region,
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each backendBuckets create 1 TerraformResource
// Need backendBuckets name as ID for terraform resource
func (g BackendBucketsGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	backendBucketsList := computeService.BackendBuckets.List(project)

	resources := g.createResources(backendBucketsList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, backendBucketsIgnoreKey, backendBucketsAllowEmptyValues, backendBucketsAdditionalFields)
	return resources, metadata, nil

}