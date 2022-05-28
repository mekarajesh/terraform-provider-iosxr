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

type dataSourceRouterBGPVRFAddressFamilyAggregateAddressType struct{}

func (t dataSourceRouterBGPVRFAddressFamilyAggregateAddressType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP VRF Address Family Aggregate Address configuration.",

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
			"as_number": {
				MarkdownDescription: "bgp as-number",
				Type:                types.StringType,
				Required:            true,
			},
			"vrf_name": {
				MarkdownDescription: "Specify a vrf name",
				Type:                types.StringType,
				Required:            true,
			},
			"af_name": {
				MarkdownDescription: "Enter Address Family command mode",
				Type:                types.StringType,
				Required:            true,
			},
			"address": {
				MarkdownDescription: "IPv6 Aggregate address and mask or masklength",
				Type:                types.StringType,
				Required:            true,
			},
			"masklength": {
				MarkdownDescription: "Network in prefix/length format (prefix part)",
				Type:                types.Int64Type,
				Required:            true,
			},
			"as_set": {
				MarkdownDescription: "Generate AS set path information",
				Type:                types.BoolType,
				Computed:            true,
			},
			"as_confed_set": {
				MarkdownDescription: "Generate AS confed set path information",
				Type:                types.BoolType,
				Computed:            true,
			},
			"summary_only": {
				MarkdownDescription: "Filter more specific routes from updates",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRouterBGPVRFAddressFamilyAggregateAddressType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRouterBGPVRFAddressFamilyAggregateAddress{
		provider: provider,
	}, diags
}

type dataSourceRouterBGPVRFAddressFamilyAggregateAddress struct {
	provider provider
}

func (d dataSourceRouterBGPVRFAddressFamilyAggregateAddress) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config RouterBGPVRFAddressFamilyAggregateAddress

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
