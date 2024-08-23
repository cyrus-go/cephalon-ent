package enums

type DeviceFaultType string

const (
	DeviceFaultTypeOK              DeviceFaultType = "ok"               //无故障
	DeviceFaultTypeNoDrive         DeviceFaultType = "drive_abnormal"   //驱动异常
	DeviceFaultTypeNoDisk          DeviceFaultType = "no_disk"          //磁盘空间不足
	DeviceFaultTypeTaskFailure     DeviceFaultType = "task_failure"     //任务连续三次执行失败
	DeviceFaultTypeRecover         DeviceFaultType = "recover"          //故障恢复
	DeviceFaultTypeHighTemperature DeviceFaultType = "high_temperature" // 高温故障
)

func (DeviceFaultType) Values() []string {
	return []string{
		string(DeviceFaultTypeOK),
		string(DeviceFaultTypeNoDrive),
		string(DeviceFaultTypeNoDisk),
		string(DeviceFaultTypeTaskFailure),
		string(DeviceFaultTypeRecover),
		string(DeviceFaultTypeHighTemperature),
	}
}

func (obj DeviceFaultType) Ptr() *DeviceFaultType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
