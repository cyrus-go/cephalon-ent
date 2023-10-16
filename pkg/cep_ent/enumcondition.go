// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/enumcondition"
)

// 任务类型枚举值对应表，计划用代码实现取代
type EnumCondition struct {
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
	// 给到前端的类型
	FrontType string `json:"front_type"`
	// 任务类型
	MissionType string `json:"mission_type"`
	// 任务调用方式
	MissionCallWay string `json:"mission_call_way"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EnumCondition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case enumcondition.FieldID, enumcondition.FieldCreatedBy, enumcondition.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case enumcondition.FieldFrontType, enumcondition.FieldMissionType, enumcondition.FieldMissionCallWay:
			values[i] = new(sql.NullString)
		case enumcondition.FieldCreatedAt, enumcondition.FieldUpdatedAt, enumcondition.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EnumCondition fields.
func (ec *EnumCondition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case enumcondition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int64(value.Int64)
		case enumcondition.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ec.CreatedBy = value.Int64
			}
		case enumcondition.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ec.UpdatedBy = value.Int64
			}
		case enumcondition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ec.CreatedAt = value.Time
			}
		case enumcondition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ec.UpdatedAt = value.Time
			}
		case enumcondition.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ec.DeletedAt = value.Time
			}
		case enumcondition.FieldFrontType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field front_type", values[i])
			} else if value.Valid {
				ec.FrontType = value.String
			}
		case enumcondition.FieldMissionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_type", values[i])
			} else if value.Valid {
				ec.MissionType = value.String
			}
		case enumcondition.FieldMissionCallWay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_call_way", values[i])
			} else if value.Valid {
				ec.MissionCallWay = value.String
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EnumCondition.
// This includes values selected through modifiers, order, etc.
func (ec *EnumCondition) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// Update returns a builder for updating this EnumCondition.
// Note that you need to call EnumCondition.Unwrap() before calling this method if this EnumCondition
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EnumCondition) Update() *EnumConditionUpdateOne {
	return NewEnumConditionClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EnumCondition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EnumCondition) Unwrap() *EnumCondition {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: EnumCondition is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EnumCondition) String() string {
	var builder strings.Builder
	builder.WriteString("EnumCondition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ec.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ec.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ec.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ec.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ec.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("front_type=")
	builder.WriteString(ec.FrontType)
	builder.WriteString(", ")
	builder.WriteString("mission_type=")
	builder.WriteString(ec.MissionType)
	builder.WriteString(", ")
	builder.WriteString("mission_call_way=")
	builder.WriteString(ec.MissionCallWay)
	builder.WriteByte(')')
	return builder.String()
}

// EnumConditions is a parsable slice of EnumCondition.
type EnumConditions []*EnumCondition
