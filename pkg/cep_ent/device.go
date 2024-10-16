// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/giftmissionconfig"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 设备，对应客户端 core，记录心跳等信息
type Device struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by,string"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by,string"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 外键用户 id
	UserID int64 `json:"user_id,string"`
	// 外键补贴任务配置 id
	GiftMissionConfigID int64 `json:"gift_mission_config_id,string"`
	// 设备唯一序列号
	SerialNumber string `json:"serial_number"`
	// 设备状态
	State device.State `json:"state,omitempty"`
	// 该设备总获得利润
	SumCep int64 `json:"sum_cep"`
	// 设备是否正在对接中
	Linking bool `json:"linking"`
	// 设备的绑定状态
	BindingStatus enums.DeviceBindingStatus `json:"binding_status"`
	// 设备状态
	Status enums.DeviceStatus `json:"status"`
	// 设备名称
	Name string `json:"name"`
	// 运维管理设备名称
	ManageName string `json:"manage_name"`
	// 设备类型
	Type enums.DeviceType `json:"type"`
	// 核心数
	CoresNumber int64 `json:"cores_number"`
	// CPU型号
	CPU string `json:"cpu,omitempty" json:cpu`
	// CPU型号
	Cpus []string `json:"cpus"`
	// 内存(单位:G)
	Memory int64 `json:"memory"`
	// 硬盘(单位:T)
	Disk float32 `json:"disk,omitempty" json:disk`
	// 延迟(单位:ms)
	Delay float64 `json:"delay"`
	// GPU 温度(单位:℃)
	GpuTemperature float64 `json:"gpu_temperature"`
	// CPU 温度(单位:℃)
	CPUTemperature float64 `json:"cpu_temperature"`
	// 设备稳定性
	Stability enums.DeviceStabilityType `json:"stability"`
	// 设备版本
	Version string `json:"version"`
	// 故障信息
	Fault enums.DeviceFaultType `json:"fault"`
	// 设备等级(目前阶段就是黑名单)
	Rank enums.DeviceRankType `json:"rank"`
	// 空闲GPU数量
	FreeGpuNum int `json:"free_gpu_num"`
	// 评定设备等级的时刻，带时区
	RankAt *time.Time `json:"rank_at"`
	// 判定稳定性的时刻，带时区
	StabilityAt *time.Time `json:"stability_at"`
	// 温度超标的时刻，带时区
	HighTemperatureAt *time.Time `json:"high_temperature_at"`
	// 托管类型，非托管/半托管/全托管等
	HostingType enums.DeviceHostingType `json:"hosting_type"`
	// 可接特殊任务类型标签
	MissionTag enums.DeviceMissionTag `json:"mission_tag"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges        DeviceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// GiftMissionConfig holds the value of the gift_mission_config edge.
	GiftMissionConfig *GiftMissionConfig `json:"gift_mission_config,omitempty"`
	// MissionProduceOrders holds the value of the mission_produce_orders edge.
	MissionProduceOrders []*MissionProduceOrder `json:"mission_produce_orders,omitempty"`
	// UserDevices holds the value of the user_devices edge.
	UserDevices []*UserDevice `json:"user_devices,omitempty"`
	// DeviceGpuMissions holds the value of the device_gpu_missions edge.
	DeviceGpuMissions []*DeviceGpuMission `json:"device_gpu_missions,omitempty"`
	// FrpcInfos holds the value of the frpc_infos edge.
	FrpcInfos []*FrpcInfo `json:"frpc_infos,omitempty"`
	// MissionOrders holds the value of the mission_orders edge.
	MissionOrders []*MissionOrder `json:"mission_orders,omitempty"`
	// MissionProductions holds the value of the mission_productions edge.
	MissionProductions []*MissionProduction `json:"mission_productions,omitempty"`
	// DeviceRebootTimes holds the value of the device_reboot_times edge.
	DeviceRebootTimes []*DeviceRebootTime `json:"device_reboot_times,omitempty"`
	// TroubleDeducts holds the value of the trouble_deducts edge.
	TroubleDeducts []*TroubleDeduct `json:"trouble_deducts,omitempty"`
	// DeviceStates holds the value of the device_states edge.
	DeviceStates []*DeviceState `json:"device_states,omitempty"`
	// DeviceOfflineRecords holds the value of the device_offline_records edge.
	DeviceOfflineRecords []*DeviceOfflineRecord `json:"device_offline_records,omitempty"`
	// MissionFailedFeedbacks holds the value of the mission_failed_feedbacks edge.
	MissionFailedFeedbacks []*MissionFailedFeedback `json:"mission_failed_feedbacks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GiftMissionConfigOrErr returns the GiftMissionConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) GiftMissionConfigOrErr() (*GiftMissionConfig, error) {
	if e.loadedTypes[1] {
		if e.GiftMissionConfig == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: giftmissionconfig.Label}
		}
		return e.GiftMissionConfig, nil
	}
	return nil, &NotLoadedError{edge: "gift_mission_config"}
}

// MissionProduceOrdersOrErr returns the MissionProduceOrders value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionProduceOrdersOrErr() ([]*MissionProduceOrder, error) {
	if e.loadedTypes[2] {
		return e.MissionProduceOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_orders"}
}

// UserDevicesOrErr returns the UserDevices value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) UserDevicesOrErr() ([]*UserDevice, error) {
	if e.loadedTypes[3] {
		return e.UserDevices, nil
	}
	return nil, &NotLoadedError{edge: "user_devices"}
}

// DeviceGpuMissionsOrErr returns the DeviceGpuMissions value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeviceGpuMissionsOrErr() ([]*DeviceGpuMission, error) {
	if e.loadedTypes[4] {
		return e.DeviceGpuMissions, nil
	}
	return nil, &NotLoadedError{edge: "device_gpu_missions"}
}

// FrpcInfosOrErr returns the FrpcInfos value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) FrpcInfosOrErr() ([]*FrpcInfo, error) {
	if e.loadedTypes[5] {
		return e.FrpcInfos, nil
	}
	return nil, &NotLoadedError{edge: "frpc_infos"}
}

// MissionOrdersOrErr returns the MissionOrders value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionOrdersOrErr() ([]*MissionOrder, error) {
	if e.loadedTypes[6] {
		return e.MissionOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_orders"}
}

// MissionProductionsOrErr returns the MissionProductions value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionProductionsOrErr() ([]*MissionProduction, error) {
	if e.loadedTypes[7] {
		return e.MissionProductions, nil
	}
	return nil, &NotLoadedError{edge: "mission_productions"}
}

// DeviceRebootTimesOrErr returns the DeviceRebootTimes value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeviceRebootTimesOrErr() ([]*DeviceRebootTime, error) {
	if e.loadedTypes[8] {
		return e.DeviceRebootTimes, nil
	}
	return nil, &NotLoadedError{edge: "device_reboot_times"}
}

// TroubleDeductsOrErr returns the TroubleDeducts value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) TroubleDeductsOrErr() ([]*TroubleDeduct, error) {
	if e.loadedTypes[9] {
		return e.TroubleDeducts, nil
	}
	return nil, &NotLoadedError{edge: "trouble_deducts"}
}

// DeviceStatesOrErr returns the DeviceStates value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeviceStatesOrErr() ([]*DeviceState, error) {
	if e.loadedTypes[10] {
		return e.DeviceStates, nil
	}
	return nil, &NotLoadedError{edge: "device_states"}
}

// DeviceOfflineRecordsOrErr returns the DeviceOfflineRecords value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeviceOfflineRecordsOrErr() ([]*DeviceOfflineRecord, error) {
	if e.loadedTypes[11] {
		return e.DeviceOfflineRecords, nil
	}
	return nil, &NotLoadedError{edge: "device_offline_records"}
}

// MissionFailedFeedbacksOrErr returns the MissionFailedFeedbacks value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionFailedFeedbacksOrErr() ([]*MissionFailedFeedback, error) {
	if e.loadedTypes[12] {
		return e.MissionFailedFeedbacks, nil
	}
	return nil, &NotLoadedError{edge: "mission_failed_feedbacks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldLinking:
			values[i] = new(sql.NullBool)
		case device.FieldDisk, device.FieldDelay, device.FieldGpuTemperature, device.FieldCPUTemperature:
			values[i] = new(sql.NullFloat64)
		case device.FieldID, device.FieldCreatedBy, device.FieldUpdatedBy, device.FieldUserID, device.FieldGiftMissionConfigID, device.FieldSumCep, device.FieldCoresNumber, device.FieldMemory, device.FieldFreeGpuNum:
			values[i] = new(sql.NullInt64)
		case device.FieldSerialNumber, device.FieldState, device.FieldBindingStatus, device.FieldStatus, device.FieldName, device.FieldManageName, device.FieldType, device.FieldCPU, device.FieldStability, device.FieldVersion, device.FieldFault, device.FieldRank, device.FieldHostingType, device.FieldMissionTag:
			values[i] = new(sql.NullString)
		case device.FieldCreatedAt, device.FieldUpdatedAt, device.FieldDeletedAt, device.FieldRankAt, device.FieldStabilityAt, device.FieldHighTemperatureAt:
			values[i] = new(sql.NullTime)
		case device.FieldCpus:
			values[i] = device.ValueScanner.Cpus.ScanValue()
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int64(value.Int64)
		case device.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				d.CreatedBy = value.Int64
			}
		case device.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				d.UpdatedBy = value.Int64
			}
		case device.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case device.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case device.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = value.Time
			}
		case device.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				d.UserID = value.Int64
			}
		case device.FieldGiftMissionConfigID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gift_mission_config_id", values[i])
			} else if value.Valid {
				d.GiftMissionConfigID = value.Int64
			}
		case device.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				d.SerialNumber = value.String
			}
		case device.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				d.State = device.State(value.String)
			}
		case device.FieldSumCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sum_cep", values[i])
			} else if value.Valid {
				d.SumCep = value.Int64
			}
		case device.FieldLinking:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field linking", values[i])
			} else if value.Valid {
				d.Linking = value.Bool
			}
		case device.FieldBindingStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field binding_status", values[i])
			} else if value.Valid {
				d.BindingStatus = enums.DeviceBindingStatus(value.String)
			}
		case device.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = enums.DeviceStatus(value.String)
			}
		case device.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case device.FieldManageName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manage_name", values[i])
			} else if value.Valid {
				d.ManageName = value.String
			}
		case device.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = enums.DeviceType(value.String)
			}
		case device.FieldCoresNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cores_number", values[i])
			} else if value.Valid {
				d.CoresNumber = value.Int64
			}
		case device.FieldCPU:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cpu", values[i])
			} else if value.Valid {
				d.CPU = value.String
			}
		case device.FieldCpus:
			if value, err := device.ValueScanner.Cpus.FromValue(values[i]); err != nil {
				return err
			} else {
				d.Cpus = value
			}
		case device.FieldMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				d.Memory = value.Int64
			}
		case device.FieldDisk:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field disk", values[i])
			} else if value.Valid {
				d.Disk = float32(value.Float64)
			}
		case device.FieldDelay:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field delay", values[i])
			} else if value.Valid {
				d.Delay = value.Float64
			}
		case device.FieldGpuTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_temperature", values[i])
			} else if value.Valid {
				d.GpuTemperature = value.Float64
			}
		case device.FieldCPUTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_temperature", values[i])
			} else if value.Valid {
				d.CPUTemperature = value.Float64
			}
		case device.FieldStability:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stability", values[i])
			} else if value.Valid {
				d.Stability = enums.DeviceStabilityType(value.String)
			}
		case device.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				d.Version = value.String
			}
		case device.FieldFault:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fault", values[i])
			} else if value.Valid {
				d.Fault = enums.DeviceFaultType(value.String)
			}
		case device.FieldRank:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				d.Rank = enums.DeviceRankType(value.String)
			}
		case device.FieldFreeGpuNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field free_gpu_num", values[i])
			} else if value.Valid {
				d.FreeGpuNum = int(value.Int64)
			}
		case device.FieldRankAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field rank_at", values[i])
			} else if value.Valid {
				d.RankAt = new(time.Time)
				*d.RankAt = value.Time
			}
		case device.FieldStabilityAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field stability_at", values[i])
			} else if value.Valid {
				d.StabilityAt = new(time.Time)
				*d.StabilityAt = value.Time
			}
		case device.FieldHighTemperatureAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field high_temperature_at", values[i])
			} else if value.Valid {
				d.HighTemperatureAt = new(time.Time)
				*d.HighTemperatureAt = value.Time
			}
		case device.FieldHostingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hosting_type", values[i])
			} else if value.Valid {
				d.HostingType = enums.DeviceHostingType(value.String)
			}
		case device.FieldMissionTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_tag", values[i])
			} else if value.Valid {
				d.MissionTag = enums.DeviceMissionTag(value.String)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Device.
// This includes values selected through modifiers, order, etc.
func (d *Device) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Device entity.
func (d *Device) QueryUser() *UserQuery {
	return NewDeviceClient(d.config).QueryUser(d)
}

// QueryGiftMissionConfig queries the "gift_mission_config" edge of the Device entity.
func (d *Device) QueryGiftMissionConfig() *GiftMissionConfigQuery {
	return NewDeviceClient(d.config).QueryGiftMissionConfig(d)
}

// QueryMissionProduceOrders queries the "mission_produce_orders" edge of the Device entity.
func (d *Device) QueryMissionProduceOrders() *MissionProduceOrderQuery {
	return NewDeviceClient(d.config).QueryMissionProduceOrders(d)
}

// QueryUserDevices queries the "user_devices" edge of the Device entity.
func (d *Device) QueryUserDevices() *UserDeviceQuery {
	return NewDeviceClient(d.config).QueryUserDevices(d)
}

// QueryDeviceGpuMissions queries the "device_gpu_missions" edge of the Device entity.
func (d *Device) QueryDeviceGpuMissions() *DeviceGpuMissionQuery {
	return NewDeviceClient(d.config).QueryDeviceGpuMissions(d)
}

// QueryFrpcInfos queries the "frpc_infos" edge of the Device entity.
func (d *Device) QueryFrpcInfos() *FrpcInfoQuery {
	return NewDeviceClient(d.config).QueryFrpcInfos(d)
}

// QueryMissionOrders queries the "mission_orders" edge of the Device entity.
func (d *Device) QueryMissionOrders() *MissionOrderQuery {
	return NewDeviceClient(d.config).QueryMissionOrders(d)
}

// QueryMissionProductions queries the "mission_productions" edge of the Device entity.
func (d *Device) QueryMissionProductions() *MissionProductionQuery {
	return NewDeviceClient(d.config).QueryMissionProductions(d)
}

// QueryDeviceRebootTimes queries the "device_reboot_times" edge of the Device entity.
func (d *Device) QueryDeviceRebootTimes() *DeviceRebootTimeQuery {
	return NewDeviceClient(d.config).QueryDeviceRebootTimes(d)
}

// QueryTroubleDeducts queries the "trouble_deducts" edge of the Device entity.
func (d *Device) QueryTroubleDeducts() *TroubleDeductQuery {
	return NewDeviceClient(d.config).QueryTroubleDeducts(d)
}

// QueryDeviceStates queries the "device_states" edge of the Device entity.
func (d *Device) QueryDeviceStates() *DeviceStateQuery {
	return NewDeviceClient(d.config).QueryDeviceStates(d)
}

// QueryDeviceOfflineRecords queries the "device_offline_records" edge of the Device entity.
func (d *Device) QueryDeviceOfflineRecords() *DeviceOfflineRecordQuery {
	return NewDeviceClient(d.config).QueryDeviceOfflineRecords(d)
}

// QueryMissionFailedFeedbacks queries the "mission_failed_feedbacks" edge of the Device entity.
func (d *Device) QueryMissionFailedFeedbacks() *MissionFailedFeedbackQuery {
	return NewDeviceClient(d.config).QueryMissionFailedFeedbacks(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return NewDeviceClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(d.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", d.UserID))
	builder.WriteString(", ")
	builder.WriteString("gift_mission_config_id=")
	builder.WriteString(fmt.Sprintf("%v", d.GiftMissionConfigID))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(d.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", d.State))
	builder.WriteString(", ")
	builder.WriteString("sum_cep=")
	builder.WriteString(fmt.Sprintf("%v", d.SumCep))
	builder.WriteString(", ")
	builder.WriteString("linking=")
	builder.WriteString(fmt.Sprintf("%v", d.Linking))
	builder.WriteString(", ")
	builder.WriteString("binding_status=")
	builder.WriteString(fmt.Sprintf("%v", d.BindingStatus))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("manage_name=")
	builder.WriteString(d.ManageName)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", d.Type))
	builder.WriteString(", ")
	builder.WriteString("cores_number=")
	builder.WriteString(fmt.Sprintf("%v", d.CoresNumber))
	builder.WriteString(", ")
	builder.WriteString("cpu=")
	builder.WriteString(d.CPU)
	builder.WriteString(", ")
	builder.WriteString("cpus=")
	builder.WriteString(fmt.Sprintf("%v", d.Cpus))
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(fmt.Sprintf("%v", d.Memory))
	builder.WriteString(", ")
	builder.WriteString("disk=")
	builder.WriteString(fmt.Sprintf("%v", d.Disk))
	builder.WriteString(", ")
	builder.WriteString("delay=")
	builder.WriteString(fmt.Sprintf("%v", d.Delay))
	builder.WriteString(", ")
	builder.WriteString("gpu_temperature=")
	builder.WriteString(fmt.Sprintf("%v", d.GpuTemperature))
	builder.WriteString(", ")
	builder.WriteString("cpu_temperature=")
	builder.WriteString(fmt.Sprintf("%v", d.CPUTemperature))
	builder.WriteString(", ")
	builder.WriteString("stability=")
	builder.WriteString(fmt.Sprintf("%v", d.Stability))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(d.Version)
	builder.WriteString(", ")
	builder.WriteString("fault=")
	builder.WriteString(fmt.Sprintf("%v", d.Fault))
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", d.Rank))
	builder.WriteString(", ")
	builder.WriteString("free_gpu_num=")
	builder.WriteString(fmt.Sprintf("%v", d.FreeGpuNum))
	builder.WriteString(", ")
	if v := d.RankAt; v != nil {
		builder.WriteString("rank_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.StabilityAt; v != nil {
		builder.WriteString("stability_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.HighTemperatureAt; v != nil {
		builder.WriteString("high_temperature_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("hosting_type=")
	builder.WriteString(fmt.Sprintf("%v", d.HostingType))
	builder.WriteString(", ")
	builder.WriteString("mission_tag=")
	builder.WriteString(fmt.Sprintf("%v", d.MissionTag))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
