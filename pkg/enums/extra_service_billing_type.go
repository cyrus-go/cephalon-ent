package enums

type ExtraServiceBillingType string

const (
	ExtraServiceBillingTypeUnknown ExtraServiceBillingType = "unknown"
	// ExtraServiceBillingTypeVPNTimePlanHour VPN 服务包时 - 按小时
	ExtraServiceBillingTypeVPNTimePlanHour ExtraServiceBillingType = "vpn_time_plan_hour"
	// ExtraServiceBillingTypeVPNTimePlanDay  VPN 服务包时 - 按天
	ExtraServiceBillingTypeVPNTimePlanDay ExtraServiceBillingType = "vpn_time_plan_day"
	// ExtraServiceBillingTypeVPNTimePlanWeek VPN 服务包时 - 按周
	ExtraServiceBillingTypeVPNTimePlanWeek ExtraServiceBillingType = "vpn_time_plan_week"
	// ExtraServiceBillingTypeVPNTimePlanMonth VPN 服务包时 - 按月
	ExtraServiceBillingTypeVPNTimePlanMonth ExtraServiceBillingType = "vpn_time_plan_month"
	// ExtraServiceBillingTypeVPNTimePlanVolume VPN 服务按量，单次任务，内部计时 * 时间单价
	ExtraServiceBillingTypeVPNTimePlanVolume ExtraServiceBillingType = "vpn_time_plan_volume"
	// ExtraServiceBillingTypeVPNHold 按位，特点是开启期间具备某些功能，结束后失去功能
	ExtraServiceBillingTypeVPNHold ExtraServiceBillingType = "hold"
)

func (obj ExtraServiceBillingType) Values() []string {
	return []string{
		string(ExtraServiceBillingTypeUnknown),
		string(ExtraServiceBillingTypeVPNTimePlanHour),
		string(ExtraServiceBillingTypeVPNTimePlanDay),
		string(ExtraServiceBillingTypeVPNTimePlanWeek),
		string(ExtraServiceBillingTypeVPNTimePlanMonth),
		string(ExtraServiceBillingTypeVPNTimePlanVolume),
		string(ExtraServiceBillingTypeVPNHold),
	}
}
