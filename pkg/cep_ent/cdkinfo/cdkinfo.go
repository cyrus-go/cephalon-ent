// Code generated by ent, DO NOT EDIT.

package cdkinfo

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the cdkinfo type in the database.
	Label = "cdk_info"
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
	// FieldIssueUserID holds the string denoting the issue_user_id field in the database.
	FieldIssueUserID = "issue_user_id"
	// FieldCdkNumber holds the string denoting the cdk_number field in the database.
	FieldCdkNumber = "cdk_number"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldGetCep holds the string denoting the get_cep field in the database.
	FieldGetCep = "get_cep"
	// FieldGetTime holds the string denoting the get_time field in the database.
	FieldGetTime = "get_time"
	// FieldBillingType holds the string denoting the billing_type field in the database.
	FieldBillingType = "billing_type"
	// FieldExpiredAt holds the string denoting the expired_at field in the database.
	FieldExpiredAt = "expired_at"
	// FieldUseTimes holds the string denoting the use_times field in the database.
	FieldUseTimes = "use_times"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeIssueUser holds the string denoting the issue_user edge name in mutations.
	EdgeIssueUser = "issue_user"
	// Table holds the table name of the cdkinfo in the database.
	Table = "cdk_infos"
	// IssueUserTable is the table that holds the issue_user relation/edge.
	IssueUserTable = "cdk_infos"
	// IssueUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	IssueUserInverseTable = "users"
	// IssueUserColumn is the table column denoting the issue_user relation/edge.
	IssueUserColumn = "issue_user_id"
)

// Columns holds all SQL columns for cdkinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIssueUserID,
	FieldCdkNumber,
	FieldType,
	FieldGetCep,
	FieldGetTime,
	FieldBillingType,
	FieldExpiredAt,
	FieldUseTimes,
	FieldStatus,
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
	// DefaultIssueUserID holds the default value on creation for the "issue_user_id" field.
	DefaultIssueUserID int64
	// DefaultCdkNumber holds the default value on creation for the "cdk_number" field.
	DefaultCdkNumber string
	// DefaultGetCep holds the default value on creation for the "get_cep" field.
	DefaultGetCep int64
	// DefaultGetTime holds the default value on creation for the "get_time" field.
	DefaultGetTime int64
	// DefaultExpiredAt holds the default value on creation for the "expired_at" field.
	DefaultExpiredAt time.Time
	// DefaultUseTimes holds the default value on creation for the "use_times" field.
	DefaultUseTimes int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultType enums.CDKType = "unknown"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.CDKType) error {
	switch _type {
	case "unknown", "get_cep", "get_gpu_use":
		return nil
	default:
		return fmt.Errorf("cdkinfo: invalid enum value for type field: %q", _type)
	}
}

const DefaultBillingType enums.MissionBillingType = "unknown"

// BillingTypeValidator is a validator for the "billing_type" field enum values. It is called by the builders before save.
func BillingTypeValidator(bt enums.MissionBillingType) error {
	switch bt {
	case "unknown", "time", "count", "hold", "volume", "time_plan_hour", "time_plan_day", "time_plan_week", "time_plan_month":
		return nil
	default:
		return fmt.Errorf("cdkinfo: invalid enum value for billing_type field: %q", bt)
	}
}

const DefaultStatus enums.CDKStatus = "unknown"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.CDKStatus) error {
	switch s {
	case "unknown", "normal", "freeze", "used", "canceled":
		return nil
	default:
		return fmt.Errorf("cdkinfo: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the CDKInfo queries.
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

// ByIssueUserID orders the results by the issue_user_id field.
func ByIssueUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssueUserID, opts...).ToFunc()
}

// ByCdkNumber orders the results by the cdk_number field.
func ByCdkNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCdkNumber, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByGetCep orders the results by the get_cep field.
func ByGetCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGetCep, opts...).ToFunc()
}

// ByGetTime orders the results by the get_time field.
func ByGetTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGetTime, opts...).ToFunc()
}

// ByBillingType orders the results by the billing_type field.
func ByBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBillingType, opts...).ToFunc()
}

// ByExpiredAt orders the results by the expired_at field.
func ByExpiredAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiredAt, opts...).ToFunc()
}

// ByUseTimes orders the results by the use_times field.
func ByUseTimes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseTimes, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByIssueUserField orders the results by issue_user field.
func ByIssueUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssueUserStep(), sql.OrderByField(field, opts...))
	}
}
func newIssueUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssueUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, IssueUserTable, IssueUserColumn),
	)
}
