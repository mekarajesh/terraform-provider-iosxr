// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterISISInterfaceAddressFamilyDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterISISInterfaceAddressFamilyDataSource{}
)

func NewRouterISISInterfaceAddressFamilyDataSource() datasource.DataSource {
	return &RouterISISInterfaceAddressFamilyDataSource{}
}

type RouterISISInterfaceAddressFamilyDataSource struct {
	client *client.Client
}

func (d *RouterISISInterfaceAddressFamilyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_isis_interface_address_family"
}

func (d *RouterISISInterfaceAddressFamilyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router ISIS Interface Address Family configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"process_id": schema.StringAttribute{
				MarkdownDescription: "Process ID",
				Required:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Enter the IS-IS interface configuration submode",
				Required:            true,
			},
			"af_name": schema.StringAttribute{
				MarkdownDescription: "Address family name",
				Required:            true,
			},
			"saf_name": schema.StringAttribute{
				MarkdownDescription: "Sub address family name",
				Required:            true,
			},
			"fast_reroute_per_prefix_levels": schema.ListNestedAttribute{
				MarkdownDescription: "Enable EPCFRR LFA for one level only",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"level_id": schema.Int64Attribute{
							MarkdownDescription: "Enable EPCFRR LFA for one level only",
							Computed:            true,
						},
						"ti_lfa": schema.BoolAttribute{
							MarkdownDescription: "Enable TI LFA computation",
							Computed:            true,
						},
					},
				},
			},
			"tag": schema.Int64Attribute{
				MarkdownDescription: "Set interface tag",
				Computed:            true,
			},
			"prefix_sid_absolute": schema.Int64Attribute{
				MarkdownDescription: "Specify the absolute value of Prefix Segement ID",
				Computed:            true,
			},
			"prefix_sid_n_flag_clear": schema.BoolAttribute{
				MarkdownDescription: "Clear N-flag for the prefix-SID ",
				Computed:            true,
			},
			"advertise_prefix_route_policy": schema.StringAttribute{
				MarkdownDescription: "Filter routes based on a route policy",
				Computed:            true,
			},
			"prefix_sid_index": schema.Int64Attribute{
				MarkdownDescription: "Specify the index of Prefix Segement ID",
				Computed:            true,
			},
			"prefix_sid_strict_spf_absolute": schema.Int64Attribute{
				MarkdownDescription: "Specify the absolute value of Prefix Segement ID",
				Computed:            true,
			},
		},
	}
}

func (d *RouterISISInterfaceAddressFamilyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterISISInterfaceAddressFamilyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterISISInterfaceAddressFamilyData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
