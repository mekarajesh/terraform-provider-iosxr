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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CDP struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	DeleteMode          types.String `tfsdk:"delete_mode"`
	Enable              types.Bool   `tfsdk:"enable"`
	Holdtime            types.Int64  `tfsdk:"holdtime"`
	Timer               types.Int64  `tfsdk:"timer"`
	AdvertiseV1         types.Bool   `tfsdk:"advertise_v1"`
	LogAdjacencyChanges types.Bool   `tfsdk:"log_adjacency_changes"`
}

type CDPData struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	Enable              types.Bool   `tfsdk:"enable"`
	Holdtime            types.Int64  `tfsdk:"holdtime"`
	Timer               types.Int64  `tfsdk:"timer"`
	AdvertiseV1         types.Bool   `tfsdk:"advertise_v1"`
	LogAdjacencyChanges types.Bool   `tfsdk:"log_adjacency_changes"`
}

func (data CDP) getPath() string {
	return "Cisco-IOS-XR-um-cdp-cfg:/cdp"
}

func (data CDPData) getPath() string {
	return "Cisco-IOS-XR-um-cdp-cfg:/cdp"
}

func (data CDP) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Enable.IsNull() && !data.Enable.IsUnknown() {
		if data.Enable.ValueBool() {
			body, _ = sjson.Set(body, "enable", map[string]string{})
		}
	}
	if !data.Holdtime.IsNull() && !data.Holdtime.IsUnknown() {
		body, _ = sjson.Set(body, "holdtime", strconv.FormatInt(data.Holdtime.ValueInt64(), 10))
	}
	if !data.Timer.IsNull() && !data.Timer.IsUnknown() {
		body, _ = sjson.Set(body, "timer", strconv.FormatInt(data.Timer.ValueInt64(), 10))
	}
	if !data.AdvertiseV1.IsNull() && !data.AdvertiseV1.IsUnknown() {
		if data.AdvertiseV1.ValueBool() {
			body, _ = sjson.Set(body, "advertise.v1", map[string]string{})
		}
	}
	if !data.LogAdjacencyChanges.IsNull() && !data.LogAdjacencyChanges.IsUnknown() {
		if data.LogAdjacencyChanges.ValueBool() {
			body, _ = sjson.Set(body, "log.adjacency.changes", map[string]string{})
		}
	}
	return body
}

func (data *CDP) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "enable"); !data.Enable.IsNull() {
		if value.Exists() {
			data.Enable = types.BoolValue(true)
		} else {
			data.Enable = types.BoolValue(false)
		}
	} else {
		data.Enable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "holdtime"); value.Exists() && !data.Holdtime.IsNull() {
		data.Holdtime = types.Int64Value(value.Int())
	} else {
		data.Holdtime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timer"); value.Exists() && !data.Timer.IsNull() {
		data.Timer = types.Int64Value(value.Int())
	} else {
		data.Timer = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "advertise.v1"); !data.AdvertiseV1.IsNull() {
		if value.Exists() {
			data.AdvertiseV1 = types.BoolValue(true)
		} else {
			data.AdvertiseV1 = types.BoolValue(false)
		}
	} else {
		data.AdvertiseV1 = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "log.adjacency.changes"); !data.LogAdjacencyChanges.IsNull() {
		if value.Exists() {
			data.LogAdjacencyChanges = types.BoolValue(true)
		} else {
			data.LogAdjacencyChanges = types.BoolValue(false)
		}
	} else {
		data.LogAdjacencyChanges = types.BoolNull()
	}
}

func (data *CDPData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "enable"); value.Exists() {
		data.Enable = types.BoolValue(true)
	} else {
		data.Enable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "holdtime"); value.Exists() {
		data.Holdtime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer"); value.Exists() {
		data.Timer = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "advertise.v1"); value.Exists() {
		data.AdvertiseV1 = types.BoolValue(true)
	} else {
		data.AdvertiseV1 = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "log.adjacency.changes"); value.Exists() {
		data.LogAdjacencyChanges = types.BoolValue(true)
	} else {
		data.LogAdjacencyChanges = types.BoolValue(false)
	}
}

func (data *CDP) getDeletedListItems(ctx context.Context, state CDP) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *CDP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Enable.IsNull() && !data.Enable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/enable", data.getPath()))
	}
	if !data.AdvertiseV1.IsNull() && !data.AdvertiseV1.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/advertise/v1", data.getPath()))
	}
	if !data.LogAdjacencyChanges.IsNull() && !data.LogAdjacencyChanges.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log/adjacency/changes", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *CDP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Enable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/enable", data.getPath()))
	}
	if !data.Holdtime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/holdtime", data.getPath()))
	}
	if !data.Timer.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timer", data.getPath()))
	}
	if !data.AdvertiseV1.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/advertise/v1", data.getPath()))
	}
	if !data.LogAdjacencyChanges.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log/adjacency/changes", data.getPath()))
	}
	return deletePaths
}
