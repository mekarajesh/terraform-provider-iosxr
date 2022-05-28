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

type dataSourceRouterOSPFType struct{}

func (t dataSourceRouterOSPFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router OSPF configuration.",

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
			"mpls_ldp_sync": {
				MarkdownDescription: "Enable LDP IGP synchronization",
				Type:                types.BoolType,
				Computed:            true,
			},
			"hello_interval": {
				MarkdownDescription: "Time between HELLO packets",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"dead_interval": {
				MarkdownDescription: "Interval after which a neighbor is declared dead",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"priority": {
				MarkdownDescription: "Router priority",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"mtu_ignore_enable": {
				MarkdownDescription: "Ignores the MTU in DBD packets",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mtu_ignore_disable": {
				MarkdownDescription: "Disable ignoring the MTU in DBD packets",
				Type:                types.BoolType,
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
			"router_id": {
				MarkdownDescription: "configure this node",
				Type:                types.StringType,
				Computed:            true,
			},
			"redistribute_connected": {
				MarkdownDescription: "Connected routes",
				Type:                types.BoolType,
				Computed:            true,
			},
			"redistribute_connected_tag": {
				MarkdownDescription: "Set tag for routes redistributed into OSPF",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"redistribute_connected_metric_type": {
				MarkdownDescription: "OSPF exterior metric type for redistributed routes",
				Type:                types.StringType,
				Computed:            true,
			},
			"redistribute_static": {
				MarkdownDescription: "Static routes",
				Type:                types.BoolType,
				Computed:            true,
			},
			"redistribute_static_tag": {
				MarkdownDescription: "Set tag for routes redistributed into OSPF",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"redistribute_static_metric_type": {
				MarkdownDescription: "OSPF exterior metric type for redistributed routes",
				Type:                types.StringType,
				Computed:            true,
			},
			"bfd_fast_detect": {
				MarkdownDescription: "Enable Fast detection",
				Type:                types.BoolType,
				Computed:            true,
			},
			"bfd_minimum_interval": {
				MarkdownDescription: "Minimum interval",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"bfd_multiplier": {
				MarkdownDescription: "Detect multiplier",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"default_information_originate": {
				MarkdownDescription: "Distribute a default route",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_information_originate_always": {
				MarkdownDescription: "Always advertise default route",
				Type:                types.BoolType,
				Computed:            true,
			},
			"default_information_originate_metric_type": {
				MarkdownDescription: "OSPF metric type for default routes",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRouterOSPFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRouterOSPF{
		provider: provider,
	}, diags
}

type dataSourceRouterOSPF struct {
	provider provider
}

func (d dataSourceRouterOSPF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config RouterOSPF

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
