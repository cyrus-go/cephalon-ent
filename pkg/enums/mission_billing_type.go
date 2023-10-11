package enums

type MissionBillingType string

const (
	MissionBillingTypeUnknown MissionBillingType = "unknown"
	// MissionBillingTypeTime 按时，时间 * 分钟单价，特点是有供应中状态 supplying
	MissionBillingTypeTime MissionBillingType = "time"
	// MissionBillingTypeCount 按次，特点是单次任务不管干啥，都算单价钱
	MissionBillingTypeCount MissionBillingType = "count"
	// MissionBillingTypeHold 按位，特点是开启期间具备某些功能，结束后失去功能
	MissionBillingTypeHold MissionBillingType = "hold"
	// MissionBillingTypeVolume 按量，单次任务，内部计时 * 时间单价，用多少算力计多少钱
	MissionBillingTypeVolume MissionBillingType = "volume"
	// MissionBillingTypeTimePlanHour 包时 - 按小时，先给钱，然后按计划的时间服务
	MissionBillingTypeTimePlanHour MissionBillingType = "time_plan_hour"
	// MissionBillingTypeTimePlanDay 包时 - 按天
	MissionBillingTypeTimePlanDay MissionBillingType = "time_plan_day"
	// MissionBillingTypeTimePlanWeek 包时 - 按周
	MissionBillingTypeTimePlanWeek MissionBillingType = "time_plan_week"
	// MissionBillingTypeTimePlanMonth 包时 - 按月
	MissionBillingTypeTimePlanMonth MissionBillingType = "time_plan_month"
)

func (obj MissionBillingType) Values() []string {
	return []string{
		string(MissionBillingTypeUnknown),
		string(MissionBillingTypeTime),
		string(MissionBillingTypeCount),
		string(MissionBillingTypeHold),
		string(MissionBillingTypeVolume),
		string(MissionBillingTypeTimePlanHour),
		string(MissionBillingTypeTimePlanDay),
		string(MissionBillingTypeTimePlanWeek),
		string(MissionBillingTypeTimePlanMonth),
	}
}

func (obj MissionBillingType) Ptr() *MissionBillingType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
