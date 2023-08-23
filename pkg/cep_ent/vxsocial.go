// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/vxsocial"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VXSocial is the model entity for the VXSocial schema.
type VXSocial struct {
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
	// 微信应用 id
	AppID string `json:"app_id"`
	// 微信身份源的 open_id
	OpenID string `json:"open_id"`
	// 微信身份源的 union_id
	UnionID string `json:"union_id"`
	// 账户的权限级别，一般为 base
	Scope vxsocial.Scope `json:"scope"`
	// 小程序专用的会话密钥，不可以返回给前端
	SessionKey string `json:"-"`
	// 微信能力访问凭证
	AccessToken string `json:"access_token"`
	// 刷新微信凭证的刷新凭证
	RefreshToken string `json:"refresh_token"`
	// 外键用户 id
	UserID int64 `json:"user_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VXSocialQuery when eager-loading is set.
	Edges        VXSocialEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VXSocialEdges holds the relations/edges for other nodes in the graph.
type VXSocialEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// RechargeOrders holds the value of the recharge_orders edge.
	RechargeOrders []*RechargeOrder `json:"recharge_orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VXSocialEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RechargeOrdersOrErr returns the RechargeOrders value or an error if the edge
// was not loaded in eager-loading.
func (e VXSocialEdges) RechargeOrdersOrErr() ([]*RechargeOrder, error) {
	if e.loadedTypes[1] {
		return e.RechargeOrders, nil
	}
	return nil, &NotLoadedError{edge: "recharge_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VXSocial) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vxsocial.FieldID, vxsocial.FieldCreatedBy, vxsocial.FieldUpdatedBy, vxsocial.FieldUserID:
			values[i] = new(sql.NullInt64)
		case vxsocial.FieldAppID, vxsocial.FieldOpenID, vxsocial.FieldUnionID, vxsocial.FieldScope, vxsocial.FieldSessionKey, vxsocial.FieldAccessToken, vxsocial.FieldRefreshToken:
			values[i] = new(sql.NullString)
		case vxsocial.FieldCreatedAt, vxsocial.FieldUpdatedAt, vxsocial.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VXSocial fields.
func (vs *VXSocial) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vxsocial.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vs.ID = int64(value.Int64)
		case vxsocial.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				vs.CreatedBy = value.Int64
			}
		case vxsocial.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				vs.UpdatedBy = value.Int64
			}
		case vxsocial.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vs.CreatedAt = value.Time
			}
		case vxsocial.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vs.UpdatedAt = value.Time
			}
		case vxsocial.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				vs.DeletedAt = value.Time
			}
		case vxsocial.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				vs.AppID = value.String
			}
		case vxsocial.FieldOpenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field open_id", values[i])
			} else if value.Valid {
				vs.OpenID = value.String
			}
		case vxsocial.FieldUnionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field union_id", values[i])
			} else if value.Valid {
				vs.UnionID = value.String
			}
		case vxsocial.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				vs.Scope = vxsocial.Scope(value.String)
			}
		case vxsocial.FieldSessionKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_key", values[i])
			} else if value.Valid {
				vs.SessionKey = value.String
			}
		case vxsocial.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				vs.AccessToken = value.String
			}
		case vxsocial.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				vs.RefreshToken = value.String
			}
		case vxsocial.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				vs.UserID = value.Int64
			}
		default:
			vs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VXSocial.
// This includes values selected through modifiers, order, etc.
func (vs *VXSocial) Value(name string) (ent.Value, error) {
	return vs.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the VXSocial entity.
func (vs *VXSocial) QueryUser() *UserQuery {
	return NewVXSocialClient(vs.config).QueryUser(vs)
}

// QueryRechargeOrders queries the "recharge_orders" edge of the VXSocial entity.
func (vs *VXSocial) QueryRechargeOrders() *RechargeOrderQuery {
	return NewVXSocialClient(vs.config).QueryRechargeOrders(vs)
}

// Update returns a builder for updating this VXSocial.
// Note that you need to call VXSocial.Unwrap() before calling this method if this VXSocial
// was returned from a transaction, and the transaction was committed or rolled back.
func (vs *VXSocial) Update() *VXSocialUpdateOne {
	return NewVXSocialClient(vs.config).UpdateOne(vs)
}

// Unwrap unwraps the VXSocial entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vs *VXSocial) Unwrap() *VXSocial {
	_tx, ok := vs.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: VXSocial is not a transactional entity")
	}
	vs.config.driver = _tx.drv
	return vs
}

// String implements the fmt.Stringer.
func (vs *VXSocial) String() string {
	var builder strings.Builder
	builder.WriteString("VXSocial(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vs.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", vs.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", vs.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(vs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(vs.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(vs.AppID)
	builder.WriteString(", ")
	builder.WriteString("open_id=")
	builder.WriteString(vs.OpenID)
	builder.WriteString(", ")
	builder.WriteString("union_id=")
	builder.WriteString(vs.UnionID)
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(fmt.Sprintf("%v", vs.Scope))
	builder.WriteString(", ")
	builder.WriteString("session_key=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(vs.AccessToken)
	builder.WriteString(", ")
	builder.WriteString("refresh_token=")
	builder.WriteString(vs.RefreshToken)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", vs.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// VXSocials is a parsable slice of VXSocial.
type VXSocials []*VXSocial
