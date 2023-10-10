// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 设备，对应客户端 core，记录心跳等信息
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at"`
	// 外键用户 id
	UserID int64 `json:"user_id,omitempty"`
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
	Status device.Status `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges        DeviceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
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
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
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

// MissionProduceOrdersOrErr returns the MissionProduceOrders value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionProduceOrdersOrErr() ([]*MissionProduceOrder, error) {
	if e.loadedTypes[1] {
		return e.MissionProduceOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_orders"}
}

// UserDevicesOrErr returns the UserDevices value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) UserDevicesOrErr() ([]*UserDevice, error) {
	if e.loadedTypes[2] {
		return e.UserDevices, nil
	}
	return nil, &NotLoadedError{edge: "user_devices"}
}

// DeviceGpuMissionsOrErr returns the DeviceGpuMissions value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) DeviceGpuMissionsOrErr() ([]*DeviceGpuMission, error) {
	if e.loadedTypes[3] {
		return e.DeviceGpuMissions, nil
	}
	return nil, &NotLoadedError{edge: "device_gpu_missions"}
}

// FrpcInfosOrErr returns the FrpcInfos value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) FrpcInfosOrErr() ([]*FrpcInfo, error) {
	if e.loadedTypes[4] {
		return e.FrpcInfos, nil
	}
	return nil, &NotLoadedError{edge: "frpc_infos"}
}

// MissionOrdersOrErr returns the MissionOrders value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) MissionOrdersOrErr() ([]*MissionOrder, error) {
	if e.loadedTypes[5] {
		return e.MissionOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldLinking:
			values[i] = new(sql.NullBool)
		case device.FieldID, device.FieldCreatedBy, device.FieldUpdatedBy, device.FieldUserID, device.FieldSumCep:
			values[i] = new(sql.NullInt64)
		case device.FieldSerialNumber, device.FieldState, device.FieldBindingStatus, device.FieldStatus:
			values[i] = new(sql.NullString)
		case device.FieldCreatedAt, device.FieldUpdatedAt, device.FieldDeletedAt:
			values[i] = new(sql.NullTime)
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
				d.Status = device.Status(value.String)
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
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
