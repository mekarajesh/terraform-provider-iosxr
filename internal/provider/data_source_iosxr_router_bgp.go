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

type dataSourceRouterBGPType struct{}

func (t dataSourceRouterBGPType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router BGP configuration.",

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
			"default_information_originate": {
				MarkdownDescription: "Distribute a default route",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_metric": {
				MarkdownDescription: "default redistributed metric",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"timers_bgp_keepalive_interval": {
				MarkdownDescription: "BGP timers",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"timers_bgp_holdtime": {
				MarkdownDescription: "Holdtime. Set 0 to disable keepalives/hold time.",
				Type:                types.StringType,
				Computed:            true,
			},
			"bfd_minimum_interval": {
				MarkdownDescription: "Hello interval",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"bfd_multiplier": {
				MarkdownDescription: "Detect multiplier",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRouterBGPType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRouterBGP{
		provider: provider,
	}, diags
}

type dataSourceRouterBGP struct {
	provider provider
}

func (d dataSourceRouterBGP) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config RouterBGP

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
