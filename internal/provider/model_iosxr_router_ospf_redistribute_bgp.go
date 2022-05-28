// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterOSPFRedistributeBGP struct {
	Device      types.String `tfsdk:"device"`
	Id          types.String `tfsdk:"id"`
	ProcessName types.String `tfsdk:"process_name"`
	AsNumber    types.String `tfsdk:"as_number"`
	Tag         types.Int64  `tfsdk:"tag"`
	MetricType  types.String `tfsdk:"metric_type"`
}

func (data RouterOSPFRedistributeBGP) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/redistribute/bgp/as[as-number=%s]", data.ProcessName.Value, data.AsNumber.Value)
}

func (data RouterOSPFRedistributeBGP) toBody() string {
	body := "{}"
	if !data.Tag.Null && !data.Tag.Unknown {
		body, _ = sjson.Set(body, "tag", strconv.FormatInt(data.Tag.Value, 10))
	}
	if !data.MetricType.Null && !data.MetricType.Unknown {
		body, _ = sjson.Set(body, "metric-type", data.MetricType.Value)
	}
	return body
}

func (data *RouterOSPFRedistributeBGP) updateFromBody(res []byte) {
	if value := gjson.GetBytes(res, "tag"); value.Exists() {
		data.Tag.Value = value.Int()
	} else {
		data.Tag.Null = true
	}
	if value := gjson.GetBytes(res, "metric-type"); value.Exists() {
		data.MetricType.Value = value.String()
	} else {
		data.MetricType.Null = true
	}
}

func (data *RouterOSPFRedistributeBGP) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "tag"); value.Exists() {
		data.Tag.Value = value.Int()
		data.Tag.Null = false
	}
	if value := gjson.GetBytes(res, "metric-type"); value.Exists() {
		data.MetricType.Value = value.String()
		data.MetricType.Null = false
	}
}

func (data *RouterOSPFRedistributeBGP) fromPlan(plan RouterOSPFRedistributeBGP) {
	data.Device = plan.Device
	data.ProcessName.Value = plan.ProcessName.Value
	data.AsNumber.Value = plan.AsNumber.Value
}

func (data *RouterOSPFRedistributeBGP) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.ProcessName.Unknown {
		data.ProcessName.Unknown = false
		data.ProcessName.Null = true
	}
	if data.AsNumber.Unknown {
		data.AsNumber.Unknown = false
		data.AsNumber.Null = true
	}
	if data.Tag.Unknown {
		data.Tag.Unknown = false
		data.Tag.Null = true
	}
	if data.MetricType.Unknown {
		data.MetricType.Unknown = false
		data.MetricType.Null = true
	}
}

func (data *RouterOSPFRedistributeBGP) getDeletedListItems(state RouterOSPFRedistributeBGP) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RouterOSPFRedistributeBGP) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
