package sso

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	client "github.com/pingidentity/terraform-provider-pingone/internal/client"
	"github.com/pingidentity/terraform-provider-pingone/internal/sdk"
	"github.com/pingidentity/terraform-provider-pingone/internal/verify"
)

func DatasourceSchema() *schema.Resource {
	return &schema.Resource{

		// This description is used by the documentation generator and the language server.
		Description: "Datasource to read PingOne schema data",

		ReadContext: datasourcePingOneSchemaRead,

		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description:      "The ID of the environment.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"schema_id": {
				Description:      "The ID of the schema.",
				Type:             schema.TypeString,
				Optional:         true,
				ConflictsWith:    []string{"name"},
				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"name": {
				Description:   "The name of the schema.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"schema_id"},
			},
			"description": {
				Description: "A description of the schema.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func datasourcePingOneSchemaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	p1Client := meta.(*client.Client)
	apiClient := p1Client.API.ManagementAPIClient
	ctx = context.WithValue(ctx, management.ContextServerVariables, map[string]string{
		"suffix": p1Client.API.Region.URLSuffix,
	})
	var diags diag.Diagnostics

	var resp management.Schema

	if v, ok := d.GetOk("name"); ok {

		respList, diags := sdk.ParseResponse(
			ctx,

			func() (interface{}, *http.Response, error) {
				return apiClient.SchemasApi.ReadAllSchemas(ctx, d.Get("environment_id").(string)).Execute()
			},
			"ReadAllSchemas",
			sdk.DefaultCustomError,
			sdk.DefaultCreateReadRetryable,
		)
		if diags.HasError() {
			return diags
		}

		if schemas, ok := respList.(*management.EntityArray).Embedded.GetSchemasOk(); ok {

			found := false
			for _, schema := range schemas {

				if schema.GetName() == v.(string) {
					resp = schema
					found = true
					break
				}
			}

			if !found {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Cannot find schema %s", v),
				})

				return diags
			}

		}

	} else if v, ok2 := d.GetOk("schema_id"); ok2 {

		schemaResp, diags := sdk.ParseResponse(
			ctx,

			func() (interface{}, *http.Response, error) {
				return apiClient.SchemasApi.ReadOneSchema(ctx, d.Get("environment_id").(string), v.(string)).Execute()
			},
			"ReadOneSchema",
			sdk.DefaultCustomError,
			sdk.DefaultCreateReadRetryable,
		)
		if diags.HasError() {
			return diags
		}

		resp = *schemaResp.(*management.Schema)

	} else {

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Neither schema_id or name are set",
			Detail:   "Neither schema_id or name are set",
		})

		return diags

	}

	d.SetId(resp.GetId())
	d.Set("schema_id", resp.GetId())
	d.Set("name", resp.GetName())
	d.Set("description", resp.GetDescription())

	return diags
}
