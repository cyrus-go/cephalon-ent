// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/profitsetting"
	"cephalon-ent/pkg/cep_ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProfitSetting is the model entity for the ProfitSetting schema.
type ProfitSetting struct {
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
	UserID int64 `json:"user_id"`
	// 分润比例
	Ratio int64 `json:"ratio"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfitSettingQuery when eager-loading is set.
	Edges        ProfitSettingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProfitSettingEdges holds the relations/edges for other nodes in the graph.
type ProfitSettingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfitSettingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProfitSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case profitsetting.FieldID, profitsetting.FieldCreatedBy, profitsetting.FieldUpdatedBy, profitsetting.FieldUserID, profitsetting.FieldRatio:
			values[i] = new(sql.NullInt64)
		case profitsetting.FieldCreatedAt, profitsetting.FieldUpdatedAt, profitsetting.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProfitSetting fields.
func (ps *ProfitSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profitsetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ps.ID = int64(value.Int64)
		case profitsetting.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ps.CreatedBy = value.Int64
			}
		case profitsetting.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ps.UpdatedBy = value.Int64
			}
		case profitsetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ps.CreatedAt = value.Time
			}
		case profitsetting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ps.UpdatedAt = value.Time
			}
		case profitsetting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ps.DeletedAt = value.Time
			}
		case profitsetting.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ps.UserID = value.Int64
			}
		case profitsetting.FieldRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ratio", values[i])
			} else if value.Valid {
				ps.Ratio = value.Int64
			}
		default:
			ps.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProfitSetting.
// This includes values selected through modifiers, order, etc.
func (ps *ProfitSetting) Value(name string) (ent.Value, error) {
	return ps.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the ProfitSetting entity.
func (ps *ProfitSetting) QueryUser() *UserQuery {
	return NewProfitSettingClient(ps.config).QueryUser(ps)
}

// Update returns a builder for updating this ProfitSetting.
// Note that you need to call ProfitSetting.Unwrap() before calling this method if this ProfitSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *ProfitSetting) Update() *ProfitSettingUpdateOne {
	return NewProfitSettingClient(ps.config).UpdateOne(ps)
}

// Unwrap unwraps the ProfitSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *ProfitSetting) Unwrap() *ProfitSetting {
	_tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: ProfitSetting is not a transactional entity")
	}
	ps.config.driver = _tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *ProfitSetting) String() string {
	var builder strings.Builder
	builder.WriteString("ProfitSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ps.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ps.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ps.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ps.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ps.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ps.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ps.UserID))
	builder.WriteString(", ")
	builder.WriteString("ratio=")
	builder.WriteString(fmt.Sprintf("%v", ps.Ratio))
	builder.WriteByte(')')
	return builder.String()
}

// ProfitSettings is a parsable slice of ProfitSetting.
type ProfitSettings []*ProfitSetting
