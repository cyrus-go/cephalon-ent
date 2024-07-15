// Code generated by ent, DO NOT EDIT.

package invite

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the invite type in the database.
	Label = "invite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldInviteCode holds the string denoting the invite_code field in the database.
	FieldInviteCode = "invite_code"
	// FieldShareCep holds the string denoting the share_cep field in the database.
	FieldShareCep = "share_cep"
	// FieldRegCep holds the string denoting the reg_cep field in the database.
	FieldRegCep = "reg_cep"
	// FieldFirstRechargeCep holds the string denoting the first_recharge_cep field in the database.
	FieldFirstRechargeCep = "first_recharge_cep"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCampaignID holds the string denoting the campaign_id field in the database.
	FieldCampaignID = "campaign_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCampaign holds the string denoting the campaign edge name in mutations.
	EdgeCampaign = "campaign"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// Table holds the table name of the invite in the database.
	Table = "invites"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "invites"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CampaignTable is the table that holds the campaign relation/edge.
	CampaignTable = "invites"
	// CampaignInverseTable is the table name for the Campaign entity.
	// It exists in this package in order to avoid circular dependency with the "campaign" package.
	CampaignInverseTable = "campaigns"
	// CampaignColumn is the table column denoting the campaign relation/edge.
	CampaignColumn = "campaign_id"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "invite_id"
)

// Columns holds all SQL columns for invite fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldInviteCode,
	FieldShareCep,
	FieldRegCep,
	FieldFirstRechargeCep,
	FieldType,
	FieldUserID,
	FieldCampaignID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultInviteCode holds the default value on creation for the "invite_code" field.
	DefaultInviteCode string
	// DefaultShareCep holds the default value on creation for the "share_cep" field.
	DefaultShareCep int64
	// DefaultRegCep holds the default value on creation for the "reg_cep" field.
	DefaultRegCep int64
	// DefaultFirstRechargeCep holds the default value on creation for the "first_recharge_cep" field.
	DefaultFirstRechargeCep int64
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultCampaignID holds the default value on creation for the "campaign_id" field.
	DefaultCampaignID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.InviteType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.InviteType) error {
	switch _type {
	case "unknown", "share_register", "app_invite":
		return nil
	default:
		return fmt.Errorf("invite: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Invite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByInviteCode orders the results by the invite_code field.
func ByInviteCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInviteCode, opts...).ToFunc()
}

// ByShareCep orders the results by the share_cep field.
func ByShareCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShareCep, opts...).ToFunc()
}

// ByRegCep orders the results by the reg_cep field.
func ByRegCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegCep, opts...).ToFunc()
}

// ByFirstRechargeCep orders the results by the first_recharge_cep field.
func ByFirstRechargeCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstRechargeCep, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCampaignID orders the results by the campaign_id field.
func ByCampaignID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCampaignID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCampaignField orders the results by campaign field.
func ByCampaignField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCampaignStep(), sql.OrderByField(field, opts...))
	}
}

// ByBillsCount orders the results by bills count.
func ByBillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBillsStep(), opts...)
	}
}

// ByBills orders the results by bills terms.
func ByBills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newCampaignStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CampaignInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CampaignTable, CampaignColumn),
	)
}
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
	)
}
