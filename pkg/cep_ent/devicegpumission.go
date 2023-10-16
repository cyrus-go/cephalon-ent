// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/devicegpumission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/gpu"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionkind"
)

// 登记设备的显卡信息，以及设备的任务执行能力配置状态
type DeviceGpuMission struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 外键设备 id
	DeviceID int64 `json:"device_id,string"`
	// 外键 gpu id
	GpuID int64 `json:"gpu_id,string"`
	// 外键任务种类 id
	MissionKindID int64 `json:"mission_kind_id,string"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceGpuMissionQuery when eager-loading is set.
	Edges        DeviceGpuMissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeviceGpuMissionEdges holds the relations/edges for other nodes in the graph.
type DeviceGpuMissionEdges struct {
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// MissionKind holds the value of the mission_kind edge.
	MissionKind *MissionKind `json:"mission_kind,omitempty"`
	// Gpu holds the value of the gpu edge.
	Gpu *Gpu `json:"gpu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceGpuMissionEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[0] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// MissionKindOrErr returns the MissionKind value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceGpuMissionEdges) MissionKindOrErr() (*MissionKind, error) {
	if e.loadedTypes[1] {
		if e.MissionKind == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionkind.Label}
		}
		return e.MissionKind, nil
	}
	return nil, &NotLoadedError{edge: "mission_kind"}
}

// GpuOrErr returns the Gpu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceGpuMissionEdges) GpuOrErr() (*Gpu, error) {
	if e.loadedTypes[2] {
		if e.Gpu == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: gpu.Label}
		}
		return e.Gpu, nil
	}
	return nil, &NotLoadedError{edge: "gpu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeviceGpuMission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case devicegpumission.FieldID, devicegpumission.FieldCreatedBy, devicegpumission.FieldUpdatedBy, devicegpumission.FieldDeviceID, devicegpumission.FieldGpuID, devicegpumission.FieldMissionKindID:
			values[i] = new(sql.NullInt64)
		case devicegpumission.FieldCreatedAt, devicegpumission.FieldUpdatedAt, devicegpumission.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeviceGpuMission fields.
func (dgm *DeviceGpuMission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case devicegpumission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dgm.ID = int64(value.Int64)
		case devicegpumission.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				dgm.CreatedBy = value.Int64
			}
		case devicegpumission.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				dgm.UpdatedBy = value.Int64
			}
		case devicegpumission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dgm.CreatedAt = value.Time
			}
		case devicegpumission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dgm.UpdatedAt = value.Time
			}
		case devicegpumission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				dgm.DeletedAt = value.Time
			}
		case devicegpumission.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				dgm.DeviceID = value.Int64
			}
		case devicegpumission.FieldGpuID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_id", values[i])
			} else if value.Valid {
				dgm.GpuID = value.Int64
			}
		case devicegpumission.FieldMissionKindID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_kind_id", values[i])
			} else if value.Valid {
				dgm.MissionKindID = value.Int64
			}
		default:
			dgm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DeviceGpuMission.
// This includes values selected through modifiers, order, etc.
func (dgm *DeviceGpuMission) Value(name string) (ent.Value, error) {
	return dgm.selectValues.Get(name)
}

// QueryDevice queries the "device" edge of the DeviceGpuMission entity.
func (dgm *DeviceGpuMission) QueryDevice() *DeviceQuery {
	return NewDeviceGpuMissionClient(dgm.config).QueryDevice(dgm)
}

// QueryMissionKind queries the "mission_kind" edge of the DeviceGpuMission entity.
func (dgm *DeviceGpuMission) QueryMissionKind() *MissionKindQuery {
	return NewDeviceGpuMissionClient(dgm.config).QueryMissionKind(dgm)
}

// QueryGpu queries the "gpu" edge of the DeviceGpuMission entity.
func (dgm *DeviceGpuMission) QueryGpu() *GpuQuery {
	return NewDeviceGpuMissionClient(dgm.config).QueryGpu(dgm)
}

// Update returns a builder for updating this DeviceGpuMission.
// Note that you need to call DeviceGpuMission.Unwrap() before calling this method if this DeviceGpuMission
// was returned from a transaction, and the transaction was committed or rolled back.
func (dgm *DeviceGpuMission) Update() *DeviceGpuMissionUpdateOne {
	return NewDeviceGpuMissionClient(dgm.config).UpdateOne(dgm)
}

// Unwrap unwraps the DeviceGpuMission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dgm *DeviceGpuMission) Unwrap() *DeviceGpuMission {
	_tx, ok := dgm.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: DeviceGpuMission is not a transactional entity")
	}
	dgm.config.driver = _tx.drv
	return dgm
}

// String implements the fmt.Stringer.
func (dgm *DeviceGpuMission) String() string {
	var builder strings.Builder
	builder.WriteString("DeviceGpuMission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dgm.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", dgm.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", dgm.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dgm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dgm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(dgm.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", dgm.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("gpu_id=")
	builder.WriteString(fmt.Sprintf("%v", dgm.GpuID))
	builder.WriteString(", ")
	builder.WriteString("mission_kind_id=")
	builder.WriteString(fmt.Sprintf("%v", dgm.MissionKindID))
	builder.WriteByte(')')
	return builder.String()
}

// DeviceGpuMissions is a parsable slice of DeviceGpuMission.
type DeviceGpuMissions []*DeviceGpuMission
