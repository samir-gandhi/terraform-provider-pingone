package base

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pingone "github.com/patrickcping/pingone-go/management"
	client "github.com/pingidentity/terraform-provider-pingone/internal/client"
)

func DatasourceEnvironment() *schema.Resource {
	return &schema.Resource{

		// This description is used by the documentation generator and the language server.
		Description: "Datasource to read PingOne environment data",

		ReadContext: datasourcePingOneEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description:   "The ID of the environment.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "license_id"},
			},
			"name": {
				Description:   "The name of the environment.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"environment_id", "license_id"},
			},
			"description": {
				Description: "A description of the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of the environment.  Options are SANDBOX for a development/testing environment and PRODUCTION for environments that require protection from deletion.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"region": {
				Description: "The region the environment is created in.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"license_id": {
				Description: "An ID of a valid license to apply to the environment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			// "solution": {
			// 	Description: "The solution context of the environment.  Blank values indicate a non-workforce solution context.  Valid options are `WORKFORCE` and `CUSTOMER`",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"service": {
				Description: "The services enabled in the environment.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "The service type.  Valid options are `SSO`, `MFA`, `RISK`, `VERIFY`, `CREDENTIALS`, `API_INTELLIGENCE`, `AUTHORIZE`, `FRAUD`, `PING_ID`, `PING_FEDERATE`, `PING_ACCESS`, `PING_DIRECTORY`, `PING_AUTHORIZE` and `PING_CENTRAL`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"console_url": {
							Description: "A custom console URL.  Generally used with services that are deployed separately to the PingOne SaaS service, such as `PING_FEDERATE`, `PING_ACCESS`, `PING_DIRECTORY`, `PING_AUTHORIZE` and `PING_CENTRAL`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"bookmark": {
							Description: "Custom bookmark links for the service.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Description: "Bookmark name.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"url": {
										Description: "Bookmark URL.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func datasourcePingOneEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	p1Client := meta.(*client.Client)
	apiClient := p1Client.API
	ctx = context.WithValue(ctx, pingone.ContextServerVariables, map[string]string{
		"suffix": p1Client.RegionSuffix,
	})
	var diags diag.Diagnostics

	var resp pingone.Environment

	if v, ok := d.GetOk("name"); ok {

		respList, r, err := apiClient.EnvironmentsApi.ReadAllEnvironments(ctx).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Error when calling `EnvironmentsApi.ReadAllEnvironments``: %v", err),
				Detail:   fmt.Sprintf("Full HTTP response: %v\n", r.Body),
			})

			return diags
		}

		if environments, ok := respList.Embedded.GetEnvironmentsOk(); ok {

			found := false
			for _, environment := range environments {

				if environment.GetName() == v.(string) {
					resp = environment
					found = true
					break
				}
			}

			if !found {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Cannot find environment %s", v),
				})

				return diags
			}

		}

	} else if v, ok2 := d.GetOk("environment_id"); ok2 {

		environmentResp, r, err := apiClient.EnvironmentsApi.ReadOneEnvironment(ctx, v.(string)).Execute()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Error when calling `EnvironmentsApi.ReadOneEnvironment``: %v", err),
				Detail:   fmt.Sprintf("Full HTTP response: %v\n", r.Body),
			})

			return diags
		}

		resp = *environmentResp

	} else {

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Neither environment_id or name are set",
			Detail:   "Neither environment_id or name are set",
		})

		return diags

	}

	d.SetId(resp.GetId())
	d.Set("environment_id", resp.GetId())
	d.Set("name", resp.GetName())

	if v, ok := resp.GetDescriptionOk(); ok {
		d.Set("description", v)
	} else {
		d.Set("description", nil)
	}

	d.Set("type", resp.GetType())
	d.Set("region", resp.GetRegion())
	d.Set("license_id", resp.GetLicense().Id)

	// The bill of materials

	servicesResp, servicesR, servicesErr := apiClient.BillOfMaterialsBOMApi.ReadOneBillOfMaterials(ctx, resp.GetId()).Execute()
	if servicesErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error when calling `EnvironmentsApi.ReadOneBillOfMaterials``: %v", servicesErr),
			Detail:   fmt.Sprintf("Full HTTP response: %v\n", servicesR),
		})

		return diags
	}

	// d.Set("solution", servicesResp.SolutionType)
	productBOMItems, err := flattenBOMProducts(servicesResp)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error mapping platform services with the configured services``: %v", err),
			Detail:   fmt.Sprintf("Platform services: %v\n", servicesResp),
		})

		return diags
	}
	d.Set("service", productBOMItems)

	return diags
}
