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
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PolicyMapQoS struct {
	Device        types.String          `tfsdk:"device"`
	Id            types.String          `tfsdk:"id"`
	PolicyMapName types.String          `tfsdk:"policy_map_name"`
	Description   types.String          `tfsdk:"description"`
	Classes       []PolicyMapQoSClasses `tfsdk:"classes"`
}

type PolicyMapQoSData struct {
	Device        types.String          `tfsdk:"device"`
	Id            types.String          `tfsdk:"id"`
	PolicyMapName types.String          `tfsdk:"policy_map_name"`
	Description   types.String          `tfsdk:"description"`
	Classes       []PolicyMapQoSClasses `tfsdk:"classes"`
}
type PolicyMapQoSClasses struct {
	Name                        types.String                     `tfsdk:"name"`
	Type                        types.String                     `tfsdk:"type"`
	SetMplsExperimentalTopmost  types.Int64                      `tfsdk:"set_mpls_experimental_topmost"`
	SetDscp                     types.String                     `tfsdk:"set_dscp"`
	PriorityLevel               types.Int64                      `tfsdk:"priority_level"`
	QueueLimits                 []PolicyMapQoSClassesQueueLimits `tfsdk:"queue_limits"`
	ServicePolicyName           types.String                     `tfsdk:"service_policy_name"`
	PoliceRateValue             types.String                     `tfsdk:"police_rate_value"`
	PoliceRateUnit              types.String                     `tfsdk:"police_rate_unit"`
	PoliceConformActionTransmit types.Bool                       `tfsdk:"police_conform_action_transmit"`
	PoliceConformActionDrop     types.Bool                       `tfsdk:"police_conform_action_drop"`
	PoliceExceedActionTransmit  types.Bool                       `tfsdk:"police_exceed_action_transmit"`
	PoliceExceedActionDrop      types.Bool                       `tfsdk:"police_exceed_action_drop"`
	PoliceViolateActionTransmit types.Bool                       `tfsdk:"police_violate_action_transmit"`
	PoliceViolateActionDrop     types.Bool                       `tfsdk:"police_violate_action_drop"`
	ShapeAverageRateValue       types.String                     `tfsdk:"shape_average_rate_value"`
	ShapeAverageRateUnit        types.String                     `tfsdk:"shape_average_rate_unit"`
	BandwidthRemainingUnit      types.String                     `tfsdk:"bandwidth_remaining_unit"`
	BandwidthRemainingValue     types.String                     `tfsdk:"bandwidth_remaining_value"`
}
type PolicyMapQoSClassesQueueLimits struct {
	Value types.String `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func (data PolicyMapQoS) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=%s]", data.PolicyMapName.ValueString())
}

func (data PolicyMapQoSData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=%s]", data.PolicyMapName.ValueString())
}

func (data PolicyMapQoS) toBody(ctx context.Context) string {
	body := "{}"
	if !data.PolicyMapName.IsNull() && !data.PolicyMapName.IsUnknown() {
		body, _ = sjson.Set(body, "policy-map-name", data.PolicyMapName.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Classes) > 0 {
		body, _ = sjson.Set(body, "class", []interface{}{})
		for index, item := range data.Classes {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Type.IsNull() && !item.Type.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"type", item.Type.ValueString())
			}
			if !item.SetMplsExperimentalTopmost.IsNull() && !item.SetMplsExperimentalTopmost.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"set.mpls.experimental.topmost", strconv.FormatInt(item.SetMplsExperimentalTopmost.ValueInt64(), 10))
			}
			if !item.SetDscp.IsNull() && !item.SetDscp.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"set.dscp", item.SetDscp.ValueString())
			}
			if !item.PriorityLevel.IsNull() && !item.PriorityLevel.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"priority.level", strconv.FormatInt(item.PriorityLevel.ValueInt64(), 10))
			}
			if !item.ServicePolicyName.IsNull() && !item.ServicePolicyName.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"service-policy.name", item.ServicePolicyName.ValueString())
			}
			if !item.PoliceRateValue.IsNull() && !item.PoliceRateValue.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.rate.value", item.PoliceRateValue.ValueString())
			}
			if !item.PoliceRateUnit.IsNull() && !item.PoliceRateUnit.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.rate.unit", item.PoliceRateUnit.ValueString())
			}
			if !item.PoliceConformActionTransmit.IsNull() && !item.PoliceConformActionTransmit.IsUnknown() {
				if item.PoliceConformActionTransmit.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.conform-action.transmit", map[string]string{})
				}
			}
			if !item.PoliceConformActionDrop.IsNull() && !item.PoliceConformActionDrop.IsUnknown() {
				if item.PoliceConformActionDrop.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.conform-action.drop", map[string]string{})
				}
			}
			if !item.PoliceExceedActionTransmit.IsNull() && !item.PoliceExceedActionTransmit.IsUnknown() {
				if item.PoliceExceedActionTransmit.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.exceed-action.transmit", map[string]string{})
				}
			}
			if !item.PoliceExceedActionDrop.IsNull() && !item.PoliceExceedActionDrop.IsUnknown() {
				if item.PoliceExceedActionDrop.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.exceed-action.drop", map[string]string{})
				}
			}
			if !item.PoliceViolateActionTransmit.IsNull() && !item.PoliceViolateActionTransmit.IsUnknown() {
				if item.PoliceViolateActionTransmit.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.violate-action.transmit", map[string]string{})
				}
			}
			if !item.PoliceViolateActionDrop.IsNull() && !item.PoliceViolateActionDrop.IsUnknown() {
				if item.PoliceViolateActionDrop.ValueBool() {
					body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"police.violate-action.drop", map[string]string{})
				}
			}
			if !item.ShapeAverageRateValue.IsNull() && !item.ShapeAverageRateValue.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"shape.average.rate.value", item.ShapeAverageRateValue.ValueString())
			}
			if !item.ShapeAverageRateUnit.IsNull() && !item.ShapeAverageRateUnit.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"shape.average.rate.unit", item.ShapeAverageRateUnit.ValueString())
			}
			if !item.BandwidthRemainingUnit.IsNull() && !item.BandwidthRemainingUnit.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"bandwidth-remaining.unit", item.BandwidthRemainingUnit.ValueString())
			}
			if !item.BandwidthRemainingValue.IsNull() && !item.BandwidthRemainingValue.IsUnknown() {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"bandwidth-remaining.value", item.BandwidthRemainingValue.ValueString())
			}
			if len(item.QueueLimits) > 0 {
				body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"queue-limits.queue-limit", []interface{}{})
				for cindex, citem := range item.QueueLimits {
					if !citem.Value.IsNull() && !citem.Value.IsUnknown() {
						body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"queue-limits.queue-limit"+"."+strconv.Itoa(cindex)+"."+"value", citem.Value.ValueString())
					}
					if !citem.Unit.IsNull() && !citem.Unit.IsUnknown() {
						body, _ = sjson.Set(body, "class"+"."+strconv.Itoa(index)+"."+"queue-limits.queue-limit"+"."+strconv.Itoa(cindex)+"."+"unit", citem.Unit.ValueString())
					}
				}
			}
		}
	}
	return body
}

func (data *PolicyMapQoS) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Classes {
		keys := [...]string{"name", "type"}
		keyValues := [...]string{data.Classes[i].Name.ValueString(), data.Classes[i].Type.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "class").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.Classes[i].Name.IsNull() {
			data.Classes[i].Name = types.StringValue(value.String())
		} else {
			data.Classes[i].Name = types.StringNull()
		}
		if value := r.Get("type"); value.Exists() && !data.Classes[i].Type.IsNull() {
			data.Classes[i].Type = types.StringValue(value.String())
		} else {
			data.Classes[i].Type = types.StringNull()
		}
		if value := r.Get("set.mpls.experimental.topmost"); value.Exists() && !data.Classes[i].SetMplsExperimentalTopmost.IsNull() {
			data.Classes[i].SetMplsExperimentalTopmost = types.Int64Value(value.Int())
		} else {
			data.Classes[i].SetMplsExperimentalTopmost = types.Int64Null()
		}
		if value := r.Get("set.dscp"); value.Exists() && !data.Classes[i].SetDscp.IsNull() {
			data.Classes[i].SetDscp = types.StringValue(value.String())
		} else {
			data.Classes[i].SetDscp = types.StringNull()
		}
		if value := r.Get("priority.level"); value.Exists() && !data.Classes[i].PriorityLevel.IsNull() {
			data.Classes[i].PriorityLevel = types.Int64Value(value.Int())
		} else {
			data.Classes[i].PriorityLevel = types.Int64Null()
		}
		for ci := range data.Classes[i].QueueLimits {
			keys := [...]string{"value", "unit"}
			keyValues := [...]string{data.Classes[i].QueueLimits[ci].Value.ValueString(), data.Classes[i].QueueLimits[ci].Unit.ValueString()}

			var cr gjson.Result
			r.Get("queue-limits.queue-limit").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("value"); value.Exists() && !data.Classes[i].QueueLimits[ci].Value.IsNull() {
				data.Classes[i].QueueLimits[ci].Value = types.StringValue(value.String())
			} else {
				data.Classes[i].QueueLimits[ci].Value = types.StringNull()
			}
			if value := cr.Get("unit"); value.Exists() && !data.Classes[i].QueueLimits[ci].Unit.IsNull() {
				data.Classes[i].QueueLimits[ci].Unit = types.StringValue(value.String())
			} else {
				data.Classes[i].QueueLimits[ci].Unit = types.StringNull()
			}
		}
		if value := r.Get("service-policy.name"); value.Exists() && !data.Classes[i].ServicePolicyName.IsNull() {
			data.Classes[i].ServicePolicyName = types.StringValue(value.String())
		} else {
			data.Classes[i].ServicePolicyName = types.StringNull()
		}
		if value := r.Get("police.rate.value"); value.Exists() && !data.Classes[i].PoliceRateValue.IsNull() {
			data.Classes[i].PoliceRateValue = types.StringValue(value.String())
		} else {
			data.Classes[i].PoliceRateValue = types.StringNull()
		}
		if value := r.Get("police.rate.unit"); value.Exists() && !data.Classes[i].PoliceRateUnit.IsNull() {
			data.Classes[i].PoliceRateUnit = types.StringValue(value.String())
		} else {
			data.Classes[i].PoliceRateUnit = types.StringNull()
		}
		if value := r.Get("police.conform-action.transmit"); !data.Classes[i].PoliceConformActionTransmit.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceConformActionTransmit = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceConformActionTransmit = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceConformActionTransmit = types.BoolNull()
		}
		if value := r.Get("police.conform-action.drop"); !data.Classes[i].PoliceConformActionDrop.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceConformActionDrop = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceConformActionDrop = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceConformActionDrop = types.BoolNull()
		}
		if value := r.Get("police.exceed-action.transmit"); !data.Classes[i].PoliceExceedActionTransmit.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceExceedActionTransmit = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceExceedActionTransmit = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceExceedActionTransmit = types.BoolNull()
		}
		if value := r.Get("police.exceed-action.drop"); !data.Classes[i].PoliceExceedActionDrop.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceExceedActionDrop = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceExceedActionDrop = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceExceedActionDrop = types.BoolNull()
		}
		if value := r.Get("police.violate-action.transmit"); !data.Classes[i].PoliceViolateActionTransmit.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceViolateActionTransmit = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceViolateActionTransmit = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceViolateActionTransmit = types.BoolNull()
		}
		if value := r.Get("police.violate-action.drop"); !data.Classes[i].PoliceViolateActionDrop.IsNull() {
			if value.Exists() {
				data.Classes[i].PoliceViolateActionDrop = types.BoolValue(true)
			} else {
				data.Classes[i].PoliceViolateActionDrop = types.BoolValue(false)
			}
		} else {
			data.Classes[i].PoliceViolateActionDrop = types.BoolNull()
		}
		if value := r.Get("shape.average.rate.value"); value.Exists() && !data.Classes[i].ShapeAverageRateValue.IsNull() {
			data.Classes[i].ShapeAverageRateValue = types.StringValue(value.String())
		} else {
			data.Classes[i].ShapeAverageRateValue = types.StringNull()
		}
		if value := r.Get("shape.average.rate.unit"); value.Exists() && !data.Classes[i].ShapeAverageRateUnit.IsNull() {
			data.Classes[i].ShapeAverageRateUnit = types.StringValue(value.String())
		} else {
			data.Classes[i].ShapeAverageRateUnit = types.StringNull()
		}
		if value := r.Get("bandwidth-remaining.unit"); value.Exists() && !data.Classes[i].BandwidthRemainingUnit.IsNull() {
			data.Classes[i].BandwidthRemainingUnit = types.StringValue(value.String())
		} else {
			data.Classes[i].BandwidthRemainingUnit = types.StringNull()
		}
		if value := r.Get("bandwidth-remaining.value"); value.Exists() && !data.Classes[i].BandwidthRemainingValue.IsNull() {
			data.Classes[i].BandwidthRemainingValue = types.StringValue(value.String())
		} else {
			data.Classes[i].BandwidthRemainingValue = types.StringNull()
		}
	}
}

func (data *PolicyMapQoSData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "class"); value.Exists() {
		data.Classes = make([]PolicyMapQoSClasses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyMapQoSClasses{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			}
			if cValue := v.Get("set.mpls.experimental.topmost"); cValue.Exists() {
				item.SetMplsExperimentalTopmost = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("set.dscp"); cValue.Exists() {
				item.SetDscp = types.StringValue(cValue.String())
			}
			if cValue := v.Get("priority.level"); cValue.Exists() {
				item.PriorityLevel = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("queue-limits.queue-limit"); cValue.Exists() {
				item.QueueLimits = make([]PolicyMapQoSClassesQueueLimits, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyMapQoSClassesQueueLimits{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("unit"); ccValue.Exists() {
						cItem.Unit = types.StringValue(ccValue.String())
					}
					item.QueueLimits = append(item.QueueLimits, cItem)
					return true
				})
			}
			if cValue := v.Get("service-policy.name"); cValue.Exists() {
				item.ServicePolicyName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("police.rate.value"); cValue.Exists() {
				item.PoliceRateValue = types.StringValue(cValue.String())
			}
			if cValue := v.Get("police.rate.unit"); cValue.Exists() {
				item.PoliceRateUnit = types.StringValue(cValue.String())
			}
			if cValue := v.Get("police.conform-action.transmit"); cValue.Exists() {
				item.PoliceConformActionTransmit = types.BoolValue(true)
			} else {
				item.PoliceConformActionTransmit = types.BoolValue(false)
			}
			if cValue := v.Get("police.conform-action.drop"); cValue.Exists() {
				item.PoliceConformActionDrop = types.BoolValue(true)
			} else {
				item.PoliceConformActionDrop = types.BoolValue(false)
			}
			if cValue := v.Get("police.exceed-action.transmit"); cValue.Exists() {
				item.PoliceExceedActionTransmit = types.BoolValue(true)
			} else {
				item.PoliceExceedActionTransmit = types.BoolValue(false)
			}
			if cValue := v.Get("police.exceed-action.drop"); cValue.Exists() {
				item.PoliceExceedActionDrop = types.BoolValue(true)
			} else {
				item.PoliceExceedActionDrop = types.BoolValue(false)
			}
			if cValue := v.Get("police.violate-action.transmit"); cValue.Exists() {
				item.PoliceViolateActionTransmit = types.BoolValue(true)
			} else {
				item.PoliceViolateActionTransmit = types.BoolValue(false)
			}
			if cValue := v.Get("police.violate-action.drop"); cValue.Exists() {
				item.PoliceViolateActionDrop = types.BoolValue(true)
			} else {
				item.PoliceViolateActionDrop = types.BoolValue(false)
			}
			if cValue := v.Get("shape.average.rate.value"); cValue.Exists() {
				item.ShapeAverageRateValue = types.StringValue(cValue.String())
			}
			if cValue := v.Get("shape.average.rate.unit"); cValue.Exists() {
				item.ShapeAverageRateUnit = types.StringValue(cValue.String())
			}
			if cValue := v.Get("bandwidth-remaining.unit"); cValue.Exists() {
				item.BandwidthRemainingUnit = types.StringValue(cValue.String())
			}
			if cValue := v.Get("bandwidth-remaining.value"); cValue.Exists() {
				item.BandwidthRemainingValue = types.StringValue(cValue.String())
			}
			data.Classes = append(data.Classes, item)
			return true
		})
	}
}

func (data *PolicyMapQoS) getDeletedListItems(ctx context.Context, state PolicyMapQoS) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Classes {
		keys := [...]string{"name", "type"}
		stateKeyValues := [...]string{state.Classes[i].Name.ValueString(), state.Classes[i].Type.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Classes[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Classes[i].Type.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Classes {
			found = true
			if state.Classes[i].Name.ValueString() != data.Classes[j].Name.ValueString() {
				found = false
			}
			if state.Classes[i].Type.ValueString() != data.Classes[j].Type.ValueString() {
				found = false
			}
			if found {
				for ci := range state.Classes[i].QueueLimits {
					ckeys := [...]string{"value", "unit"}
					cstateKeyValues := [...]string{state.Classes[i].QueueLimits[ci].Value.ValueString(), state.Classes[i].QueueLimits[ci].Unit.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.Classes[i].QueueLimits[ci].Value.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if !reflect.ValueOf(state.Classes[i].QueueLimits[ci].Unit.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.Classes[j].QueueLimits {
						found = true
						if state.Classes[i].QueueLimits[ci].Value.ValueString() != data.Classes[j].QueueLimits[cj].Value.ValueString() {
							found = false
						}
						if state.Classes[i].QueueLimits[ci].Unit.ValueString() != data.Classes[j].QueueLimits[cj].Unit.ValueString() {
							found = false
						}
						if found {
							break
						}
					}
					if !found {
						keyString := ""
						for ki := range keys {
							keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
						}
						ckeyString := ""
						for cki := range ckeys {
							ckeyString += "[" + ckeys[cki] + "=" + cstateKeyValues[cki] + "]"
						}
						deletedListItems = append(deletedListItems, fmt.Sprintf("%v/class%v/queue-limits/queue-limit%v", state.getPath(), keyString, ckeyString))
					}
				}
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/class%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *PolicyMapQoS) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Classes {
		keys := [...]string{"name", "type"}
		keyValues := [...]string{data.Classes[i].Name.ValueString(), data.Classes[i].Type.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		for ci := range data.Classes[i].QueueLimits {
			ckeys := [...]string{"value", "unit"}
			ckeyValues := [...]string{data.Classes[i].QueueLimits[ci].Value.ValueString(), data.Classes[i].QueueLimits[ci].Unit.ValueString()}
			ckeyString := ""
			for cki := range ckeys {
				ckeyString += "[" + ckeys[cki] + "=" + ckeyValues[cki] + "]"
			}
		}
		if !data.Classes[i].PoliceConformActionTransmit.IsNull() && !data.Classes[i].PoliceConformActionTransmit.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/conform-action/transmit", data.getPath(), keyString))
		}
		if !data.Classes[i].PoliceConformActionDrop.IsNull() && !data.Classes[i].PoliceConformActionDrop.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/conform-action/drop", data.getPath(), keyString))
		}
		if !data.Classes[i].PoliceExceedActionTransmit.IsNull() && !data.Classes[i].PoliceExceedActionTransmit.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/exceed-action/transmit", data.getPath(), keyString))
		}
		if !data.Classes[i].PoliceExceedActionDrop.IsNull() && !data.Classes[i].PoliceExceedActionDrop.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/exceed-action/drop", data.getPath(), keyString))
		}
		if !data.Classes[i].PoliceViolateActionTransmit.IsNull() && !data.Classes[i].PoliceViolateActionTransmit.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/violate-action/transmit", data.getPath(), keyString))
		}
		if !data.Classes[i].PoliceViolateActionDrop.IsNull() && !data.Classes[i].PoliceViolateActionDrop.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class%v/police/violate-action/drop", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

func (data *PolicyMapQoS) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	for i := range data.Classes {
		keys := [...]string{"name", "type"}
		keyValues := [...]string{data.Classes[i].Name.ValueString(), data.Classes[i].Type.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/class%v", data.getPath(), keyString))
	}
	return deletePaths
}
