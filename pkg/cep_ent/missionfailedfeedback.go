// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionfailedfeedback"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 应用启动失败反馈信息表
type MissionFailedFeedback struct {
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
	// 外键，反馈的用户 ID
	UserID int64 `json:"user_id,string"`
	// 外键，反馈关联的设备 ID
	DeviceID int64 `json:"device_id,string"`
	// 外键，反馈关联的任务 ID
	MissionID int64 `json:"mission_id,string"`
	// 应用名称
	MissionName string `json:"mission_name"`
	// 状态
	Status enums.MissionFailedFeedbackStatus `json:"status"`
	// 任务失败的原因
	Reason string `json:"reason"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionFailedFeedbackQuery when eager-loading is set.
	Edges        MissionFailedFeedbackEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionFailedFeedbackEdges holds the relations/edges for other nodes in the graph.
type MissionFailedFeedbackEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionFailedFeedbackEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionFailedFeedbackEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[1] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionFailedFeedbackEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[2] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionFailedFeedback) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionfailedfeedback.FieldID, missionfailedfeedback.FieldCreatedBy, missionfailedfeedback.FieldUpdatedBy, missionfailedfeedback.FieldUserID, missionfailedfeedback.FieldDeviceID, missionfailedfeedback.FieldMissionID:
			values[i] = new(sql.NullInt64)
		case missionfailedfeedback.FieldMissionName, missionfailedfeedback.FieldStatus, missionfailedfeedback.FieldReason:
			values[i] = new(sql.NullString)
		case missionfailedfeedback.FieldCreatedAt, missionfailedfeedback.FieldUpdatedAt, missionfailedfeedback.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionFailedFeedback fields.
func (mff *MissionFailedFeedback) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionfailedfeedback.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mff.ID = int64(value.Int64)
		case missionfailedfeedback.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mff.CreatedBy = value.Int64
			}
		case missionfailedfeedback.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mff.UpdatedBy = value.Int64
			}
		case missionfailedfeedback.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mff.CreatedAt = value.Time
			}
		case missionfailedfeedback.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mff.UpdatedAt = value.Time
			}
		case missionfailedfeedback.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mff.DeletedAt = value.Time
			}
		case missionfailedfeedback.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				mff.UserID = value.Int64
			}
		case missionfailedfeedback.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				mff.DeviceID = value.Int64
			}
		case missionfailedfeedback.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				mff.MissionID = value.Int64
			}
		case missionfailedfeedback.FieldMissionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_name", values[i])
			} else if value.Valid {
				mff.MissionName = value.String
			}
		case missionfailedfeedback.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mff.Status = enums.MissionFailedFeedbackStatus(value.String)
			}
		case missionfailedfeedback.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				mff.Reason = value.String
			}
		default:
			mff.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionFailedFeedback.
// This includes values selected through modifiers, order, etc.
func (mff *MissionFailedFeedback) Value(name string) (ent.Value, error) {
	return mff.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the MissionFailedFeedback entity.
func (mff *MissionFailedFeedback) QueryUser() *UserQuery {
	return NewMissionFailedFeedbackClient(mff.config).QueryUser(mff)
}

// QueryDevice queries the "device" edge of the MissionFailedFeedback entity.
func (mff *MissionFailedFeedback) QueryDevice() *DeviceQuery {
	return NewMissionFailedFeedbackClient(mff.config).QueryDevice(mff)
}

// QueryMission queries the "mission" edge of the MissionFailedFeedback entity.
func (mff *MissionFailedFeedback) QueryMission() *MissionQuery {
	return NewMissionFailedFeedbackClient(mff.config).QueryMission(mff)
}

// Update returns a builder for updating this MissionFailedFeedback.
// Note that you need to call MissionFailedFeedback.Unwrap() before calling this method if this MissionFailedFeedback
// was returned from a transaction, and the transaction was committed or rolled back.
func (mff *MissionFailedFeedback) Update() *MissionFailedFeedbackUpdateOne {
	return NewMissionFailedFeedbackClient(mff.config).UpdateOne(mff)
}

// Unwrap unwraps the MissionFailedFeedback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mff *MissionFailedFeedback) Unwrap() *MissionFailedFeedback {
	_tx, ok := mff.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionFailedFeedback is not a transactional entity")
	}
	mff.config.driver = _tx.drv
	return mff
}

// String implements the fmt.Stringer.
func (mff *MissionFailedFeedback) String() string {
	var builder strings.Builder
	builder.WriteString("MissionFailedFeedback(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mff.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mff.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mff.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mff.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mff.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mff.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", mff.UserID))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", mff.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", mff.MissionID))
	builder.WriteString(", ")
	builder.WriteString("mission_name=")
	builder.WriteString(mff.MissionName)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mff.Status))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(mff.Reason)
	builder.WriteByte(')')
	return builder.String()
}

// MissionFailedFeedbacks is a parsable slice of MissionFailedFeedback.
type MissionFailedFeedbacks []*MissionFailedFeedback
