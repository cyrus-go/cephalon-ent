// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/renewalagreement"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 自动续费协议
type RenewalAgreement struct {
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
	// 下次扣款时间
	NextPayTime time.Time `json:"next_pay_time"`
	// 自动续费类型（按小时、按天等）
	Type enums.RenewalType `json:"type"`
	// 订阅自动续费状态（订阅中、已结束等）
	SubStatus enums.RenewalSubStatus `json:"sub_status"`
	// 支付状态（待支付、已支付、支付失败等）
	PayStatus enums.RenewalPayStatus `json:"pay_status"`
	// 币种 id
	SymbolID int64 `json:"symbol_id,string"`
	// 首次扣款价格
	FirstPay int64 `json:"first_pay"`
	// 后续扣款价格
	AfterPay int64 `json:"after_pay"`
	// 最后一次预警时间
	LastWarningTime time.Time `json:"last_warning_time"`
	// 订阅自动续费结束时间
	SubFinishedTime time.Time `json:"sub_finished_time"`
	// 外键用户 id
	UserID int64 `json:"user_id,string"`
	// 外键任务 id
	MissionID int64 `json:"mission_id,string"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RenewalAgreementQuery when eager-loading is set.
	Edges        RenewalAgreementEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RenewalAgreementEdges holds the relations/edges for other nodes in the graph.
type RenewalAgreementEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Mission holds the value of the mission edge.
	Mission *Mission `json:"mission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RenewalAgreementEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MissionOrErr returns the Mission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RenewalAgreementEdges) MissionOrErr() (*Mission, error) {
	if e.loadedTypes[1] {
		if e.Mission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.Mission, nil
	}
	return nil, &NotLoadedError{edge: "mission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RenewalAgreement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case renewalagreement.FieldID, renewalagreement.FieldCreatedBy, renewalagreement.FieldUpdatedBy, renewalagreement.FieldSymbolID, renewalagreement.FieldFirstPay, renewalagreement.FieldAfterPay, renewalagreement.FieldUserID, renewalagreement.FieldMissionID:
			values[i] = new(sql.NullInt64)
		case renewalagreement.FieldType, renewalagreement.FieldSubStatus, renewalagreement.FieldPayStatus:
			values[i] = new(sql.NullString)
		case renewalagreement.FieldCreatedAt, renewalagreement.FieldUpdatedAt, renewalagreement.FieldDeletedAt, renewalagreement.FieldNextPayTime, renewalagreement.FieldLastWarningTime, renewalagreement.FieldSubFinishedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RenewalAgreement fields.
func (ra *RenewalAgreement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case renewalagreement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ra.ID = int64(value.Int64)
		case renewalagreement.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ra.CreatedBy = value.Int64
			}
		case renewalagreement.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ra.UpdatedBy = value.Int64
			}
		case renewalagreement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ra.CreatedAt = value.Time
			}
		case renewalagreement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ra.UpdatedAt = value.Time
			}
		case renewalagreement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ra.DeletedAt = value.Time
			}
		case renewalagreement.FieldNextPayTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field next_pay_time", values[i])
			} else if value.Valid {
				ra.NextPayTime = value.Time
			}
		case renewalagreement.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ra.Type = enums.RenewalType(value.String)
			}
		case renewalagreement.FieldSubStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub_status", values[i])
			} else if value.Valid {
				ra.SubStatus = enums.RenewalSubStatus(value.String)
			}
		case renewalagreement.FieldPayStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pay_status", values[i])
			} else if value.Valid {
				ra.PayStatus = enums.RenewalPayStatus(value.String)
			}
		case renewalagreement.FieldSymbolID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field symbol_id", values[i])
			} else if value.Valid {
				ra.SymbolID = value.Int64
			}
		case renewalagreement.FieldFirstPay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field first_pay", values[i])
			} else if value.Valid {
				ra.FirstPay = value.Int64
			}
		case renewalagreement.FieldAfterPay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field after_pay", values[i])
			} else if value.Valid {
				ra.AfterPay = value.Int64
			}
		case renewalagreement.FieldLastWarningTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_warning_time", values[i])
			} else if value.Valid {
				ra.LastWarningTime = value.Time
			}
		case renewalagreement.FieldSubFinishedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sub_finished_time", values[i])
			} else if value.Valid {
				ra.SubFinishedTime = value.Time
			}
		case renewalagreement.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ra.UserID = value.Int64
			}
		case renewalagreement.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				ra.MissionID = value.Int64
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RenewalAgreement.
// This includes values selected through modifiers, order, etc.
func (ra *RenewalAgreement) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the RenewalAgreement entity.
func (ra *RenewalAgreement) QueryUser() *UserQuery {
	return NewRenewalAgreementClient(ra.config).QueryUser(ra)
}

// QueryMission queries the "mission" edge of the RenewalAgreement entity.
func (ra *RenewalAgreement) QueryMission() *MissionQuery {
	return NewRenewalAgreementClient(ra.config).QueryMission(ra)
}

// Update returns a builder for updating this RenewalAgreement.
// Note that you need to call RenewalAgreement.Unwrap() before calling this method if this RenewalAgreement
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *RenewalAgreement) Update() *RenewalAgreementUpdateOne {
	return NewRenewalAgreementClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the RenewalAgreement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *RenewalAgreement) Unwrap() *RenewalAgreement {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: RenewalAgreement is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *RenewalAgreement) String() string {
	var builder strings.Builder
	builder.WriteString("RenewalAgreement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ra.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ra.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ra.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ra.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ra.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("next_pay_time=")
	builder.WriteString(ra.NextPayTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ra.Type))
	builder.WriteString(", ")
	builder.WriteString("sub_status=")
	builder.WriteString(fmt.Sprintf("%v", ra.SubStatus))
	builder.WriteString(", ")
	builder.WriteString("pay_status=")
	builder.WriteString(fmt.Sprintf("%v", ra.PayStatus))
	builder.WriteString(", ")
	builder.WriteString("symbol_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.SymbolID))
	builder.WriteString(", ")
	builder.WriteString("first_pay=")
	builder.WriteString(fmt.Sprintf("%v", ra.FirstPay))
	builder.WriteString(", ")
	builder.WriteString("after_pay=")
	builder.WriteString(fmt.Sprintf("%v", ra.AfterPay))
	builder.WriteString(", ")
	builder.WriteString("last_warning_time=")
	builder.WriteString(ra.LastWarningTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sub_finished_time=")
	builder.WriteString(ra.SubFinishedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.UserID))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.MissionID))
	builder.WriteByte(')')
	return builder.String()
}

// RenewalAgreements is a parsable slice of RenewalAgreement.
type RenewalAgreements []*RenewalAgreement
