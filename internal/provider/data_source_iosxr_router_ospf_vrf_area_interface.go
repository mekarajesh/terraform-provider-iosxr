// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceRouterOSPFVRFAreaInterfaceType struct{}

func (t dataSourceRouterOSPFVRFAreaInterfaceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router OSPF VRF Area Interface configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"process_name": {
				MarkdownDescription: "Name for this OSPF process",
				Type:                types.StringType,
				Required:            true,
			},
			"vrf_name": {
				MarkdownDescription: "Name for this OSPF vrf",
				Type:                types.StringType,
				Required:            true,
			},
			"area_id": {
				MarkdownDescription: "Enter the OSPF area configuration submode",
				Type:                types.StringType,
				Required:            true,
			},
			"interface_name": {
				MarkdownDescription: "Enable routing on an interface ",
				Type:                types.StringType,
				Required:            true,
			},
			"network_broadcast": {
				MarkdownDescription: "Specify OSPF broadcast multi-access network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_non_broadcast": {
				MarkdownDescription: "Specify OSPF NBMA network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_point_to_point": {
				MarkdownDescription: "Specify OSPF point-to-point network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_point_to_multipoint": {
				MarkdownDescription: "Specify OSPF point-to-multipoint network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"cost": {
				MarkdownDescription: "Interface cost",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"priority": {
				MarkdownDescription: "Router priority",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"passive_enable": {
				MarkdownDescription: "Enable passive",
				Type:                types.BoolType,
				Computed:            true,
			},
			"passive_disable": {
				MarkdownDescription: "Disable passive",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRouterOSPFVRFAreaInterfaceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRouterOSPFVRFAreaInterface{
		provider: provider,
	}, diags
}

type dataSourceRouterOSPFVRFAreaInterface struct {
	provider provider
}

func (d dataSourceRouterOSPFVRFAreaInterface) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config RouterOSPFVRFAreaInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.provider.client.Get(ctx, config.Device.Value, config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.String{Value: config.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
