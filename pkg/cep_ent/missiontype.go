// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/missiontype"
	"cephalon-ent/pkg/enums"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 任务类型等信息，给任务定价归类等
type MissionType struct {
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
	// 任务类型
	Type enums.MissionType `json:"type"`
	// 显卡型号
	Gpu enums.GPU `json:"gpu"`
	// 单价消耗 cep
	Cep int64 `json:"cep"`
	// 是否计时任务
	IsTime bool `json:"is_time"`
	// 任务种类，SD，Jupyter 等
	Category     enums.MissionCategory `json:"category"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missiontype.FieldIsTime:
			values[i] = new(sql.NullBool)
		case missiontype.FieldID, missiontype.FieldCreatedBy, missiontype.FieldUpdatedBy, missiontype.FieldCep:
			values[i] = new(sql.NullInt64)
		case missiontype.FieldType, missiontype.FieldGpu, missiontype.FieldCategory:
			values[i] = new(sql.NullString)
		case missiontype.FieldCreatedAt, missiontype.FieldUpdatedAt, missiontype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionType fields.
func (mt *MissionType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missiontype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mt.ID = int64(value.Int64)
		case missiontype.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mt.CreatedBy = value.Int64
			}
		case missiontype.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mt.UpdatedBy = value.Int64
			}
		case missiontype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mt.CreatedAt = value.Time
			}
		case missiontype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mt.UpdatedAt = value.Time
			}
		case missiontype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mt.DeletedAt = value.Time
			}
		case missiontype.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mt.Type = enums.MissionType(value.String)
			}
		case missiontype.FieldGpu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gpu", values[i])
			} else if value.Valid {
				mt.Gpu = enums.GPU(value.String)
			}
		case missiontype.FieldCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cep", values[i])
			} else if value.Valid {
				mt.Cep = value.Int64
			}
		case missiontype.FieldIsTime:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_time", values[i])
			} else if value.Valid {
				mt.IsTime = value.Bool
			}
		case missiontype.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				mt.Category = enums.MissionCategory(value.String)
			}
		default:
			mt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionType.
// This includes values selected through modifiers, order, etc.
func (mt *MissionType) Value(name string) (ent.Value, error) {
	return mt.selectValues.Get(name)
}

// Update returns a builder for updating this MissionType.
// Note that you need to call MissionType.Unwrap() before calling this method if this MissionType
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MissionType) Update() *MissionTypeUpdateOne {
	return NewMissionTypeClient(mt.config).UpdateOne(mt)
}

// Unwrap unwraps the MissionType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mt *MissionType) Unwrap() *MissionType {
	_tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionType is not a transactional entity")
	}
	mt.config.driver = _tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MissionType) String() string {
	var builder strings.Builder
	builder.WriteString("MissionType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mt.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mt.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mt.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mt.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", mt.Type))
	builder.WriteString(", ")
	builder.WriteString("gpu=")
	builder.WriteString(fmt.Sprintf("%v", mt.Gpu))
	builder.WriteString(", ")
	builder.WriteString("cep=")
	builder.WriteString(fmt.Sprintf("%v", mt.Cep))
	builder.WriteString(", ")
	builder.WriteString("is_time=")
	builder.WriteString(fmt.Sprintf("%v", mt.IsTime))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", mt.Category))
	builder.WriteByte(')')
	return builder.String()
}

// MissionTypes is a parsable slice of MissionType.
type MissionTypes []*MissionType
