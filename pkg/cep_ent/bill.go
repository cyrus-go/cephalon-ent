// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/bill"
	"cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"cephalon-ent/pkg/cep_ent/missionproduceorder"
	"cephalon-ent/pkg/cep_ent/platformwallet"
	"cephalon-ent/pkg/cep_ent/rechargeorder"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/wallet"
	"cephalon-ent/pkg/enums"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 记录额度账户的变动的主流水账单记录
type Bill struct {
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
	// 次级账单流水的类型，充值或者任务消耗或任务收益
	Type enums.BillType `json:"type"`
	// 是否增加余额，布尔值默认为否
	IsAdd bool `json:"is_add"`
	// 外键用户 id
	UserID int64 `json:"user_id"`
	// 账单序列号
	SerialNumber string `json:"serial_number"`
	// 外键钱包 id
	WalletID int64 `json:"wallet_id"`
	// 消费或收益多少 cep
	Cep int64 `json:"cep"`
	// 关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单
	ReasonID int64 `json:"reason_id"`
	// 消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费
	Status enums.BillStatus `json:"status"`
	// 营销流水 id
	MarketBillID int64 `json:"market_bill_id"`
	// 平台分润钱包 id
	PlatformWalletID int64 `json:"-"`
	// 平台收取多少本金余额
	PlatformCep int64 `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillQuery when eager-loading is set.
	Edges        BillEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BillEdges holds the relations/edges for other nodes in the graph.
type BillEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Wallet holds the value of the wallet edge.
	Wallet *Wallet `json:"wallet,omitempty"`
	// PlatformWallet holds the value of the platform_wallet edge.
	PlatformWallet *PlatformWallet `json:"platform_wallet,omitempty"`
	// RechargeOrder holds the value of the recharge_order edge.
	RechargeOrder *RechargeOrder `json:"recharge_order,omitempty"`
	// MissionConsumeOrder holds the value of the mission_consume_order edge.
	MissionConsumeOrder *MissionConsumeOrder `json:"mission_consume_order,omitempty"`
	// MissionProduceOrder holds the value of the mission_produce_order edge.
	MissionProduceOrder *MissionProduceOrder `json:"mission_produce_order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// WalletOrErr returns the Wallet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) WalletOrErr() (*Wallet, error) {
	if e.loadedTypes[1] {
		if e.Wallet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: wallet.Label}
		}
		return e.Wallet, nil
	}
	return nil, &NotLoadedError{edge: "wallet"}
}

// PlatformWalletOrErr returns the PlatformWallet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) PlatformWalletOrErr() (*PlatformWallet, error) {
	if e.loadedTypes[2] {
		if e.PlatformWallet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: platformwallet.Label}
		}
		return e.PlatformWallet, nil
	}
	return nil, &NotLoadedError{edge: "platform_wallet"}
}

// RechargeOrderOrErr returns the RechargeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) RechargeOrderOrErr() (*RechargeOrder, error) {
	if e.loadedTypes[3] {
		if e.RechargeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rechargeorder.Label}
		}
		return e.RechargeOrder, nil
	}
	return nil, &NotLoadedError{edge: "recharge_order"}
}

// MissionConsumeOrderOrErr returns the MissionConsumeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) MissionConsumeOrderOrErr() (*MissionConsumeOrder, error) {
	if e.loadedTypes[4] {
		if e.MissionConsumeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionconsumeorder.Label}
		}
		return e.MissionConsumeOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_order"}
}

// MissionProduceOrderOrErr returns the MissionProduceOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) MissionProduceOrderOrErr() (*MissionProduceOrder, error) {
	if e.loadedTypes[5] {
		if e.MissionProduceOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionproduceorder.Label}
		}
		return e.MissionProduceOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bill.FieldIsAdd:
			values[i] = new(sql.NullBool)
		case bill.FieldID, bill.FieldCreatedBy, bill.FieldUpdatedBy, bill.FieldUserID, bill.FieldWalletID, bill.FieldCep, bill.FieldReasonID, bill.FieldMarketBillID, bill.FieldPlatformWalletID, bill.FieldPlatformCep:
			values[i] = new(sql.NullInt64)
		case bill.FieldType, bill.FieldSerialNumber, bill.FieldStatus:
			values[i] = new(sql.NullString)
		case bill.FieldCreatedAt, bill.FieldUpdatedAt, bill.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bill fields.
func (b *Bill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case bill.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				b.CreatedBy = value.Int64
			}
		case bill.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				b.UpdatedBy = value.Int64
			}
		case bill.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bill.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case bill.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = value.Time
			}
		case bill.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				b.Type = enums.BillType(value.String)
			}
		case bill.FieldIsAdd:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_add", values[i])
			} else if value.Valid {
				b.IsAdd = value.Bool
			}
		case bill.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = value.Int64
			}
		case bill.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				b.SerialNumber = value.String
			}
		case bill.FieldWalletID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_id", values[i])
			} else if value.Valid {
				b.WalletID = value.Int64
			}
		case bill.FieldCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cep", values[i])
			} else if value.Valid {
				b.Cep = value.Int64
			}
		case bill.FieldReasonID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reason_id", values[i])
			} else if value.Valid {
				b.ReasonID = value.Int64
			}
		case bill.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = enums.BillStatus(value.String)
			}
		case bill.FieldMarketBillID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field market_bill_id", values[i])
			} else if value.Valid {
				b.MarketBillID = value.Int64
			}
		case bill.FieldPlatformWalletID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field platform_wallet_id", values[i])
			} else if value.Valid {
				b.PlatformWalletID = value.Int64
			}
		case bill.FieldPlatformCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field platform_cep", values[i])
			} else if value.Valid {
				b.PlatformCep = value.Int64
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bill.
// This includes values selected through modifiers, order, etc.
func (b *Bill) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Bill entity.
func (b *Bill) QueryUser() *UserQuery {
	return NewBillClient(b.config).QueryUser(b)
}

// QueryWallet queries the "wallet" edge of the Bill entity.
func (b *Bill) QueryWallet() *WalletQuery {
	return NewBillClient(b.config).QueryWallet(b)
}

// QueryPlatformWallet queries the "platform_wallet" edge of the Bill entity.
func (b *Bill) QueryPlatformWallet() *PlatformWalletQuery {
	return NewBillClient(b.config).QueryPlatformWallet(b)
}

// QueryRechargeOrder queries the "recharge_order" edge of the Bill entity.
func (b *Bill) QueryRechargeOrder() *RechargeOrderQuery {
	return NewBillClient(b.config).QueryRechargeOrder(b)
}

// QueryMissionConsumeOrder queries the "mission_consume_order" edge of the Bill entity.
func (b *Bill) QueryMissionConsumeOrder() *MissionConsumeOrderQuery {
	return NewBillClient(b.config).QueryMissionConsumeOrder(b)
}

// QueryMissionProduceOrder queries the "mission_produce_order" edge of the Bill entity.
func (b *Bill) QueryMissionProduceOrder() *MissionProduceOrderQuery {
	return NewBillClient(b.config).QueryMissionProduceOrder(b)
}

// Update returns a builder for updating this Bill.
// Note that you need to call Bill.Unwrap() before calling this method if this Bill
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bill) Update() *BillUpdateOne {
	return NewBillClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bill) Unwrap() *Bill {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Bill is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bill) String() string {
	var builder strings.Builder
	builder.WriteString("Bill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", b.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", b.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(b.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", b.Type))
	builder.WriteString(", ")
	builder.WriteString("is_add=")
	builder.WriteString(fmt.Sprintf("%v", b.IsAdd))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", b.UserID))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(b.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("wallet_id=")
	builder.WriteString(fmt.Sprintf("%v", b.WalletID))
	builder.WriteString(", ")
	builder.WriteString("cep=")
	builder.WriteString(fmt.Sprintf("%v", b.Cep))
	builder.WriteString(", ")
	builder.WriteString("reason_id=")
	builder.WriteString(fmt.Sprintf("%v", b.ReasonID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("market_bill_id=")
	builder.WriteString(fmt.Sprintf("%v", b.MarketBillID))
	builder.WriteString(", ")
	builder.WriteString("platform_wallet_id=")
	builder.WriteString(fmt.Sprintf("%v", b.PlatformWalletID))
	builder.WriteString(", ")
	builder.WriteString("platform_cep=")
	builder.WriteString(fmt.Sprintf("%v", b.PlatformCep))
	builder.WriteByte(')')
	return builder.String()
}

// Bills is a parsable slice of Bill.
type Bills []*Bill
