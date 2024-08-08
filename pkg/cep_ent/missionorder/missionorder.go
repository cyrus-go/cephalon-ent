// Code generated by ent, DO NOT EDIT.

package missionorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the missionorder type in the database.
	Label = "mission_order"
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
	// FieldMissionID holds the string denoting the mission_id field in the database.
	FieldMissionID = "mission_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSymbolID holds the string denoting the symbol_id field in the database.
	FieldSymbolID = "symbol_id"
	// FieldConsumeUserID holds the string denoting the consume_user_id field in the database.
	FieldConsumeUserID = "consume_user_id"
	// FieldConsumeAmount holds the string denoting the consume_amount field in the database.
	FieldConsumeAmount = "consume_amount"
	// FieldProduceUserID holds the string denoting the produce_user_id field in the database.
	FieldProduceUserID = "produce_user_id"
	// FieldProduceAmount holds the string denoting the produce_amount field in the database.
	FieldProduceAmount = "produce_amount"
	// FieldGasAmount holds the string denoting the gas_amount field in the database.
	FieldGasAmount = "gas_amount"
	// FieldMissionType holds the string denoting the mission_type field in the database.
	FieldMissionType = "mission_type"
	// FieldMissionBillingType holds the string denoting the mission_billing_type field in the database.
	FieldMissionBillingType = "mission_billing_type"
	// FieldCallWay holds the string denoting the call_way field in the database.
	FieldCallWay = "call_way"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldBuyDuration holds the string denoting the buy_duration field in the database.
	FieldBuyDuration = "buy_duration"
	// FieldPlanStartedAt holds the string denoting the plan_started_at field in the database.
	FieldPlanStartedAt = "plan_started_at"
	// FieldPlanFinishedAt holds the string denoting the plan_finished_at field in the database.
	FieldPlanFinishedAt = "plan_finished_at"
	// FieldExpiredWarningTime holds the string denoting the expired_warning_time field in the database.
	FieldExpiredWarningTime = "expired_warning_time"
	// FieldMissionBatchID holds the string denoting the mission_batch_id field in the database.
	FieldMissionBatchID = "mission_batch_id"
	// FieldMissionBatchNumber holds the string denoting the mission_batch_number field in the database.
	FieldMissionBatchNumber = "mission_batch_number"
	// FieldDeviceID holds the string denoting the device_id field in the database.
	FieldDeviceID = "device_id"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldSettledAmount holds the string denoting the settled_amount field in the database.
	FieldSettledAmount = "settled_amount"
	// FieldSettledCount holds the string denoting the settled_count field in the database.
	FieldSettledCount = "settled_count"
	// FieldTotalSettleCount holds the string denoting the total_settle_count field in the database.
	FieldTotalSettleCount = "total_settle_count"
	// FieldLatelySettledAt holds the string denoting the lately_settled_at field in the database.
	FieldLatelySettledAt = "lately_settled_at"
	// EdgeConsumeUser holds the string denoting the consume_user edge name in mutations.
	EdgeConsumeUser = "consume_user"
	// EdgeProduceUser holds the string denoting the produce_user edge name in mutations.
	EdgeProduceUser = "produce_user"
	// EdgeSymbol holds the string denoting the symbol edge name in mutations.
	EdgeSymbol = "symbol"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// EdgeMissionBatch holds the string denoting the mission_batch edge name in mutations.
	EdgeMissionBatch = "mission_batch"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// EdgeExtraServiceOrders holds the string denoting the extra_service_orders edge name in mutations.
	EdgeExtraServiceOrders = "extra_service_orders"
	// Table holds the table name of the missionorder in the database.
	Table = "mission_orders"
	// ConsumeUserTable is the table that holds the consume_user relation/edge.
	ConsumeUserTable = "mission_orders"
	// ConsumeUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ConsumeUserInverseTable = "users"
	// ConsumeUserColumn is the table column denoting the consume_user relation/edge.
	ConsumeUserColumn = "consume_user_id"
	// ProduceUserTable is the table that holds the produce_user relation/edge.
	ProduceUserTable = "mission_orders"
	// ProduceUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ProduceUserInverseTable = "users"
	// ProduceUserColumn is the table column denoting the produce_user relation/edge.
	ProduceUserColumn = "produce_user_id"
	// SymbolTable is the table that holds the symbol relation/edge.
	SymbolTable = "mission_orders"
	// SymbolInverseTable is the table name for the Symbol entity.
	// It exists in this package in order to avoid circular dependency with the "symbol" package.
	SymbolInverseTable = "symbols"
	// SymbolColumn is the table column denoting the symbol relation/edge.
	SymbolColumn = "symbol_id"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "order_id"
	// MissionBatchTable is the table that holds the mission_batch relation/edge.
	MissionBatchTable = "mission_orders"
	// MissionBatchInverseTable is the table name for the MissionBatch entity.
	// It exists in this package in order to avoid circular dependency with the "missionbatch" package.
	MissionBatchInverseTable = "mission_batches"
	// MissionBatchColumn is the table column denoting the mission_batch relation/edge.
	MissionBatchColumn = "mission_batch_id"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "mission_orders"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "mission_orders"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "device_id"
	// ExtraServiceOrdersTable is the table that holds the extra_service_orders relation/edge.
	ExtraServiceOrdersTable = "extra_service_orders"
	// ExtraServiceOrdersInverseTable is the table name for the ExtraServiceOrder entity.
	// It exists in this package in order to avoid circular dependency with the "extraserviceorder" package.
	ExtraServiceOrdersInverseTable = "extra_service_orders"
	// ExtraServiceOrdersColumn is the table column denoting the extra_service_orders relation/edge.
	ExtraServiceOrdersColumn = "mission_order_id"
)

// Columns holds all SQL columns for missionorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMissionID,
	FieldStatus,
	FieldSymbolID,
	FieldConsumeUserID,
	FieldConsumeAmount,
	FieldProduceUserID,
	FieldProduceAmount,
	FieldGasAmount,
	FieldMissionType,
	FieldMissionBillingType,
	FieldCallWay,
	FieldSerialNumber,
	FieldStartedAt,
	FieldFinishedAt,
	FieldBuyDuration,
	FieldPlanStartedAt,
	FieldPlanFinishedAt,
	FieldExpiredWarningTime,
	FieldMissionBatchID,
	FieldMissionBatchNumber,
	FieldDeviceID,
	FieldTotalAmount,
	FieldSettledAmount,
	FieldSettledCount,
	FieldTotalSettleCount,
	FieldLatelySettledAt,
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
	// DefaultMissionID holds the default value on creation for the "mission_id" field.
	DefaultMissionID int64
	// DefaultSymbolID holds the default value on creation for the "symbol_id" field.
	DefaultSymbolID int64
	// DefaultConsumeUserID holds the default value on creation for the "consume_user_id" field.
	DefaultConsumeUserID int64
	// DefaultConsumeAmount holds the default value on creation for the "consume_amount" field.
	DefaultConsumeAmount int64
	// DefaultProduceUserID holds the default value on creation for the "produce_user_id" field.
	DefaultProduceUserID int64
	// DefaultProduceAmount holds the default value on creation for the "produce_amount" field.
	DefaultProduceAmount int64
	// DefaultGasAmount holds the default value on creation for the "gas_amount" field.
	DefaultGasAmount int64
	// DefaultSerialNumber holds the default value on creation for the "serial_number" field.
	DefaultSerialNumber string
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt time.Time
	// DefaultFinishedAt holds the default value on creation for the "finished_at" field.
	DefaultFinishedAt time.Time
	// DefaultBuyDuration holds the default value on creation for the "buy_duration" field.
	DefaultBuyDuration int64
	// DefaultPlanStartedAt holds the default value on creation for the "plan_started_at" field.
	DefaultPlanStartedAt time.Time
	// DefaultPlanFinishedAt holds the default value on creation for the "plan_finished_at" field.
	DefaultPlanFinishedAt time.Time
	// DefaultExpiredWarningTime holds the default value on creation for the "expired_warning_time" field.
	DefaultExpiredWarningTime time.Time
	// DefaultMissionBatchID holds the default value on creation for the "mission_batch_id" field.
	DefaultMissionBatchID int64
	// DefaultMissionBatchNumber holds the default value on creation for the "mission_batch_number" field.
	DefaultMissionBatchNumber string
	// DefaultDeviceID holds the default value on creation for the "device_id" field.
	DefaultDeviceID int64
	// DefaultTotalAmount holds the default value on creation for the "total_amount" field.
	DefaultTotalAmount int64
	// DefaultSettledAmount holds the default value on creation for the "settled_amount" field.
	DefaultSettledAmount int64
	// DefaultSettledCount holds the default value on creation for the "settled_count" field.
	DefaultSettledCount int64
	// DefaultTotalSettleCount holds the default value on creation for the "total_settle_count" field.
	DefaultTotalSettleCount int64
	// DefaultLatelySettledAt holds the default value on creation for the "lately_settled_at" field.
	DefaultLatelySettledAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultStatus enums.MissionOrderStatus = "unknown"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s enums.MissionOrderStatus) error {
	switch s {
	case "unknown", "waiting", "canceled", "doing", "supplying", "failed", "succeed":
		return nil
	default:
		return fmt.Errorf("missionorder: invalid enum value for status field: %q", s)
	}
}

const DefaultMissionType enums.MissionType = "unknown"

// MissionTypeValidator is a validator for the "mission_type" field enum values. It is called by the builders before save.
func MissionTypeValidator(mt enums.MissionType) error {
	switch mt {
	case "unknown", "sd_time", "sd_pro_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sd_time_plan", "sd_pro_time_plan", "wt_time_plan", "jp_time_plan", "jp_dk_time_plan", "ssh_time_plan", "sd_tomato_time", "sd_tomato_time_plan", "sd_cmd_time", "sd_cmd_time_plan", "sd_tian_time", "sd_tian_time_plan", "sd_bingo_time", "sd_bingo_time_plan", "fooocus_time", "fooocus_time_plan", "fooocus_lan_que_time", "fooocus_lan_que_time_plan", "fooocus_sha_api_time", "fooocus_sha_api_time_plan", "tabby_time", "tabby_time_plan", "ollama_time", "ollama_time_plan", "jp_conda_time", "jp_conda_time_plan", "jp_ml_time", "jp_ml_time_plan", "sd_cat_time", "sd_cat_time_plan", "sd_fire_time", "sd_fire_time_plan", "comfyui_time", "comfyui_time_plan", "comfyui_pro_time", "comfyui_pro_time_plan", "comfyui_advance_time", "comfyui_advance_time_plan", "jp_dk_3_time", "jp_dk_3_time_plan", "sd_xl_time", "sd_xl_time_plan", "sd_chick_time", "sd_chick_time_plan", "ascend_time", "ascend_time_plan", "sd_wu_shan_time", "sd_wu_shan_time_plan", "sd_lang_time", "sd_lang_time_plan", "comfyui_ke_time", "comfyui_ke_time_plan", "chatchat_time", "chatchat_time_plan", "lora_time", "lora_time_plan", "fooocus_wu_time", "fooocus_wu_time_plan", "svd_back_time", "svd_back_time_plan", "sd_ji_time", "sd_ji_time_plan", "sd_shang_jin_time", "sd_shang_jin_time_plan", "sd_xiao_chun_time", "sd_xiao_chun_time_plan", "comfyui_wu_time", "comfyui_wu_time_plan", "comfyui_liu_time", "comfyui_liu_time_plan", "comfyui_li_time", "comfyui_li_time_plan", "comfyui_nenly_time", "comfyui_nenly_time_plan", "waiting_time", "waiting_time_plan":
		return nil
	default:
		return fmt.Errorf("missionorder: invalid enum value for mission_type field: %q", mt)
	}
}

const DefaultMissionBillingType enums.MissionBillingType = "unknown"

// MissionBillingTypeValidator is a validator for the "mission_billing_type" field enum values. It is called by the builders before save.
func MissionBillingTypeValidator(mbt enums.MissionBillingType) error {
	switch mbt {
	case "unknown", "time", "count", "hold", "volume", "time_plan_hour", "time_plan_day", "time_plan_week", "time_plan_month":
		return nil
	default:
		return fmt.Errorf("missionorder: invalid enum value for mission_billing_type field: %q", mbt)
	}
}

const DefaultCallWay enums.MissionCallWay = "unknown"

// CallWayValidator is a validator for the "call_way" field enum values. It is called by the builders before save.
func CallWayValidator(cw enums.MissionCallWay) error {
	switch cw {
	case "unknown", "api", "yuan_hui", "dev_platform":
		return nil
	default:
		return fmt.Errorf("missionorder: invalid enum value for call_way field: %q", cw)
	}
}

// OrderOption defines the ordering options for the MissionOrder queries.
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

// ByMissionID orders the results by the mission_id field.
func ByMissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySymbolID orders the results by the symbol_id field.
func BySymbolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbolID, opts...).ToFunc()
}

// ByConsumeUserID orders the results by the consume_user_id field.
func ByConsumeUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsumeUserID, opts...).ToFunc()
}

// ByConsumeAmount orders the results by the consume_amount field.
func ByConsumeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsumeAmount, opts...).ToFunc()
}

// ByProduceUserID orders the results by the produce_user_id field.
func ByProduceUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProduceUserID, opts...).ToFunc()
}

// ByProduceAmount orders the results by the produce_amount field.
func ByProduceAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProduceAmount, opts...).ToFunc()
}

// ByGasAmount orders the results by the gas_amount field.
func ByGasAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGasAmount, opts...).ToFunc()
}

// ByMissionType orders the results by the mission_type field.
func ByMissionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionType, opts...).ToFunc()
}

// ByMissionBillingType orders the results by the mission_billing_type field.
func ByMissionBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBillingType, opts...).ToFunc()
}

// ByCallWay orders the results by the call_way field.
func ByCallWay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallWay, opts...).ToFunc()
}

// BySerialNumber orders the results by the serial_number field.
func BySerialNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSerialNumber, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByBuyDuration orders the results by the buy_duration field.
func ByBuyDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuyDuration, opts...).ToFunc()
}

// ByPlanStartedAt orders the results by the plan_started_at field.
func ByPlanStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlanStartedAt, opts...).ToFunc()
}

// ByPlanFinishedAt orders the results by the plan_finished_at field.
func ByPlanFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlanFinishedAt, opts...).ToFunc()
}

// ByExpiredWarningTime orders the results by the expired_warning_time field.
func ByExpiredWarningTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiredWarningTime, opts...).ToFunc()
}

// ByMissionBatchID orders the results by the mission_batch_id field.
func ByMissionBatchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchID, opts...).ToFunc()
}

// ByMissionBatchNumber orders the results by the mission_batch_number field.
func ByMissionBatchNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBatchNumber, opts...).ToFunc()
}

// ByDeviceID orders the results by the device_id field.
func ByDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceID, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// BySettledAmount orders the results by the settled_amount field.
func BySettledAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettledAmount, opts...).ToFunc()
}

// BySettledCount orders the results by the settled_count field.
func BySettledCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettledCount, opts...).ToFunc()
}

// ByTotalSettleCount orders the results by the total_settle_count field.
func ByTotalSettleCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalSettleCount, opts...).ToFunc()
}

// ByLatelySettledAt orders the results by the lately_settled_at field.
func ByLatelySettledAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatelySettledAt, opts...).ToFunc()
}

// ByConsumeUserField orders the results by consume_user field.
func ByConsumeUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConsumeUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByProduceUserField orders the results by produce_user field.
func ByProduceUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProduceUserStep(), sql.OrderByField(field, opts...))
	}
}

// BySymbolField orders the results by symbol field.
func BySymbolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSymbolStep(), sql.OrderByField(field, opts...))
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

// ByMissionBatchField orders the results by mission_batch field.
func ByMissionBatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionBatchStep(), sql.OrderByField(field, opts...))
	}
}

// ByMissionField orders the results by mission field.
func ByMissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMissionStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeviceField orders the results by device field.
func ByDeviceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStep(), sql.OrderByField(field, opts...))
	}
}

// ByExtraServiceOrdersCount orders the results by extra_service_orders count.
func ByExtraServiceOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExtraServiceOrdersStep(), opts...)
	}
}

// ByExtraServiceOrders orders the results by extra_service_orders terms.
func ByExtraServiceOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExtraServiceOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newConsumeUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConsumeUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ConsumeUserTable, ConsumeUserColumn),
	)
}
func newProduceUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProduceUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProduceUserTable, ProduceUserColumn),
	)
}
func newSymbolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SymbolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SymbolTable, SymbolColumn),
	)
}
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillsTable, BillsColumn),
	)
}
func newMissionBatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionBatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionBatchTable, MissionBatchColumn),
	)
}
func newMissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MissionTable, MissionColumn),
	)
}
func newDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
	)
}
func newExtraServiceOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExtraServiceOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExtraServiceOrdersTable, ExtraServiceOrdersColumn),
	)
}
