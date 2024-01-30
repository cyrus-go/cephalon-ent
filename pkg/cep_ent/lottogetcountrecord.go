// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottogetcountrecord"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// 用户获得抽奖次数记录表
type LottoGetCountRecord struct {
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
	// 外键：用户 ID
	UserID int64 `json:"user_id"`
	// 外键：抽奖活动 ID
	LottoID int64 `json:"lotto_id"`
	// 此次奖励的抽奖次数
	Count int64 `json:"count"`
	// 抽奖结果
	Type lottogetcountrecord.Type `json:"type"`
	// 充值金额，类型为充值时才有数据
	RechargeAmount int64 `json:"recharge_amount"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LottoGetCountRecordQuery when eager-loading is set.
	Edges        LottoGetCountRecordEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LottoGetCountRecordEdges holds the relations/edges for other nodes in the graph.
type LottoGetCountRecordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Lotto holds the value of the lotto edge.
	Lotto *Lotto `json:"lotto,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoGetCountRecordEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// LottoOrErr returns the Lotto value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LottoGetCountRecordEdges) LottoOrErr() (*Lotto, error) {
	if e.loadedTypes[1] {
		if e.Lotto == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lotto.Label}
		}
		return e.Lotto, nil
	}
	return nil, &NotLoadedError{edge: "lotto"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LottoGetCountRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lottogetcountrecord.FieldID, lottogetcountrecord.FieldCreatedBy, lottogetcountrecord.FieldUpdatedBy, lottogetcountrecord.FieldUserID, lottogetcountrecord.FieldLottoID, lottogetcountrecord.FieldCount, lottogetcountrecord.FieldRechargeAmount:
			values[i] = new(sql.NullInt64)
		case lottogetcountrecord.FieldType:
			values[i] = new(sql.NullString)
		case lottogetcountrecord.FieldCreatedAt, lottogetcountrecord.FieldUpdatedAt, lottogetcountrecord.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LottoGetCountRecord fields.
func (lgcr *LottoGetCountRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lottogetcountrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lgcr.ID = int64(value.Int64)
		case lottogetcountrecord.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				lgcr.CreatedBy = value.Int64
			}
		case lottogetcountrecord.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				lgcr.UpdatedBy = value.Int64
			}
		case lottogetcountrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lgcr.CreatedAt = value.Time
			}
		case lottogetcountrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lgcr.UpdatedAt = value.Time
			}
		case lottogetcountrecord.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				lgcr.DeletedAt = value.Time
			}
		case lottogetcountrecord.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				lgcr.UserID = value.Int64
			}
		case lottogetcountrecord.FieldLottoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lotto_id", values[i])
			} else if value.Valid {
				lgcr.LottoID = value.Int64
			}
		case lottogetcountrecord.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				lgcr.Count = value.Int64
			}
		case lottogetcountrecord.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				lgcr.Type = lottogetcountrecord.Type(value.String)
			}
		case lottogetcountrecord.FieldRechargeAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recharge_amount", values[i])
			} else if value.Valid {
				lgcr.RechargeAmount = value.Int64
			}
		default:
			lgcr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LottoGetCountRecord.
// This includes values selected through modifiers, order, etc.
func (lgcr *LottoGetCountRecord) Value(name string) (ent.Value, error) {
	return lgcr.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the LottoGetCountRecord entity.
func (lgcr *LottoGetCountRecord) QueryUser() *UserQuery {
	return NewLottoGetCountRecordClient(lgcr.config).QueryUser(lgcr)
}

// QueryLotto queries the "lotto" edge of the LottoGetCountRecord entity.
func (lgcr *LottoGetCountRecord) QueryLotto() *LottoQuery {
	return NewLottoGetCountRecordClient(lgcr.config).QueryLotto(lgcr)
}

// Update returns a builder for updating this LottoGetCountRecord.
// Note that you need to call LottoGetCountRecord.Unwrap() before calling this method if this LottoGetCountRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (lgcr *LottoGetCountRecord) Update() *LottoGetCountRecordUpdateOne {
	return NewLottoGetCountRecordClient(lgcr.config).UpdateOne(lgcr)
}

// Unwrap unwraps the LottoGetCountRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lgcr *LottoGetCountRecord) Unwrap() *LottoGetCountRecord {
	_tx, ok := lgcr.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: LottoGetCountRecord is not a transactional entity")
	}
	lgcr.config.driver = _tx.drv
	return lgcr
}

// String implements the fmt.Stringer.
func (lgcr *LottoGetCountRecord) String() string {
	var builder strings.Builder
	builder.WriteString("LottoGetCountRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lgcr.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(lgcr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lgcr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(lgcr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.UserID))
	builder.WriteString(", ")
	builder.WriteString("lotto_id=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.LottoID))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.Count))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.Type))
	builder.WriteString(", ")
	builder.WriteString("recharge_amount=")
	builder.WriteString(fmt.Sprintf("%v", lgcr.RechargeAmount))
	builder.WriteByte(')')
	return builder.String()
}

// LottoGetCountRecords is a parsable slice of LottoGetCountRecord.
type LottoGetCountRecords []*LottoGetCountRecord
