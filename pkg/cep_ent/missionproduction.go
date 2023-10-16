// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduction"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务执行情况，记录谁接了什么任务，用什么接，做得怎么样
type MissionProduction struct {
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
	// 任务 ID
	MissionID int64 `json:"mission_id,string"`
	// 用户 ID
	UserID int64 `json:"user_id,string"`
	// 任务开始时刻
	StartedAt time.Time `json:"started_at"`
	// 任务完成时刻
	FinishedAt time.Time `json:"finished_at"`
	// 任务执行状态情况
	State enums.MissionState `json:"state"`
	// 领到任务的设备 ID
	DeviceID int64 `json:"device_id,string"`
	// 任务使用什么显卡在执行
	GpuVersion enums.GpuVersion `json:"gpu_version"`
	// 任务结果链接列表，json 序列化后存储
	Urls string `json:"urls"`
	// 内部功能返回码
	RespStatusCode int32 `json:"resp_status_code"`
	// 返回内容体 json 转 string
	RespBody string `json:"resp_body"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionProductionQuery when eager-loading is set.
	Edges        MissionProductionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionProductionEdges holds the relations/edges for other nodes in the graph.
type MissionProductionEdges struct {
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// MissionProduceOrder holds the value of the mission_produce_order edge.
	MissionProduceOrder *MissionProduceOrder `json:"mission_produce_order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProductionEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[0] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProductionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MissionProduceOrderOrErr returns the MissionProduceOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProductionEdges) MissionProduceOrderOrErr() (*MissionProduceOrder, error) {
	if e.loadedTypes[2] {
		if e.MissionProduceOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionproduceorder.Label}
		}
		return e.MissionProduceOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionProduction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionproduction.FieldID, missionproduction.FieldCreatedBy, missionproduction.FieldUpdatedBy, missionproduction.FieldMissionID, missionproduction.FieldUserID, missionproduction.FieldDeviceID, missionproduction.FieldRespStatusCode:
			values[i] = new(sql.NullInt64)
		case missionproduction.FieldState, missionproduction.FieldGpuVersion, missionproduction.FieldUrls, missionproduction.FieldRespBody:
			values[i] = new(sql.NullString)
		case missionproduction.FieldCreatedAt, missionproduction.FieldUpdatedAt, missionproduction.FieldDeletedAt, missionproduction.FieldStartedAt, missionproduction.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionProduction fields.
func (mp *MissionProduction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionproduction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mp.ID = int64(value.Int64)
		case missionproduction.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mp.CreatedBy = value.Int64
			}
		case missionproduction.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mp.UpdatedBy = value.Int64
			}
		case missionproduction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mp.CreatedAt = value.Time
			}
		case missionproduction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mp.UpdatedAt = value.Time
			}
		case missionproduction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mp.DeletedAt = value.Time
			}
		case missionproduction.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				mp.MissionID = value.Int64
			}
		case missionproduction.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				mp.UserID = value.Int64
			}
		case missionproduction.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				mp.StartedAt = value.Time
			}
		case missionproduction.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				mp.FinishedAt = value.Time
			}
		case missionproduction.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				mp.State = enums.MissionState(value.String)
			}
		case missionproduction.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				mp.DeviceID = value.Int64
			}
		case missionproduction.FieldGpuVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_version", values[i])
			} else if value.Valid {
				mp.GpuVersion = enums.GpuVersion(value.String)
			}
		case missionproduction.FieldUrls:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field urls", values[i])
			} else if value.Valid {
				mp.Urls = value.String
			}
		case missionproduction.FieldRespStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resp_status_code", values[i])
			} else if value.Valid {
				mp.RespStatusCode = int32(value.Int64)
			}
		case missionproduction.FieldRespBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resp_body", values[i])
			} else if value.Valid {
				mp.RespBody = value.String
			}
		default:
			mp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionProduction.
// This includes values selected through modifiers, order, etc.
func (mp *MissionProduction) Value(name string) (ent.Value, error) {
	return mp.selectValues.Get(name)
}

// QueryMission queries the "mission" edge of the MissionProduction entity.
func (mp *MissionProduction) QueryMission() *MissionQuery {
	return NewMissionProductionClient(mp.config).QueryMission(mp)
}

// QueryUser queries the "user" edge of the MissionProduction entity.
func (mp *MissionProduction) QueryUser() *UserQuery {
	return NewMissionProductionClient(mp.config).QueryUser(mp)
}

// QueryMissionProduceOrder queries the "mission_produce_order" edge of the MissionProduction entity.
func (mp *MissionProduction) QueryMissionProduceOrder() *MissionProduceOrderQuery {
	return NewMissionProductionClient(mp.config).QueryMissionProduceOrder(mp)
}

// Update returns a builder for updating this MissionProduction.
// Note that you need to call MissionProduction.Unwrap() before calling this method if this MissionProduction
// was returned from a transaction, and the transaction was committed or rolled back.
func (mp *MissionProduction) Update() *MissionProductionUpdateOne {
	return NewMissionProductionClient(mp.config).UpdateOne(mp)
}

// Unwrap unwraps the MissionProduction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mp *MissionProduction) Unwrap() *MissionProduction {
	_tx, ok := mp.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionProduction is not a transactional entity")
	}
	mp.config.driver = _tx.drv
	return mp
}

// String implements the fmt.Stringer.
func (mp *MissionProduction) String() string {
	var builder strings.Builder
	builder.WriteString("MissionProduction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.MissionID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.UserID))
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(mp.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(mp.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", mp.State))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("gpu_version=")
	builder.WriteString(fmt.Sprintf("%v", mp.GpuVersion))
	builder.WriteString(", ")
	builder.WriteString("urls=")
	builder.WriteString(mp.Urls)
	builder.WriteString(", ")
	builder.WriteString("resp_status_code=")
	builder.WriteString(fmt.Sprintf("%v", mp.RespStatusCode))
	builder.WriteString(", ")
	builder.WriteString("resp_body=")
	builder.WriteString(mp.RespBody)
	builder.WriteByte(')')
	return builder.String()
}

// MissionProductions is a parsable slice of MissionProduction.
type MissionProductions []*MissionProduction
