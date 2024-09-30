// Code generated by ent, DO NOT EDIT.

package lotto

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the lotto type in the database.
	Label = "lotto"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTotalWeight holds the string denoting the total_weight field in the database.
	FieldTotalWeight = "total_weight"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the ended_at field in the database.
	FieldEndedAt = "ended_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// EdgeLottoPrizes holds the string denoting the lotto_prizes edge name in mutations.
	EdgeLottoPrizes = "lotto_prizes"
	// EdgeLottoRecords holds the string denoting the lotto_records edge name in mutations.
	EdgeLottoRecords = "lotto_records"
	// EdgeLottoUserCounts holds the string denoting the lotto_user_counts edge name in mutations.
	EdgeLottoUserCounts = "lotto_user_counts"
	// EdgeLottoGetCountRecords holds the string denoting the lotto_get_count_records edge name in mutations.
	EdgeLottoGetCountRecords = "lotto_get_count_records"
	// EdgeLottoChangeRules holds the string denoting the lotto_change_rules edge name in mutations.
	EdgeLottoChangeRules = "lotto_Change_rules"
	// Table holds the table name of the lotto in the database.
	Table = "lottos"
	// LottoPrizesTable is the table that holds the lotto_prizes relation/edge.
	LottoPrizesTable = "prizes"
	// LottoPrizesInverseTable is the table name for the LottoPrize entity.
	// It exists in this package in order to avoid circular dependency with the "lottoprize" package.
	LottoPrizesInverseTable = "prizes"
	// LottoPrizesColumn is the table column denoting the lotto_prizes relation/edge.
	LottoPrizesColumn = "lotto_id"
	// LottoRecordsTable is the table that holds the lotto_records relation/edge.
	LottoRecordsTable = "lotto_records"
	// LottoRecordsInverseTable is the table name for the LottoRecord entity.
	// It exists in this package in order to avoid circular dependency with the "lottorecord" package.
	LottoRecordsInverseTable = "lotto_records"
	// LottoRecordsColumn is the table column denoting the lotto_records relation/edge.
	LottoRecordsColumn = "lotto_id"
	// LottoUserCountsTable is the table that holds the lotto_user_counts relation/edge.
	LottoUserCountsTable = "lotto_user_counts"
	// LottoUserCountsInverseTable is the table name for the LottoUserCount entity.
	// It exists in this package in order to avoid circular dependency with the "lottousercount" package.
	LottoUserCountsInverseTable = "lotto_user_counts"
	// LottoUserCountsColumn is the table column denoting the lotto_user_counts relation/edge.
	LottoUserCountsColumn = "lotto_id"
	// LottoGetCountRecordsTable is the table that holds the lotto_get_count_records relation/edge.
	LottoGetCountRecordsTable = "lotto_get_count_records"
	// LottoGetCountRecordsInverseTable is the table name for the LottoGetCountRecord entity.
	// It exists in this package in order to avoid circular dependency with the "lottogetcountrecord" package.
	LottoGetCountRecordsInverseTable = "lotto_get_count_records"
	// LottoGetCountRecordsColumn is the table column denoting the lotto_get_count_records relation/edge.
	LottoGetCountRecordsColumn = "lotto_id"
	// LottoChangeRulesTable is the table that holds the lotto_Change_rules relation/edge.
	LottoChangeRulesTable = "lotto_chance_rules"
	// LottoChangeRulesInverseTable is the table name for the LottoChanceRule entity.
	// It exists in this package in order to avoid circular dependency with the "lottochancerule" package.
	LottoChangeRulesInverseTable = "lotto_chance_rules"
	// LottoChangeRulesColumn is the table column denoting the lotto_Change_rules relation/edge.
	LottoChangeRulesColumn = "lotto_id"
)

// Columns holds all SQL columns for lotto fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldTotalWeight,
	FieldStartedAt,
	FieldEndedAt,
	FieldStatus,
	FieldRemark,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultTotalWeight holds the default value on creation for the "total_weight" field.
	DefaultTotalWeight int64
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt time.Time
	// DefaultEndedAt holds the default value on creation for the "ended_at" field.
	DefaultEndedAt time.Time
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultStatus enums.LottoStatus = "unknown"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.LottoStatus) error {
	switch s {
	case "unknown", "normal", "canceled":
		return nil
	default:
		return fmt.Errorf("lotto: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Lotto queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTotalWeight orders the results by the total_weight field.
func ByTotalWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalWeight, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByEndedAt orders the results by the ended_at field.
func ByEndedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByLottoPrizesCount orders the results by lotto_prizes count.
func ByLottoPrizesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLottoPrizesStep(), opts...)
	}
}

// ByLottoPrizes orders the results by lotto_prizes terms.
func ByLottoPrizes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLottoPrizesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLottoRecordsCount orders the results by lotto_records count.
func ByLottoRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLottoRecordsStep(), opts...)
	}
}

// ByLottoRecords orders the results by lotto_records terms.
func ByLottoRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLottoRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLottoUserCountsCount orders the results by lotto_user_counts count.
func ByLottoUserCountsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLottoUserCountsStep(), opts...)
	}
}

// ByLottoUserCounts orders the results by lotto_user_counts terms.
func ByLottoUserCounts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLottoUserCountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLottoGetCountRecordsCount orders the results by lotto_get_count_records count.
func ByLottoGetCountRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLottoGetCountRecordsStep(), opts...)
	}
}

// ByLottoGetCountRecords orders the results by lotto_get_count_records terms.
func ByLottoGetCountRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLottoGetCountRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLottoChangeRulesCount orders the results by lotto_Change_rules count.
func ByLottoChangeRulesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLottoChangeRulesStep(), opts...)
	}
}

// ByLottoChangeRules orders the results by lotto_Change_rules terms.
func ByLottoChangeRules(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLottoChangeRulesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLottoPrizesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LottoPrizesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LottoPrizesTable, LottoPrizesColumn),
	)
}
func newLottoRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LottoRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LottoRecordsTable, LottoRecordsColumn),
	)
}
func newLottoUserCountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LottoUserCountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LottoUserCountsTable, LottoUserCountsColumn),
	)
}
func newLottoGetCountRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LottoGetCountRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LottoGetCountRecordsTable, LottoGetCountRecordsColumn),
	)
}
func newLottoChangeRulesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LottoChangeRulesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LottoChangeRulesTable, LottoChangeRulesColumn),
	)
}
