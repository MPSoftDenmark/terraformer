// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"strings"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var targetTcpProxiesIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,

	"proxy_id": true,
}

var targetTcpProxiesAllowEmptyValues = map[string]bool{}

var targetTcpProxiesAdditionalFields = map[string]string{
	"project": "waze-development",
}

type TargetTcpProxiesGenerator struct {
	gcp_generator.BasicGenerator
}

func (TargetTcpProxiesGenerator) createResources(targetTcpProxiesList *compute.TargetTcpProxiesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := targetTcpProxiesList.Pages(ctx, func(page *compute.TargetTcpProxyList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_target_tcp_proxy",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": "waze-development",
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

func (g TargetTcpProxiesGenerator) Generate(zone string) error {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	targetTcpProxiesList := computeService.TargetTcpProxies.List(project)

	resources := g.createResources(targetTcpProxiesList, ctx, region, zone)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, targetTcpProxiesIgnoreKey, targetTcpProxiesAllowEmptyValues, targetTcpProxiesAdditionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "targetTcpProxies", region, "google")
	if err != nil {
		return err
	}
	return nil

}