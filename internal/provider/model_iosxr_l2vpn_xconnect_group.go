// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type L2VPNXconnectGroup struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	GroupName types.String `tfsdk:"group_name"`
}

func (data L2VPNXconnectGroup) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]", data.GroupName.Value)
}

func (data L2VPNXconnectGroup) toBody() string {
	body := "{}"
	return body
}

func (data *L2VPNXconnectGroup) updateFromBody(res []byte) {
}

func (data *L2VPNXconnectGroup) fromBody(res []byte) {
}

func (data *L2VPNXconnectGroup) fromPlan(plan L2VPNXconnectGroup) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
}

func (data *L2VPNXconnectGroup) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.GroupName.Unknown {
		data.GroupName.Unknown = false
		data.GroupName.Null = true
	}
}

func (data *L2VPNXconnectGroup) getDeletedListItems(state L2VPNXconnectGroup) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *L2VPNXconnectGroup) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
