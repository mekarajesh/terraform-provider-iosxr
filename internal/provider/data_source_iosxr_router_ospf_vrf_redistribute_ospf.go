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

type dataSourceRouterOSPFVRFRedistributeOSPFType struct{}

func (t dataSourceRouterOSPFVRFRedistributeOSPFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router OSPF VRF Redistribute OSPF configuration.",

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
			"instance_name": {
				MarkdownDescription: "Open Shortest Path First (OSPF)",
				Type:                types.StringType,
				Required:            true,
			},
			"match_internal": {
				MarkdownDescription: "Redistribute OSPF internal routes",
				Type:                types.BoolType,
				Computed:            true,
			},
			"match_external": {
				MarkdownDescription: "Redistribute OSPF external routes",
				Type:                types.BoolType,
				Computed:            true,
			},
			"match_nssa_external": {
				MarkdownDescription: "Redistribute OSPF NSSA external routes",
				Type:                types.BoolType,
				Computed:            true,
			},
			"tag": {
				MarkdownDescription: "Set tag for routes redistributed into OSPF",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"metric_type": {
				MarkdownDescription: "OSPF exterior metric type for redistributed routes",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceRouterOSPFVRFRedistributeOSPFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRouterOSPFVRFRedistributeOSPF{
		provider: provider,
	}, diags
}

type dataSourceRouterOSPFVRFRedistributeOSPF struct {
	provider provider
}

func (d dataSourceRouterOSPFVRFRedistributeOSPF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config RouterOSPFVRFRedistributeOSPF

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
