// Code generated by ent, DO NOT EDIT.

package incomewalletoperate

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the incomewalletoperate type in the database.
	Label = "income_wallet_operate"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldCurrentBalance holds the string denoting the current_balance field in the database.
	FieldCurrentBalance = "current_balance"
	// FieldLastEditedAt holds the string denoting the last_edited_at field in the database.
	FieldLastEditedAt = "last_edited_at"
	// FieldRejectReason holds the string denoting the reject_reason field in the database.
	FieldRejectReason = "reject_reason"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldApproveUserID holds the string denoting the approve_user_id field in the database.
	FieldApproveUserID = "approve_user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeApproveUser holds the string denoting the approve_user edge name in mutations.
	EdgeApproveUser = "approve_user"
	// Table holds the table name of the incomewalletoperate in the database.
	Table = "income_wallet_operates"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "income_wallet_operates"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ApproveUserTable is the table that holds the approve_user relation/edge.
	ApproveUserTable = "income_wallet_operates"
	// ApproveUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ApproveUserInverseTable = "users"
	// ApproveUserColumn is the table column denoting the approve_user relation/edge.
	ApproveUserColumn = "approve_user_id"
)

// Columns holds all SQL columns for incomewalletoperate fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldPhone,
	FieldType,
	FieldAmount,
	FieldReason,
	FieldCurrentBalance,
	FieldLastEditedAt,
	FieldRejectReason,
	FieldStatus,
	FieldApproveUserID,
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
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// DefaultReason holds the default value on creation for the "reason" field.
	DefaultReason string
	// DefaultCurrentBalance holds the default value on creation for the "current_balance" field.
	DefaultCurrentBalance int64
	// DefaultLastEditedAt holds the default value on creation for the "last_edited_at" field.
	DefaultLastEditedAt time.Time
	// DefaultRejectReason holds the default value on creation for the "reject_reason" field.
	DefaultRejectReason string
	// DefaultApproveUserID holds the default value on creation for the "approve_user_id" field.
	DefaultApproveUserID int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.IncomeWalletOperateType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.IncomeWalletOperateType) error {
	switch _type {
	case "unknown", "income_replacement", "income_deduct":
		return nil
	default:
		return fmt.Errorf("incomewalletoperate: invalid enum value for type field: %q", _type)
	}
}

const DefaultStatus enums.IncomeWalletOperateStatus = "pending"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.IncomeWalletOperateStatus) error {
	switch s {
	case "pending", "canceled", "succeed", "failed", "reject":
		return nil
	default:
		return fmt.Errorf("incomewalletoperate: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the IncomeWalletOperate queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByReason orders the results by the reason field.
func ByReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReason, opts...).ToFunc()
}

// ByCurrentBalance orders the results by the current_balance field.
func ByCurrentBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentBalance, opts...).ToFunc()
}

// ByLastEditedAt orders the results by the last_edited_at field.
func ByLastEditedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastEditedAt, opts...).ToFunc()
}

// ByRejectReason orders the results by the reject_reason field.
func ByRejectReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRejectReason, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByApproveUserID orders the results by the approve_user_id field.
func ByApproveUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApproveUserID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByApproveUserField orders the results by approve_user field.
func ByApproveUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApproveUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newApproveUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApproveUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ApproveUserTable, ApproveUserColumn),
	)
}
