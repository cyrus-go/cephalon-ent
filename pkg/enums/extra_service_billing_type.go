package enums

type ExtraServiceBillingType string

const (
	ExtraServiceBillingTypeUnknown ExtraServiceBillingType = "unknown"
	// ExtraServiceBillingTypeTimePlanHour VPN 服务包时 - 按小时
	ExtraServiceBillingTypeTimePlanHour ExtraServiceBillingType = "time_plan_hour"
	// ExtraServiceBillingTypeTimePlanDay  VPN 服务包时 - 按天
	ExtraServiceBillingTypeTimePlanDay ExtraServiceBillingType = "time_plan_day"
	// ExtraServiceBillingTypeTimePlanWeek VPN 服务包时 - 按周
	ExtraServiceBillingTypeTimePlanWeek ExtraServiceBillingType = "time_plan_week"
	// ExtraServiceBillingTypeTimePlanMonth VPN 服务包时 - 按月
	ExtraServiceBillingTypeTimePlanMonth ExtraServiceBillingType = "time_plan_month"
	// ExtraServiceBillingTypeTimePlanVolume VPN 服务按量，单次任务，内部计时 * 时间单价
	ExtraServiceBillingTypeTimePlanVolume ExtraServiceBillingType = "time_plan_volume"
	// ExtraServiceBillingTypeHold 按位，特点是开启期间具备某些功能，结束后失去功能
	ExtraServiceBillingTypeHold ExtraServiceBillingType = "hold"
)

func (obj ExtraServiceBillingType) Values() []string {
	return []string{
		string(ExtraServiceBillingTypeUnknown),
		string(ExtraServiceBillingTypeTimePlanHour),
		string(ExtraServiceBillingTypeTimePlanDay),
		string(ExtraServiceBillingTypeTimePlanWeek),
		string(ExtraServiceBillingTypeTimePlanMonth),
		string(ExtraServiceBillingTypeTimePlanVolume),
		string(ExtraServiceBillingTypeHold),
	}
}
