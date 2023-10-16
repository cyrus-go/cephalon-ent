// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargecampaignrule"
)

// 充值活动的规则，死表，为特定充值赠送逻辑服务
type RechargeCampaignRule struct {
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
	// 充值范围下限
	LittleValue int64 `json:"little_value"`
	// 充值范围上限
	LargeValue int64 `json:"large_value"`
	// 赠送的比例
	GiftPercent  int64 `json:"gift_percent"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RechargeCampaignRule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rechargecampaignrule.FieldID, rechargecampaignrule.FieldCreatedBy, rechargecampaignrule.FieldUpdatedBy, rechargecampaignrule.FieldLittleValue, rechargecampaignrule.FieldLargeValue, rechargecampaignrule.FieldGiftPercent:
			values[i] = new(sql.NullInt64)
		case rechargecampaignrule.FieldCreatedAt, rechargecampaignrule.FieldUpdatedAt, rechargecampaignrule.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RechargeCampaignRule fields.
func (rcr *RechargeCampaignRule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rechargecampaignrule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rcr.ID = int64(value.Int64)
		case rechargecampaignrule.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				rcr.CreatedBy = value.Int64
			}
		case rechargecampaignrule.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				rcr.UpdatedBy = value.Int64
			}
		case rechargecampaignrule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rcr.CreatedAt = value.Time
			}
		case rechargecampaignrule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rcr.UpdatedAt = value.Time
			}
		case rechargecampaignrule.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rcr.DeletedAt = value.Time
			}
		case rechargecampaignrule.FieldLittleValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field little_value", values[i])
			} else if value.Valid {
				rcr.LittleValue = value.Int64
			}
		case rechargecampaignrule.FieldLargeValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field large_value", values[i])
			} else if value.Valid {
				rcr.LargeValue = value.Int64
			}
		case rechargecampaignrule.FieldGiftPercent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gift_percent", values[i])
			} else if value.Valid {
				rcr.GiftPercent = value.Int64
			}
		default:
			rcr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RechargeCampaignRule.
// This includes values selected through modifiers, order, etc.
func (rcr *RechargeCampaignRule) Value(name string) (ent.Value, error) {
	return rcr.selectValues.Get(name)
}

// Update returns a builder for updating this RechargeCampaignRule.
// Note that you need to call RechargeCampaignRule.Unwrap() before calling this method if this RechargeCampaignRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (rcr *RechargeCampaignRule) Update() *RechargeCampaignRuleUpdateOne {
	return NewRechargeCampaignRuleClient(rcr.config).UpdateOne(rcr)
}

// Unwrap unwraps the RechargeCampaignRule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rcr *RechargeCampaignRule) Unwrap() *RechargeCampaignRule {
	_tx, ok := rcr.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: RechargeCampaignRule is not a transactional entity")
	}
	rcr.config.driver = _tx.drv
	return rcr
}

// String implements the fmt.Stringer.
func (rcr *RechargeCampaignRule) String() string {
	var builder strings.Builder
	builder.WriteString("RechargeCampaignRule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rcr.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", rcr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", rcr.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rcr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rcr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(rcr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("little_value=")
	builder.WriteString(fmt.Sprintf("%v", rcr.LittleValue))
	builder.WriteString(", ")
	builder.WriteString("large_value=")
	builder.WriteString(fmt.Sprintf("%v", rcr.LargeValue))
	builder.WriteString(", ")
	builder.WriteString("gift_percent=")
	builder.WriteString(fmt.Sprintf("%v", rcr.GiftPercent))
	builder.WriteByte(')')
	return builder.String()
}

// RechargeCampaignRules is a parsable slice of RechargeCampaignRule.
type RechargeCampaignRules []*RechargeCampaignRule
