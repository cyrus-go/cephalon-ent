// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/userdevice"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserDevice is the model entity for the UserDevice schema.
type UserDevice struct {
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
	// 外键设备 id
	DeviceID int64 `json:"device_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserDeviceQuery when eager-loading is set.
	Edges        UserDeviceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserDeviceEdges holds the relations/edges for other nodes in the graph.
type UserDeviceEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserDeviceEdges) UserOrErr() (*User, error) {
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
func (e UserDeviceEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[1] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserDevice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userdevice.FieldID, userdevice.FieldCreatedBy, userdevice.FieldUpdatedBy, userdevice.FieldUserID, userdevice.FieldDeviceID:
			values[i] = new(sql.NullInt64)
		case userdevice.FieldCreatedAt, userdevice.FieldUpdatedAt, userdevice.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserDevice fields.
func (ud *UserDevice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userdevice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ud.ID = int64(value.Int64)
		case userdevice.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ud.CreatedBy = value.Int64
			}
		case userdevice.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ud.UpdatedBy = value.Int64
			}
		case userdevice.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ud.CreatedAt = value.Time
			}
		case userdevice.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ud.UpdatedAt = value.Time
			}
		case userdevice.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ud.DeletedAt = value.Time
			}
		case userdevice.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ud.UserID = value.Int64
			}
		case userdevice.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				ud.DeviceID = value.Int64
			}
		default:
			ud.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserDevice.
// This includes values selected through modifiers, order, etc.
func (ud *UserDevice) Value(name string) (ent.Value, error) {
	return ud.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserDevice entity.
func (ud *UserDevice) QueryUser() *UserQuery {
	return NewUserDeviceClient(ud.config).QueryUser(ud)
}

// QueryDevice queries the "device" edge of the UserDevice entity.
func (ud *UserDevice) QueryDevice() *DeviceQuery {
	return NewUserDeviceClient(ud.config).QueryDevice(ud)
}

// Update returns a builder for updating this UserDevice.
// Note that you need to call UserDevice.Unwrap() before calling this method if this UserDevice
// was returned from a transaction, and the transaction was committed or rolled back.
func (ud *UserDevice) Update() *UserDeviceUpdateOne {
	return NewUserDeviceClient(ud.config).UpdateOne(ud)
}

// Unwrap unwraps the UserDevice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ud *UserDevice) Unwrap() *UserDevice {
	_tx, ok := ud.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: UserDevice is not a transactional entity")
	}
	ud.config.driver = _tx.drv
	return ud
}

// String implements the fmt.Stringer.
func (ud *UserDevice) String() string {
	var builder strings.Builder
	builder.WriteString("UserDevice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ud.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ud.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ud.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ud.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ud.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ud.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.UserID))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.DeviceID))
	builder.WriteByte(')')
	return builder.String()
}

// UserDevices is a parsable slice of UserDevice.
type UserDevices []*UserDevice
