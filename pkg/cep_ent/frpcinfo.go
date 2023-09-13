// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpcinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpsinfo"
)

// FrpcInfo is the model entity for the FrpcInfo schema.
type FrpcInfo struct {
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
	// ini 文件客户端 tag
	Tag string `json:"tag"`
	// frpc 通讯类型
	Type string `json:"type"`
	// frpc 本地地址
	LocalIP string `json:"local_ip"`
	// frpc 本地要使用端口
	LocalPort int `json:"local_port"`
	// frpc 本地要使用端口对应的远程端口
	RemotePort int `json:"remote_port"`
	// 端口是否已经在使用
	IsUsing bool `json:"is_using"`
	// 外键 frps id
	FrpsID int64 `json:"frps_id"`
	// 外键设备 id
	DeviceID int64 `json:"device_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FrpcInfoQuery when eager-loading is set.
	Edges        FrpcInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FrpcInfoEdges holds the relations/edges for other nodes in the graph.
type FrpcInfoEdges struct {
	// FrpsInfo holds the value of the frps_info edge.
	FrpsInfo *FrpsInfo `json:"frps_info,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FrpsInfoOrErr returns the FrpsInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FrpcInfoEdges) FrpsInfoOrErr() (*FrpsInfo, error) {
	if e.loadedTypes[0] {
		if e.FrpsInfo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: frpsinfo.Label}
		}
		return e.FrpsInfo, nil
	}
	return nil, &NotLoadedError{edge: "frps_info"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FrpcInfoEdges) DeviceOrErr() (*Device, error) {
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
func (*FrpcInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case frpcinfo.FieldIsUsing:
			values[i] = new(sql.NullBool)
		case frpcinfo.FieldID, frpcinfo.FieldCreatedBy, frpcinfo.FieldUpdatedBy, frpcinfo.FieldLocalPort, frpcinfo.FieldRemotePort, frpcinfo.FieldFrpsID, frpcinfo.FieldDeviceID:
			values[i] = new(sql.NullInt64)
		case frpcinfo.FieldTag, frpcinfo.FieldType, frpcinfo.FieldLocalIP:
			values[i] = new(sql.NullString)
		case frpcinfo.FieldCreatedAt, frpcinfo.FieldUpdatedAt, frpcinfo.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FrpcInfo fields.
func (fi *FrpcInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case frpcinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fi.ID = int64(value.Int64)
		case frpcinfo.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				fi.CreatedBy = value.Int64
			}
		case frpcinfo.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				fi.UpdatedBy = value.Int64
			}
		case frpcinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fi.CreatedAt = value.Time
			}
		case frpcinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fi.UpdatedAt = value.Time
			}
		case frpcinfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fi.DeletedAt = value.Time
			}
		case frpcinfo.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				fi.Tag = value.String
			}
		case frpcinfo.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				fi.Type = value.String
			}
		case frpcinfo.FieldLocalIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field local_ip", values[i])
			} else if value.Valid {
				fi.LocalIP = value.String
			}
		case frpcinfo.FieldLocalPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field local_port", values[i])
			} else if value.Valid {
				fi.LocalPort = int(value.Int64)
			}
		case frpcinfo.FieldRemotePort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remote_port", values[i])
			} else if value.Valid {
				fi.RemotePort = int(value.Int64)
			}
		case frpcinfo.FieldIsUsing:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_using", values[i])
			} else if value.Valid {
				fi.IsUsing = value.Bool
			}
		case frpcinfo.FieldFrpsID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field frps_id", values[i])
			} else if value.Valid {
				fi.FrpsID = value.Int64
			}
		case frpcinfo.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				fi.DeviceID = value.Int64
			}
		default:
			fi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FrpcInfo.
// This includes values selected through modifiers, order, etc.
func (fi *FrpcInfo) Value(name string) (ent.Value, error) {
	return fi.selectValues.Get(name)
}

// QueryFrpsInfo queries the "frps_info" edge of the FrpcInfo entity.
func (fi *FrpcInfo) QueryFrpsInfo() *FrpsInfoQuery {
	return NewFrpcInfoClient(fi.config).QueryFrpsInfo(fi)
}

// QueryDevice queries the "device" edge of the FrpcInfo entity.
func (fi *FrpcInfo) QueryDevice() *DeviceQuery {
	return NewFrpcInfoClient(fi.config).QueryDevice(fi)
}

// Update returns a builder for updating this FrpcInfo.
// Note that you need to call FrpcInfo.Unwrap() before calling this method if this FrpcInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (fi *FrpcInfo) Update() *FrpcInfoUpdateOne {
	return NewFrpcInfoClient(fi.config).UpdateOne(fi)
}

// Unwrap unwraps the FrpcInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fi *FrpcInfo) Unwrap() *FrpcInfo {
	_tx, ok := fi.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: FrpcInfo is not a transactional entity")
	}
	fi.config.driver = _tx.drv
	return fi
}

// String implements the fmt.Stringer.
func (fi *FrpcInfo) String() string {
	var builder strings.Builder
	builder.WriteString("FrpcInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fi.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", fi.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", fi.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fi.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(fi.Tag)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fi.Type)
	builder.WriteString(", ")
	builder.WriteString("local_ip=")
	builder.WriteString(fi.LocalIP)
	builder.WriteString(", ")
	builder.WriteString("local_port=")
	builder.WriteString(fmt.Sprintf("%v", fi.LocalPort))
	builder.WriteString(", ")
	builder.WriteString("remote_port=")
	builder.WriteString(fmt.Sprintf("%v", fi.RemotePort))
	builder.WriteString(", ")
	builder.WriteString("is_using=")
	builder.WriteString(fmt.Sprintf("%v", fi.IsUsing))
	builder.WriteString(", ")
	builder.WriteString("frps_id=")
	builder.WriteString(fmt.Sprintf("%v", fi.FrpsID))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", fi.DeviceID))
	builder.WriteByte(')')
	return builder.String()
}

// FrpcInfos is a parsable slice of FrpcInfo.
type FrpcInfos []*FrpcInfo