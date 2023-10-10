// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costaccount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/profitaccount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// 用户表
type User struct {
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
	// 用户名
	Name string `json:"name"`
	// 头像
	JpgURL string `json:"jpg_url"`
	// 用户在任务中枢密钥对的键
	Key string `json:"key"`
	// 用户在任务中枢的密钥对的值
	Secret string `json:"-"`
	// 用户的手机号
	Phone string `json:"phone"`
	// 密码
	Password string `json:"-"`
	// 是否冻结
	IsFrozen bool `json:"is_frozen"`
	// 是否充值过
	IsRecharge bool `json:"is_recharge"`
	// 用户类型
	UserType user.UserType `json:"user_type"`
	// 邀请人用户 id
	ParentID int64 `json:"parent_id"`
	// 用户最新弹窗版本
	PopVersion string `json:"pop_version"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// VxAccounts holds the value of the vx_accounts edge.
	VxAccounts []*VXAccount `json:"vx_accounts,omitempty"`
	// Collects holds the value of the collects edge.
	Collects []*Collect `json:"collects,omitempty"`
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// ProfitSettings holds the value of the profit_settings edge.
	ProfitSettings []*ProfitSetting `json:"profit_settings,omitempty"`
	// CostAccount holds the value of the cost_account edge.
	CostAccount *CostAccount `json:"cost_account,omitempty"`
	// ProfitAccount holds the value of the profit_account edge.
	ProfitAccount *ProfitAccount `json:"profit_account,omitempty"`
	// CostBills holds the value of the cost_bills edge.
	CostBills []*CostBill `json:"cost_bills,omitempty"`
	// EarnBills holds the value of the earn_bills edge.
	EarnBills []*EarnBill `json:"earn_bills,omitempty"`
	// MissionConsumeOrders holds the value of the mission_consume_orders edge.
	MissionConsumeOrders []*MissionConsumeOrder `json:"mission_consume_orders,omitempty"`
	// MissionProduceOrders holds the value of the mission_produce_orders edge.
	MissionProduceOrders []*MissionProduceOrder `json:"mission_produce_orders,omitempty"`
	// RechargeOrders holds the value of the recharge_orders edge.
	RechargeOrders []*RechargeOrder `json:"recharge_orders,omitempty"`
	// VxSocials holds the value of the vx_socials edge.
	VxSocials []*VXSocial `json:"vx_socials,omitempty"`
	// MissionBatches holds the value of the mission_batches edge.
	MissionBatches []*MissionBatch `json:"mission_batches,omitempty"`
	// UserDevices holds the value of the user_devices edge.
	UserDevices []*UserDevice `json:"user_devices,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *User `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*User `json:"children,omitempty"`
	// Invites holds the value of the invites edge.
	Invites []*Invite `json:"invites,omitempty"`
	// CampaignOrders holds the value of the campaign_orders edge.
	CampaignOrders []*CampaignOrder `json:"campaign_orders,omitempty"`
	// Wallets holds the value of the wallets edge.
	Wallets []*Wallet `json:"wallets,omitempty"`
	// IncomeBills holds the value of the income_bills edge.
	IncomeBills []*Bill `json:"income_bills,omitempty"`
	// OutcomeBills holds the value of the outcome_bills edge.
	OutcomeBills []*Bill `json:"outcome_bills,omitempty"`
	// MissionProductions holds the value of the mission_productions edge.
	MissionProductions []*MissionProduction `json:"mission_productions,omitempty"`
	// Missions holds the value of the missions edge.
	Missions []*Mission `json:"missions,omitempty"`
	// IncomeTransferOrders holds the value of the income_transfer_orders edge.
	IncomeTransferOrders []*TransferOrder `json:"income_transfer_orders,omitempty"`
	// OutcomeTransferOrders holds the value of the outcome_transfer_orders edge.
	OutcomeTransferOrders []*TransferOrder `json:"outcome_transfer_orders,omitempty"`
	// ConsumeMissionOrders holds the value of the consume_mission_orders edge.
	ConsumeMissionOrders []*MissionOrder `json:"consume_mission_orders,omitempty"`
	// ProduceMissionOrders holds the value of the produce_mission_orders edge.
	ProduceMissionOrders []*MissionOrder `json:"produce_mission_orders,omitempty"`
	// LoginRecords holds the value of the login_records edge.
	LoginRecords []*LoginRecord `json:"login_records,omitempty"`
	// RenewalAgreements holds the value of the renewal_agreements edge.
	RenewalAgreements []*RenewalAgreement `json:"renewal_agreements,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [29]bool
}

// VxAccountsOrErr returns the VxAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VxAccountsOrErr() ([]*VXAccount, error) {
	if e.loadedTypes[0] {
		return e.VxAccounts, nil
	}
	return nil, &NotLoadedError{edge: "vx_accounts"}
}

// CollectsOrErr returns the Collects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CollectsOrErr() ([]*Collect, error) {
	if e.loadedTypes[1] {
		return e.Collects, nil
	}
	return nil, &NotLoadedError{edge: "collects"}
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[2] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// ProfitSettingsOrErr returns the ProfitSettings value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProfitSettingsOrErr() ([]*ProfitSetting, error) {
	if e.loadedTypes[3] {
		return e.ProfitSettings, nil
	}
	return nil, &NotLoadedError{edge: "profit_settings"}
}

// CostAccountOrErr returns the CostAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CostAccountOrErr() (*CostAccount, error) {
	if e.loadedTypes[4] {
		if e.CostAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: costaccount.Label}
		}
		return e.CostAccount, nil
	}
	return nil, &NotLoadedError{edge: "cost_account"}
}

// ProfitAccountOrErr returns the ProfitAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfitAccountOrErr() (*ProfitAccount, error) {
	if e.loadedTypes[5] {
		if e.ProfitAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: profitaccount.Label}
		}
		return e.ProfitAccount, nil
	}
	return nil, &NotLoadedError{edge: "profit_account"}
}

// CostBillsOrErr returns the CostBills value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CostBillsOrErr() ([]*CostBill, error) {
	if e.loadedTypes[6] {
		return e.CostBills, nil
	}
	return nil, &NotLoadedError{edge: "cost_bills"}
}

// EarnBillsOrErr returns the EarnBills value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EarnBillsOrErr() ([]*EarnBill, error) {
	if e.loadedTypes[7] {
		return e.EarnBills, nil
	}
	return nil, &NotLoadedError{edge: "earn_bills"}
}

// MissionConsumeOrdersOrErr returns the MissionConsumeOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MissionConsumeOrdersOrErr() ([]*MissionConsumeOrder, error) {
	if e.loadedTypes[8] {
		return e.MissionConsumeOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_orders"}
}

// MissionProduceOrdersOrErr returns the MissionProduceOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MissionProduceOrdersOrErr() ([]*MissionProduceOrder, error) {
	if e.loadedTypes[9] {
		return e.MissionProduceOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_orders"}
}

// RechargeOrdersOrErr returns the RechargeOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RechargeOrdersOrErr() ([]*RechargeOrder, error) {
	if e.loadedTypes[10] {
		return e.RechargeOrders, nil
	}
	return nil, &NotLoadedError{edge: "recharge_orders"}
}

// VxSocialsOrErr returns the VxSocials value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VxSocialsOrErr() ([]*VXSocial, error) {
	if e.loadedTypes[11] {
		return e.VxSocials, nil
	}
	return nil, &NotLoadedError{edge: "vx_socials"}
}

// MissionBatchesOrErr returns the MissionBatches value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MissionBatchesOrErr() ([]*MissionBatch, error) {
	if e.loadedTypes[12] {
		return e.MissionBatches, nil
	}
	return nil, &NotLoadedError{edge: "mission_batches"}
}

// UserDevicesOrErr returns the UserDevices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserDevicesOrErr() ([]*UserDevice, error) {
	if e.loadedTypes[13] {
		return e.UserDevices, nil
	}
	return nil, &NotLoadedError{edge: "user_devices"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.loadedTypes[14] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[15] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// InvitesOrErr returns the Invites value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InvitesOrErr() ([]*Invite, error) {
	if e.loadedTypes[16] {
		return e.Invites, nil
	}
	return nil, &NotLoadedError{edge: "invites"}
}

// CampaignOrdersOrErr returns the CampaignOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CampaignOrdersOrErr() ([]*CampaignOrder, error) {
	if e.loadedTypes[17] {
		return e.CampaignOrders, nil
	}
	return nil, &NotLoadedError{edge: "campaign_orders"}
}

// WalletsOrErr returns the Wallets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WalletsOrErr() ([]*Wallet, error) {
	if e.loadedTypes[18] {
		return e.Wallets, nil
	}
	return nil, &NotLoadedError{edge: "wallets"}
}

// IncomeBillsOrErr returns the IncomeBills value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IncomeBillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[19] {
		return e.IncomeBills, nil
	}
	return nil, &NotLoadedError{edge: "income_bills"}
}

// OutcomeBillsOrErr returns the OutcomeBills value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OutcomeBillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[20] {
		return e.OutcomeBills, nil
	}
	return nil, &NotLoadedError{edge: "outcome_bills"}
}

// MissionProductionsOrErr returns the MissionProductions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MissionProductionsOrErr() ([]*MissionProduction, error) {
	if e.loadedTypes[21] {
		return e.MissionProductions, nil
	}
	return nil, &NotLoadedError{edge: "mission_productions"}
}

// MissionsOrErr returns the Missions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MissionsOrErr() ([]*Mission, error) {
	if e.loadedTypes[22] {
		return e.Missions, nil
	}
	return nil, &NotLoadedError{edge: "missions"}
}

// IncomeTransferOrdersOrErr returns the IncomeTransferOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IncomeTransferOrdersOrErr() ([]*TransferOrder, error) {
	if e.loadedTypes[23] {
		return e.IncomeTransferOrders, nil
	}
	return nil, &NotLoadedError{edge: "income_transfer_orders"}
}

// OutcomeTransferOrdersOrErr returns the OutcomeTransferOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OutcomeTransferOrdersOrErr() ([]*TransferOrder, error) {
	if e.loadedTypes[24] {
		return e.OutcomeTransferOrders, nil
	}
	return nil, &NotLoadedError{edge: "outcome_transfer_orders"}
}

// ConsumeMissionOrdersOrErr returns the ConsumeMissionOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ConsumeMissionOrdersOrErr() ([]*MissionOrder, error) {
	if e.loadedTypes[25] {
		return e.ConsumeMissionOrders, nil
	}
	return nil, &NotLoadedError{edge: "consume_mission_orders"}
}

// ProduceMissionOrdersOrErr returns the ProduceMissionOrders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProduceMissionOrdersOrErr() ([]*MissionOrder, error) {
	if e.loadedTypes[26] {
		return e.ProduceMissionOrders, nil
	}
	return nil, &NotLoadedError{edge: "produce_mission_orders"}
}

// LoginRecordsOrErr returns the LoginRecords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LoginRecordsOrErr() ([]*LoginRecord, error) {
	if e.loadedTypes[27] {
		return e.LoginRecords, nil
	}
	return nil, &NotLoadedError{edge: "login_records"}
}

// RenewalAgreementsOrErr returns the RenewalAgreements value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RenewalAgreementsOrErr() ([]*RenewalAgreement, error) {
	if e.loadedTypes[28] {
		return e.RenewalAgreements, nil
	}
	return nil, &NotLoadedError{edge: "renewal_agreements"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsFrozen, user.FieldIsRecharge:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy, user.FieldParentID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldJpgURL, user.FieldKey, user.FieldSecret, user.FieldPhone, user.FieldPassword, user.FieldUserType, user.FieldPopVersion:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = value.Int64
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = value.Int64
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldJpgURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jpg_url", values[i])
			} else if value.Valid {
				u.JpgURL = value.String
			}
		case user.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				u.Key = value.String
			}
		case user.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				u.Secret = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldIsFrozen:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_frozen", values[i])
			} else if value.Valid {
				u.IsFrozen = value.Bool
			}
		case user.FieldIsRecharge:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_recharge", values[i])
			} else if value.Valid {
				u.IsRecharge = value.Bool
			}
		case user.FieldUserType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value.Valid {
				u.UserType = user.UserType(value.String)
			}
		case user.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				u.ParentID = value.Int64
			}
		case user.FieldPopVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pop_version", values[i])
			} else if value.Valid {
				u.PopVersion = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryVxAccounts queries the "vx_accounts" edge of the User entity.
func (u *User) QueryVxAccounts() *VXAccountQuery {
	return NewUserClient(u.config).QueryVxAccounts(u)
}

// QueryCollects queries the "collects" edge of the User entity.
func (u *User) QueryCollects() *CollectQuery {
	return NewUserClient(u.config).QueryCollects(u)
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *DeviceQuery {
	return NewUserClient(u.config).QueryDevices(u)
}

// QueryProfitSettings queries the "profit_settings" edge of the User entity.
func (u *User) QueryProfitSettings() *ProfitSettingQuery {
	return NewUserClient(u.config).QueryProfitSettings(u)
}

// QueryCostAccount queries the "cost_account" edge of the User entity.
func (u *User) QueryCostAccount() *CostAccountQuery {
	return NewUserClient(u.config).QueryCostAccount(u)
}

// QueryProfitAccount queries the "profit_account" edge of the User entity.
func (u *User) QueryProfitAccount() *ProfitAccountQuery {
	return NewUserClient(u.config).QueryProfitAccount(u)
}

// QueryCostBills queries the "cost_bills" edge of the User entity.
func (u *User) QueryCostBills() *CostBillQuery {
	return NewUserClient(u.config).QueryCostBills(u)
}

// QueryEarnBills queries the "earn_bills" edge of the User entity.
func (u *User) QueryEarnBills() *EarnBillQuery {
	return NewUserClient(u.config).QueryEarnBills(u)
}

// QueryMissionConsumeOrders queries the "mission_consume_orders" edge of the User entity.
func (u *User) QueryMissionConsumeOrders() *MissionConsumeOrderQuery {
	return NewUserClient(u.config).QueryMissionConsumeOrders(u)
}

// QueryMissionProduceOrders queries the "mission_produce_orders" edge of the User entity.
func (u *User) QueryMissionProduceOrders() *MissionProduceOrderQuery {
	return NewUserClient(u.config).QueryMissionProduceOrders(u)
}

// QueryRechargeOrders queries the "recharge_orders" edge of the User entity.
func (u *User) QueryRechargeOrders() *RechargeOrderQuery {
	return NewUserClient(u.config).QueryRechargeOrders(u)
}

// QueryVxSocials queries the "vx_socials" edge of the User entity.
func (u *User) QueryVxSocials() *VXSocialQuery {
	return NewUserClient(u.config).QueryVxSocials(u)
}

// QueryMissionBatches queries the "mission_batches" edge of the User entity.
func (u *User) QueryMissionBatches() *MissionBatchQuery {
	return NewUserClient(u.config).QueryMissionBatches(u)
}

// QueryUserDevices queries the "user_devices" edge of the User entity.
func (u *User) QueryUserDevices() *UserDeviceQuery {
	return NewUserClient(u.config).QueryUserDevices(u)
}

// QueryParent queries the "parent" edge of the User entity.
func (u *User) QueryParent() *UserQuery {
	return NewUserClient(u.config).QueryParent(u)
}

// QueryChildren queries the "children" edge of the User entity.
func (u *User) QueryChildren() *UserQuery {
	return NewUserClient(u.config).QueryChildren(u)
}

// QueryInvites queries the "invites" edge of the User entity.
func (u *User) QueryInvites() *InviteQuery {
	return NewUserClient(u.config).QueryInvites(u)
}

// QueryCampaignOrders queries the "campaign_orders" edge of the User entity.
func (u *User) QueryCampaignOrders() *CampaignOrderQuery {
	return NewUserClient(u.config).QueryCampaignOrders(u)
}

// QueryWallets queries the "wallets" edge of the User entity.
func (u *User) QueryWallets() *WalletQuery {
	return NewUserClient(u.config).QueryWallets(u)
}

// QueryIncomeBills queries the "income_bills" edge of the User entity.
func (u *User) QueryIncomeBills() *BillQuery {
	return NewUserClient(u.config).QueryIncomeBills(u)
}

// QueryOutcomeBills queries the "outcome_bills" edge of the User entity.
func (u *User) QueryOutcomeBills() *BillQuery {
	return NewUserClient(u.config).QueryOutcomeBills(u)
}

// QueryMissionProductions queries the "mission_productions" edge of the User entity.
func (u *User) QueryMissionProductions() *MissionProductionQuery {
	return NewUserClient(u.config).QueryMissionProductions(u)
}

// QueryMissions queries the "missions" edge of the User entity.
func (u *User) QueryMissions() *MissionQuery {
	return NewUserClient(u.config).QueryMissions(u)
}

// QueryIncomeTransferOrders queries the "income_transfer_orders" edge of the User entity.
func (u *User) QueryIncomeTransferOrders() *TransferOrderQuery {
	return NewUserClient(u.config).QueryIncomeTransferOrders(u)
}

// QueryOutcomeTransferOrders queries the "outcome_transfer_orders" edge of the User entity.
func (u *User) QueryOutcomeTransferOrders() *TransferOrderQuery {
	return NewUserClient(u.config).QueryOutcomeTransferOrders(u)
}

// QueryConsumeMissionOrders queries the "consume_mission_orders" edge of the User entity.
func (u *User) QueryConsumeMissionOrders() *MissionOrderQuery {
	return NewUserClient(u.config).QueryConsumeMissionOrders(u)
}

// QueryProduceMissionOrders queries the "produce_mission_orders" edge of the User entity.
func (u *User) QueryProduceMissionOrders() *MissionOrderQuery {
	return NewUserClient(u.config).QueryProduceMissionOrders(u)
}

// QueryLoginRecords queries the "login_records" edge of the User entity.
func (u *User) QueryLoginRecords() *LoginRecordQuery {
	return NewUserClient(u.config).QueryLoginRecords(u)
}

// QueryRenewalAgreements queries the "renewal_agreements" edge of the User entity.
func (u *User) QueryRenewalAgreements() *RenewalAgreementQuery {
	return NewUserClient(u.config).QueryRenewalAgreements(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("jpg_url=")
	builder.WriteString(u.JpgURL)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(u.Key)
	builder.WriteString(", ")
	builder.WriteString("secret=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("is_frozen=")
	builder.WriteString(fmt.Sprintf("%v", u.IsFrozen))
	builder.WriteString(", ")
	builder.WriteString("is_recharge=")
	builder.WriteString(fmt.Sprintf("%v", u.IsRecharge))
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", u.UserType))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ParentID))
	builder.WriteString(", ")
	builder.WriteString("pop_version=")
	builder.WriteString(u.PopVersion)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
