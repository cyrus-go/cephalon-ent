// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/missionbatch"
	"cephalon-ent/pkg/cep_ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MissionBatch is the model entity for the MissionBatch schema.
type MissionBatch struct {
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
	// 批次号码
	Number string `json:"number"`
	// 创建者
	UserID int64 `json:"user_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionBatchQuery when eager-loading is set.
	Edges        MissionBatchEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MissionBatchEdges holds the relations/edges for other nodes in the graph.
type MissionBatchEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// MissionConsumeOrders holds the value of the mission_consume_orders edge.
	MissionConsumeOrders []*MissionConsumeOrder `json:"mission_consume_orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionBatchEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MissionConsumeOrdersOrErr returns the MissionConsumeOrders value or an error if the edge
// was not loaded in eager-loading.
func (e MissionBatchEdges) MissionConsumeOrdersOrErr() ([]*MissionConsumeOrder, error) {
	if e.loadedTypes[1] {
		return e.MissionConsumeOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionBatch) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionbatch.FieldID, missionbatch.FieldCreatedBy, missionbatch.FieldUpdatedBy, missionbatch.FieldUserID:
			values[i] = new(sql.NullInt64)
		case missionbatch.FieldNumber:
			values[i] = new(sql.NullString)
		case missionbatch.FieldCreatedAt, missionbatch.FieldUpdatedAt, missionbatch.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionBatch fields.
func (mb *MissionBatch) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionbatch.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mb.ID = int64(value.Int64)
		case missionbatch.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mb.CreatedBy = value.Int64
			}
		case missionbatch.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mb.UpdatedBy = value.Int64
			}
		case missionbatch.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mb.CreatedAt = value.Time
			}
		case missionbatch.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mb.UpdatedAt = value.Time
			}
		case missionbatch.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mb.DeletedAt = value.Time
			}
		case missionbatch.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				mb.Number = value.String
			}
		case missionbatch.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				mb.UserID = value.Int64
			}
		default:
			mb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionBatch.
// This includes values selected through modifiers, order, etc.
func (mb *MissionBatch) Value(name string) (ent.Value, error) {
	return mb.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the MissionBatch entity.
func (mb *MissionBatch) QueryUser() *UserQuery {
	return NewMissionBatchClient(mb.config).QueryUser(mb)
}

// QueryMissionConsumeOrders queries the "mission_consume_orders" edge of the MissionBatch entity.
func (mb *MissionBatch) QueryMissionConsumeOrders() *MissionConsumeOrderQuery {
	return NewMissionBatchClient(mb.config).QueryMissionConsumeOrders(mb)
}

// Update returns a builder for updating this MissionBatch.
// Note that you need to call MissionBatch.Unwrap() before calling this method if this MissionBatch
// was returned from a transaction, and the transaction was committed or rolled back.
func (mb *MissionBatch) Update() *MissionBatchUpdateOne {
	return NewMissionBatchClient(mb.config).UpdateOne(mb)
}

// Unwrap unwraps the MissionBatch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mb *MissionBatch) Unwrap() *MissionBatch {
	_tx, ok := mb.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionBatch is not a transactional entity")
	}
	mb.config.driver = _tx.drv
	return mb
}

// String implements the fmt.Stringer.
func (mb *MissionBatch) String() string {
	var builder strings.Builder
	builder.WriteString("MissionBatch(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mb.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mb.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mb.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mb.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(mb.Number)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", mb.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// MissionBatches is a parsable slice of MissionBatch.
type MissionBatches []*MissionBatch
