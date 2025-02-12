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

type RouterISISInterface struct {
	Device                types.String `tfsdk:"device"`
	Id                    types.String `tfsdk:"id"`
	DeleteMode            types.String `tfsdk:"delete_mode"`
	ProcessId             types.String `tfsdk:"process_id"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	CircuitType           types.String `tfsdk:"circuit_type"`
	HelloPaddingDisable   types.Bool   `tfsdk:"hello_padding_disable"`
	HelloPaddingSometimes types.Bool   `tfsdk:"hello_padding_sometimes"`
	Priority              types.Int64  `tfsdk:"priority"`
	PointToPoint          types.Bool   `tfsdk:"point_to_point"`
	Passive               types.Bool   `tfsdk:"passive"`
	Suppressed            types.Bool   `tfsdk:"suppressed"`
	Shutdown              types.Bool   `tfsdk:"shutdown"`
	HelloPasswordText     types.String `tfsdk:"hello_password_text"`
	HelloPasswordHmacMd5  types.String `tfsdk:"hello_password_hmac_md5"`
	HelloPasswordKeychain types.String `tfsdk:"hello_password_keychain"`
	BfdFastDetectIpv6     types.Bool   `tfsdk:"bfd_fast_detect_ipv6"`
}

type RouterISISInterfaceData struct {
	Device                types.String `tfsdk:"device"`
	Id                    types.String `tfsdk:"id"`
	ProcessId             types.String `tfsdk:"process_id"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	CircuitType           types.String `tfsdk:"circuit_type"`
	HelloPaddingDisable   types.Bool   `tfsdk:"hello_padding_disable"`
	HelloPaddingSometimes types.Bool   `tfsdk:"hello_padding_sometimes"`
	Priority              types.Int64  `tfsdk:"priority"`
	PointToPoint          types.Bool   `tfsdk:"point_to_point"`
	Passive               types.Bool   `tfsdk:"passive"`
	Suppressed            types.Bool   `tfsdk:"suppressed"`
	Shutdown              types.Bool   `tfsdk:"shutdown"`
	HelloPasswordText     types.String `tfsdk:"hello_password_text"`
	HelloPasswordHmacMd5  types.String `tfsdk:"hello_password_hmac_md5"`
	HelloPasswordKeychain types.String `tfsdk:"hello_password_keychain"`
	BfdFastDetectIpv6     types.Bool   `tfsdk:"bfd_fast_detect_ipv6"`
}

func (data RouterISISInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]", data.ProcessId.ValueString(), data.InterfaceName.ValueString())
}

func (data RouterISISInterfaceData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]", data.ProcessId.ValueString(), data.InterfaceName.ValueString())
}

func (data RouterISISInterface) toBody(ctx context.Context) string {
	body := "{}"
	if !data.InterfaceName.IsNull() && !data.InterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "interface-name", data.InterfaceName.ValueString())
	}
	if !data.CircuitType.IsNull() && !data.CircuitType.IsUnknown() {
		body, _ = sjson.Set(body, "circuit-type", data.CircuitType.ValueString())
	}
	if !data.HelloPaddingDisable.IsNull() && !data.HelloPaddingDisable.IsUnknown() {
		if data.HelloPaddingDisable.ValueBool() {
			body, _ = sjson.Set(body, "hello-padding.disable", map[string]string{})
		}
	}
	if !data.HelloPaddingSometimes.IsNull() && !data.HelloPaddingSometimes.IsUnknown() {
		if data.HelloPaddingSometimes.ValueBool() {
			body, _ = sjson.Set(body, "hello-padding.sometimes", map[string]string{})
		}
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		body, _ = sjson.Set(body, "priority.priority-value", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}
	if !data.PointToPoint.IsNull() && !data.PointToPoint.IsUnknown() {
		if data.PointToPoint.ValueBool() {
			body, _ = sjson.Set(body, "point-to-point", map[string]string{})
		}
	}
	if !data.Passive.IsNull() && !data.Passive.IsUnknown() {
		if data.Passive.ValueBool() {
			body, _ = sjson.Set(body, "passive", map[string]string{})
		}
	}
	if !data.Suppressed.IsNull() && !data.Suppressed.IsUnknown() {
		if data.Suppressed.ValueBool() {
			body, _ = sjson.Set(body, "suppressed", map[string]string{})
		}
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, "shutdown", map[string]string{})
		}
	}
	if !data.HelloPasswordText.IsNull() && !data.HelloPasswordText.IsUnknown() {
		body, _ = sjson.Set(body, "hello-password.text.encrypted", data.HelloPasswordText.ValueString())
	}
	if !data.HelloPasswordHmacMd5.IsNull() && !data.HelloPasswordHmacMd5.IsUnknown() {
		body, _ = sjson.Set(body, "hello-password.hmac-md5.encrypted", data.HelloPasswordHmacMd5.ValueString())
	}
	if !data.HelloPasswordKeychain.IsNull() && !data.HelloPasswordKeychain.IsUnknown() {
		body, _ = sjson.Set(body, "hello-password.keychain.keychain-name", data.HelloPasswordKeychain.ValueString())
	}
	if !data.BfdFastDetectIpv6.IsNull() && !data.BfdFastDetectIpv6.IsUnknown() {
		if data.BfdFastDetectIpv6.ValueBool() {
			body, _ = sjson.Set(body, "bfd.fast-detect.ipv6", map[string]string{})
		}
	}
	return body
}

func (data *RouterISISInterface) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "circuit-type"); value.Exists() && !data.CircuitType.IsNull() {
		data.CircuitType = types.StringValue(value.String())
	} else {
		data.CircuitType = types.StringNull()
	}
	if value := gjson.GetBytes(res, "hello-padding.disable"); !data.HelloPaddingDisable.IsNull() {
		if value.Exists() {
			data.HelloPaddingDisable = types.BoolValue(true)
		} else {
			data.HelloPaddingDisable = types.BoolValue(false)
		}
	} else {
		data.HelloPaddingDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "hello-padding.sometimes"); !data.HelloPaddingSometimes.IsNull() {
		if value.Exists() {
			data.HelloPaddingSometimes = types.BoolValue(true)
		} else {
			data.HelloPaddingSometimes = types.BoolValue(false)
		}
	} else {
		data.HelloPaddingSometimes = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "priority.priority-value"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "point-to-point"); !data.PointToPoint.IsNull() {
		if value.Exists() {
			data.PointToPoint = types.BoolValue(true)
		} else {
			data.PointToPoint = types.BoolValue(false)
		}
	} else {
		data.PointToPoint = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "passive"); !data.Passive.IsNull() {
		if value.Exists() {
			data.Passive = types.BoolValue(true)
		} else {
			data.Passive = types.BoolValue(false)
		}
	} else {
		data.Passive = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "suppressed"); !data.Suppressed.IsNull() {
		if value.Exists() {
			data.Suppressed = types.BoolValue(true)
		} else {
			data.Suppressed = types.BoolValue(false)
		}
	} else {
		data.Suppressed = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "hello-password.text.encrypted"); value.Exists() && !data.HelloPasswordText.IsNull() {
		data.HelloPasswordText = types.StringValue(value.String())
	} else {
		data.HelloPasswordText = types.StringNull()
	}
	if value := gjson.GetBytes(res, "hello-password.hmac-md5.encrypted"); value.Exists() && !data.HelloPasswordHmacMd5.IsNull() {
		data.HelloPasswordHmacMd5 = types.StringValue(value.String())
	} else {
		data.HelloPasswordHmacMd5 = types.StringNull()
	}
	if value := gjson.GetBytes(res, "hello-password.keychain.keychain-name"); value.Exists() && !data.HelloPasswordKeychain.IsNull() {
		data.HelloPasswordKeychain = types.StringValue(value.String())
	} else {
		data.HelloPasswordKeychain = types.StringNull()
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.ipv6"); !data.BfdFastDetectIpv6.IsNull() {
		if value.Exists() {
			data.BfdFastDetectIpv6 = types.BoolValue(true)
		} else {
			data.BfdFastDetectIpv6 = types.BoolValue(false)
		}
	} else {
		data.BfdFastDetectIpv6 = types.BoolNull()
	}
}

func (data *RouterISISInterfaceData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "circuit-type"); value.Exists() {
		data.CircuitType = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "hello-padding.disable"); value.Exists() {
		data.HelloPaddingDisable = types.BoolValue(true)
	} else {
		data.HelloPaddingDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "hello-padding.sometimes"); value.Exists() {
		data.HelloPaddingSometimes = types.BoolValue(true)
	} else {
		data.HelloPaddingSometimes = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "priority.priority-value"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "point-to-point"); value.Exists() {
		data.PointToPoint = types.BoolValue(true)
	} else {
		data.PointToPoint = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "passive"); value.Exists() {
		data.Passive = types.BoolValue(true)
	} else {
		data.Passive = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "suppressed"); value.Exists() {
		data.Suppressed = types.BoolValue(true)
	} else {
		data.Suppressed = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "hello-password.text.encrypted"); value.Exists() {
		data.HelloPasswordText = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "hello-password.hmac-md5.encrypted"); value.Exists() {
		data.HelloPasswordHmacMd5 = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "hello-password.keychain.keychain-name"); value.Exists() {
		data.HelloPasswordKeychain = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "bfd.fast-detect.ipv6"); value.Exists() {
		data.BfdFastDetectIpv6 = types.BoolValue(true)
	} else {
		data.BfdFastDetectIpv6 = types.BoolValue(false)
	}
}

func (data *RouterISISInterface) getDeletedListItems(ctx context.Context, state RouterISISInterface) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RouterISISInterface) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.HelloPaddingDisable.IsNull() && !data.HelloPaddingDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/hello-padding/disable", data.getPath()))
	}
	if !data.HelloPaddingSometimes.IsNull() && !data.HelloPaddingSometimes.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/hello-padding/sometimes", data.getPath()))
	}
	if !data.PointToPoint.IsNull() && !data.PointToPoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/point-to-point", data.getPath()))
	}
	if !data.Passive.IsNull() && !data.Passive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/passive", data.getPath()))
	}
	if !data.Suppressed.IsNull() && !data.Suppressed.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/suppressed", data.getPath()))
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.BfdFastDetectIpv6.IsNull() && !data.BfdFastDetectIpv6.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/bfd/fast-detect/ipv6", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *RouterISISInterface) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.CircuitType.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/circuit-type", data.getPath()))
	}
	if !data.HelloPaddingDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hello-padding/disable", data.getPath()))
	}
	if !data.HelloPaddingSometimes.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hello-padding/sometimes", data.getPath()))
	}
	if !data.Priority.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priority/priority-value", data.getPath()))
	}
	if !data.PointToPoint.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/point-to-point", data.getPath()))
	}
	if !data.Passive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/passive", data.getPath()))
	}
	if !data.Suppressed.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/suppressed", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.BfdFastDetectIpv6.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/fast-detect/ipv6", data.getPath()))
	}
	return deletePaths
}
