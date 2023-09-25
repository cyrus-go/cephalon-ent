// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduction"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务生产订单，记录任务生产情况对应的订单业务，比如分润情况，分润比例等信息
type MissionProduceOrder struct {
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
	// 外键关联用户 id
	UserID int64 `json:"user_id"`
	// 任务 id，关联任务中枢的任务
	MissionID int64 `json:"mission_id"`
	// 任务执行情况 id
	MissionProductionID *int64 `json:"mission_production_id"`
	// 任务订单的状态，注意不强关联任务的状态
	Status enums.MissionOrderStatus `json:"status"`
	// 任务收益的本金 cep 量
	PureCep int64 `json:"pure_cep"`
	// 任务收益的赠送 cep 量
	GiftCep int64 `json:"gift_cep"`
	// 币种 id
	SymbolID int64 `json:"symbol_id""`
	// 任务订单分润
	Amount int64 `json:"amount"`
	// 任务类型，计时或者次数任务
	Type enums.MissionType `json:"type"`
	// 是否为计时类型任务
	IsTime bool `json:"is_time"`
	// 生产者接该任务用的设备 id
	DeviceID int64 `json:"device_id"`
	// 订单序列号
	SerialNumber string `json:"serial_number"`
	// 外键任务消费订单，一个任务消费订单可能会对应多个任务生产订单
	MissionConsumeOrderID int64 `json:"mission_consume_order_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionProduceOrderQuery when eager-loading is set.
	Edges                          MissionProduceOrderEdges `json:"edges"`
	mission_mission_produce_orders *int64
	selectValues                   sql.SelectValues
}

// MissionProduceOrderEdges holds the relations/edges for other nodes in the graph.
type MissionProduceOrderEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// EarnBills holds the value of the earn_bills edge.
	EarnBills []*EarnBill `json:"earn_bills,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// MissionConsumeOrder holds the value of the mission_consume_order edge.
	MissionConsumeOrder *MissionConsumeOrder `json:"mission_consume_order,omitempty"`
	// MissionProduction holds the value of the mission_production edge.
	MissionProduction *MissionProduction `json:"mission_production,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProduceOrderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// EarnBillsOrErr returns the EarnBills value or an error if the edge
// was not loaded in eager-loading.
func (e MissionProduceOrderEdges) EarnBillsOrErr() ([]*EarnBill, error) {
	if e.loadedTypes[1] {
		return e.EarnBills, nil
	}
	return nil, &NotLoadedError{edge: "earn_bills"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProduceOrderEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[2] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// MissionConsumeOrderOrErr returns the MissionConsumeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProduceOrderEdges) MissionConsumeOrderOrErr() (*MissionConsumeOrder, error) {
	if e.loadedTypes[3] {
		if e.MissionConsumeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionconsumeorder.Label}
		}
		return e.MissionConsumeOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_order"}
}

// MissionProductionOrErr returns the MissionProduction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionProduceOrderEdges) MissionProductionOrErr() (*MissionProduction, error) {
	if e.loadedTypes[4] {
		if e.MissionProduction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionproduction.Label}
		}
		return e.MissionProduction, nil
	}
	return nil, &NotLoadedError{edge: "mission_production"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MissionProduceOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case missionproduceorder.FieldIsTime:
			values[i] = new(sql.NullBool)
		case missionproduceorder.FieldID, missionproduceorder.FieldCreatedBy, missionproduceorder.FieldUpdatedBy, missionproduceorder.FieldUserID, missionproduceorder.FieldMissionID, missionproduceorder.FieldMissionProductionID, missionproduceorder.FieldPureCep, missionproduceorder.FieldGiftCep, missionproduceorder.FieldSymbolID, missionproduceorder.FieldAmount, missionproduceorder.FieldDeviceID, missionproduceorder.FieldMissionConsumeOrderID:
			values[i] = new(sql.NullInt64)
		case missionproduceorder.FieldStatus, missionproduceorder.FieldType, missionproduceorder.FieldSerialNumber:
			values[i] = new(sql.NullString)
		case missionproduceorder.FieldCreatedAt, missionproduceorder.FieldUpdatedAt, missionproduceorder.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case missionproduceorder.ForeignKeys[0]: // mission_mission_produce_orders
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MissionProduceOrder fields.
func (mpo *MissionProduceOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case missionproduceorder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mpo.ID = int64(value.Int64)
		case missionproduceorder.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mpo.CreatedBy = value.Int64
			}
		case missionproduceorder.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mpo.UpdatedBy = value.Int64
			}
		case missionproduceorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mpo.CreatedAt = value.Time
			}
		case missionproduceorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mpo.UpdatedAt = value.Time
			}
		case missionproduceorder.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mpo.DeletedAt = value.Time
			}
		case missionproduceorder.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				mpo.UserID = value.Int64
			}
		case missionproduceorder.FieldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_id", values[i])
			} else if value.Valid {
				mpo.MissionID = value.Int64
			}
		case missionproduceorder.FieldMissionProductionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_production_id", values[i])
			} else if value.Valid {
				mpo.MissionProductionID = new(int64)
				*mpo.MissionProductionID = value.Int64
			}
		case missionproduceorder.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mpo.Status = enums.MissionOrderStatus(value.String)
			}
		case missionproduceorder.FieldPureCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pure_cep", values[i])
			} else if value.Valid {
				mpo.PureCep = value.Int64
			}
		case missionproduceorder.FieldGiftCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gift_cep", values[i])
			} else if value.Valid {
				mpo.GiftCep = value.Int64
			}
		case missionproduceorder.FieldSymbolID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field symbol_id", values[i])
			} else if value.Valid {
				mpo.SymbolID = value.Int64
			}
		case missionproduceorder.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				mpo.Amount = value.Int64
			}
		case missionproduceorder.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mpo.Type = enums.MissionType(value.String)
			}
		case missionproduceorder.FieldIsTime:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_time", values[i])
			} else if value.Valid {
				mpo.IsTime = value.Bool
			}
		case missionproduceorder.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				mpo.DeviceID = value.Int64
			}
		case missionproduceorder.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				mpo.SerialNumber = value.String
			}
		case missionproduceorder.FieldMissionConsumeOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_consume_order_id", values[i])
			} else if value.Valid {
				mpo.MissionConsumeOrderID = value.Int64
			}
		case missionproduceorder.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mission_mission_produce_orders", value)
			} else if value.Valid {
				mpo.mission_mission_produce_orders = new(int64)
				*mpo.mission_mission_produce_orders = int64(value.Int64)
			}
		default:
			mpo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MissionProduceOrder.
// This includes values selected through modifiers, order, etc.
func (mpo *MissionProduceOrder) Value(name string) (ent.Value, error) {
	return mpo.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the MissionProduceOrder entity.
func (mpo *MissionProduceOrder) QueryUser() *UserQuery {
	return NewMissionProduceOrderClient(mpo.config).QueryUser(mpo)
}

// QueryEarnBills queries the "earn_bills" edge of the MissionProduceOrder entity.
func (mpo *MissionProduceOrder) QueryEarnBills() *EarnBillQuery {
	return NewMissionProduceOrderClient(mpo.config).QueryEarnBills(mpo)
}

// QueryDevice queries the "device" edge of the MissionProduceOrder entity.
func (mpo *MissionProduceOrder) QueryDevice() *DeviceQuery {
	return NewMissionProduceOrderClient(mpo.config).QueryDevice(mpo)
}

// QueryMissionConsumeOrder queries the "mission_consume_order" edge of the MissionProduceOrder entity.
func (mpo *MissionProduceOrder) QueryMissionConsumeOrder() *MissionConsumeOrderQuery {
	return NewMissionProduceOrderClient(mpo.config).QueryMissionConsumeOrder(mpo)
}

// QueryMissionProduction queries the "mission_production" edge of the MissionProduceOrder entity.
func (mpo *MissionProduceOrder) QueryMissionProduction() *MissionProductionQuery {
	return NewMissionProduceOrderClient(mpo.config).QueryMissionProduction(mpo)
}

// Update returns a builder for updating this MissionProduceOrder.
// Note that you need to call MissionProduceOrder.Unwrap() before calling this method if this MissionProduceOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (mpo *MissionProduceOrder) Update() *MissionProduceOrderUpdateOne {
	return NewMissionProduceOrderClient(mpo.config).UpdateOne(mpo)
}

// Unwrap unwraps the MissionProduceOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mpo *MissionProduceOrder) Unwrap() *MissionProduceOrder {
	_tx, ok := mpo.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: MissionProduceOrder is not a transactional entity")
	}
	mpo.config.driver = _tx.drv
	return mpo
}

// String implements the fmt.Stringer.
func (mpo *MissionProduceOrder) String() string {
	var builder strings.Builder
	builder.WriteString("MissionProduceOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mpo.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mpo.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mpo.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mpo.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mpo.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mpo.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", mpo.UserID))
	builder.WriteString(", ")
	builder.WriteString("mission_id=")
	builder.WriteString(fmt.Sprintf("%v", mpo.MissionID))
	builder.WriteString(", ")
	if v := mpo.MissionProductionID; v != nil {
		builder.WriteString("mission_production_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mpo.Status))
	builder.WriteString(", ")
	builder.WriteString("pure_cep=")
	builder.WriteString(fmt.Sprintf("%v", mpo.PureCep))
	builder.WriteString(", ")
	builder.WriteString("gift_cep=")
	builder.WriteString(fmt.Sprintf("%v", mpo.GiftCep))
	builder.WriteString(", ")
	builder.WriteString("symbol_id=")
	builder.WriteString(fmt.Sprintf("%v", mpo.SymbolID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", mpo.Amount))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", mpo.Type))
	builder.WriteString(", ")
	builder.WriteString("is_time=")
	builder.WriteString(fmt.Sprintf("%v", mpo.IsTime))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", mpo.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(mpo.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("mission_consume_order_id=")
	builder.WriteString(fmt.Sprintf("%v", mpo.MissionConsumeOrderID))
	builder.WriteByte(')')
	return builder.String()
}

// MissionProduceOrders is a parsable slice of MissionProduceOrder.
type MissionProduceOrders []*MissionProduceOrder
