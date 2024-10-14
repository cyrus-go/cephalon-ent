package enums

type MissionLoadBalanceState string

const (
	MissionLoadBalanceStateUnknown     MissionLoadBalanceState = "unknown"
	MissionLoadBalanceStateWaiting     MissionLoadBalanceState = "waiting"   // 实际有用
	MissionLoadBalanceStateCanceled    MissionLoadBalanceState = "canceled"  // 实际有用
	MissionLoadBalanceStateDoing       MissionLoadBalanceState = "doing"     // 实际有用
	MissionLoadBalanceStateSupplying   MissionLoadBalanceState = "supplying" // 实际有用
	MissionLoadBalanceStateClosing     MissionLoadBalanceState = "closing"   // 实际有用
	MissionLoadBalanceStateSucceed     MissionLoadBalanceState = "succeed"   // 实际有用
	MissionLoadBalanceStateFailed      MissionLoadBalanceState = "failed"
	MissionLoadBalanceStatePaused      MissionLoadBalanceState = "paused"
	MissionLoadBalanceStateResuming    MissionLoadBalanceState = "resuming"
	MissionLoadBalanceStateContracting MissionLoadBalanceState = "contracting" // 收缩中
	MissionLoadBalanceStateExpanding   MissionLoadBalanceState = "expanding"   // 扩展中
	MissionLoadBalanceStateRecovering  MissionLoadBalanceState = "recovering"  // 故障恢复中

)

func (obj MissionLoadBalanceState) Values() []string {
	return []string{
		string(MissionLoadBalanceStateUnknown),
		string(MissionLoadBalanceStateWaiting),
		string(MissionLoadBalanceStateCanceled),
		string(MissionLoadBalanceStateDoing),
		string(MissionLoadBalanceStateSupplying),
		string(MissionLoadBalanceStateClosing),
		string(MissionLoadBalanceStateSucceed),
		string(MissionLoadBalanceStateFailed),
		string(MissionLoadBalanceStatePaused),
		string(MissionLoadBalanceStateResuming),
		string(MissionLoadBalanceStateContracting),
		string(MissionLoadBalanceStateExpanding),
		string(MissionLoadBalanceStateRecovering),
	}
}

func (obj MissionLoadBalanceState) Ptr() *MissionLoadBalanceState {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
