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

type RouterISISAddressFamily struct {
	Device                                    types.String                                 `tfsdk:"device"`
	Id                                        types.String                                 `tfsdk:"id"`
	ProcessId                                 types.String                                 `tfsdk:"process_id"`
	AfName                                    types.String                                 `tfsdk:"af_name"`
	SafName                                   types.String                                 `tfsdk:"saf_name"`
	MetricStyleNarrow                         types.Bool                                   `tfsdk:"metric_style_narrow"`
	MetricStyleWide                           types.Bool                                   `tfsdk:"metric_style_wide"`
	MetricStyleTransition                     types.Bool                                   `tfsdk:"metric_style_transition"`
	MetricStyleLevels                         []RouterISISAddressFamilyMetricStyleLevels   `tfsdk:"metric_style_levels"`
	RouterIdInterfaceName                     types.String                                 `tfsdk:"router_id_interface_name"`
	RouterIdIpAddress                         types.String                                 `tfsdk:"router_id_ip_address"`
	DefaultInformationOriginate               types.Bool                                   `tfsdk:"default_information_originate"`
	FastRerouteDelayInterval                  types.Int64                                  `tfsdk:"fast_reroute_delay_interval"`
	FastReroutePerLinkPriorityLimitCritical   types.Bool                                   `tfsdk:"fast_reroute_per_link_priority_limit_critical"`
	FastReroutePerLinkPriorityLimitHigh       types.Bool                                   `tfsdk:"fast_reroute_per_link_priority_limit_high"`
	FastReroutePerLinkPriorityLimitMedium     types.Bool                                   `tfsdk:"fast_reroute_per_link_priority_limit_medium"`
	FastReroutePerPrefixPriorityLimitCritical types.Bool                                   `tfsdk:"fast_reroute_per_prefix_priority_limit_critical"`
	FastReroutePerPrefixPriorityLimitHigh     types.Bool                                   `tfsdk:"fast_reroute_per_prefix_priority_limit_high"`
	FastReroutePerPrefixPriorityLimitMedium   types.Bool                                   `tfsdk:"fast_reroute_per_prefix_priority_limit_medium"`
	MicroloopAvoidanceProtected               types.Bool                                   `tfsdk:"microloop_avoidance_protected"`
	MicroloopAvoidanceSegmentRouting          types.Bool                                   `tfsdk:"microloop_avoidance_segment_routing"`
	AdvertisePassiveOnly                      types.Bool                                   `tfsdk:"advertise_passive_only"`
	MplsLdpAutoConfig                         types.Bool                                   `tfsdk:"mpls_ldp_auto_config"`
	MplsTrafficEngRouterIdIpAddress           types.String                                 `tfsdk:"mpls_traffic_eng_router_id_ip_address"`
	MplsTrafficEngRouterIdInterface           types.String                                 `tfsdk:"mpls_traffic_eng_router_id_interface"`
	MplsTrafficEngLevel12                     types.Bool                                   `tfsdk:"mpls_traffic_eng_level_1_2"`
	MplsTrafficEngLevel1                      types.Bool                                   `tfsdk:"mpls_traffic_eng_level_1"`
	MplsTrafficEngLevel2Only                  types.Bool                                   `tfsdk:"mpls_traffic_eng_level_2_only"`
	SpfIntervalMaximumWait                    types.Int64                                  `tfsdk:"spf_interval_maximum_wait"`
	SpfIntervalInitialWait                    types.Int64                                  `tfsdk:"spf_interval_initial_wait"`
	SpfIntervalSecondaryWait                  types.Int64                                  `tfsdk:"spf_interval_secondary_wait"`
	SpfPrefixPriorities                       []RouterISISAddressFamilySpfPrefixPriorities `tfsdk:"spf_prefix_priorities"`
	SegmentRoutingMplsSrPrefer                types.Bool                                   `tfsdk:"segment_routing_mpls_sr_prefer"`
}
type RouterISISAddressFamilyMetricStyleLevels struct {
	LevelId    types.Int64 `tfsdk:"level_id"`
	Narrow     types.Bool  `tfsdk:"narrow"`
	Wide       types.Bool  `tfsdk:"wide"`
	Transition types.Bool  `tfsdk:"transition"`
}
type RouterISISAddressFamilySpfPrefixPriorities struct {
	Priority       types.String `tfsdk:"priority"`
	Tag            types.Int64  `tfsdk:"tag"`
	AccessListName types.String `tfsdk:"access_list_name"`
}

func (data RouterISISAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:router/isis/processes/process[process-id=%s]/address-families/address-family[af-name=%s][saf-name=%s]", data.ProcessId.ValueString(), data.AfName.ValueString(), data.SafName.ValueString())
}

func (data RouterISISAddressFamily) toBody(ctx context.Context) string {
	body := "{}"
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, "af-name", data.AfName.ValueString())
	}
	if !data.SafName.IsNull() && !data.SafName.IsUnknown() {
		body, _ = sjson.Set(body, "saf-name", data.SafName.ValueString())
	}
	if !data.MetricStyleNarrow.IsNull() && !data.MetricStyleNarrow.IsUnknown() {
		if data.MetricStyleNarrow.ValueBool() {
			body, _ = sjson.Set(body, "metric-style.narrow", map[string]string{})
		}
	}
	if !data.MetricStyleWide.IsNull() && !data.MetricStyleWide.IsUnknown() {
		if data.MetricStyleWide.ValueBool() {
			body, _ = sjson.Set(body, "metric-style.wide", map[string]string{})
		}
	}
	if !data.MetricStyleTransition.IsNull() && !data.MetricStyleTransition.IsUnknown() {
		if data.MetricStyleTransition.ValueBool() {
			body, _ = sjson.Set(body, "metric-style.transition", map[string]string{})
		}
	}
	if !data.RouterIdInterfaceName.IsNull() && !data.RouterIdInterfaceName.IsUnknown() {
		body, _ = sjson.Set(body, "router-id.interface-name", data.RouterIdInterfaceName.ValueString())
	}
	if !data.RouterIdIpAddress.IsNull() && !data.RouterIdIpAddress.IsUnknown() {
		body, _ = sjson.Set(body, "router-id.ip-address", data.RouterIdIpAddress.ValueString())
	}
	if !data.DefaultInformationOriginate.IsNull() && !data.DefaultInformationOriginate.IsUnknown() {
		if data.DefaultInformationOriginate.ValueBool() {
			body, _ = sjson.Set(body, "default-information.originate", map[string]string{})
		}
	}
	if !data.FastRerouteDelayInterval.IsNull() && !data.FastRerouteDelayInterval.IsUnknown() {
		body, _ = sjson.Set(body, "fast-reroute.delay-interval", strconv.FormatInt(data.FastRerouteDelayInterval.ValueInt64(), 10))
	}
	if !data.FastReroutePerLinkPriorityLimitCritical.IsNull() && !data.FastReroutePerLinkPriorityLimitCritical.IsUnknown() {
		if data.FastReroutePerLinkPriorityLimitCritical.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-link.priority-limit.critical", map[string]string{})
		}
	}
	if !data.FastReroutePerLinkPriorityLimitHigh.IsNull() && !data.FastReroutePerLinkPriorityLimitHigh.IsUnknown() {
		if data.FastReroutePerLinkPriorityLimitHigh.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-link.priority-limit.high", map[string]string{})
		}
	}
	if !data.FastReroutePerLinkPriorityLimitMedium.IsNull() && !data.FastReroutePerLinkPriorityLimitMedium.IsUnknown() {
		if data.FastReroutePerLinkPriorityLimitMedium.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-link.priority-limit.medium", map[string]string{})
		}
	}
	if !data.FastReroutePerPrefixPriorityLimitCritical.IsNull() && !data.FastReroutePerPrefixPriorityLimitCritical.IsUnknown() {
		if data.FastReroutePerPrefixPriorityLimitCritical.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-prefix.priority-limit.critical", map[string]string{})
		}
	}
	if !data.FastReroutePerPrefixPriorityLimitHigh.IsNull() && !data.FastReroutePerPrefixPriorityLimitHigh.IsUnknown() {
		if data.FastReroutePerPrefixPriorityLimitHigh.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-prefix.priority-limit.high", map[string]string{})
		}
	}
	if !data.FastReroutePerPrefixPriorityLimitMedium.IsNull() && !data.FastReroutePerPrefixPriorityLimitMedium.IsUnknown() {
		if data.FastReroutePerPrefixPriorityLimitMedium.ValueBool() {
			body, _ = sjson.Set(body, "fast-reroute.per-prefix.priority-limit.medium", map[string]string{})
		}
	}
	if !data.MicroloopAvoidanceProtected.IsNull() && !data.MicroloopAvoidanceProtected.IsUnknown() {
		if data.MicroloopAvoidanceProtected.ValueBool() {
			body, _ = sjson.Set(body, "microloop.avoidance.protected", map[string]string{})
		}
	}
	if !data.MicroloopAvoidanceSegmentRouting.IsNull() && !data.MicroloopAvoidanceSegmentRouting.IsUnknown() {
		if data.MicroloopAvoidanceSegmentRouting.ValueBool() {
			body, _ = sjson.Set(body, "microloop.avoidance.segment-routing", map[string]string{})
		}
	}
	if !data.AdvertisePassiveOnly.IsNull() && !data.AdvertisePassiveOnly.IsUnknown() {
		if data.AdvertisePassiveOnly.ValueBool() {
			body, _ = sjson.Set(body, "advertise.passive-only", map[string]string{})
		}
	}
	if !data.MplsLdpAutoConfig.IsNull() && !data.MplsLdpAutoConfig.IsUnknown() {
		if data.MplsLdpAutoConfig.ValueBool() {
			body, _ = sjson.Set(body, "mpls.ldp.auto-config", map[string]string{})
		}
	}
	if !data.MplsTrafficEngRouterIdIpAddress.IsNull() && !data.MplsTrafficEngRouterIdIpAddress.IsUnknown() {
		body, _ = sjson.Set(body, "mpls.traffic-eng.router-id.ip-address", data.MplsTrafficEngRouterIdIpAddress.ValueString())
	}
	if !data.MplsTrafficEngRouterIdInterface.IsNull() && !data.MplsTrafficEngRouterIdInterface.IsUnknown() {
		body, _ = sjson.Set(body, "mpls.traffic-eng.router-id.interface", data.MplsTrafficEngRouterIdInterface.ValueString())
	}
	if !data.MplsTrafficEngLevel12.IsNull() && !data.MplsTrafficEngLevel12.IsUnknown() {
		if data.MplsTrafficEngLevel12.ValueBool() {
			body, _ = sjson.Set(body, "mpls.traffic-eng.level-1-2", map[string]string{})
		}
	}
	if !data.MplsTrafficEngLevel1.IsNull() && !data.MplsTrafficEngLevel1.IsUnknown() {
		if data.MplsTrafficEngLevel1.ValueBool() {
			body, _ = sjson.Set(body, "mpls.traffic-eng.level-1", map[string]string{})
		}
	}
	if !data.MplsTrafficEngLevel2Only.IsNull() && !data.MplsTrafficEngLevel2Only.IsUnknown() {
		if data.MplsTrafficEngLevel2Only.ValueBool() {
			body, _ = sjson.Set(body, "mpls.traffic-eng.level-2-only", map[string]string{})
		}
	}
	if !data.SpfIntervalMaximumWait.IsNull() && !data.SpfIntervalMaximumWait.IsUnknown() {
		body, _ = sjson.Set(body, "spf-interval.maximum-wait.maximum-wait-time", strconv.FormatInt(data.SpfIntervalMaximumWait.ValueInt64(), 10))
	}
	if !data.SpfIntervalInitialWait.IsNull() && !data.SpfIntervalInitialWait.IsUnknown() {
		body, _ = sjson.Set(body, "spf-interval.initial-wait.initial-wait-time", strconv.FormatInt(data.SpfIntervalInitialWait.ValueInt64(), 10))
	}
	if !data.SpfIntervalSecondaryWait.IsNull() && !data.SpfIntervalSecondaryWait.IsUnknown() {
		body, _ = sjson.Set(body, "spf-interval.secondary-wait.secondary-wait-time", strconv.FormatInt(data.SpfIntervalSecondaryWait.ValueInt64(), 10))
	}
	if !data.SegmentRoutingMplsSrPrefer.IsNull() && !data.SegmentRoutingMplsSrPrefer.IsUnknown() {
		body, _ = sjson.Set(body, "segment-routing.mpls.sr-prefer", data.SegmentRoutingMplsSrPrefer.ValueBool())
	}
	if len(data.MetricStyleLevels) > 0 {
		body, _ = sjson.Set(body, "metric-style.levels.level", []interface{}{})
		for index, item := range data.MetricStyleLevels {
			if !item.LevelId.IsNull() && !item.LevelId.IsUnknown() {
				body, _ = sjson.Set(body, "metric-style.levels.level"+"."+strconv.Itoa(index)+"."+"level-id", strconv.FormatInt(item.LevelId.ValueInt64(), 10))
			}
			if !item.Narrow.IsNull() && !item.Narrow.IsUnknown() {
				if item.Narrow.ValueBool() {
					body, _ = sjson.Set(body, "metric-style.levels.level"+"."+strconv.Itoa(index)+"."+"narrow", map[string]string{})
				}
			}
			if !item.Wide.IsNull() && !item.Wide.IsUnknown() {
				if item.Wide.ValueBool() {
					body, _ = sjson.Set(body, "metric-style.levels.level"+"."+strconv.Itoa(index)+"."+"wide", map[string]string{})
				}
			}
			if !item.Transition.IsNull() && !item.Transition.IsUnknown() {
				if item.Transition.ValueBool() {
					body, _ = sjson.Set(body, "metric-style.levels.level"+"."+strconv.Itoa(index)+"."+"transition", map[string]string{})
				}
			}
		}
	}
	if len(data.SpfPrefixPriorities) > 0 {
		body, _ = sjson.Set(body, "spf.prefix-priority.prefix-priority", []interface{}{})
		for index, item := range data.SpfPrefixPriorities {
			if !item.Priority.IsNull() && !item.Priority.IsUnknown() {
				body, _ = sjson.Set(body, "spf.prefix-priority.prefix-priority"+"."+strconv.Itoa(index)+"."+"priority", item.Priority.ValueString())
			}
			if !item.Tag.IsNull() && !item.Tag.IsUnknown() {
				body, _ = sjson.Set(body, "spf.prefix-priority.prefix-priority"+"."+strconv.Itoa(index)+"."+"tag", strconv.FormatInt(item.Tag.ValueInt64(), 10))
			}
			if !item.AccessListName.IsNull() && !item.AccessListName.IsUnknown() {
				body, _ = sjson.Set(body, "spf.prefix-priority.prefix-priority"+"."+strconv.Itoa(index)+"."+"access-list-name", item.AccessListName.ValueString())
			}
		}
	}
	return body
}

func (data *RouterISISAddressFamily) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "metric-style.narrow"); !data.MetricStyleNarrow.IsNull() {
		if value.Exists() {
			data.MetricStyleNarrow = types.BoolValue(true)
		} else {
			data.MetricStyleNarrow = types.BoolValue(false)
		}
	} else {
		data.MetricStyleNarrow = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "metric-style.wide"); !data.MetricStyleWide.IsNull() {
		if value.Exists() {
			data.MetricStyleWide = types.BoolValue(true)
		} else {
			data.MetricStyleWide = types.BoolValue(false)
		}
	} else {
		data.MetricStyleWide = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "metric-style.transition"); !data.MetricStyleTransition.IsNull() {
		if value.Exists() {
			data.MetricStyleTransition = types.BoolValue(true)
		} else {
			data.MetricStyleTransition = types.BoolValue(false)
		}
	} else {
		data.MetricStyleTransition = types.BoolNull()
	}
	for i := range data.MetricStyleLevels {
		keys := [...]string{"level-id"}
		keyValues := [...]string{strconv.FormatInt(data.MetricStyleLevels[i].LevelId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "metric-style.levels.level").ForEach(
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
		if value := r.Get("level-id"); value.Exists() && !data.MetricStyleLevels[i].LevelId.IsNull() {
			data.MetricStyleLevels[i].LevelId = types.Int64Value(value.Int())
		} else {
			data.MetricStyleLevels[i].LevelId = types.Int64Null()
		}
		if value := r.Get("narrow"); !data.MetricStyleLevels[i].Narrow.IsNull() {
			if value.Exists() {
				data.MetricStyleLevels[i].Narrow = types.BoolValue(true)
			} else {
				data.MetricStyleLevels[i].Narrow = types.BoolValue(false)
			}
		} else {
			data.MetricStyleLevels[i].Narrow = types.BoolNull()
		}
		if value := r.Get("wide"); !data.MetricStyleLevels[i].Wide.IsNull() {
			if value.Exists() {
				data.MetricStyleLevels[i].Wide = types.BoolValue(true)
			} else {
				data.MetricStyleLevels[i].Wide = types.BoolValue(false)
			}
		} else {
			data.MetricStyleLevels[i].Wide = types.BoolNull()
		}
		if value := r.Get("transition"); !data.MetricStyleLevels[i].Transition.IsNull() {
			if value.Exists() {
				data.MetricStyleLevels[i].Transition = types.BoolValue(true)
			} else {
				data.MetricStyleLevels[i].Transition = types.BoolValue(false)
			}
		} else {
			data.MetricStyleLevels[i].Transition = types.BoolNull()
		}
	}
	if value := gjson.GetBytes(res, "router-id.interface-name"); value.Exists() && !data.RouterIdInterfaceName.IsNull() {
		data.RouterIdInterfaceName = types.StringValue(value.String())
	} else {
		data.RouterIdInterfaceName = types.StringNull()
	}
	if value := gjson.GetBytes(res, "router-id.ip-address"); value.Exists() && !data.RouterIdIpAddress.IsNull() {
		data.RouterIdIpAddress = types.StringValue(value.String())
	} else {
		data.RouterIdIpAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "default-information.originate"); !data.DefaultInformationOriginate.IsNull() {
		if value.Exists() {
			data.DefaultInformationOriginate = types.BoolValue(true)
		} else {
			data.DefaultInformationOriginate = types.BoolValue(false)
		}
	} else {
		data.DefaultInformationOriginate = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.delay-interval"); value.Exists() && !data.FastRerouteDelayInterval.IsNull() {
		data.FastRerouteDelayInterval = types.Int64Value(value.Int())
	} else {
		data.FastRerouteDelayInterval = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.critical"); !data.FastReroutePerLinkPriorityLimitCritical.IsNull() {
		if value.Exists() {
			data.FastReroutePerLinkPriorityLimitCritical = types.BoolValue(true)
		} else {
			data.FastReroutePerLinkPriorityLimitCritical = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerLinkPriorityLimitCritical = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.high"); !data.FastReroutePerLinkPriorityLimitHigh.IsNull() {
		if value.Exists() {
			data.FastReroutePerLinkPriorityLimitHigh = types.BoolValue(true)
		} else {
			data.FastReroutePerLinkPriorityLimitHigh = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerLinkPriorityLimitHigh = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.medium"); !data.FastReroutePerLinkPriorityLimitMedium.IsNull() {
		if value.Exists() {
			data.FastReroutePerLinkPriorityLimitMedium = types.BoolValue(true)
		} else {
			data.FastReroutePerLinkPriorityLimitMedium = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerLinkPriorityLimitMedium = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.critical"); !data.FastReroutePerPrefixPriorityLimitCritical.IsNull() {
		if value.Exists() {
			data.FastReroutePerPrefixPriorityLimitCritical = types.BoolValue(true)
		} else {
			data.FastReroutePerPrefixPriorityLimitCritical = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerPrefixPriorityLimitCritical = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.high"); !data.FastReroutePerPrefixPriorityLimitHigh.IsNull() {
		if value.Exists() {
			data.FastReroutePerPrefixPriorityLimitHigh = types.BoolValue(true)
		} else {
			data.FastReroutePerPrefixPriorityLimitHigh = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerPrefixPriorityLimitHigh = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.medium"); !data.FastReroutePerPrefixPriorityLimitMedium.IsNull() {
		if value.Exists() {
			data.FastReroutePerPrefixPriorityLimitMedium = types.BoolValue(true)
		} else {
			data.FastReroutePerPrefixPriorityLimitMedium = types.BoolValue(false)
		}
	} else {
		data.FastReroutePerPrefixPriorityLimitMedium = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "microloop.avoidance.protected"); !data.MicroloopAvoidanceProtected.IsNull() {
		if value.Exists() {
			data.MicroloopAvoidanceProtected = types.BoolValue(true)
		} else {
			data.MicroloopAvoidanceProtected = types.BoolValue(false)
		}
	} else {
		data.MicroloopAvoidanceProtected = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "microloop.avoidance.segment-routing"); !data.MicroloopAvoidanceSegmentRouting.IsNull() {
		if value.Exists() {
			data.MicroloopAvoidanceSegmentRouting = types.BoolValue(true)
		} else {
			data.MicroloopAvoidanceSegmentRouting = types.BoolValue(false)
		}
	} else {
		data.MicroloopAvoidanceSegmentRouting = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "advertise.passive-only"); !data.AdvertisePassiveOnly.IsNull() {
		if value.Exists() {
			data.AdvertisePassiveOnly = types.BoolValue(true)
		} else {
			data.AdvertisePassiveOnly = types.BoolValue(false)
		}
	} else {
		data.AdvertisePassiveOnly = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mpls.ldp.auto-config"); !data.MplsLdpAutoConfig.IsNull() {
		if value.Exists() {
			data.MplsLdpAutoConfig = types.BoolValue(true)
		} else {
			data.MplsLdpAutoConfig = types.BoolValue(false)
		}
	} else {
		data.MplsLdpAutoConfig = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.router-id.ip-address"); value.Exists() && !data.MplsTrafficEngRouterIdIpAddress.IsNull() {
		data.MplsTrafficEngRouterIdIpAddress = types.StringValue(value.String())
	} else {
		data.MplsTrafficEngRouterIdIpAddress = types.StringNull()
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.router-id.interface"); value.Exists() && !data.MplsTrafficEngRouterIdInterface.IsNull() {
		data.MplsTrafficEngRouterIdInterface = types.StringValue(value.String())
	} else {
		data.MplsTrafficEngRouterIdInterface = types.StringNull()
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-1-2"); !data.MplsTrafficEngLevel12.IsNull() {
		if value.Exists() {
			data.MplsTrafficEngLevel12 = types.BoolValue(true)
		} else {
			data.MplsTrafficEngLevel12 = types.BoolValue(false)
		}
	} else {
		data.MplsTrafficEngLevel12 = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-1"); !data.MplsTrafficEngLevel1.IsNull() {
		if value.Exists() {
			data.MplsTrafficEngLevel1 = types.BoolValue(true)
		} else {
			data.MplsTrafficEngLevel1 = types.BoolValue(false)
		}
	} else {
		data.MplsTrafficEngLevel1 = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-2-only"); !data.MplsTrafficEngLevel2Only.IsNull() {
		if value.Exists() {
			data.MplsTrafficEngLevel2Only = types.BoolValue(true)
		} else {
			data.MplsTrafficEngLevel2Only = types.BoolValue(false)
		}
	} else {
		data.MplsTrafficEngLevel2Only = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "spf-interval.maximum-wait.maximum-wait-time"); value.Exists() && !data.SpfIntervalMaximumWait.IsNull() {
		data.SpfIntervalMaximumWait = types.Int64Value(value.Int())
	} else {
		data.SpfIntervalMaximumWait = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "spf-interval.initial-wait.initial-wait-time"); value.Exists() && !data.SpfIntervalInitialWait.IsNull() {
		data.SpfIntervalInitialWait = types.Int64Value(value.Int())
	} else {
		data.SpfIntervalInitialWait = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "spf-interval.secondary-wait.secondary-wait-time"); value.Exists() && !data.SpfIntervalSecondaryWait.IsNull() {
		data.SpfIntervalSecondaryWait = types.Int64Value(value.Int())
	} else {
		data.SpfIntervalSecondaryWait = types.Int64Null()
	}
	for i := range data.SpfPrefixPriorities {
		keys := [...]string{"priority"}
		keyValues := [...]string{data.SpfPrefixPriorities[i].Priority.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "spf.prefix-priority.prefix-priority").ForEach(
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
		if value := r.Get("priority"); value.Exists() && !data.SpfPrefixPriorities[i].Priority.IsNull() {
			data.SpfPrefixPriorities[i].Priority = types.StringValue(value.String())
		} else {
			data.SpfPrefixPriorities[i].Priority = types.StringNull()
		}
		if value := r.Get("tag"); value.Exists() && !data.SpfPrefixPriorities[i].Tag.IsNull() {
			data.SpfPrefixPriorities[i].Tag = types.Int64Value(value.Int())
		} else {
			data.SpfPrefixPriorities[i].Tag = types.Int64Null()
		}
		if value := r.Get("access-list-name"); value.Exists() && !data.SpfPrefixPriorities[i].AccessListName.IsNull() {
			data.SpfPrefixPriorities[i].AccessListName = types.StringValue(value.String())
		} else {
			data.SpfPrefixPriorities[i].AccessListName = types.StringNull()
		}
	}
	if value := gjson.GetBytes(res, "segment-routing.mpls.sr-prefer"); !data.SegmentRoutingMplsSrPrefer.IsNull() {
		if value.Exists() {
			data.SegmentRoutingMplsSrPrefer = types.BoolValue(value.Bool())
		}
	} else {
		data.SegmentRoutingMplsSrPrefer = types.BoolNull()
	}
}

func (data *RouterISISAddressFamily) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "metric-style.narrow"); value.Exists() {
		data.MetricStyleNarrow = types.BoolValue(true)
	} else {
		data.MetricStyleNarrow = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "metric-style.wide"); value.Exists() {
		data.MetricStyleWide = types.BoolValue(true)
	} else {
		data.MetricStyleWide = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "metric-style.transition"); value.Exists() {
		data.MetricStyleTransition = types.BoolValue(true)
	} else {
		data.MetricStyleTransition = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "metric-style.levels.level"); value.Exists() {
		data.MetricStyleLevels = make([]RouterISISAddressFamilyMetricStyleLevels, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISAddressFamilyMetricStyleLevels{}
			if cValue := v.Get("level-id"); cValue.Exists() {
				item.LevelId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("narrow"); cValue.Exists() {
				item.Narrow = types.BoolValue(true)
			} else {
				item.Narrow = types.BoolValue(false)
			}
			if cValue := v.Get("wide"); cValue.Exists() {
				item.Wide = types.BoolValue(true)
			} else {
				item.Wide = types.BoolValue(false)
			}
			if cValue := v.Get("transition"); cValue.Exists() {
				item.Transition = types.BoolValue(true)
			} else {
				item.Transition = types.BoolValue(false)
			}
			data.MetricStyleLevels = append(data.MetricStyleLevels, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "router-id.interface-name"); value.Exists() {
		data.RouterIdInterfaceName = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "router-id.ip-address"); value.Exists() {
		data.RouterIdIpAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate = types.BoolValue(true)
	} else {
		data.DefaultInformationOriginate = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.delay-interval"); value.Exists() {
		data.FastRerouteDelayInterval = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.critical"); value.Exists() {
		data.FastReroutePerLinkPriorityLimitCritical = types.BoolValue(true)
	} else {
		data.FastReroutePerLinkPriorityLimitCritical = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.high"); value.Exists() {
		data.FastReroutePerLinkPriorityLimitHigh = types.BoolValue(true)
	} else {
		data.FastReroutePerLinkPriorityLimitHigh = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-link.priority-limit.medium"); value.Exists() {
		data.FastReroutePerLinkPriorityLimitMedium = types.BoolValue(true)
	} else {
		data.FastReroutePerLinkPriorityLimitMedium = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.critical"); value.Exists() {
		data.FastReroutePerPrefixPriorityLimitCritical = types.BoolValue(true)
	} else {
		data.FastReroutePerPrefixPriorityLimitCritical = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.high"); value.Exists() {
		data.FastReroutePerPrefixPriorityLimitHigh = types.BoolValue(true)
	} else {
		data.FastReroutePerPrefixPriorityLimitHigh = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "fast-reroute.per-prefix.priority-limit.medium"); value.Exists() {
		data.FastReroutePerPrefixPriorityLimitMedium = types.BoolValue(true)
	} else {
		data.FastReroutePerPrefixPriorityLimitMedium = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "microloop.avoidance.protected"); value.Exists() {
		data.MicroloopAvoidanceProtected = types.BoolValue(true)
	} else {
		data.MicroloopAvoidanceProtected = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "microloop.avoidance.segment-routing"); value.Exists() {
		data.MicroloopAvoidanceSegmentRouting = types.BoolValue(true)
	} else {
		data.MicroloopAvoidanceSegmentRouting = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "advertise.passive-only"); value.Exists() {
		data.AdvertisePassiveOnly = types.BoolValue(true)
	} else {
		data.AdvertisePassiveOnly = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mpls.ldp.auto-config"); value.Exists() {
		data.MplsLdpAutoConfig = types.BoolValue(true)
	} else {
		data.MplsLdpAutoConfig = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.router-id.ip-address"); value.Exists() {
		data.MplsTrafficEngRouterIdIpAddress = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.router-id.interface"); value.Exists() {
		data.MplsTrafficEngRouterIdInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-1-2"); value.Exists() {
		data.MplsTrafficEngLevel12 = types.BoolValue(true)
	} else {
		data.MplsTrafficEngLevel12 = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-1"); value.Exists() {
		data.MplsTrafficEngLevel1 = types.BoolValue(true)
	} else {
		data.MplsTrafficEngLevel1 = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "mpls.traffic-eng.level-2-only"); value.Exists() {
		data.MplsTrafficEngLevel2Only = types.BoolValue(true)
	} else {
		data.MplsTrafficEngLevel2Only = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "spf-interval.maximum-wait.maximum-wait-time"); value.Exists() {
		data.SpfIntervalMaximumWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "spf-interval.initial-wait.initial-wait-time"); value.Exists() {
		data.SpfIntervalInitialWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "spf-interval.secondary-wait.secondary-wait-time"); value.Exists() {
		data.SpfIntervalSecondaryWait = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "spf.prefix-priority.prefix-priority"); value.Exists() {
		data.SpfPrefixPriorities = make([]RouterISISAddressFamilySpfPrefixPriorities, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouterISISAddressFamilySpfPrefixPriorities{}
			if cValue := v.Get("priority"); cValue.Exists() {
				item.Priority = types.StringValue(cValue.String())
			}
			if cValue := v.Get("tag"); cValue.Exists() {
				item.Tag = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("access-list-name"); cValue.Exists() {
				item.AccessListName = types.StringValue(cValue.String())
			}
			data.SpfPrefixPriorities = append(data.SpfPrefixPriorities, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "segment-routing.mpls.sr-prefer"); value.Exists() {
		data.SegmentRoutingMplsSrPrefer = types.BoolValue(value.Bool())
	} else {
		data.SegmentRoutingMplsSrPrefer = types.BoolValue(false)
	}
}

func (data *RouterISISAddressFamily) fromPlan(ctx context.Context, plan RouterISISAddressFamily) {
	data.Device = plan.Device
	data.ProcessId = types.StringValue(plan.ProcessId.ValueString())
	data.AfName = types.StringValue(plan.AfName.ValueString())
	data.SafName = types.StringValue(plan.SafName.ValueString())
}

func (data *RouterISISAddressFamily) getDeletedListItems(ctx context.Context, state RouterISISAddressFamily) []string {
	deletedListItems := make([]string, 0)
	for i := range state.MetricStyleLevels {
		keys := [...]string{"level-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.MetricStyleLevels[i].LevelId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.MetricStyleLevels[i].LevelId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MetricStyleLevels {
			found = true
			if state.MetricStyleLevels[i].LevelId.ValueInt64() != data.MetricStyleLevels[j].LevelId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/metric-style/levels/level%v", state.getPath(), keyString))
		}
	}
	for i := range state.SpfPrefixPriorities {
		keys := [...]string{"priority"}
		stateKeyValues := [...]string{state.SpfPrefixPriorities[i].Priority.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.SpfPrefixPriorities[i].Priority.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SpfPrefixPriorities {
			found = true
			if state.SpfPrefixPriorities[i].Priority.ValueString() != data.SpfPrefixPriorities[j].Priority.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/spf/prefix-priority/prefix-priority%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *RouterISISAddressFamily) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}
